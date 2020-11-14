package model

type Source struct {
	SourcesCode string `xorm:"VARCHAR(100)"` //源的编码
	SourcesName string `xorm:"VARCHAR(100)"` //源的名称
}

//SourceConfig 小说源配置
type SourceConfig struct {
	SourcesCode string `json:"SourcesCode"`
	SourcesName string `json:"SourcesName"`
	Weburl      string `json:"Weburl"`
	Searchurl   string `json:"Searchurl"`
	Charset     string `json:"Charset"`
	Search      `json:"Search"`
	Catalog     `json:"Catalog"`
	Article     `json:"Article"`
}

//Search 小说搜索配置
type Search struct {
	DataRange       string `json:"Range"`
	Replace         `json:"Replace"`
	Repair          `json:"Repair"`
	BookName        Attribute `json:"BookName"`
	BookUrl         Attribute `json:"BookUrl"`
	BookImg         Attribute `json:"BookImg"`
	BookAuthor      Attribute `json:"BookAuthor"`
	BookDescription Attribute `json:"BookDescription"`
}

//Catalog 小说目录配置
type Catalog struct {
	DataRange    string `json:"Range"`
	Repair       `json:"Repair"`
	NextFiltered string    `json:"NextFiltered"`
	Find         string    `json:"Find"`
	Title        Attribute `json:"Title"`
	Url          Attribute `json:"Url"`
}

//Article 小说章节内容配置
type Article struct {
	DataRange string    `json:"Range"`
	Title     Attribute `json:"Title"`
	Doc       Attribute `json:"Doc"`
}

//Attribute 解析规则
type Attribute struct {
	Type   string `json:"Type"`
	Rule   string `json:"rule"`
	Filter string `json:"filter"`
}

//Replace 指定字段替换配置
type Replace struct {
	Field string `json:"Field"`
	Old   string `json:"Old"`
	New   string `json:"New"`
}

//Repair 指定字段左或是右进行补充内容
type Repair struct {
	Field    string `json:"Field"`
	Position string `json:"Position"`
	Value    string `json:"Value"`
}
