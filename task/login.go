package task

import (
	"context"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"log"
	"spider/config"
	"time"
)

type Task func(context.Context) error

// createTask task for chromedp
func createTask(task Task) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		return task(ctx)
	}
}

// newTargetToShowQrCode new tab to WeChat login
func newTargetToShowQrCode(ctx context.Context) error {
	log.Println("showQrCode to logging...")
	ch := chromedp.WaitNewTarget(ctx, func(info *target.Info) bool {
		return info.URL != ""
	})
	newCtx, _ := chromedp.NewContext(ctx, chromedp.WithTargetID(<-ch))

	var urlStr string
	if err := chromedp.Run(newCtx, chromedp.Tasks{
		chromedp.Location(&urlStr),
		createTask(getWeChatLoginQrCode),
	}); err != nil {
		log.Fatal(err)
	}
	return nil
}

// WeChatLogin login entry
func WeChatLogin() {
	log.Println("WeChat logging...")
	ctx, _ := chromedp.NewExecAllocator(
		context.Background(),
		append(
			chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", config.ChromeDpHeadless),
		)...,
	)

	ctx, _ = context.WithTimeout(ctx, 1*time.Minute)

	ctx, _ = chromedp.NewContext(
		ctx,
		chromedp.WithLogf(log.Printf),
	)

	if err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(config.URL),
		chromedp.Click(`.login-button`, chromedp.ByQuery),
		chromedp.Click(`.prompt-box .clickable`, chromedp.ByQuery),
		chromedp.Click(`.oauth-bg .oauth-btn[title="微信"]`, chromedp.ByQuery),
		createTask(newTargetToShowQrCode),
		chromedp.WaitVisible(
			`#juejin > div.view-container.container > div > header > div > nav > ul > ul > li.nav-item.menu > div > img`,
			chromedp.ByQuery,
		),
		createTask(saveCookie),
	}); err != nil {
		log.Fatal(err)
	}
}
