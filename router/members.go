package router

import (
	"go-admin/apis/members"
	"go-admin/middleware"
	jwt "go-admin/pkg/jwtauth"

	"github.com/gin-gonic/gin"
)

// 需认证的路由代码
func registerMembersRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("/members").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/:id", members.GetMembers)
		r.POST("", members.InsertMembers)
		r.PUT("", members.UpdateMembers)
		r.DELETE("/:id", members.DeleteMembers)
	}

	l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		l.GET("/membersList", members.GetMembersList)
	}

}
