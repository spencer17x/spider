package engine

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// saveFile save file with content
func saveFile(dir, filename, content string) error {
	if _, err := ioutil.ReadDir(dir); err != nil {
		if osMkdirErr := os.MkdirAll(dir, os.ModePerm); osMkdirErr != nil {
			return osMkdirErr
		}
	}
	log.Println("dir", dir)
	f, err := os.Create(
		fmt.Sprintf(`%s/%s`, dir, filename),
	)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
