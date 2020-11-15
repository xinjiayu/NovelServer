package service

import (
	"NovelServer/app/model"
	"NovelServer/library/utils"
	"encoding/base64"
	"github.com/PuerkitoBio/goquery"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

func (bs *BookService) BookCatalog(sourceCode, bookURL string) []model.BookCatalog {

	//传统的URL是经过编码的需要解码
	tmpURL, _ := base64.URLEncoding.DecodeString(bookURL)
	bookURL = string(tmpURL)
	var catalogListData []model.BookCatalog

	config := new(model.SourceConfig)
	config = bs.SourceConfigInfo[sourceCode]
	glog.Info(bookURL)

	//解析JSON数据

	if config.Catalog.DataType == "json" {

		jsonData := getRemoteJsonData(bookURL)
		catalogList := jsonData.GetArray("items")
		for _, data := range catalogList {
			catalog := gconv.Map(data)
			var bookCatalog model.BookCatalog

			bookCatalog.SourcesCode = config.SourcesCode
			//通过文字模板的处理，进行参数替换配置
			paramDataMap := make(map[string]interface{})
			paramDataMap["CatalogUrlParam"] = catalog[config.Catalog.CatalogUrlParam]
			bookCatalog.Url = utils.StringLiteralTemplate(config.Catalog.CatalogUrl, paramDataMap)
			bookCatalog.Chapter = gconv.Int(catalog["chapter_id"])
			bookCatalog.Title = gconv.String(catalog[config.Catalog.Title.Rule])

			bookCatalog.Url = base64.URLEncoding.EncodeToString([]byte(bookCatalog.Url))

			//进行正则过滤处理
			bookCatalog.Title = utils.NormFormat(bookCatalog.Title, config.Catalog.Title.Filter)
			bookCatalog.Url = utils.NormFormat(bookCatalog.Url, config.Catalog.Url.Filter)

			catalogListData = append(catalogListData, bookCatalog)
		}

		return catalogListData
	}

	//解析HTML页面数据
	doc := getHtmlDoc(bookURL)
	bs.formURL = bookURL

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
func getRemoteJsonData(webUrl string) *gjson.Json {
	c := g.Client()
	c.SetHeader("Span-Id", "0.0.1")
	c.SetHeader("Trace-Id", "NBC56410N97LJ016FQA")
	if r, e := c.Get(webUrl); e != nil {
		glog.Error(e)
		return nil
	} else {
		defer r.Close()
		jsonData := gjson.New(r.ReadAllString())
		return jsonData
	}
}
