package task

import (
	"context"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"os"
	"spider/config"
	"strings"
)

// SaveCookie save cookie
func SaveCookie() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		log.Println("save cookie...")
		if err = chromedp.WaitReady(`body`, chromedp.ByQuery).Do(ctx); err != nil {
			return
		}

		cookies, err := network.GetAllCookies().Do(ctx)
		if err != nil {
			return err
		}

		cookiesData, err := network.GetAllCookiesReturns{Cookies: cookies}.MarshalJSON()
		if err != nil {
			return err
		}

		if err = ioutil.WriteFile(config.CookieFile, cookiesData, 0755); err != nil {
			return err
		}
		return nil
	}
}

// CheckLoginStatus check login status
func CheckLoginStatus() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		var url string
		if err = chromedp.Evaluate(`window.location.href`, &url).Do(ctx); err != nil {
			return
		}
		if strings.Contains(url, "https://account.wps.cn/usercenter/apps") {
			log.Println("已经使用cookies登陆")
			chromedp.Stop()
		}
		return
	}
}

// LoadCookie load cookie
func LoadCookie() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		if _, _err := os.Stat("cookies.tmp"); os.IsNotExist(_err) {
			return
		}

		cookiesData, err := ioutil.ReadFile("cookies.tmp")
		if err != nil {
			return
		}

		cookiesParams := network.SetCookiesParams{}
		if err = cookiesParams.UnmarshalJSON(cookiesData); err != nil {
			return
		}

		return network.SetCookies(cookiesParams.Cookies).Do(ctx)
	}
}
