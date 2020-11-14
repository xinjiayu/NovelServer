package book

import (
	"NovelServer/app/service"
	"NovelServer/library/response"
	"github.com/gogf/gf/net/ghttp"
)

func Source(r *ghttp.Request) {
	bs := service.NewSetupData()
	response.JsonExit(r, 0, "当前支持的源列表", bs.SourceList)

}

func SearchBook(r *ghttp.Request) {

	bookName := r.GetString("bookName")
	sourceCode := r.GetString("sourceCode")

	bs := service.NewSetupData()
	data := bs.SearchBook(sourceCode, bookName)
	response.JsonExit(r, 0, "小说搜索结果", data)

}

func Catalog(r *ghttp.Request) {
	bookUrl := r.GetString("bookUrl")
	sourceCode := r.GetString("sourceCode")

	bs := service.NewSetupData()
	data := bs.BookCatalog(sourceCode, bookUrl)
	response.JsonExit(r, 0, "小说章节", data)
}

func Article(r *ghttp.Request) {
	bookArticleUrl := r.GetString("bookArticleUrl")
	sourceCode := r.GetString("sourceCode")

	bs := service.NewSetupData()
	data := bs.BookArticle(sourceCode, bookArticleUrl)
	response.JsonExit(r, 0, "小说文章内容", data)
}
