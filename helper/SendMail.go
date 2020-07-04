package helper

import (
	"context"
	"fmt"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

// SendVefiryMail send mail with verify link
func SendVefiryMail(link, mail string) error {
	domain := Config().MailDomain
	mg := mailgun.NewMailgun(domain, Config().MailAPIKey)
	sender := "no-reply@" + domain
	subject := "SoundRequest 인증메일"
	body := "이메일주소를 인증하여 서비스를 이용하실 수 있습니다.\n본인이 아니라면 메일을 무시하셔도 좋습니다.\n인증링크: " + link

	message := mg.NewMessage(sender, subject, body, mail)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	a, b, err := mg.Send(ctx, message)
	fmt.Println(a + "     " + b)
	return err
}
