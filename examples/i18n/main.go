package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/examples/i18n/code"
)

func main() {
	engine := gin.I18nDefault(code.NewErrAction())
	// 测试绑定和校验数据
	engine.GET("/api/bind", func(ctx *gin.Context) {
		reqData := new(LoginReq)
		if ctx.ValidRequest(reqData) {
			return
		}
		ctx.JSON4Error(gin.ParamsBindCode)
	})
	// 测试返回数据
	engine.GET("/api/reply", func(ctx *gin.Context) {
		ctx.JSON4Item(&struct {
			A string `json:"a"`
			B string `json:"b"`
		}{
			A: "aaaa",
			B: "bbbb",
		})
	})
	// 测试分页数据
	engine.GET("/api/page1", func(ctx *gin.Context) {
		reqData := &struct {
			Account  string `form:"account" valid:"Must;Min(5);Max(20);ErrorCode(8201)"`
			Password string `form:"password" valid:"Must;Min(5);Max(20);ErrorCode(8202)"`
			Name     string `form:"name" valid:"Must;Min(1);Max(20);ErrorCode(7002)"`
			RoleId   uint   `form:"roleId" valid:"Must;Min(2);ErrorCode(7003)"`
		}{}
		finish, pageParams := ctx.ValidPageRequest(reqData)
		if finish {
			return
		}
		fmt.Println(pageParams)
		xx := make([]Resp, 0)
		xx = append(xx,
			Resp{
				Id:     1,
				Status: 1,
			}, Resp{
				Id:     1,
				Status: 2,
			}, Resp{
				Id:     2,
				Status: 1,
			},
		)
		ctx.JSON4Pagination(xx, 50)
	})
	// 测试分页数据
	engine.GET("/api/page2", func(ctx *gin.Context) {
		finish, pageParams := ctx.ValidPageRequest()
		if finish {
			return
		}
		fmt.Println(pageParams)
		xx := make([]Resp, 0)
		xx = append(xx,
			Resp{
				Id:     1,
				Status: 1,
			}, Resp{
				Id:     1,
				Status: 2,
			}, Resp{
				Id:     2,
				Status: 1,
			},
		)
		ctx.JSON4Pagination(&struct {
			TotalTitle string      `json:"totalTitle"`
			Data       interface{} `json:"data"`
		}{
			TotalTitle: "随便取的",
			Data:       xx,
		}, 50)
	})
	engine.Run(":8080")
}

type LoginReq struct {
	Account  string `form:"account" valid:"Must;Min(5);Max(20);ErrorCode(8201)"`
	Password string `form:"password" valid:"Must;Min(5);Max(20);ErrorCode(8202)"`
}

type Resp struct {
	Id     uint `json:"id" `
	Status uint `json:"status"`
}
