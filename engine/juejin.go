package engine

import (
	"io/ioutil"
	"log"
	"spider/config"
	"spider/questioner"
	"spider/task"
)

func Run() {
	if _, err := ioutil.ReadFile(config.CookieFile); err != nil {
		task.WeChatLogin()
	}

	answers, err := questioner.Ask()
	if err != nil {
		panic(err)
	}
	log.Printf("BookId: %s", answers.BookId)

	if err := downloadBooklet(answers.BookId); err != nil {
		panic(err)
	}
}
