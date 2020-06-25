package controller

import (
	"UtilsApi/model"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"
	"github.com/zhshch2002/goreq"
	"net/http"
	"net/url"
	"strings"
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

// @Summary 获取网站图标
// @Description 获取网站图标
// @Tags Basic
// @ID favicon
// @Produce  image/*
// @Param url query string true "目标网站网址"
// @Success 200 {object} []byte
// @Failure 200 {object} controller.JsonResult
// @Router /v1/favicon [get]
func favicon(ctx *gin.Context) {
	u := ctx.Query("url")
	if u == "" {
		Ret(ctx).MissParam("缺少url参数")
		return
	}
	a, err := url.Parse(u)
	if err != nil {
		Ret(ctx).ServerErr(err)
		return
	}
	h, err := goreq.Get(u).Do().HTML()
	if err != nil {
		Ret(ctx).ServerErr(err)
		return
	}
	b, _ := a.Parse("/favicon.ico")
	icon := b.String()

	h.Find("link[rel]").EachWithBreak(func(i int, sel *goquery.Selection) bool {
		rels := strings.Split(sel.AttrOr("rel", ""), " ")
		for _, v := range rels {
			if v == "icon" && sel.AttrOr("href", "") != "" {
				b, err := a.Parse(sel.AttrOr("href", ""))
				if err != nil {
					continue
				}
				icon = b.String()
				return false
			}
		}
		return true
	})
	resp, err := goreq.Get(icon).Do().Resp()
	if err != nil {
		Ret(ctx).ServerErr(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		Ret(ctx).Err("非200响应")
		return
	}
	ctx.Header("content-type", resp.Header.Get("content-type"))
	ctx.Writer.WriteHeader(http.StatusOK)
	_, _ = ctx.Writer.Write(resp.NoDecodeBody)
}
