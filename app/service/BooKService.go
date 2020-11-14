package service

import (
	"NovelServer/app/model"
	"github.com/PuerkitoBio/goquery"
	"github.com/gogf/gf/os/glog"
	"net/http"
)

type BookService struct {
	formURL          string
	SourceConfigInfo map[string]*model.SourceConfig
	SourceList       []*model.Source
}

func NewSetupData() *BookService {

	bookService := new(BookService)

	bookService.SourceConfigInfo, bookService.SourceList = initSourceData()

	return bookService
}

//getHtmlDoc 获取远程html内容
func getHtmlDoc(url string) *goquery.Document {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		glog.Error(err)
		return nil
	}

	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		glog.Error(err)
		return nil
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		glog.Error(err)
	}

	return doc
}
