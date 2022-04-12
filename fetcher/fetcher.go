package fetcher

import (
	"encoding/json"
	"fmt"
	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/chromedp/cdproto/network"
	"io/ioutil"
	"net/http"
	"spider/config"
	"strings"
)

type CookieJson struct {
	Cookies []*network.Cookie `json:"cookies"`
}

// FetchSectionList get section list by bookId
func FetchSectionList(bookId string) ([]byte, error) {
	url := config.RequestUrl + config.BookletApi
	method := "POST"

	payload := strings.NewReader(
		fmt.Sprintf(`{"booklet_id":"%s"}`, bookId),
	)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("user-agent", browser.Chrome())

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func FetchSectionContent(sectionId string) ([]byte, error) {
	url := config.RequestUrl + config.SectionApi
	method := "POST"

	payload := strings.NewReader(
		fmt.Sprintf(`{"section_id":"%s"}`, sectionId),
	)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")

	var cookieJson *CookieJson
	var sessionId string
	bytes, _ := ioutil.ReadFile(config.CookieFile)
	if err := json.Unmarshal(bytes, &cookieJson); err != nil {
		return nil, err
	}
	for _, cookie := range cookieJson.Cookies {
		if cookie.Name == "sessionid" {
			sessionId = cookie.Value
		}
	}
	if sessionId == "" {
		return nil, fmt.Errorf("sessionid is empty")
	}
	req.Header.Add("cookie", fmt.Sprintf("sessionid=%s", sessionId))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
