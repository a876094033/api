package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/apis/borrow"
	"go-admin/middleware"
	jwt "go-admin/pkg/jwtauth"
)

// 需认证的路由代码
func registerBorrowRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("/borrow").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/:id", borrow.GetBorrow)
		r.POST("", borrow.InsertBorrow)
		r.PUT("", borrow.UpdateBorrow)
		r.DELETE("/:id", borrow.DeleteBorrow)
		r.POST("/avatar", borrow.InsetBorrowAvatar)
	}

	l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		l.GET("/borrowList", borrow.GetBorrowList)
	}

}
