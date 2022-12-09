/**
 * @Time    :2022/12/9 9:21
 * @Author  :Xiaoyu.Zhang
 */

package mailmodel

const (
	WaitSend    = "00" // 待发送
	SendSuccess = "01" // 发送成功
	SendFail    = "02" // 发送失败
)

// EMailSendRcd 邮件发送记录
type EMailSendRcd struct {
	ID          string `json:"id"                         gorm:"column:id;primary_key;type:varchar(36)"`
	CreateTime  string `json:"createTime"                 gorm:"column:create_time;type:varchar(19);index;"`
	UpdateTime  string `json:"updateTime,omitempty"       gorm:"column:update_time;type:varchar(19);"`
	From        string `json:"from"                       gorm:"column:from;comment:发件人;type:text;"`
	NickName    string `json:"nickName"                   gorm:"column:nick_name;comment:昵称;type:varchar(255);"`
	To          string `json:"to"                         gorm:"column:to;comment:收件人;type:text;"`
	CC          string `json:"cc"                         gorm:"column:cc;comment:抄送;type:text;"`
	BCC         string `json:"bcc"                        gorm:"column:bcc;comment:暗抄;type:text;"`
	Subject     string `json:"subject"                    gorm:"column:subject;comment:主题;type:text;"`
	ContentType string `json:"contentType"                gorm:"column:content_type;comment:内容类型;type:varchar(255);"`
	Body        string `json:"body"                       gorm:"column:body;comment:内容;type:text;"`
	SendStatus  string `json:"sendStatus"                gorm:"column:send_status;comment:发送状态;type:varchar(2);default:00"`
	Remark      string `json:"remark"                     gorm:"column:remark;comment:备注;type:text;"`
}

// TableName 自定义表名
func (EMailSendRcd) TableName() string {
	return "email_send_rcd"
}
