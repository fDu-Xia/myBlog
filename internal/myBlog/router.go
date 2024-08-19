package myBlog

import (
	"github.com/gin-gonic/gin"
	"myBlog/internal/middleware"
	"myBlog/internal/myBlog/controller/v1"
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

	userCtrl := v1.NewUserCtrl(store.S, authz)
	postCtrl := v1.NewPostCtrl(store.S)

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

		// 创建 posts 路由分组
		p1 := v1.Group("/posts", middleware.Authn())
		{
			p1.POST("", postCtrl.Create)             // 创建博客
			p1.GET(":postID", postCtrl.Get)          // 获取博客详情
			p1.PUT(":postID", postCtrl.Update)       // 更新用户
			p1.DELETE("", postCtrl.DeleteCollection) // 批量删除博客
			p1.GET("", postCtrl.List)                // 获取博客列表
			p1.DELETE(":postID", postCtrl.Delete)    // 删除博客
		}
	}

	return nil
}
