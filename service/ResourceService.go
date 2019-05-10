package service

import (
	"errors"
	"restgo-admin/entity"
	"restgo-admin/restgo"
)

//资源服务层
type ResourceService struct{}

//查询全部资源
func (service *ResourceService) All() []entity.Resource {
	var res []entity.Resource = make([]entity.Resource, 0)
	orm := restgo.OrmEngin()
	orm.Where("id>0").Find(&res)
	return res
}

//添加新资源,不绑定权限
func (service *ResourceService) Add(res entity.Resource) (id int64, err error) {
	if len(res.Patern) == 0 {
		err = errors.New("请输入资源格式")
		return
	}
	if len(res.Name) == 0 {
		err = errors.New("请输入资源名称")
		return
	}
	if len(res.ResType) == 0 {
		err = errors.New("请输入资源类型")
		return
	}

	res.Stat = 1
	orm := restgo.OrmEngin()

	ret, _ := orm.Where("patern = ?", res.Patern).Count(&res)
	if ret > 0 {
		err = errors.New("该资源已经存在")
		return
	}

	id, err = orm.InsertOne(&res)
	return

}

//搜索资源
func (service *ResourceService) RoleAuth(roleId int) []entity.Resource {

	orm := restgo.OrmEngin()
	var ress []entity.Resource

	orm.Where("1=1").Find(&ress)
	return ress

}

//添加新资源,同时绑定权限
func (service *ResourceService) AddWithRoleIds(res entity.Resource, roleIds []int) (id int64, err error) {
	if len(res.Patern) == 0 {
		err = errors.New("请输入资源格式")
		return
	}
	if len(res.Name) == 0 {
		err = errors.New("请输入资源名称")
		return
	}
	if len(res.ResType) == 0 {
		err = errors.New("请输入资源类型")
		return
	}

	res.Stat = 1
	orm := restgo.OrmEngin()

	ret, _ := orm.Where("patern = ?", res.Patern).Count(&res)
	if ret > 0 {
		err = errors.New("该资源已经存在")
		return
	}
	//如果没有需要绑定的权限
	if len(roleIds) == 0 {
		id, err = orm.InsertOne(&res)
		return
	}
	//开启事务
	session := orm.NewSession()
	defer session.Close()
	id, err = session.InsertOne(&res)
	for _, r := range roleIds {
		_, err = session.InsertOne(&entity.RefRoleRes{RoleId: int(r), ResId: int(id)})
	}
	//如果有错误,那么rollback
	if err != nil {
		session.Rollback()
	} else {
		//否则提交
		session.Commit()
	}

	return
}

//删除新资源
func (service *ResourceService) Delete(res entity.Resource) (n int64, err error) {
	if res.ID == 0 {
		err = errors.New("请选择资源")
		return
	}
	orm0 := restgo.OrmEngin()
	res.Stat = 0
	n, err = orm0.Id(res.ID).Update(res)
	return
}

//恢复资源
func (service *ResourceService) Reback(res entity.Resource) (n int64, err error) {
	if res.ID == 0 {
		err = errors.New("请选择资源")
		return
	}
	orm0 := restgo.OrmEngin()
	res.Stat = 1
	n, err = orm0.Id(res.ID).Update(&res)
	return
}
