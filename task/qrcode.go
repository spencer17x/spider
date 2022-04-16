package task

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

// printQrCode print qr code to terminal
func printQrCode(code []byte) (err error) {
	img, _, err := image.Decode(bytes.NewReader(code))
	if err != nil {
		return err
	}

	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return err
	}

	res, err := qrcode.NewQRCodeReader().Decode(bmp, nil)
	if err != nil {
		return err
	}
	qr, err := goQrcode.New(res.String(), goQrcode.High)
	if err != nil {
		return err
	}

	fmt.Println(qr.ToSmallString(false))
	return nil
}

// getWeChatLoginQrCode get wechat login qr code
func getWeChatLoginQrCode(ctx context.Context) (err error) {
	log.Println("getWeChatLoginQrCode...")
	var code []byte

	if err = chromedp.Screenshot(`.web_qrcode_img`, &code, chromedp.ByQuery).Do(ctx); err != nil {
		return err
	}

	if err = printQrCode(code); err != nil {
		return err
	}
	return nil
}
