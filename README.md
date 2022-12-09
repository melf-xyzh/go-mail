# go-mail

Go email的便捷封装

## 安装

```bash
go get -u github.com/melf-xyzh/go-mail
```

## 发送邮件

### DEMO

```go
package main

import (
	"github.com/melf-xyzh/go-email/email"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {
	//参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456789@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 创建一个发件人配置
	client := email.CreateCustomEMail("smtp.126.com",465,"xxx...@126.com","CJ*******YCIW","昵称",true)
	client.SetDB(db)

	// 创建邮件内容
	//body := email.NewTextMailBody("验证码：123456" + "时间：" + time.Now().Format("2006-01-02 15:04:05"))
	body := email.NewHTMLMailBody("<h1>验证码：123456</h1>" + "<h2>时间：" + time.Now().Format("2006-01-02 15:04:05") + "</h2>")

	// 发送邮件
	err = client.SendMassage(
		[]string{"xyzh.****@petalmail.com"},
		"测试",
		body,
		[]string{"zhuan*****1999@126.com"},
		[]string{"9*****90@qq.com"},
		nil,
		"",
	)
	if err != nil {
		panic(err)
	} else {
		log.Println("邮件发送成功")
	}
}
```

#### 126邮箱

```go
client := email.Create126EMail(eMailAddr, Password, nickName)
```

#### 163邮箱

```go
client := email.Create163EMail(eMailAddr, Password, nickName)
```

#### QQ邮箱

```go
client := email.CreateQQEMail(eMailAddr, Password, nickName)
```

#### 企业微信邮箱

```go
client := email.CreateWorkWXEMail(eMailAddr, Password, nickName)
```

