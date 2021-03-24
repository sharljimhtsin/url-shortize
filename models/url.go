package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
	"math/rand"
)

type Url struct {
	A string
	B string `orm:"pk"`
}

func (m *Url) TableName() string {
	return "url"
}

func init() {
	//DB
	_ = orm.RegisterDriver("sqlite3", orm.DRSqlite)
	_ = orm.RegisterDataBase("default", "sqlite3", "db/url-shortize.db")
	orm.RegisterModel(new(Url))
	orm.Debug = false
}

func GetAllUrls() []*Url {
	o := orm.NewOrm()
	m := new(Url)
	query := o.QueryTable(m)
	var urls []*Url
	_, _ = query.Limit(100).All(&urls)
	return urls
}

func GetUrlByHash(hash string) Url {
	o := orm.NewOrm()
	m := Url{B: hash}
	err := o.Read(&m)
	if err != nil {
		fmt.Println(err)
	}
	return m
}

func AddUrl(url string) (bool, *Url) {
	o := orm.NewOrm()
	var m Url
	m.A = url
	m.B = RandStringBytesRemainder(5)
	_, err := o.Insert(&m)
	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	return true, &m
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandStringBytesRemainder(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
