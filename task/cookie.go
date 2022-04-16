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

func saveCookie(ctx context.Context) error {
	log.Println("saveCookie...")
	if err := chromedp.WaitReady(`body`, chromedp.ByQuery).Do(ctx); err != nil {
		return err
	}

	cookies, err := network.GetAllCookies().Do(ctx)
	if err != nil {
		return err
	}

	cookiesData, err := network.GetAllCookiesReturns{Cookies: cookies}.MarshalJSON()
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(config.CookieFile, cookiesData, 0755); err != nil {
		return err
	}
	return nil
}

// checkLoginStatus check login status
func checkLoginStatus(ctx context.Context) (err error) {
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

// loadCookie load cookie
func loadCookie(ctx context.Context) (err error) {
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
