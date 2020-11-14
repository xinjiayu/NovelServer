package router

import (
	"NovelServer/app/api/book"
	"NovelServer/app/service/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/v1", func(group *ghttp.RouterGroup) {

		group.Middleware(middleware.CORS, middleware.White, middleware.Auth)

		group.GET("/source", book.Source)
		group.GET("/searchbook", book.SearchBook)
		group.GET("/catalog", book.Catalog)
		group.GET("/article", book.Article)

	})
}
