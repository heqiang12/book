package controllers

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
)

func init() {
	orm.RegisterModel(new(Source))
}

// Operations about Users
type SourceController struct {
	beego.Controller
}

type Source struct {
	Id               int
	Name             string
	Url              string
	SearchUrl        string
	SearchCodeBefore string
	SearchCode       string
	SearchRule       string
	AuthorRule       string
	ContentRule      string
	ChapterRule      string
	Status           int
}

//@router /index [get]
func (u *SourceController) Index() {
	sid := u.GetString("sid")
	//u.Ctx.WriteString("传达室")
	id, _ := strconv.Atoi(sid)
	//models.GetSource(1)
	//u.Ctx.WriteString(id)
	o := orm.NewOrm()
	source := Source{Id: id}

	err := o.Read(&source)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		//fmt.Println(source.Id, source.Name)
		u.Ctx.WriteString(source.Name)
	}
}

//@router /data [get]
func (u *SourceController) Data() {
	//u.Ctx.WriteString("年水电费")
	url := "https://www.daocaorenshuwu.com/plus/search.php?q=鬼"
	//html := httplib.Get(url)
	//dom,err:=goquery.NewDocumentFromReader(strings.NewReader(html))
	dom, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatalln(err)
	}
	dom.Find(".table-condensed tbody tr td:first-child > a[class=orange]").Each(func(i int, selection *goquery.Selection) {
		//fmt.Println(selection.Text())
		//u.Ctx.WriteString(selection.Text())
		title, _ := selection.Attr("title")
		url, _ := selection.Attr("href")
		fmt.Println(title)
		fmt.Println(url)
	})
}

//@router /search [get]
//@Description 获取搜索关键词，找出需要的书籍
func (u *SourceController) Search() {
	kwd := u.GetString("kwd")
	//kwd := "圣墟"

	//从库中获取搜索源
	o := orm.NewOrm()
	// 获取 QuerySeter 对象，user 为表名
	qs := o.QueryTable("source")
	qs = qs.Filter("status", 1)
	var sources []*Source
	//查询且赋值
	qs.All(&sources)
	for _, value := range sources {
		searchUrl := u.UrlCoding(value.SearchUrl, value.SearchCode, value.SearchCodeBefore , kwd)
		//fmt.Println(searchUrl)
		u.getBookList(searchUrl,value.SearchRule,value.AuthorRule)
	}

}

//@Description 组装搜索地址
func (u *SourceController) UrlCoding(url string, code string, beforeCode string, kwd string) string {
	//var endCode string
	//if(code != "" && beforeCode != ""){
	//	endCode = u.ConvertToString(kwd,"utf-8","gbk")
	//	fmt.Println(endCode)
	//}else{
	//	endCode = kwd
	//}

	//SearchUrl := url+endCode
	SearchUrl := url+kwd
	return SearchUrl
}

//@Description 编码转换
//func (u *SourceController) ConvertToString(src string, srcCode string, tagCode string) string {
//	srcCoder := mahonia.NewDecoder(srcCode)
//	srcResult := srcCoder.ConvertString(src)
//	tagCoder := mahonia.NewDecoder(tagCode)
//	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
//	result := string(cdata)
//	return result
//}

//@Description 爬取搜索结果页，返回前三个书籍
func (u *SourceController) getBookList(url string , rule string , authorRule string)  {
	dom, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatalln(err)
	}
	//小说名称+链接+作者
	//var data []string
	data :=  make(map[string]map[string]string)
	dom.Find(rule).Each(func(i int, selection *goquery.Selection) {
		//书名
		bookName := selection.Text()
		//作者
		authorName := selection.Parent().NextAllFiltered(authorRule).Text()
		//书籍链接
		bookUrl,_ := selection.Attr("href")

		data[bookName] = make(map[string]string)
		data[bookName]["bookName"] = bookName
		data[bookName]["authorName"] = authorName
		data[bookName]["bookUrl"] = bookUrl
	})
}
