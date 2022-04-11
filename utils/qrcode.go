package utils

import (
	"bytes"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	goQrcode "github.com/skip2/go-qrcode"
	"image"
	"log"
)

func PrintQRCode(code []byte) (err error) {
	img, _, err := image.Decode(bytes.NewReader(code))
	if err != nil {
		return
	}

	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return
	}

	res, err := qrcode.NewQRCodeReader().Decode(bmp, nil)
	if err != nil {
		return
	}
	qr, err := goQrcode.New(res.String(), goQrcode.High)
	if err != nil {
		return
	}

	fmt.Println(qr.ToSmallString(false))
	return
}

func GetQRCode() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		log.Println("get qrcode...")
		var code []byte

		if err = chromedp.Screenshot(`.web_qrcode_img`, &code, chromedp.ByQuery).Do(ctx); err != nil {
			return
		}

		if err = PrintQRCode(code); err != nil {
			return err
		}
		return
	}
}
