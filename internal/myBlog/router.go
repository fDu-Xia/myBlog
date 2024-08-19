package myBlog

import (
	"github.com/gin-gonic/gin"
	"myBlog/internal/middleware"
	"myBlog/internal/myBlog/controller/v1/user"
	"myBlog/internal/myBlog/store"
	"myBlog/internal/pkg/core"
	"myBlog/internal/pkg/errno"
	"myBlog/internal/pkg/log"
	"myBlog/pkg/auth"
)

// installRouters 安装 myBlog 接口路由.
func installRouters(g *gin.Engine) error {
	authz, err := auth.NewAuthz(store.S.DB())
	if err != nil {
		return err
	}
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// 注册 /ping handler.
	g.GET("/ping", func(c *gin.Context) {
		log.C(c).Infow("ping function called")
		core.WriteResponse(c, nil, gin.H{"status": "ok"})
	})

	userCtrl := user.New(store.S, authz)

	g.POST("/login", userCtrl.Login)

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		// 创建 users 路由分组
		u1 := v1.Group("/users")
		{
			u1.POST("/create", userCtrl.Create)
			u1.Use(middleware.Authn(), middleware.Authz(authz))
			u1.PUT(":name/change-password", userCtrl.ChangePassword)
			u1.GET(":name", userCtrl.Get)
		}
	}

	return nil
}
