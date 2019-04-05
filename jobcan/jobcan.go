package jobcan

import (
	"github.com/sclevine/agouti"
)

const (
	signInPageUrl    = "https://id.jobcan.jp/users/sign_in?app_key=atd"
	seleniumEndpoint = "http://localhost:4444/wd/hub"
)

func Touch(email, password string) error {
	options := []agouti.Option{agouti.Browser("chrome")}
	page, err := agouti.NewPage(seleniumEndpoint, options...)
	if err != nil {
		return err
	}

	_ = page.Navigate(signInPageUrl)
	_ = page.Find("#user_email").Fill(email)
	_ = page.Find("#user_password").Fill(password)
	_ = page.Find("#new_user > input.form__login").Click()

	err = page.Find("#adit-button-push").Click()
	if err != nil {
		return err
	}

	return nil
}
