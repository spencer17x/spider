package engine

import (
	"encoding/json"
	"fmt"
	"log"
	"spider/config"
	"spider/fetcher"
)

type Section struct {
	Title     string `json:"title"`
	SectionId string `json:"section_id"`
	BookTitle string `json:"book_title"`
	Index     int    `json:"index"`
}

// downloadSection a downloader for a section
func downloadSection(section *Section) error {
	bytes, err := fetcher.FetchSectionContent(section.SectionId)
	if err != nil {
		return err
	}

	log.Printf("Section: %s", section.Title)
	log.Printf("SectionId: %s", section.SectionId)
	dir := fmt.Sprintf(`%s/%s`, config.SavePath, section.BookTitle)
	filename := fmt.Sprintf(`%d.%s.md`, section.Index, section.Title)

	var response struct {
		Data struct {
			Section struct {
				MarkdownShow string `json:"markdown_show"`
			}
		}
	}
	if err := json.Unmarshal(bytes, &response); err != nil {
		return err
	}
	if err := saveFile(
		dir,
		filename,
		response.Data.Section.MarkdownShow,
	); err != nil {
		return err
	}
	log.Println("done.")
	return nil
}

// downloadBooklet a downloader for a booklet
func downloadBooklet(id string) error {
	bytes, err := fetcher.FetchSectionList(id)
	if err != nil {
		return err
	}

	var response struct {
		Data struct {
			Booklet struct {
				BaseInfo struct {
					Title string `json:"title"`
				} `json:"base_info"`
			} `json:"booklet"`
			Sections []*Section `json:"sections"`
		} `json:"data"`
	}
	if err := json.Unmarshal(bytes, &response); err != nil {
		return err
	}

	bookTitle := response.Data.Booklet.BaseInfo.Title

	log.Printf("bookTitle: %s", bookTitle)
	for index, section := range response.Data.Sections {
		section.BookTitle = bookTitle
		section.Index = index + 1
		log.Printf("SectionIndex: %d", section.Index)
		if err := downloadSection(section); err != nil {
			log.Printf("downloadSection error: %s", err)
		}
	}

	return nil
}
