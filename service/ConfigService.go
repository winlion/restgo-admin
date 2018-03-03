package service

import (
	"restgo-admin/entity"
	"restgo-admin/model"
	"restgo-admin/restgo"
)

type ConfigService struct{}

//根据userId 获取用户编号
func (service *ConfigService) FindOne(name string) entity.Config {
	var obj entity.Config
	orm := restgo.OrmEngin()
	orm.Id(name).Get(&obj)
	return obj
}

func (service *ConfigService) Query(arg model.ConfigArg) []entity.Config {
	var objs []entity.Config = make([]entity.Config, 0)
	orm := restgo.OrmEngin()
	t := orm.Where("1=1")
	if 0 < len(arg.Kword) {
		t = t.Where("name like ? or label like ?", "%"+arg.Kword+"%", "%"+arg.Kword+"%")
	}
	t.Limit(arg.GetPageFrom()).Find(&objs)
	return objs
}

func (service *ConfigService) All() []entity.Config {
	var objs []entity.Config = make([]entity.Config, 0)
	orm := restgo.OrmEngin()
	orm.Where("1=1").Find(&objs)
	return objs
}

func (service *ConfigService) Update(name string, value string) (int64, error) {
	var obj entity.Config
	obj.Name = name
	obj.Value = value
	orm := restgo.OrmEngin()
	r, e := orm.ID(name).Update(&obj)
	return r, e
}

func (service *ConfigService) Add(obj entity.Config) (int64, error) {

	orm := restgo.OrmEngin()
	r, e := orm.InsertOne(&obj)
	return r, e
}
