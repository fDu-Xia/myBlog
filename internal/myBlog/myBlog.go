package myBlog

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"myBlog/internal/middleware"
	"myBlog/internal/pkg/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "myBlog",
	Short: "myBlog is a blog system",
	Long: `myBlog is a blog system written by golang
                Complete documentation is available at https://github.com/fDu-Xia/myBlog`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//初始化全局Logger配置
		log.Init(logOptions())
		defer log.Sync()
		return run()
	},
	SilenceUsage: true,
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&configFile, "configs", "c", "", "The path to the blog system configuration file. Empty string for no configuration file.")
}

func run() error {
	gin.SetMode(viper.GetString("run-mode"))

	// 创建 Gin 引擎
	g := gin.New()
	g.Use(gin.Recovery(), middleware.RequestID(), middleware.NoCache, middleware.Cors())

	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 10003, "message": "Page not found."})
	})

	// 注册 /ping handler.
	g.GET("/ping", func(c *gin.Context) {
		log.C(c).Infow("ping function called")
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	// 创建Http服务器实例
	httpServer := &http.Server{Addr: viper.GetString("addr"), Handler: g}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalw(err.Error())
		}
	}()

	//捕获关闭信号，优雅关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Infow("Shutting down server ...")

	//设置10s时间处理未处理完的请求
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Errorw("Insecure Server forced to shutdown", "err", err)
		return err
	}

	log.Infow("Server Exit")

	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
