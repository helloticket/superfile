package file

import (
	"fmt"
	"os"
	"strings"

	"github.com/helloticket/superfile/field"
	"github.com/helloticket/superfile/model"
)

type RemessaDebug struct {
	model   *model.Remessa
	encoder *Encoder
}

func NewRemessaDebug(remessa *model.Remessa) *RemessaDebug {
	return &RemessaDebug{
		model:   remessa,
		encoder: NewEncoder(remessa),
	}
}

func (w *RemessaDebug) Write() *os.File {
	w.encodeFileHeader()
	w.encodeLotes()
	w.encodeFileTrailer()
	return nil
}

func (w *RemessaDebug) encodeFileHeader() {
	config := w.model.GetRemessaLayout()
	layout := config["header_arquivo"].(map[interface{}]interface{})
	linhas := w.encoder.Parse("header_arquivo", w.model.Header, layout)
	w.output("remessa.Header", linhas)
}

func (w *RemessaDebug) encodeLotes() {
	fmt.Println()
	lote := w.model.NovoLoteSegmentoVazio(1)

	if w.model.GetLayout() == model.LayoutCNAB240 {
		headerEncoded := w.encoder.Parse("header_lote", lote.Header, w.getLayoutFor("header_lote"))
		w.output("lote.Header", headerEncoded)
	}

	fmt.Println()
	lote.InserirDetalhe(lote.NovoDetalhe())

	for _, registro := range lote.Segmentos() {
		block := fmt.Sprintf("detalhes.%s", registro.Nome)
		layout := w.getLayoutFor("detalhes")
		layout = layout[registro.Nome].(map[interface{}]interface{})
		bodyEncoded := w.encoder.Parse(block, registro.Valores, layout)
		w.output(fmt.Sprintf("detalhe[\"%s\"]", registro.Nome), bodyEncoded)
	}

	fmt.Println()

	if w.model.GetLayout() == model.LayoutCNAB240 {
		trailerEncoded := w.encoder.Parse("trailer_lote", lote.Trailer, w.getLayoutFor("trailer_lote"))
		w.output("lote.Trailer", trailerEncoded)
	}
}

func (w *RemessaDebug) encodeFileTrailer() {
	fmt.Println()
	config := w.model.GetRemessaLayout()
	layout := config["trailer_arquivo"].(map[interface{}]interface{})
	linhas := w.encoder.Parse("trailer_arquivo", w.model.Header, layout)
	w.output("remessa.Trailer", linhas)
}

func (w *RemessaDebug) getLayoutFor(name string) map[interface{}]interface{} {
	config := w.model.GetRemessaLayout()
	return config[name].(map[interface{}]interface{})
}

func (w *RemessaDebug) output(blockName string, linhas []model.Linha) {
	picture := field.NewPicture()

	for _, l := range linhas {
		if l.DefaultValue != nil {
			continue
		}

		var newValue interface{}
		if strings.HasPrefix(l.Picture, "X") {
			newValue = "0"
		} else {
			newValue = 0
		}

		clone := l.Clone(newValue)
		out := strings.TrimSpace(clone.Encode(picture))
		if strings.HasPrefix(l.Picture, "X") {
			fmt.Printf("%s[\"%s\"] = \"%v\"\r\n", blockName, clone.Name, out)
		} else {
			fmt.Printf("%s[\"%s\"] = %v\r\n", blockName, clone.Name, out)
		}
	}
}
