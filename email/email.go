/**
 * @Time    :2022/12/8 11:59
 * @Author  :Xiaoyu.Zhang
 */

package email

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/melf-xyzh/go-email/commons"
	mailmodel "github.com/melf-xyzh/go-email/model"
	"gopkg.in/gomail.v2"
	"log"
	"mime"
	"path"
	"time"
)

// MailAttach 邮件附件
type MailAttach struct {
	Url      string // 附件地址
	FileName string // 附件别名
}

// SendMassage
/**
 *  @Description: 发送邮件
 *  @receiver client
 *  @param toEMailAddresses 收件地址
 *  @param subject 主题
 *  @param ccEMailAddresses 抄送
 *  @param bccEMailAddresses 暗送
 *  @return err 错误
 */
func (client *Client) SendMassage(toEMailAddresses []string, subject string, body MailBody, ccEMailAddresses []string, bccEMailAddresses []string, attaches []MailAttach, table string) (err error) {
	sendRcd := &mailmodel.EMailSendRcd{}
	sendRcd.SendStatus = mailmodel.WaitSend
	defer func() {
		errDB := client.saveRcd(sendRcd, table)
		if errDB != nil {
			log.Println(err.Error())
		}
	}()

	m := gomail.NewMessage(gomail.SetEncoding(gomail.Unencoded))
	// 发件人
	if client.NickName == "" {
		m.SetHeader("From", client.EMailAddr)
		sendRcd.From = client.EMailAddr
	} else {
		// 使用别名
		m.SetHeader("From", m.FormatAddress(client.EMailAddr, client.NickName))
		sendRcd.From = fmt.Sprintf("%s<%s>", client.NickName, client.EMailAddr)
		sendRcd.NickName = client.NickName
	}
	// 邮件主题
	m.SetHeader("Subject", subject)
	sendRcd.Subject = subject
	if len(toEMailAddresses) == 0 {
		err = errors.New("收件人地址为空")
		return
	}
	// 收件人
	m.SetHeader("To", toEMailAddresses...)
	sendRcd.To = commons.ArrayToStr(toEMailAddresses)

	// 抄送
	if len(ccEMailAddresses) > 0 {
		m.SetHeader("Cc", ccEMailAddresses...)
		sendRcd.CC = commons.ArrayToStr(ccEMailAddresses)
	}
	// 暗送
	if len(bccEMailAddresses) > 0 {
		m.SetHeader("Bcc", bccEMailAddresses...)
		sendRcd.BCC = commons.ArrayToStr(bccEMailAddresses)
	}
	// 邮件内容
	m.SetBody(body.contentType, body.body)
	sendRcd.ContentType = body.contentType
	sendRcd.Body = body.body
	// 附件
	if len(attaches) > 0 {
		for _, attach := range attaches {
			if attach.FileName != "" {
				// 对邮件进行重命名（设置格式为UTF-8，防止乱码）
				m.Attach(attach.Url, gomail.Rename(mime.QEncoding.Encode("UTF-8", path.Base(attach.FileName))))
			} else {
				m.Attach(attach.Url)
			}
		}
	}
	// 创建SMTP客户端，连接到远程的邮件服务器
	d := gomail.NewDialer(
		client.Host,
		client.Port,
		client.EMailAddr,
		client.Password,
	)
	// 关闭SSL协议认证
	if client.InsecureSkipVerify {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}
	// 发送邮件
	err = d.DialAndSend(m)
	if err != nil {
		sendRcd.SendStatus = mailmodel.SendFail
		sendRcd.Remark = err.Error()
	} else {
		sendRcd.SendStatus = mailmodel.SendSuccess
	}
	return
}

// saveRcd
/**
 *  @Description: 保存Email发送记录
 *  @receiver client
 *  @param sendRcd
 *  @param tableName
 *  @return err
 */
func (client *Client) saveRcd(sendRcd *mailmodel.EMailSendRcd, tableName string) (err error) {
	if client.db == nil {
		err = errors.New("未设置数据库链接")
		return
	}
	if tableName == "" {
		tableName = sendRcd.TableName()
	}
	if client.tableNameMap == nil {
		client.tableNameMap = make(map[string]struct{})
	}
	_, ok := client.tableNameMap[tableName]
	if !ok {
		// 没有表,创建表
		err = client.db.Table(tableName).AutoMigrate(&mailmodel.EMailSendRcd{})
		if err != nil {
			err = errors.New("创建数据库表失败：" + err.Error())
			return
		}
	}
	if sendRcd.ID == "" {
		sendRcd.ID = commons.UUID()
		sendRcd.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		err = client.db.Table(tableName).Create(&sendRcd).Error
		if err != nil {
			err = errors.New("创建数据失败：" + err.Error())
			return
		}
	} else {
		sendRcd.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
		err = client.db.Table(tableName).Updates(&sendRcd).Error
		if err != nil {
			err = errors.New("更新数据失败：" + err.Error())
			return
		}
	}
	return
}
