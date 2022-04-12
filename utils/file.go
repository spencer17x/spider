package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"spider/config"
)

// SaveFile save file with content
func SaveFile(dir, filename, content string) error {
	if _, err := ioutil.ReadDir(dir); err != nil {
		if osMkdirErr := os.Mkdir(dir, os.ModePerm); osMkdirErr != nil {
			return osMkdirErr
		}
	}
	f, err := os.Create(
		fmt.Sprintf(`%s/%s`, config.SavePath, filename),
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
