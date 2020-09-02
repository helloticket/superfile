package file

import (
	"log"
	"os"
)

func writeString(f *os.File, buff string) {
	if f == nil {
		return
	}

	_, err := f.WriteString(buff)

	if err != nil {
		log.Println(err)
	}
}
