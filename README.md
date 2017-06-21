[Asofdate Project Release](https://github.com/hzwy23/asofdate/releases)

#### golang版本建议:
```
go sdk >= 1.8
```

**Java版本地址**
[https://github.com/hzwy23/asofdate-etl](https://github.com/hzwy23/asofdate-etl)

## 获取项目源代码方法. **请确保设置了GOPATH环境变量**
```golang
go get github.com/hzwy23/asofdate
#上边命令,将会在GOPATH/bin目录中生成asofate可执行文件,由于系统运行需要配置文件和静态页面
#所系请将GOPATH/bin的asofdate可执行文件,复制到$GOPATH/src/github.com/hzwy23/asofdate目录中
#或者使用项目中提供的build.sh脚本编译,再次生成可执行文件.
#请按照下边的操作,导入数据库.
```

## asofdate项目简介
每一个项目,都有截止日期,为了实现快速开发目标,我们以beego为基础，开发出了一款快速开发平台。这个平台内部集成了菜单管理、用户管理、角色管理、授权管理、日志管理、机构管理、路由管理、域定义管理等等。在这个平台的基础上，可以快速的开发自己的应用,以响应瞬息万变的市场需求。

## 项目目标
打造一款安全，稳定，易拓展的快速开发平台.在这个平台的基础上，能够迅速的开发出市场上需要的应用产品，省去系统基础服务开发测试工作量。

## 特点介绍

1. 去session化，采用jwt标准管理用户连接信息，易于分布式环境部署.
2. 菜单页面采用metro风格,简洁明了.
3. 权限控制到按钮级别，有效的对系统API服务进行控制.
4. 快速添加应用程序，只需要在菜单资源管理页面中注册新应用的菜单、路由信息，便可便捷的扩展新应用.
5. 用户操作记录十分精细，有效的记录用户每一个API请求.
6. 后台服务代码,提供国际化服务,轻松实现国际化
7. 系统帮助,提供swagger ui界面,方便管理系统API.

## 系统简介

系统管理是整个产品的核心功能部分，系统中菜单资源是整个系统的公有资源，其余的资源，都是建立在各自的域中。

每个域中特有的信息是：机构、用户、角色，所以，在这个开发平台中，可以轻松的构建出一个适用于不同群体的应用产品，不同的群体信息相互隔离，同一个群体内信息共享。在应用系统中，当新增一个用户群体时，只需要新建一个域，便可实现这个功能。

## 安装方法

**1. 导入数据库信息**

创建数据库用户，导入数据文件，目前支持mysql，mariadb。oracle版本属于商业版，暂时不开源，有需求可以联系。

导入数据文件方法，请修改下边“数据库名”为你的数据库中存在的数据库名
```shell
mysql -uroot -p 数据库名 < ./init_hauth.sql
```
提示：init_hauth.sql在src/github.com/hzwy23/hauth/script目录中

**2. 编译asofdate代码，生成可执行文件**


**A. 直接以安装包的方式编译**

执行下边命令，在执行命令前，请确保您已经安装了go sdk

```shell
## cd 到asofdate的解压目录，然后执行下边命令
./build.sh
## 上边这种模式编译会生成一个可执行文件asofdate，
```
这个命令将会在hauth的解压目录下生成hauth可执行文件。

**B. 采用build编译main.go文件方式**

main.go文件在hauth解压的根目录中，编译方法如下：
```
# cd 到asofdate解压后的根目录
go get github.com/hzwy23/asofdate
go build -i main.go
```

使用liteide的童鞋，采用第二种方式比较好调试，只需要设置GOPATH环境变量后，就可以直接打开main.go，然后点击BuildAndRun按钮，既可以启动服务。

**3 修改配置文件**

配置文件在conf目录中，app.conf是beego的配置文件，主要涉及到服务端口号等等，另外一个是asofdate.conf配置文件，这个里边主要是是=数据库连接信息与日志管理信息配置。

beeog的配置方法，请在beego项目中查阅，请移步：beego.me。下边来讲讲asofdate.conf中数据库的配置方法。

```
DB.type=mysql
DB.tns = "tcp(localhost:3306)/test"
DB.user = root
DB.passwd="xzPEh+SfFL3aimN0zGNB9w=="
```

注意: 修改的文件必须保存为utf-8编码,否则可能会出现异常，DB.type=mysql，这个值请不要修改，因为当前项目中提供的数据库脚本是针对于mysql和mariadb的。

1. 修改DB.tns中对应的数据库地址，端口号，数据库名称。

2. 修改DB.user成相应的数据库用户名

3. 修改DB.passwd成上边用户所对应的密码，系统启动后会自动加密，在此输入密码明文即可。

## 启动方法
```shell
## linux上，请执行。此外需要注意的是：linux上开启1024以下端口号需要管理员权限。
nohup ./asofdate &

## Mac上，
sudo ./asofdate

## windows上，
## 请直接双击asofdate.exe可执行文件
```

打开浏览器,访问:https://localhost:8090

登录系统用户名是：demo，密码是：123456, 

管理员用户: admin, 密码: hzwy23


![系统管理界面](./system_manage.png)

## 交流方式

E-mail： hzwy23@163.com

demo演示地址：https://www.asofdate.com:8090

用户名: demo

密  码: 123456

## 代码贡献列表

github账号 | 
---|
cp542524698 |
xingyuejun | 
hzwy23 |
