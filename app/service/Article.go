package service

import (
	"NovelServer/app/model"
	"NovelServer/library/utils"
	"encoding/base64"
	"github.com/PuerkitoBio/goquery"
	"github.com/gogf/gf/os/glog"
)

func (bs *BookService) BookArticle(sourceCode, bookArticleURL string) model.BookArticle {

	//传统的URL是经过编码的需要解码
	tmpURL, _ := base64.URLEncoding.DecodeString(bookArticleURL)
	bookArticleURL = string(tmpURL)

	config := new(model.SourceConfig)
	config = bs.SourceConfigInfo[sourceCode]
	glog.Info(bookArticleURL)
	doc := getHtmlDoc(bookArticleURL)

	var article model.BookArticle

	doc.Find(config.Article.DataRange).Each(func(i int, s *goquery.Selection) {
		if config.Catalog.Title.Type == "text" {
			article.Title = s.Find(config.Article.Title.Rule).Text()
		} else {
			article.Title, _ = s.Find(config.Article.Title.Rule).Attr(config.Article.Title.Type)
		}

		if config.Catalog.Title.Type == "text" {
			article.Doc = s.Find(config.Article.Doc.Rule).Text()
		} else {
			article.Doc, _ = s.Find(config.Article.Doc.Rule).Attr(config.Article.Doc.Type)
		}

		//进行正则过滤处理
		article.Title = utils.NormFormat(article.Title, config.Article.Title.Filter)
		article.Doc = utils.NormFormat(article.Doc, config.Article.Doc.Filter)
	})

	return article
}
