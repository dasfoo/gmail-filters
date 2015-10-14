package writer

import (
	"log"
	"os"
)

func WriteToFile(filename string, text string) {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		log.Println(err)
		return
	}

	file.WriteString(text)

}
