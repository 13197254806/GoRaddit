//package helloworld

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"test.com/helloworld/pkgs/snowflake"
	"test.com/helloworld/routes"

	"test.com/helloworld/dao/redis"

	"test.com/helloworld/dao/mysql"

	"test.com/helloworld/logger"
	"test.com/helloworld/settings"
)

func main() {
	// Init Config
	if err := settings.Init(); err != nil {
		fmt.Println(err)
	}

	// Init Zap
	if err := logger.Init(settings.Conf.LoggerConfig); err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("%#v", settings.Conf)

	// Init Mysql
	if err := mysql.Init(settings.Conf.MysqlConfig); err != nil {
		fmt.Println(err)
	}
	defer mysql.CloseDB()

	// Init Redis
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Println(err)
	}
	defer redis.CloseRedis()

	// Init SnowFlake
	if err := snowflake.Init(settings.Conf.StartTime, int64(settings.Conf.MachineID)); err != nil {
		fmt.Println(err)
	}
	//testID := snowflake.GenerateID()
	//fmt.Println(testID)

	// register routes
	r := routes.SetUp()

	// smoothing turn-off
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: r,
	}

	go func() {
		// create a goroutine for listening
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Println("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
}
