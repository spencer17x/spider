package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"spider/config"
	"spider/fetcher"
	"spider/questioner"
	"spider/task"
	"spider/utils"
)

type Section struct {
	Title     string `json:"title"`
	SectionId string `json:"section_id"`
}

func main() {

	if _, err := ioutil.ReadFile(config.CookieFile); err != nil {
		task.Login()
	}

	answers, answersErr := questioner.Ask()
	if answersErr != nil {
		panic(answersErr)
	}
	log.Printf("BookId: %s", answers.BookId)

	sectionListBytes, sectionListBytesErr := fetcher.FetchSectionList(answers.BookId)
	if sectionListBytesErr != nil {
		panic(sectionListBytesErr)
	}

	var sectionListResponse struct {
		Data struct {
			Booklet struct {
				BaseInfo struct {
					Title string `json:"title"`
				} `json:"base_info"`
			} `json:"booklet"`
			Sections []*Section
		}
	}
	if err := json.Unmarshal(sectionListBytes, &sectionListResponse); err != nil {
		panic(err)
	}

	bookTitle := sectionListResponse.Data.Booklet.BaseInfo.Title

	for sectionIndex, section := range sectionListResponse.Data.Sections {
		bytes, sectionResponseErr := fetcher.FetchSectionContent(section.SectionId)
		if sectionResponseErr != nil {
			panic(sectionResponseErr)
		}
		log.Printf("SectionIndex: %d", sectionIndex)
		log.Printf("Section: %s", section.Title)
		log.Printf("SectionId: %s", section.SectionId)
		dir := fmt.Sprintf(`%s/%s`, config.SavePath, bookTitle)
		filename := fmt.Sprintf(`%d.%s.md`, sectionIndex+1, section.Title)

		var sectionContentResponse struct {
			Data struct {
				Section struct {
					MarkdownShow string `json:"markdown_show"`
				}
			}
		}
		if err := json.Unmarshal(bytes, &sectionContentResponse); err != nil {
			panic(err)
		}
		if saveErr := utils.SaveFile(
			dir,
			filename,
			sectionContentResponse.Data.Section.MarkdownShow,
		); saveErr != nil {
			panic(saveErr)
		}
		log.Println("done.")
	}
}
