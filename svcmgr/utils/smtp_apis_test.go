package utils

import "testing"

func TestSendMail(t *testing.T) {
	SendMailTest()
}

var smtpServer = SmtpServer{
	Host:     "smtp.gmail.com",
	Port:     "587",
	User:     "nubesbh@gmail.com",
	Password: "tycp zngl ehop smvy",
}

var svcmgrAddress = "localhost"

func TestSendMail2(t *testing.T) {
	uuid, _ := NewUUID()
	emailmsg := MailMsg{
		To:       "bhjung@nubes-bridge.com",
		Header:   "콘텐츠브릿지 로그인 Email 인증",
		ServerIp: svcmgrAddress,
		Uuid:     uuid,
		UserId:   "nubesbr",
		Text:     "계정에 대한 이메일 인증을 위해서 아래 URL을 클릭하시기 바랍니다.",
	}

	SendMail(smtpServer, emailmsg)
}
