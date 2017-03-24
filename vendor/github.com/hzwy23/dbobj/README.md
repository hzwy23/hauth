## 介绍
* dbobj 提供一个抽象层，通过sql的方式操作数据库

* golang数据库接口

> 这个包,封装了golang与数据库之间的接口,目前支持oracle,mysql数据库

* 获取方式：

> go get github.com/hzwy23/dbobj

## 使用方法:

### oracle数据库 

1. 如果选择的是oracle数据库,请按照go-oci8包的要求配置pkgconfig和oracle instantclient.
2. oci8.pc在vendor/github.com/mattn/go-oci8中.请按照要求,修改oci8.pc文件,然后修改oracle.go文件，**将第14行的注释去掉**。
3. 请设置环境变量.HBIGDATA_HOME.这个变量中创建目录conf.然后将dbobj中的system.properties复制到conf中.

### mysql，mariadb数据库
1. 请设置环境变量.HBIGDATA_HOME.这个变量中创建目录conf.然后将dbobj中的system.properties复制到conf中.

### 创建目录

```shell
    export HBIGDATA_HOME=/opt/go/hcloud
    mkdir $HBIGDATA_HOME/conf
    cp system.properties $HBIGDATA_HOME/conf   #将system.properties文件复制到conf目录中.
```

### 工程目录样式:
```
$HBIGDATA_HOME
            ----bin

            ----src

            --------github.com

            ------------hzwy23

            ----------------dbobj

            ----conf

            --------system.properties
```

在指定的配置文件目录中创建配置文件,配置文件名称指定为:system.properties,在文件中输入下面信息:

* mysql配置文件

```
    DB.type=mysql
    DB.tns = "tcp(localhost:3306)/bigdata"
    DB.user = root
    DB.passwd= huang
```
* oracle配置文件

```
    DB.type=oracle
    DB.tns = "192.168.1.101:1521/orcl"
    DB.user = test
    DB.passwd= huang
	
```

* 系统启动后,会默认自动对密码进行加密.

### 例子
```go
package main

import (
    "fmt"

    "github.com/hzwy23/dbobj"
)

func main() {

    rows, err := dbobj.Query("SELECT user_id,user_name FROM sys_user_info where user_id = ?", "admin")
    defer rows.Close()
    if err != nil {
        fmt.Println("query failed.")
        return
    }
    for rows.Next() {
        var userId string
        var userName string
        err = rows.Scan(&userId, &userName)
        if err != nil {
            fmt.Println("query failed. scan failed.")
            return
        }
        fmt.Println("user id is :", userId, "user name is :", userName)
    }
}
```