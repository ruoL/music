package models

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id     int64
	Name   string    `orm:"size(8)"`
	Mobile string    `orm:"size(16)"`
	Time   time.Time `orm:"index;auto_now_add;type(datetime)"`
}

type Ticket struct {
	Id   int64
	Name string `orm:"size(32)"`
	Get  int
}

const DB_NAME = "root:28a316c6@/music?charset=utf8"

func RegisterDB() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysl", DB_NAME, 10)
	orm.RegisterModel(new(User))
	err := CreateTable()
	if err != nil {
		beego.Error(err)
	}
}

func CreateTable() error {
	name := "default"
	force := false
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)
	return err
}

func AddUser(name, mobile string) error {
	o := orm.NewOrm()
	user := &User{
		Name:   name,
		Mobile: mobile,
	}
	created, _, err := o.ReadOrCreate(user, "Name", "Mobile")
	if err == nil {
		if created {
			return nil
		} else {
			return errors.New("已经提交过了!赶紧去领取门票吧!")
		}
	}
	return err
}

func GetCountTicket() (int64, error) {
	o := orm.NewOrm()
	cnt, err := o.QueryTable("Ticket").Count()
	return cnt, err
}

func GetTicket()(*Ticket, error) {
    o:=orm.NewOrm()

}

func