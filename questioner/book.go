package questioner

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
)

type Answer struct {
	BookId string `json:"bookId"`
}

// the questions to ask
var qs = []*survey.Question{
	{
		Name:     "bookId",
		Prompt:   &survey.Input{Message: "Book ID"},
		Validate: survey.Required,
	},
}

func Ask() (*Answer, error) {
	answers := Answer{}
	// perform the questions
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &answers, nil
}
