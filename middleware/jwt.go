package middleware

import (
	"context"
	"net/http"
	"todo_list/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

// JWTAuth中间件 ： 验证用户是否登录
func JWTAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 从Header获取Token
		token := string(c.GetHeader("Authorization"))

		if token == "" {
			utils.JSONError(c, http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(token)
		if err != nil {
			utils.JSONError(c, http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		// 将解析出来的用户信息塞进ctx
		c.Set("user_id", claims.ID)
		c.Set("user_name", claims.UserName)
		c.Next(ctx)
	}
}
