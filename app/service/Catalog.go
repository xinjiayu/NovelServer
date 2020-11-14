package service

import (
	"NovelServer/app/model"
	"NovelServer/library/utils"
	"encoding/base64"
	"github.com/PuerkitoBio/goquery"
	"github.com/gogf/gf/os/glog"
)

func (bs *BookService) BookCatalog(sourceCode, bookURL string) []model.BookCatalog {

	//传统的URL是经过编码的需要解码
	tmpURL, _ := base64.URLEncoding.DecodeString(bookURL)
	bookURL = string(tmpURL)

	config := new(model.SourceConfig)
	config = bs.SourceConfigInfo[sourceCode]

	glog.Info(bookURL)
	doc := getHtmlDoc(bookURL)
	bs.formURL = bookURL

	var catalogListData []model.BookCatalog

	var bookCatalog model.BookCatalog

	if config.Catalog.NextFiltered != "" {

		doc.Find(config.Catalog.DataRange).NextFiltered(config.Catalog.NextFiltered).Find(config.Catalog.Find).Each(func(i int, s *goquery.Selection) {
			bookCatalog = bs.analysis(config, s)
			if bookCatalog.Title != "" {
				catalogListData = append(catalogListData, bookCatalog)

			}
		})

	} else {
		doc.Find(config.Catalog.DataRange).Each(func(i int, s *goquery.Selection) {

			bookCatalog = bs.analysis(config, s)
			if bookCatalog.Title != "" {
				catalogListData = append(catalogListData, bookCatalog)

			}
		})

	}

	return catalogListData
}

func (bs *BookService) analysis(config *model.SourceConfig, s *goquery.Selection) (bookCatalog model.BookCatalog) {
	bookCatalog.SourcesCode = config.SourcesCode
	if config.Catalog.Title.Type == "text" {
		bookCatalog.Title = s.Find(config.Catalog.Title.Rule).Text()
	} else {
		bookCatalog.Title, _ = s.Find(config.Catalog.Title.Rule).Attr(config.Catalog.Title.Type)
	}

	if config.Catalog.Url.Type == "text" {
		bookCatalog.Url = s.Find(config.Catalog.Url.Rule).Text()
	} else {
		if config.Catalog.Url.Rule == "" {
			bookCatalog.Url, _ = s.Attr(config.Catalog.Url.Type)

		} else {
			bookCatalog.Url, _ = s.Find(config.Catalog.Url.Rule).Attr(config.Catalog.Url.Type)

		}
	}

	if config.Catalog.Repair.Field != "" {
		switch config.Catalog.Repair.Field {
		case "Url":

			if config.Catalog.Repair.Value == "FormURL" {
				bookCatalog.Url = bs.formURL + bookCatalog.Url

			} else {
				bookCatalog.Url = config.Catalog.Repair.Value + bookCatalog.Url

			}
		}
	}

	if config.Weburl != "" {
		bookCatalog.Url = config.Weburl + bookCatalog.Url
	}

	bookCatalog.Url = base64.URLEncoding.EncodeToString([]byte(bookCatalog.Url))

	//进行正则过滤处理
	bookCatalog.Title = utils.NormFormat(bookCatalog.Title, config.Catalog.Title.Filter)
	bookCatalog.Url = utils.NormFormat(bookCatalog.Url, config.Catalog.Url.Filter)
	return
}
