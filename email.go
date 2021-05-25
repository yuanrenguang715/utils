package utils

import (
	"regexp"
	"strconv"

	"gopkg.in/gomail.v2"
)

func CheckEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

/* mailTo := []string{
	"1870997052@qq.com",
}
//邮件主题为"Hello"
subject := "Wolf.Color.com"
// 邮件正文
body := "<h4>您的验证码是:</h4><h2>456123</h2>"
fmt.Println(SendMail(mailTo, subject, body)) */
func SendMail(mailTo []string, subject string, body string) error {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"user": "Wolf.Color@outlook.com",
		"pass": "yuan-5201",
		"host": "smtp.office365.com",
		"port": "587",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()
	m.SetHeader("From", "XD Game"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)                            //发送给多个用户
	m.SetHeader("Subject", subject)                         //设置邮件主题
	m.SetBody("text/html", body)                            //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}

func Count(year int, month int) (days int) {
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30

		} else {
			days = 31
		}
	} else {
		if (((year % 4) == 0 && (year % 100) != 0) || (year % 400) == 0) {
			days = 29
		} else {
			days = 28
		}
	}
	return
}