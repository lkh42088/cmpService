package utils

import (
	"fmt"
	"net/smtp"
)

type SmtpServer struct {
	Host     string
	Port     string
	User     string
	Password string
}

func (s *SmtpServer) Address() string {
	return s.Host + ":" + s.Port
}

type MailMsg struct {
	To string
	Header string
	UserId string
	Text   string
	ServerIp string
	Uuid   string
}

func (m *MailMsg) GetMessage() []byte {
	header := fmt.Sprintf("Subject: %s \r\n", m.Header)
	body := fmt.Sprintf( "%s %s \r\n", m.UserId, m.Text)
	body += fmt.Sprintf( "http://%s/log_in/emailconfirm/%s \r\n", m.ServerIp, m.Uuid)
	msg := header + "\r\n" + body
	return []byte(msg)
}

func SendMail(server SmtpServer, msg MailMsg) error {
	from := server.User
	to := []string{
		msg.To,
	}
	auth := smtp.PlainAuth("", from, server.Password, server.Host)
	return smtp.SendMail(server.Address(), auth, from, to, msg.GetMessage())
}

func SendMailTest() {
	smtpSvr := SmtpServer{
		"smtp.gmail.com",
		"587" ,
		"nubesbh@gmail.com",
		"tycp zngl ehop smvy",
	}

	mailAccount_from := smtpSvr.User
	mailAccount_to := []string{
		"byeonghwa.jung@gmail.com",
	}

	auth := smtp.PlainAuth("", mailAccount_from, smtpSvr.Password, smtpSvr.Host)

	headerSubject := "Subject: TEST2.....\r\n"
	headerBlank := "\r\n"
	body := "메일 시험입니다...\r\n"
	msg := []byte(headerSubject + headerBlank + body)

	err :=smtp.SendMail(smtpSvr.Address(), auth, mailAccount_from, mailAccount_to, msg)
	if err != nil {
		fmt.Println("err: ", err)
	}
}