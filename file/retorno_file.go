package file

import (
	"bufio"
	"io"
	"log"

	"github.com/helderfarias/cnab-go/model"
)

type RetornoFile struct {
	content io.Reader
	layout  model.Layout
	decoder *Decoder
}

func NewRetornoFile(layout model.Layout, content io.Reader) (*RetornoFile, error) {
	return &RetornoFile{
		layout:  layout,
		content: content,
		decoder: NewDecoder(),
	}, nil
}

func (r *RetornoFile) Read() model.Remessa {
	remessa := model.Remessa{}
	reader := bufio.NewScanner(r.content)

	if reader.Scan() {
		header := r.decodeFileHeader(reader.Text())
		log.Println(header)
	}

	for reader.Scan() {
		// log.Println(reader.Text())
	}

	return remessa
}

func (r *RetornoFile) decodeFileHeader(row string) map[string]interface{} {
	header := map[string]interface{}{}

	config := r.layout.GetRetornoLayout()
	layout := config["header_arquivo"].(map[interface{}]interface{})
	linhas := r.decoder.Parse("header_arquivo", row, layout)

	for _, linha := range linhas {
		header[linha.Name] = r.decoder.Decode(linha.Block, linha)
	}

	return header
}
