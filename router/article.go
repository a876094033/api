package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/apis/article"
	"go-admin/middleware"
	jwt "go-admin/pkg/jwtauth"
)

// 需认证的路由代码
func registerArticleRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("/article").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/:id", article.GetArticle)
		r.POST("", article.InsertArticle)
		r.PUT("", article.UpdateArticle)
		r.DELETE("/:id", article.DeleteArticle)
	}

	l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		l.GET("/articleList", article.GetArticleList)
	}

}
