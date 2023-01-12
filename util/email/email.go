package email

import (
	"crypto/tls"
	"github.com/go-gomail/gomail"
	"github.com/sirupsen/logrus"
)

type EmailInfo struct {
	ServerHost string // 邮箱服务器地址
	ServerPort int    // 邮箱服务器端口
	FromName   string // 发件人别名
	FromEmail  string // 发件人邮箱地址
	FromPasswd string // 发件人邮箱密码

	Recipient []string //收件人邮箱
	CC        []string //抄送

}

func NewAduEmailSender(rec []string) *EmailInfo {
	return &EmailInfo{}
}

/**
 * @Author: czh
 * @Date: 2020-06-06 15:45:55
 * @Description: 发送邮件
 * @Param : subject[主题]、body[内容]、enclosure[附件地址]
 * @Return:
 */

func (emailInfo *EmailInfo) SendEmail(subject, body, enclosure string) {
	if len(emailInfo.Recipient) == 0 {
		logrus.Errorf("邮件发送失败,操作人为空!")
		return
	}

	mes := gomail.NewMessage()
	//设置收件人
	mes.SetHeader("To", emailInfo.Recipient...)
	//设置抄送列表
	if len(emailInfo.CC) != 0 {
		mes.SetHeader("Cc", emailInfo.CC...)
	}
	// 发件人别名
	mes.SetAddressHeader("From", emailInfo.FromEmail, emailInfo.FromName)

	// 主题
	mes.SetHeader("Subject", subject)

	// 正文
	mes.SetBody("text/html", body)

	// 附件
	if enclosure != "" {
		//name := "附件.csv"
		mes.Attach(enclosure)
	}

	d := gomail.NewDialer(emailInfo.ServerHost, emailInfo.ServerPort, emailInfo.FromEmail, emailInfo.FromPasswd)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(mes)
	if err != nil {
		logrus.Errorf("发送邮件失败:%s", err)
	}
}
