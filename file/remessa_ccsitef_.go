package file

import (
	"fmt"
	"os"

	"github.com/helderfarias/superfile/model"
)

type remessaCCSITEFFile struct {
	model    *model.Remessa
	fileName string
	encoder  *Encoder
}

func (w *remessaCCSITEFFile) Write() *os.File {
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

func (w *remessaCCSITEFFile) encodeLotes() []string {
	encoded := []string{}

	for _, lote := range w.model.Lotes {
		headerEncoded := w.encoder.ParseAndEncode("header_lote", lote.Header, w.getLayoutFor("header_lote"))
		encoded = append(encoded, headerEncoded)

		for _, registro := range lote.Segmentos() {
			block := fmt.Sprintf("detalhes.%s", registro.Nome)
			layout := w.getLayoutFor("detalhes")
			layout = layout[registro.Nome].(map[interface{}]interface{})
			bodyEncoded := w.encoder.ParseAndEncode(block, registro.Valores, layout)
			encoded = append(encoded, bodyEncoded)
		}

		trailerEncoded := w.encoder.ParseAndEncode("trailer_lote", lote.Trailer, w.getLayoutFor("trailer_lote"))
		encoded = append(encoded, trailerEncoded)
	}

	return encoded
}

func (w *remessaCCSITEFFile) encodeFileHeader() string {
	config := w.model.GetRemessaLayout()
	layout := config["header_arquivo"].(map[interface{}]interface{})
	return w.encoder.ParseAndEncode("header_arquivo", w.model.Header, layout)
}

func (w *remessaCCSITEFFile) encodeFileTrailer() string {
	config := w.model.GetRemessaLayout()
	layout := config["trailer_arquivo"].(map[interface{}]interface{})
	return w.encoder.ParseAndEncode("trailer_arquivo", w.model.Trailer, layout)
}

func (w *remessaCCSITEFFile) getLayoutFor(name string) map[interface{}]interface{} {
	config := w.model.GetRemessaLayout()
	return config[name].(map[interface{}]interface{})
}
