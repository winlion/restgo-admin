package service

import (
	"restgo-admin/entity"
	"restgo-admin/restgo"
	"restgo-admin/model"
	"github.com/gin-gonic/gin"
	"errors"
	"time"
)

type UserService struct {}
//根据userId 获取用户编号
func (service *UserService)FindOne(userId int64)(entity.User){
	var user entity.User
	orm := restgo.OrmEngin()
	orm.Id(userId).Get(&user)
	return  user
}

func (service *UserService)Query(arg model.UserArg)([]entity.User){
	var users []entity.User = make([]entity.User , 0)
	orm := restgo.OrmEngin()
	t := orm.Where("id>0")
	if (0<len(arg.Kword)){
		t = t.Where("name like ?","%"+arg.Kword+"%")
	}

	if (!arg.Datefrom.IsZero()){
		t = t.Where("create_at >= ?",arg.Datefrom)
	}
	if (!arg.Dateto.IsZero()){
		t = t.Where("create_at <= ?",arg.Dateto)
	}
	t.Limit(arg.GetPageFrom()).Find(&users)
	return  users
}

func (service *UserService)UpdateStat(id int64,stat int)(int64,error){
	var user entity.User
	user.ID=id
	user.Stat=stat
	orm := restgo.OrmEngin()
	 r,e:=orm.ID(id).Cols("stat").Update(&user)
	 return r,e
}

func (service *UserService)Count(arg model.UserArg)(n int64){
	var user entity.User
	orm := restgo.OrmEngin()
	t := orm.Where("id>0")
	if (0<len(arg.Kword)){
		t = t.Where("name like ?","%"+arg.Kword+"%")
	}

	if (!arg.Datefrom.IsZero()){
		t = t.Where("create_at >= ?",arg.Datefrom)
	}
	if (!arg.Dateto.IsZero()){
		t = t.Where("create_at <= ?",arg.Dateto)
	}
	n,_=t.Count(&user)
	return
}

//登录服务,通过手机号/邮箱/用户名登录
func (service *UserService)Login(ctx *gin.Context,kword string,passwd string)(u entity.User,err error){
	ismobile := restgo.IsMobile(kword)
	isemail := restgo.IsEmail(kword)
	var user entity.User
	orm := restgo.OrmEngin()
	t := orm.Where("id>0")
	if ismobile{
		t.Where("mobile=?",kword)
	}else if(isemail){
		t.Where("email=?",kword)
	}else{
		t.Where("account=?",kword)
	}
	t.Get(&user)
	if user.ID==0{
		err = errors.New("该用户不存在")
		return
	}
	if (restgo.Md5encode(passwd)!=user.Passwd){
		err = errors.New("密码不正确,请重试")
		return
	}
	u = user
	restgo.SaveRoleId(ctx,u.RoleId)
	return
}

//注册服务,注册后自动登录
func (service *UserService)Register(ctx *gin.Context,user *entity.User)(p *entity.User,err error){

	isemail := restgo.IsEmail(user.Email)
	if !isemail{
		err = errors.New("email格式不正确")
		return
	}
	if len(user.Passwd)<6{
		err = errors.New("注册失败,太短了")
		return
	}
    var u entity.User

	orm := restgo.OrmEngin()
	t := orm.Where("id>0")

	t.Where("email=?",user.Email)

	t.Get(&u)
	if u.ID>0{
		err = errors.New("该账户已存在")
		return
	}
	user.Passwd = restgo.Md5encode(user.Passwd)
	user.Stat=1
	user.CreateAt = restgo.JsonTime(time.Now())
	user.ID,err = orm.InsertOne(user)
	restgo.SaveUser(ctx,user)
	p = user
	return
}