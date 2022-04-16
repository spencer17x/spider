package engine

import (
	"encoding/json"
	"fmt"
	"log"
	"spider/config"
	"spider/fetcher"
	"spider/utils"
)

type Section struct {
	Title     string `json:"title"`
	SectionId string `json:"section_id"`
	BookTitle string `json:"book_title"`
	Index     int    `json:"index"`
}

// downloadSection download section
func downloadSection(section *Section) error {
	bytes, err := fetcher.FetchSectionContent(section.SectionId)
	if err != nil {
		return err
	}

	log.Printf("Section: %s", section.Title)
	log.Printf("SectionId: %s", section.SectionId)
	dir := fmt.Sprintf(`%s/%s`, config.SavePath, section.BookTitle)
	filename := fmt.Sprintf(`%d.%s.md`, section.Index, section.Title)

	var sectionContentResponse struct {
		Data struct {
			Section struct {
				MarkdownShow string `json:"markdown_show"`
			}
		}
	}
	if err := json.Unmarshal(bytes, &sectionContentResponse); err != nil {
		return err
	}
	if err := utils.SaveFile(
		dir,
		filename,
		sectionContentResponse.Data.Section.MarkdownShow,
	); err != nil {
		return err
	}
	log.Println("done.")
	return nil
}
