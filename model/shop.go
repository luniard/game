package model

import (
	"github.com/astaxie/beego/orm"
)

type Shop struct {
	Id    int
	Name  string
	Goods string
}

func init() {
	orm.RegisterModel(new(Shop))
}

func AddShop(name string, goods string) error {
	o := orm.NewOrm()

	shop := &Shop{Name: name, Goods: goods}

	_, err := o.Insert(shop)
	if err != nil {
		return err
	}
	return nil
}
