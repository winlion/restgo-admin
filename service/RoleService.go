package service

import (
	"errors"
	"restgo-admin/entity"
	"restgo-admin/restgo"
)

type RoleService struct{}

//查询全部角色
func (service *RoleService) All() []entity.Role {

	var roles []entity.Role = make([]entity.Role, 0)
	orm := restgo.OrmEngin()
	orm.Where("id>0").Find(&roles)
	return roles
}

//添加一个新的角色
func (service *RoleService) Add(role entity.Role) (id int64, err error) {
	//确保角色名称不为空
	if len(role.Name) == 0 {
		err = errors.New("请输入角色名称")
		return
	}
	//校验角色是否存在
	all := service.All()
	for _, r := range all {
		if r.Name == role.Name {
			err = errors.New("该角色已经存在")
			return
		}
	}
	//设置成可用
	role.Stat = 1

	orm := restgo.OrmEngin()
	id, err = orm.InsertOne(&role)
	return
}

//添加一个新的角色
func (service *RoleService) LoadAuth(roleid int64) (ret []entity.Resource) {
	orm := restgo.OrmEngin()
	var res []entity.Resource = make([]entity.Resource, 0)
	orm.Where("id in (select resid from ref_role_res where roleid= ?)", roleid).Find(&res)
	ret = res
	return
}

//添加一个新的角色
func (service *RoleService) LoadAllAuth() (ret []entity.Resource) {
	orm := restgo.OrmEngin()
	var res []entity.Resource = make([]entity.Resource, 0)
	orm.Where("id >0").Find(&res)
	ret = res
	return
}

//添加一个新的角色
func (service *RoleService) Delete(role entity.Role) (n int64, err error) {
	if role.ID == 0 {
		err = errors.New("请选择角色")
		return
	}
	orm0 := restgo.OrmEngin()
	role.Stat = 0
	n, err = orm0.Id(role.ID).Update(role)
	return
}

//恢复数据
func (service *RoleService) Reback(role entity.Role) (n int64, err error) {
	if role.ID == 0 {
		err = errors.New("请选择角色")
		return
	}
	orm0 := restgo.OrmEngin()
	role.Stat = 1
	n, err = orm0.Id(role.ID).Update(role)
	return
}

//取消授权
func (service *RoleService) RevokeAuth(roleid int, resid int) (n int64, err error) {
	orm := restgo.OrmEngin()
	var refRoleRes entity.RefRoleRes

	n, err = orm.Where("roleid=? and resid=?", roleid, resid).Delete(&refRoleRes)
	return
}

//授权
func (service *RoleService) GrantAuth(roleid int, resid int) (n int64, err error) {
	orm := restgo.OrmEngin()
	var refRoleRes entity.RefRoleRes
	n, err = orm.Where("roleid=? and resid=?", roleid, resid).Delete(&refRoleRes)

	refRoleRes.ResId = resid
	refRoleRes.RoleId = roleid
	orm = restgo.OrmEngin()
	n, err = orm.InsertOne(&refRoleRes)
	return
}
