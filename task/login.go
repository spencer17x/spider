package task

import (
	"context"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"log"
	"spider/config"
	"spider/utils"
	"time"
)

// CreateTasks task to run
func CreateTasks() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(
			config.URL,
		),
		chromedp.Click(`.login-button`, chromedp.ByQuery),
		chromedp.Click(`.prompt-box .clickable`, chromedp.ByQuery),
		chromedp.Click(`.oauth-bg .oauth-btn[title="微信"]`, chromedp.ByQuery),
		WeChatLoginTask(),
		chromedp.WaitVisible(
			`#juejin > div.view-container.container > div > header > div > nav > ul > ul > li.nav-item.menu > div > img`,
			chromedp.ByQuery,
		),
		SaveCookie(),
	}
}

// WeChatLoginTask new tab to weChat login
func WeChatLoginTask() chromedp.ActionFunc {
	return func(ctx context.Context) error {
		log.Println("weChat login...")
		ch := chromedp.WaitNewTarget(ctx, func(info *target.Info) bool {
			return info.URL != ""
		})
		newCtx, _ := chromedp.NewContext(ctx, chromedp.WithTargetID(<-ch))

		var urlStr string
		if err := chromedp.Run(newCtx, chromedp.Tasks{
			chromedp.Location(&urlStr),
			utils.GetQRCode(),
		}); err != nil {
			log.Fatal(err)
		}
		return nil
	}
}

// Login entry
func Login() {
	log.Println("login...")
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

	if err := chromedp.Run(ctx, CreateTasks()); err != nil {
		log.Fatal(err)
	}
}
