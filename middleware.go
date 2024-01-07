package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

// InitGinCorsMiddleware cors跨域问题解决方案
func InitGinCorsMiddleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		cors.New(cors.Config{
			AllowHeaders: []string{"Content-Type", "Authorization"},
			//允许前端访问后端响应中带的头部
			ExposeHeaders:    []string{"X-Jwt-Token"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				if strings.HasPrefix(origin, "http://localhost") {
					return true
				}
				//【修改】修改为公司域名
				return strings.Contains(origin, "公司域名.com")
			},
			MaxAge: 12 * time.Hour,
		}),
	}
}
