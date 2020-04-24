package file

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/helloticket/superfile/helper"
	"github.com/helloticket/superfile/model"
)

type retornoCNAB400File struct {
	content io.Reader
	layout  model.Layout
	decoder *Decoder
}

func (r *retornoCNAB400File) Read() *model.Retorno {
	const registroHeaderArquivo = 0
	const registroTrailerArquivo = 9

	numeroSegmentos := int64(0)
	retorno := model.NewRetorno(r.layout)
	loteCorrente := retorno.NovoLote(1)
	detalheCorrente := loteCorrente.NovoDetalhe()
	loteCorrente.InserirDetalhe(detalheCorrente)
	retorno.InserirLote(loteCorrente)

	reader := bufio.NewScanner(r.content)

	for reader.Scan() {
		linha := reader.Text()
		tipoRegistro := helper.ToSafeInt(linha[0:1])

		if registroHeaderArquivo == tipoRegistro {
			retorno.Header = r.decodeFileHeader(linha)
		} else if registroTrailerArquivo == tipoRegistro {
			retorno.Trailer = r.decodeFileTrailer(linha)
		} else {
			numeroSegmentos++
			segmento := r.decodeSegmento(linha)
			detalheCorrente[fmt.Sprintf("%d.%s", numeroSegmentos, segmento.Nome)] = segmento.Valores
		}
	}

	return retorno
}

func (r *retornoCNAB400File) decodeSegmento(row string) model.Segmento {
	segmento := fmt.Sprintf("segmento_%s", strings.ToLower(row[0:1]))
	layout := r.getLayoutFor("detalhes")
	layout = layout[segmento].(map[interface{}]interface{})
	block := fmt.Sprintf("detalhes.%s", segmento)
	linhas := r.decoder.Parse(block, row, layout)

	valores := model.RecordMap{}

	for _, l := range linhas {
		valores[l.Name] = l.Value
	}

	return model.Segmento{
		Nome:    segmento,
		Valores: valores,
	}
}

func (r *retornoCNAB400File) decodeFileHeader(row string) map[string]interface{} {
	header := map[string]interface{}{}

	linhas := r.decoder.Parse("header_arquivo", row, r.getLayoutFor("header_arquivo"))

	for _, linha := range linhas {
		header[linha.Name] = r.decoder.Decode(linha.Block, linha)
	}

	return header
}

func (r *retornoCNAB400File) decodeFileTrailer(row string) map[string]interface{} {
	trailer := map[string]interface{}{}

	linhas := r.decoder.Parse("trailer_arquivo", row, r.getLayoutFor("trailer_arquivo"))

	for _, linha := range linhas {
		trailer[linha.Name] = r.decoder.Decode(linha.Block, linha)
	}

	return trailer
}

func (r *retornoCNAB400File) getLayoutFor(name string) map[interface{}]interface{} {
	config := r.layout.GetRetornoLayout()
	return config[name].(map[interface{}]interface{})
}
