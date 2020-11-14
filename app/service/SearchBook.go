package service

import (
	"NovelServer/app/model"
	"NovelServer/library/utils"
	"encoding/base64"
	"github.com/PuerkitoBio/goquery"
)

//SearchBook 搜索小说
func (bs *BookService) SearchBook(sourceCode, bookName string) []model.Book {

	config := new(model.SourceConfig)
	config = bs.SourceConfigInfo[sourceCode]

	url := config.Searchurl + bookName
	doc := getHtmlDoc(url)

	var bookList []model.Book
	doc.Find(config.Search.DataRange).Each(func(i int, s *goquery.Selection) {

		var book model.Book

		if config.Search.BookName.Type == "text" {
			book.BookName = s.Find(config.Search.BookName.Rule).Text()
		} else {
			book.BookName, _ = s.Find(config.Search.BookName.Rule).Attr(config.Search.BookName.Type)
		}

		if config.Search.BookUrl.Type == "text" {
			book.Url = s.Find(config.Search.BookUrl.Rule).Text()
		} else {
			book.Url, _ = s.Find(config.Search.BookUrl.Rule).Attr(config.Search.BookUrl.Type)
		}

		if config.Search.BookImg.Type == "text" {
			book.Img = s.Find(config.Search.BookImg.Rule).Text()
		} else {
			book.Img, _ = s.Find(config.Search.BookImg.Rule).Attr(config.Search.BookImg.Type)
		}

		if config.Search.BookAuthor.Type == "text" {
			book.Author = s.Find(config.Search.BookAuthor.Rule).Text()
		} else {
			book.Author, _ = s.Find(config.Search.BookAuthor.Rule).Attr(config.Search.BookAuthor.Type)
		}

		if config.Search.BookDescription.Type == "text" {
			book.Description = s.Find(config.Search.BookDescription.Rule).Text()
		} else {
			book.Description, _ = s.Find(config.Search.BookDescription.Rule).Attr(config.Search.BookDescription.Type)
		}

		//进行正则过滤处理
		book.BookName = utils.NormFormat(book.BookName, config.Search.BookName.Filter)
		book.Url = utils.NormFormat(book.Url, config.Search.BookUrl.Filter)
		book.Img = utils.NormFormat(book.Img, config.Search.BookImg.Filter)
		book.Author = utils.NormFormat(book.Author, config.Search.BookAuthor.Filter)
		book.Description = utils.NormFormat(book.Description, config.Search.BookDescription.Filter)

		if config.Search.Replace.Field != "" {
			switch config.Search.Replace.Field {
			case "BookUrl":
				book.Url = utils.FindAndReplace(book.Url, config.Search.Replace.Old, config.Search.Replace.New)
			}

		}

		if config.Weburl != "" {
			book.Url = config.Weburl + book.Url
		}

		if config.Search.Repair.Field != "" {
			switch config.Search.Repair.Field {
			case "BookUrl":
				if book.Url != "" {
					if config.Search.Repair.Position == "r" {
						book.Url = book.Url + config.Search.Repair.Value
					}
					if config.Search.Repair.Position == "l" {
						book.Url = config.Search.Repair.Value + book.Url
					}
				}

			}
		}

		book.Url = base64.URLEncoding.EncodeToString([]byte(book.Url))

		//获取图书唯一标识码，base64编码
		book.SourcesCode = config.SourcesCode
		bookId := book.BookName + "|" + book.Author
		book.BookId = base64.URLEncoding.EncodeToString([]byte(bookId))

		if book.BookName != "" && book.Url != "" {
			bookList = append(bookList, book)
		}

	})

	return bookList
}
