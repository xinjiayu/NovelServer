package model

//Book 图书模型
type Book struct {
	BookId      string `xorm:"VARCHAR(100)"` //书的标识ID：书名|作者 进行base64进行编码
	BookName    string `xorm:"VARCHAR(100)"` //书的名称
	SourcesCode string `xorm:"VARCHAR(255)"` //源标识码 多个采用逗号分隔
	Author      string `xorm:"VARCHAR(50)"`  //作者
	ClassCode   string `xorm:"VARCHAR(50)"`  //分类
	Url         string `xorm:"VARCHAR(255)"` //书的可用来源
	Status      string `xorm:"CHAR(1)"`      //书的状态0，不可用；1，连接；2，完结
	Tag         string `xorm:"VARCHAR(255)"` //书的标签
	Description string `xorm:"VARCHAR(255)"` //书的介绍
	Img         string `xorm:"VARCHAR(255)"` //书的封面
	UpdateTime  int    `xorm:"INT(11)"`      //更新时间
	LastTitle   string `xorm:"VARCHAR(255)"` //最后一章的标题
	LastUrl     string `xorm:"VARCHAR(255)"` //最后一章的地址

}

//Catalog 书的目录
type BookCatalog struct {
	SourcesCode string `xorm:"VARCHAR(50)"`  //小说源标识码
	Url         string `xorm:"VARCHAR(200)"` //章节url
	Chapter     int    `xorm:"INT(11)"`      //章节编号
	Title       string `xorm:"VARCHAR(255)"` //文章标题
}

//Article struct
type BookArticle struct {
	Url     string `xorm:"VARCHAR(200)"` //文章标题
	Chapter int64  `xorm:"BIGINT(20)"`   //文章的ID
	Title   string `xorm:"VARCHAR(255)"` //文章标题
	Doc     string `xorm:"TEXT"`         //文章内容
}
