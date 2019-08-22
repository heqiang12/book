package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
	Id       int
	Name	 string
	Url		 string
	SearchUrl string
	ContentRule string
	ChapterRule string
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
func (u *SourceController) Data()  {
	u.Ctx.WriteString("年水电费")
	//url := "http://www.xbiquge.la/13/13959/"
	////html := httplib.Get(url)
	////dom,err:=goquery.NewDocumentFromReader(strings.NewReader(html))
	//dom,err := goquery.NewDocument(url)
	//if err!=nil{
	//	log.Fatalln(err)
	//}
	//dom.Find("#info > h1").Each(func(i int, selection *goquery.Selection) {
	//	//fmt.Println(selection.Text())
	//	u.Ctx.WriteString(selection.Text())
	//})
}
