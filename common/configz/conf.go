package configz

import (
	"bytes"
	"io/ioutil"
	"log"
)

func MustMergeFile(files ...string) string {
	file, err := mergeFile(files...)
	if err != nil {
		log.Fatalf("error: merge file, %s", err.Error())
	}
	return file
}

func mergeFile(files ...string) (string, error) {
	tempFile, err := ioutil.TempFile("", "config-*.yaml")
	if err != nil {
		return "", err
	}
	defer func() {
		tempFile.Close()
	}()

	allContent := bytes.Buffer{}
	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			continue
		}
		allContent.WriteString(string(content))
		allContent.WriteString("\n")
	}
	_, err = tempFile.Write(allContent.Bytes())
	if err != nil {
		return "", err
	}
	return tempFile.Name(), err
}
