package service

import (
	"NovelServer/app/model"
	"github.com/gogf/gf/os/gcache"
)

//cacheBookInfo 在缓存中保持一份完整的小说图书信息
func (bs *BookService) cacheBookInfo(bKey string, book *model.Book) {

	cacheData, _ := gcache.Get(bKey)
	if cacheData == nil {
		gcache.Set(bKey, book, 0)
		bs.BookList[bKey] = book

	} else {
		bookInfo := (cacheData).(*model.Book)
		if book.Img != "" {
			bookInfo.Img = book.Img
		}
		if book.LastUrl != "" {
			bookInfo.LastUrl = book.LastUrl
		}
		if book.LastTitle != "" {
			bookInfo.LastTitle = book.LastTitle
		}
		if book.Description != "" {
			bookInfo.Description = book.Description
		}
		gcache.Update(bKey, bookInfo)
		bs.BookList[bKey] = bookInfo

	}

}

func (bs *BookService) BookInfo(bKey string) *model.Book {
	return bs.BookList[bKey]
}
