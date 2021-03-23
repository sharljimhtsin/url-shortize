package models

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
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
	m.B = randStr(5)
	_, err := o.Insert(&m)
	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	return true, &m
}

func randStr(len int) string {
	buff := make([]byte, len)
	_, _ = rand.Read(buff)
	str := base64.StdEncoding.EncodeToString(buff)
	// Base 64 can be longer than len
	return str[:len]
}
