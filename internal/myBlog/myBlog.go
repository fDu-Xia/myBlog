package myBlog

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"myBlog/internal/middleware"
	"myBlog/internal/pkg/log"
	"net/http"
	"os"
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
	g.Use(gin.Recovery(), middleware.RequestID())

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
	if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalw(err.Error())
	}
	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
