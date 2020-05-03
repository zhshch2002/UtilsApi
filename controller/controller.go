package controller

import (
	"UtilsApi/model"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/mmcdole/gofeed"
	"log"
)

func Run() {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})
	v1 := app.Group("/v1")
	v1.Get("/rss2json", func(ctx *fiber.Ctx) {
		rssAddr := ctx.Query("rss")
		if rssAddr == "" {
			ctx.JSON(model.Dict{
				"code": -1,
				"msg":  "缺少rss参数",
			})
			return
		}
		fp := gofeed.NewParser()
		feed, err := fp.ParseURL(rssAddr)
		if err != nil {
			ctx.JSON(model.Dict{
				"code": -1,
				"msg":  "解析rss错误",
			})
			log.Println(err)
			return
		}
		ctx.JSON(model.Dict{
			"code": 0,
			"msg":  "",
			"data": model.Dict{
				"title": feed.Title,
				"items": feed.Items,
			},
		})
	})

	err := app.Listen(4000)
	if err != nil {
		panic(err)
	}
}
