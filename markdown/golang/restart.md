# Golang Http 服务优雅重启

## 什么是优雅关机

优雅关机就是服务端关机命令发出后不是立即关机，而是等待当前还在处理的请求全部处理完毕后再退出程序，是一种对客户端友好的关机方式。而执行 Ctrl+C 关闭服务端时，会强制结束进程导致正在访问的请求出现问题。

## 如何实现优雅关机

http.Server 内置的 Shutdown\(\) 方法就支持优雅地关机，具体示例如下：

```go
package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        time.Sleep(5 * time.Second)
        c.String(http.StatusOK, "Welcome Gin Server")
    })

    srv := &http.Server{
        Addr:    ":8080",
        Handler: router,
    }

    go func() {
        // 开启一个goroutine启动服务
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("listen: %s\n", err)
        }
    }()

    // 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
    quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
    // kill 默认会发送 syscall.SIGTERM 信号
    // kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
    // kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
    // signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)  // 此处不会阻塞
    <-quit  // 阻塞在此，当接收到上述两种信号时才会往下执行
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
```

## 验证

* 上面的代码运行后会在本地的 8080 端口开启一个 web 服务，它只注册了一条路由/，后端服务会先 sleep 5 秒钟然后才返回响应信息。
* 我们按下 Ctrl+C 时会发送 syscall.SIGINT 来通知程序优雅地关机，具体做法如下： 1. 打开终端，编译并执行上面的代码 1. 打开一个浏览器，访问 127.0.0.1:8080/，此时浏览器白屏等待服务端返回响应。 1. 在终端迅速执行 Ctrl+C 命令给程序发送 syscall.SIGINT 信号 1. 此时程序并不立即退出而是等我们第 2 步的响应返回之后再退出，从而实现优雅关机。

