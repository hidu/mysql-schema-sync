package internal

import (
	"log"
	"strconv"
	"strings"

	"gopkg.in/gomail.v2"
)

// EmailStruct email conf info
type EmailStruct struct {
	SMTPHost     string `json:"smtp_host"`
	From         string `json:"from"`
	Password     string `json:"password"`
	To           string `json:"to"`
	SendMailAble bool   `json:"send_mail"`
}

const tableStyle = `
<sTyle type='text/css'>
      table {border-collapse: collapse;border-spacing: 0;}
     .tb_1{border:1px solid #cccccc;table-layout:fixed;word-break:break-all;width: 100%;background:#ffffff;margin-bottom:5px}
     .tb_1 caption{text-align: center;background: #F0F4F6;font-weight: bold;padding-top: 5px;height: 25px;border:1px solid #cccccc;border-bottom:none}
     .tb_1 a{margin:0 3px 0 3px}
     .tb_1 tr th,.tb_1 tr td{padding: 3px;border:1px solid #cccccc;line-height:20px}
     .tb_1 thead tr th{font-weight:bold;text-align: center;background:#e3eaee}
     .tb_1 tbody tr th{text-align: right;background:#f0f4f6;padding-right:5px}
     .tb_1 tfoot{color:#cccccc}
     .td_c td{text-align: center}
     .td_r td{text-align: right}
     .t_c{text-align: center !important;}
     .t_r{text-align: right !important;}
     .t_l{text-align: left !important;}
</stYle>
`

// SendMail send mail
func (m *EmailStruct) SendMail(title string, body string) {
	if !m.SendMailAble {
		log.Println("send email : no")
		return
	}
	if len(m.SMTPHost) == 0 || len(m.From) == 0 || len(m.To) == 0 {
		log.Println("smtp_host, from,to is empty")
		return
	}
	addrInfo := strings.Split(m.SMTPHost, ":")
	if len(addrInfo) != 2 {
		log.Println("smtp_host wrong, eg: host_name:25")
		return
	}
	var sendTo []string
	for _, _to := range strings.Split(m.To, ";") {
		_to = strings.TrimSpace(_to)
		if len(_to) != 0 && strings.Contains(_to, "@") {
			sendTo = append(sendTo, _to)
		}
	}

	if len(sendTo) < 1 {
		log.Println("mail receiver is empty")
		return
	}

	body = mailBody(body)

	a := gomail.NewMessage()
	a.SetHeader("From", m.From)
	a.SetHeader("To", sendTo...)  // 发送给多个用户
	a.SetHeader("Subject", title) // 设置邮件主题
	a.SetBody("text/html", body)  // 设置邮件正文
	port, _ := strconv.Atoi(addrInfo[1])

	d := gomail.NewDialer(addrInfo[0], port, m.From, m.Password)

	err := d.DialAndSend(a)
	if err == nil {
		log.Println("send mail success")
	} else {
		log.Println("send mail failed, err:", err)
	}
}

func mailBody(body string) string {
	body = tableStyle + "\n" + body
	body += "<br/><hr style='border:none;border-top:1px solid #ccc'/>" +
		"<center>Powered by <a href='" + AppURL + "'>mysql-schema-sync</a>&nbsp;" + Version + "</center>"
	return body
}
