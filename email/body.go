/**
 * @Time    :2022/12/8 11:57
 * @Author  :Xiaoyu.Zhang
 */

package email

// MailBody 邮件内容
type MailBody struct {
	contentType string
	body        string
}

func NewTextMailBody(body string) MailBody {
	return MailBody{
		contentType: "text/plain",
		body:        body,
	}
}

func NewHTMLMailBody(body string) MailBody {
	return MailBody{
		contentType: "text/html",
		body:        body,
	}
}

