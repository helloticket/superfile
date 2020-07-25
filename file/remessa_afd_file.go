package file

import (
	"fmt"
	"os"

	"github.com/helloticket/superfile/model"
)

type remessaAFDFile struct {
	model    *model.Remessa
	fileName string
	encoder  *Encoder
}

func (w *remessaAFDFile) Write() *os.File {
	file, err := os.Create(w.fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	header := w.encodeFileHeader()
	file.WriteString(header)
	file.WriteString("\r\n")

	for _, lote := range w.encodeLotes() {
		file.WriteString(lote)
		file.WriteString("\r\n")
	}

	trailer := w.encodeFileTrailer()
	file.WriteString(trailer)
	file.WriteString("\r\n")

	return file
}

func (w *remessaAFDFile) encodeLotes() []string {
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

func (w *remessaAFDFile) encodeFileHeader() string {
	config := w.model.GetRemessaLayout()
	layout := config["header_arquivo"].(map[interface{}]interface{})
	return w.encoder.ParseAndEncode("header_arquivo", w.model.Header, layout)
}

func (w *remessaAFDFile) encodeFileTrailer() string {
	config := w.model.GetRemessaLayout()
	layout := config["trailer_arquivo"].(map[interface{}]interface{})
	return w.encoder.ParseAndEncode("trailer_arquivo", w.model.Trailer, layout)
}

func (w *remessaAFDFile) getLayoutFor(name string) map[interface{}]interface{} {
	config := w.model.GetRemessaLayout()
	return config[name].(map[interface{}]interface{})
}
