package jobcan

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sclevine/agouti"
	"net/http"
	"os"
	"time"
)

const (
	signInPageUrl    = "https://id.jobcan.jp/users/sign_in?app_key=atd"
	seleniumEndpoint = "http://localhost:4444/wd/hub"
)

const (
	CheckIn  = "check_in"
	CheckOut = "check_out"
)

func Touch(email, password string, opts ...string) error {
	var checkType = ""
	if len(opts) > 0 {
		checkType = opts[0]
	}

	options := []agouti.Option{agouti.Browser("chrome")}
	page, err := agouti.NewPage(seleniumEndpoint, options...)
	if err != nil {
		return err
	}

	_ = page.Navigate(signInPageUrl)
	_ = page.Find("#user_email").Fill(email)
	_ = page.Find("#user_password").Fill(password)
	_ = page.Find("#new_user > input.form__login").Click()

	typ, err := page.Find("#working_status").Text()
	if err != nil {
		return err
	}
	if len(checkType) > 0 {
		if checkType == CheckIn && typ == "勤務中" {
			_ = slack("出勤:打刻済です")
			return nil
		}
		if checkType == CheckOut && typ == "退出中" {
			_ = slack("退勤:打刻済です")
			return nil
		}
	}

	err = page.Find("#adit-button-push").Click()
	if err != nil {
		return err
	}

	loc, _ := time.LoadLocation("Asia/Tokyo")
	_ = slack(checkType + ": 打刻しました: " + time.Now().In(loc).String())

	return nil
}

func slack(msg string) error {
	p := struct {
		Text string `json:"text"`
	}{
		Text: msg,
	}
	b, err := json.Marshal(p)
	if err != nil {
		return err
	}
	url := os.Getenv("SLACK_URL")
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return err
	}
	cli := &http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("slack error!")
	}
	return nil
}
