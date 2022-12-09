/**
 * @Time    :2022/12/8 9:35
 * @Author  :Xiaoyu.Zhang
 */

package email

import (
	"github.com/melf-xyzh/go-mail/constants"
	"gorm.io/gorm"
)

// 参考文档
// https://blog.csdn.net/qq_40227117/article/details/117988459
// https://blog.csdn.net/weixin_31702225/article/details/116772164

type Client struct {
	Host               string              // 邮件服务器
	Port               int                 // 服务端口
	NickName           string              // 昵称
	EMailAddr          string              // 发件人邮箱
	Password           string              // 发件人授权码
	InsecureSkipVerify bool                // 是否关闭SSL校验
	db                 *gorm.DB            // 数据库链接
	tableNameMap       map[string]struct{} // 数据库表Map（防止重复建表）
}

// CreateQQEMail
/**
 *  @Description: 创建一个QQ邮箱客户端
 *  @param eMailAddr
 *  @param Password
 *  @param nickName
 *  @return client
 */
func CreateQQEMail(eMailAddr, Password, nickName string) (client Client) {
	client = Client{
		Host:               mailconstant.STMPServerQQ,
		Port:               mailconstant.STMPPortQQ,
		InsecureSkipVerify: false,
		NickName:           nickName,
		EMailAddr:          eMailAddr,
		Password:           Password,
	}
	return
}

// Create126EMail
/**
 *  @Description: 创建一个126邮箱客户端
 *  @param eMailAddr
 *  @param Password
 *  @param nickName
 *  @return client
 */
func Create126EMail(eMailAddr, Password, nickName string) (client Client) {
	client = Client{
		Host:               mailconstant.STMPServer126,
		Port:               mailconstant.STMPPort126,
		InsecureSkipVerify: false,
		NickName:           nickName,
		EMailAddr:          eMailAddr,
		Password:           Password,
	}
	return
}

// Create163EMail
/**
 *  @Description: 创建一个163邮箱客户端
 *  @param eMailAddr
 *  @param Password
 *  @param nickName
 *  @return client
 */
func Create163EMail(eMailAddr, Password, nickName string) (client Client) {
	client = Client{
		Host:               mailconstant.STMPServer163,
		Port:               mailconstant.STMPPort163,
		InsecureSkipVerify: false,
		NickName:           nickName,
		EMailAddr:          eMailAddr,
		Password:           Password,
	}
	return
}

// CreateWorkWXEMail
/**
 *  @Description: 创建一个企业邮箱客户端
 *  @param eMailAddr
 *  @param Password
 *  @param nickName
 *  @return client
 */
func CreateWorkWXEMail(eMailAddr, Password, nickName string) (client Client) {
	client = Client{
		Host:               mailconstant.STMPServerWorkWX,
		Port:               mailconstant.STMPPortWorkWX,
		InsecureSkipVerify: false,
		NickName:           nickName,
		EMailAddr:          eMailAddr,
		Password:           Password,
	}
	return
}

// CreateCustomEMail
/**
 *  @Description: 创建一个自定义邮箱客户端
 *  @param host STMP服务器地址
 *  @param port 端口
 *  @param eMailAddr 发件人地址
 *  @param Password 密码/授权码
 *  @param nickName 发件人昵称
 *  @param insecureSkipVerify 是否关闭SSL校验
 *  @return client
 */
func CreateCustomEMail(host string, port int, eMailAddr, Password, nickName string, insecureSkipVerify bool) (client Client) {
	client = Client{
		Host:               host,
		Port:               port,
		InsecureSkipVerify: insecureSkipVerify,
		NickName:           nickName,
		EMailAddr:          eMailAddr,
		Password:           Password,
	}
	return
}

// SetDB
/**
 *  @Description: 设置数据库链接
 *  @receiver client
 *  @param db
 */
func (client *Client) SetDB(db *gorm.DB) {
	client.db = db
}
