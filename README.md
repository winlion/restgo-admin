# 1、首先上效果
![登录页面](https://github.com/winlion/restgo-admin/blob/master/asset/images/1.png)
![登入首页](https://github.com/winlion/restgo-admin/blob/master/asset/images/2.png)
![自由配置角色和权限](https://github.com/winlion/restgo-admin/blob/master/asset/images/3.png)
![支持系统自定义参数](https://github.com/winlion/restgo-admin/blob/master/asset/images/4.png)
# 2、如何使用
## 2.1、使用如下指令克隆
```
cd $GOPATH/src
git clone https://github.com/winlion/restgo-admin.git  
```
你将得到restgo-admin 目录  
进入目录  
```
cd restgo-admin
go run  main.go
```

## 2.2、数据库
新建数据库名称为restgo-admin,编码为utf-8  
将restgo-admin.sql导入到数据库中   
修改conf/app.properties文件24行数据库配置
restgo.datasource.default.dataSourceName=root:root@/restgo-admin?charset=utf8  
数据库配置方式有如下几种,详细请自行百度  

```
user@unix(/path/to/socket)/dbname?charset=utf8
user:password@tcp(localhost:5555)/dbname?charset=utf8
user:password@/dbname
user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
```
## 2.3、初始化依赖包
使用前先使用如下指令安装指令安装文件
```
go get github.com/go-sql-driver/mysql  
go get -v -u github.com/alecthomas/log4go  
go get github.com/gin-gonic/gin  
go get github.com/go-sql-driver/mysql  
go get github.com/go-xorm/xorm  
go get github.com/tommy351/gin-sessions  
```

## 2.4、启动
使用前先使用如下指令启动应用  
```
go run main.go  
```
浏览器使用http://localhost/即可访问   
初始用户admin,密码rootme@1   
使用账号18600000000或admin@qq.com,密码rootme@1也可以   

## 2.5、打包
window下使用如下指令打包应用
```
build-windows.bat  
```
linux/unix下使用如下指令
```
#chmod +x ./build-unix
#./build-unix
```
打包前请认真阅读资源配置章节
[资源控制器https://blog.csdn.net/keytounix/article/details/79336554](https://blog.csdn.net/keytounix/article/details/79336554 "资源控制器")
## 2.6、参数配置
相关参数配置请移步
[我的csdn博客https://blog.csdn.net/keytounix](https://blog.csdn.net/keytounix "csdn")
```
https://blog.csdn.net/keytounix
```
# 3、FAQ
## 3.1 如何安装开发环境
```
go get github.com/nsf/gocode  
go get github.com/uudashr/gopkgs/cmd/gopkgs  
go get github.com/fatih/gomodifytags  
go get github.com/haya14busa/goplay/cmd/goplay  
go get github.com/derekparker/delve/cmd/dlv  
```
如果你使用的是vscode,安装问题请访问
https://www.cnblogs.com/Leo_wl/p/8242628.html

## 3.2 如何联系我
我的微信 jiepool-winlion  
我的qq 271151388
