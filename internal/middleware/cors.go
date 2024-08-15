package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"myBlog/internal/pkg/log"
)

func Cors() gin.HandlerFunc {
	var config cors.Config
	if err := viper.UnmarshalKey("Cors", &config); err != nil {
		log.Errorw("fail to decode cors config")
	}
	return cors.New(config)
}
