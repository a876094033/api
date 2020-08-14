package router

import (
	"go-admin/apis/borrowdiya"
	"go-admin/middleware"
	jwt "go-admin/pkg/jwtauth"

	"github.com/gin-gonic/gin"
)

// 需认证的路由代码
func registerBorrowDiyaRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("/borrowdiya").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/:diyaId", borrowdiya.GetBorrowDiya)
		r.POST("", borrowdiya.InsertBorrowDiya)
		r.PUT("", borrowdiya.UpdateBorrowDiya)
		r.DELETE("/:diyaId", borrowdiya.DeleteBorrowDiya)
	}

	l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		l.GET("/borrowdiyaList", borrowdiya.GetBorrowDiyaList)
	}

}
