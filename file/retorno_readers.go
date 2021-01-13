package file

import (
	"bufio"
	"io"
	"log"
	"strings"
)

type ReaderRetorno struct {
	reader *bufio.Scanner
}

func NewReader(r io.Reader) *ReaderRetorno {
	reader := bufio.NewScanner(r)
	return &ReaderRetorno{reader: reader}
}

func (r *ReaderRetorno) Mapper(closure func(pos int64, raw string)) {
	var pos int64 = 0

	for r.reader.Scan() {
		pos++
		linha := r.reader.Text()

		if strings.TrimSpace(linha) == "" {
			log.Printf("Ignorando linha(%v) vazia \n", pos)
			continue
		}

		closure(pos, linha)
	}
}
