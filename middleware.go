package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

// GinCors cors跨域问题解决方案
func GinCors() []gin.HandlerFunc {
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

// 解析 Authorization
func ParseUserToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header
		token, ok := header["Authorization"]
		if ok {
			global := puppy_data.GGlobalManager.GetGlobalByName(login.AuthSignKey)
			if global == nil {
				return
			}
			authSignKey := global.Value

			jwtClaims, err := login.ParseJwt(token[0], authSignKey)
			if jwtClaims != nil {
				if c.Keys == nil {
					c.Keys = make(map[string]interface{})
				}
				// 判断token 是否是登陆近期登陆获取到的
				UserToken := puppy_data.GUserTokenManagementManager.GetUserTokenManagementByUid(jwtClaims.Uid)
				if UserToken != nil {
					if UserToken.Token == token[0] {
						c.Set("uid", jwtClaims.Uid)
						c.Set("auth", jwtClaims.Auth)
						c.Next()
						return
					}
				}
				resp := &puppy_protocol.RespGeneral{}
				resp.SetGeneral(false, puppy_protocol.TokenDataErr, "")
				c.Abort()
				c.JSON(http.StatusOK, resp)
			} else {
				resp := &puppy_protocol.RespGeneral{}
				resp.SetGeneral(false, puppy_protocol.TokenDataErr, err.Error())
				c.Abort()
				c.JSON(http.StatusOK, resp)
			}
		} else {
			resp := &puppy_protocol.RespGeneral{}
			resp.SetGeneral(false, puppy_protocol.HeadError, "HeadError")
			c.Abort()
			c.JSON(http.StatusOK, resp)
		}
	}
}

// 权限
func ParseUserPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		lock := sync.Mutex{}
		lock.Lock()
		defer lock.Unlock()
		url := c.Request.URL

		keyUid, getOk := c.Keys["uid"]
		uid, assertOk := keyUid.(int32)
		if uid == 10001 {
			// admin无需验证
			c.Next()
			return
		}
		if getOk && assertOk {
			user := puppy_data.GUserManager.GetUserByUid(uid)
			if user.Rid == puppy_protocol.SuperAdmin {
				// 超级管理员无需验证权限
				c.Next()
				return
			}
			rolePer := puppy_data.GRolePermissionAccessManager.GetRolePermissionAccessByRoleIdPerPath(user.Rid, url.Path)
			per := puppy_data.GPermissionManager.GetPermissionByAPI(url.Path)
			if rolePer != nil && rolePer.Access == puppy_protocol.AccessTrue || per == nil {
				// rolePer != nil && rolePer.Access == puppy_protocol.AccessTrue 保证角色存在权限
				// per == nil 这个url未进行管理
				c.Next()
				return
			} else {
				log.Println("rolePer == nil || rolePer.Access != 1 || per != nil")
				resp := &puppy_protocol.RespGeneral{}
				resp.SetGeneral(false, puppy_protocol.PowerNotThrough, "")
				c.Abort()
				c.JSON(http.StatusOK, resp)
			}
		} else {
			log.Println("getOk && assertOk")
			resp := &puppy_protocol.RespGeneral{}
			resp.SetGeneral(false, puppy_protocol.PowerNotThrough, "")
			c.Abort()
			c.JSON(http.StatusOK, resp)
		}
	}
}
