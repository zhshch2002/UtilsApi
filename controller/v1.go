package controller

import (
	"UtilsApi/model"
	"github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"
	"time"
)

// @Summary RSS转JSON
// @Description 将RSS地址返回的结果转换为JSON
// @Tags Basic
// @ID rss2json
// @Produce  json
// @Param rss query string true "RSS地址"
// @Success 200 {object} controller.JsonResult
// @Failure 200 {object} controller.JsonResult
// @Router /v1/rss2json [get]
func rss2json(ctx *gin.Context) {
	rssAddr := ctx.Query("rss")
	if rssAddr == "" {
		Ret(ctx).MissParam("缺少rss参数")
		return
	}
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(rssAddr)
	if err != nil {
		Ret(ctx).ServerErr(err)
		return
	}
	Ret(ctx).OK(model.Dict{
		"title": feed.Title,
		"items": feed.Items,
	})
}

// @Summary 获取请求者IP
// @Description 将RSS地址返回的结果转换为JSON
// @Tags Basic
// @ID ip
// @Produce  json
// @Success 200 {object} controller.JsonResult
// @Failure 200 {object} controller.JsonResult
// @Router /v1/ip [get]
func ip(ctx *gin.Context) {
	Ret(ctx).OK(model.Dict{
		"ip": ctx.ClientIP(),
	})
}

// @Summary 获取时间
// @Description 获取时间
// @Tags Basic
// @ID time
// @Produce  json
// @Success 200 {object} controller.JsonResult
// @Failure 200 {object} controller.JsonResult
// @Router /v1/time [get]
func getTime(ctx *gin.Context) {
	now := time.Now()
	Ret(ctx).OK(model.Dict{
		"unix":      now.Unix(),
		"unix_nano": now.UnixNano(),
		"second":    now.Second(),
		"minute":    now.Minute(),
		"hour":      now.Hour(),
		"week":      now.Weekday(),
		"day":       now.Day(),
		"month":     now.Month(),
		"year":      now.Year(),
	})
}
