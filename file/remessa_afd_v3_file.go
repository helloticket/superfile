package file

import (
	"fmt"
	"log"
	"os"

	"github.com/helloticket/superfile/model"
)

type remessaAFDV3File struct {
	model    *model.Remessa
	fileName string
	encoder  *Encoder
}

func (w *remessaAFDV3File) Write() *os.File {
	file, err := os.Create(w.fileName)
	if err != nil {
		panic(err)
	}

	header := w.encodeFileHeader()
	writeString(file, header)
	writeString(file, "\r\n")

	for _, lote := range w.encodeLotes() {
		writeString(file, lote)
		writeString(file, "\r\n")
	}

	trailer := w.encodeFileTrailer()
	writeString(file, trailer)
	writeString(file, "\r\n")

	if err := file.Close(); err != nil {
		log.Println(err)
	}

	return file
}

func (w *remessaAFDV3File) encodeLotes() []string {
	encoded := []string{}

	for _, lote := range w.model.Lotes {
		for _, registro := range lote.Segmentos() {
			block := fmt.Sprintf("detalhes.%s", registro.Nome)
			layout := w.getLayoutFor("detalhes")
			layout = layout[registro.Nome].(map[interface{}]interface{})
			bodyEncoded := w.encoder.ParseAndEncode(block, registro.Valores, layout)
			encoded = append(encoded, bodyEncoded)
		}
	}

	return encoded
}

func (w *remessaAFDV3File) encodeFileHeader() string {
	config := w.model.GetRemessaLayout()
	layout := config["header_arquivo"].(map[interface{}]interface{})
	return w.encoder.ParseAndEncode("header_arquivo", w.model.Header, layout)
}

func (w *remessaAFDV3File) encodeFileTrailer() string {
	config := w.model.GetRemessaLayout()
	layout := config["trailer_arquivo"].(map[interface{}]interface{})
	return w.encoder.ParseAndEncode("trailer_arquivo", w.model.Trailer, layout)
}

func (w *remessaAFDV3File) getLayoutFor(name string) map[interface{}]interface{} {
	config := w.model.GetRemessaLayout()
	return config[name].(map[interface{}]interface{})
}
