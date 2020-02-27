package api

import (
	"github.com/fanyiheng/go-web-demo/er"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	OK      = "ok"
	SUCCESS = "200"
	FAIL    = "400"
	ERROR   = "500"
)

type Resp interface{}

type CustomerHandlerFunc func(ctx *gin.Context) (Resp, error)

type HandlerGroup struct {
	Name string
}

func (g *HandlerGroup) Wrapper(handler CustomerHandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//todo 可能需要异常处理
		resp, err := handler(ctx)
		if resp == nil && err == nil {
			return
		}
		code, msg := SUCCESS, OK
		if err != nil {
			if usr, ok := err.(*er.ErrUsr); ok {
				msg = usr.Msg
				if usr.Log == true {
					code = ERROR
					logrus.WithFields(logrus.Fields{
						"module":    g.Name,
						"source":    usr.Source.Error(),
						"ip":        ctx.ClientIP(),
						"full_path": ctx.FullPath(),
					}).Error(msg)
				} else {
					code = FAIL
				}
			} else {
				code = ERROR
				msg = err.Error()
				logrus.WithFields(logrus.Fields{
					"module": g.Name,
					"ip":        ctx.ClientIP(),
					"full_path": ctx.FullPath(),
				}).Error(msg)
			}
			resp = nil
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  msg,
			"data": resp,
		})
	}
}
