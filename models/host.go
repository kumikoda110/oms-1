package models

import (
	"github.com/astaxie/beego/orm"
	"oms/logger"
)

// Model Struct
type Host struct {
	Id       int
	Name     string `orm:"size(100)"`
	Addr     string `orm:"null"`
	Port     int    `orm:"default(22)"`
	PassWord string `orm:"null" json:"-"`
	KeyFile  string `orm:"null" json:"-"`

	Group *Group `orm:"rel(fk);null;on_delete(set_null)"`
	Tags  []*Tag `orm:"rel(m2m);null;rel_table(Tag)"`
}

func GetAllHost() []Host {
	var o = orm.NewOrm()
	host := new(Host)
	var hosts []Host
	_, err := o.QueryTable(host).RelatedSel().All(&hosts)
	if err != nil {
		logger.Logger.Println(err)
	}
	// 获取tags
	for i := 0; i < len(hosts); i++ {
		o.LoadRelated(&hosts[i], "Tags")
	}
	return hosts
}
