package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

//var (
//	SourceList map[string]*Source
//)

func init() {
	//SourceList = make(map[string]*Source)
	//u := User{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
	//UserList["user_11111"] = &u
	//orm.RegisterModel(new(Source))
}

type Source struct {
	Id       int
	Name	 string
	Url		 string
	SearchUrl string
	ContentRule string
	ChapterRule string
}

func GetSource(id int)(s *Source) {
	//if u, ok := SourceList[id]; ok {
	//	return u, nil
	//}
	//return nil, errors.New("源不存在")
	o := orm.NewOrm()
	source := Source{Id: id}

	err := o.Read(&source)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(source.Id, source.Name)
	}
	return s
}
//
//func GetAllSources() map[string]*Source {
//	return SourceList
//}

