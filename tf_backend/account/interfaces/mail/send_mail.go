package mail

import (
	"fmt"
	"os"

	"github.com/TrendFindProject/tf_backend/account/interfaces/logger"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendRecoverUrl(email string, name string, token string) error {
	from := mail.NewEmail("TrendFinder", "admin@trend-find.work")
	subject := "【TrendFinder】パスワード再設定のご案内"
	to := mail.NewEmail(name, email)

	plainTextContent := `いつもTrendFinderをご利用いただき、ありがとうございます。
パスワード再設定用のURLを送付いたしますので、こちらより設定をお願いいたします。

` + `https://trend-find.work/recover_password/edit?rt=` + token + `

本URLは発行時より24時間有効です。
再設定にはメール送信時に発行された認証キーが必要です。

もしこのメールの心当たりがない場合、大変お手数ではございますが、以下の連絡先までお知らせいただければ幸いです。
今後とも何卒よろしくお願いいたします。

------------------------------------------------------------------
TrendFinder https://trend-find.work/

お問い合わせ: admin@trend-find.work
運営: TrendFinderProject
------------------------------------------------------------------
`

	message := mail.NewV3MailInit(from, subject, to, mail.NewContent("text/plain", plainTextContent))
	client := sendgrid.NewSendClient(os.Getenv("SEND_GRID_API_KEY"))

	res, err := client.Send(message)
	if err != nil {
		return err
	}

	logger.Common.Info(fmt.Sprintf("send recover email to %s, result code is %d", email, res.StatusCode))

	return nil
}

func SendRecoverComplete(email string, name string) error {
	from := mail.NewEmail("TrendFinder", "admin@trend-find.work")
	subject := "【TrendFinder】パスワード再設定のご案内"
	to := mail.NewEmail(name, email)

	plainTextContent := `パスワード再設定が完了しましたのでご連絡します。
引き続き、TrendFinderをご利用ください。

もしこのメールの心当たりがない場合、大変お手数ではございますが、以下の連絡先までお知らせいただければ幸いです。
今後とも何卒よろしくお願いいたします。

------------------------------------------------------------------
TrendFinder https://trend-find.work/

お問い合わせ: admin@trend-find.work
運営: TrendFinderProject
------------------------------------------------------------------
`

	message := mail.NewV3MailInit(from, subject, to, mail.NewContent("text/plain", plainTextContent))
	client := sendgrid.NewSendClient(os.Getenv("SEND_GRID_API_KEY"))

	res, err := client.Send(message)
	if err != nil {
		return err
	}

	logger.Common.Info(fmt.Sprintf("send recover complete email to %s, result code is %d", email, res.StatusCode))

	return nil
}
