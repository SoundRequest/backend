package helper

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

// SendVefiryMail send mail with verify link
func SendVefiryMail(verifyCode, mail string) error {

	_mailHTML, errLoadMailHTML := ioutil.ReadFile("templates/mail/sendverify.html")
	if errLoadMailHTML != nil {
		log.Println(errLoadMailHTML)
	}
	mailHTML := string(_mailHTML)
	mailHTML = strings.Replace(mailHTML, "${verifyCode}", verifyCode, -1)
	mailHTML = strings.Replace(mailHTML, "${verifyCodeLink}", "https://api.soundrequest.xyz/auth/verify/"+verifyCode, -1)

	domain := Config().MailDomain
	mg := mailgun.NewMailgun(domain, Config().MailAPIKey)
	sender := "no-reply@" + domain
	subject := "SoundRequest 인증메일"
	body := ""

	message := mg.NewMessage(sender, subject, body, mail)
	message.SetHtml(mailHTML)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	a, b, errSendMail := mg.Send(ctx, message)
	fmt.Println(a + "     " + b)
	return errSendMail
}

// SendResetPasswordMail returns error(Success of fail)
func SendResetPasswordMail(verifyCode, mail string) error {
	_mailHTML, errLoadMailHTML := ioutil.ReadFile("templates/mail/verifypassword.html")
	if errLoadMailHTML != nil {
		log.Println(errLoadMailHTML)
	}
	mailHTML := string(_mailHTML)
	mailHTML = strings.Replace(mailHTML, "${verifyCode}", verifyCode, -1)

	domain := Config().MailDomain
	mg := mailgun.NewMailgun(domain, Config().MailAPIKey)
	sender := "no-reply@" + domain
	subject := "SoundRequest 인증메일"
	body := ""

	message := mg.NewMessage(sender, subject, body, mail)
	message.SetHtml(mailHTML)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	a, b, errSendMail := mg.Send(ctx, message)
	fmt.Println(a + "     " + b)
	return errSendMail
}
