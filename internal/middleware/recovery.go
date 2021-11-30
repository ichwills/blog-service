package middleware

import (
	"Pro/blog-service/global"
	"Pro/blog-service/pkg/app"
	"Pro/blog-service/pkg/email"
	"Pro/blog-service/pkg/errcode"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	defailMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf(c, "panic recover err: %v", err)
				err := defailMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("error throw: %d", time.Now().Unix()),
					fmt.Sprintf("error message: %v", err),
				)
				if err != nil {
					global.Logger.Panicf(c, "mail.SendMail err: %v", err)
				}

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
