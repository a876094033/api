package router

import (
"github.com/gin-gonic/gin"
"go-admin/middleware"
"go-admin/apis/invest"
jwt "go-admin/pkg/jwtauth"
)

// 需认证的路由代码
func registerInvestRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

r := v1.Group("/invest").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
{
r.GET("/:investId", invest.GetInvest)
r.POST("", invest.InsertInvest)
r.PUT("", invest.UpdateInvest)
r.DELETE("/:investId", invest.DeleteInvest)
}

l := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
{
l.GET("/investList",invest.GetInvestList)
}

}
