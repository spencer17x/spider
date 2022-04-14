package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// SaveFile save file with content
func SaveFile(dir, filename, content string) error {
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
