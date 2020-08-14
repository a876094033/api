package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/apis/area"
	"go-admin/middleware"
	jwt "go-admin/pkg/jwtauth"
)

// 需认证的路由代码
func registerAreaRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("/area").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/:areaId", area.GetArea)
		r.POST("", area.InsertArea)
		r.PUT("", area.UpdateArea)
		r.DELETE("/:areaId", area.DeleteArea)
	}

	l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		l.GET("/areaList", area.GetAreaList)
	}

}
