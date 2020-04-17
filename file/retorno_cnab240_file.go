package file

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/helderfarias/cnab-go/helper"
	"github.com/helderfarias/cnab-go/model"
)

type retornoCNAB240File struct {
	content io.Reader
	layout  model.Layout
	decoder *Decoder
}

func (r *retornoCNAB240File) Read() *model.Retorno {
	const registroHeaderArquivo = 0
	const registroHeaderLote = 1
	const registroDetalhes = 3
	const registroTrailerLote = 5
	const registroTrailerArquivo = 9

	numeroLote := int64(0)
	numeroSegmentos := int64(0)
	retorno := model.NewRetorno(r.layout)
	loteCorrente := retorno.NovoLote(numeroLote)
	detalheCorrente := loteCorrente.NovoDetalhe()

	reader := bufio.NewScanner(r.content)

	for reader.Scan() {
		linha := reader.Text()
		tipoRegistro := helper.ToSafeInt(linha[7:8])

		if registroHeaderArquivo == tipoRegistro {
			retorno.Header = r.decodeFileHeader(linha)
		}

		if registroHeaderLote == tipoRegistro {
			numeroLote++
			loteCorrente = retorno.NovoLote(numeroLote)
			detalheCorrente = loteCorrente.NovoDetalhe()
			loteCorrente.Header = r.decodeLoteHeader(linha)
		}

		if registroDetalhes == tipoRegistro {
			numeroSegmentos++
			segmento := r.decodeSegmento(linha)
			detalheCorrente[fmt.Sprintf("%d.%s", numeroSegmentos, segmento.Nome)] = segmento.Valores
		}

		if registroTrailerLote == tipoRegistro {
			loteCorrente.Trailer = r.decodeLoteTrailer(linha)
			loteCorrente.InserirDetalhe(detalheCorrente)
			retorno.InserirLote(loteCorrente)
			loteCorrente = nil
			detalheCorrente = nil
		}

		if registroTrailerArquivo == tipoRegistro {
			retorno.Trailer = r.decodeFileTrailer(reader.Text())
		}
	}

	return retorno
}

func (r *retornoCNAB240File) decodeSegmento(row string) model.Segmento {
	segmento := fmt.Sprintf("segmento_%s", strings.ToLower(row[13:14]))
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

func (r *retornoCNAB240File) decodeLoteTrailer(row string) map[string]interface{} {
	trailer := map[string]interface{}{}

	linhas := r.decoder.Parse("trailer_lote", row, r.getLayoutFor("trailer_lote"))

	for _, linha := range linhas {
		trailer[linha.Name] = r.decoder.Decode(linha.Block, linha)
	}

	return trailer
}

func (r *retornoCNAB240File) decodeLoteHeader(row string) map[string]interface{} {
	header := map[string]interface{}{}

	linhas := r.decoder.Parse("header_lote", row, r.getLayoutFor("header_lote"))

	for _, linha := range linhas {
		header[linha.Name] = r.decoder.Decode(linha.Block, linha)
	}

	return header
}

func (r *retornoCNAB240File) decodeFileHeader(row string) map[string]interface{} {
	header := map[string]interface{}{}

	linhas := r.decoder.Parse("header_arquivo", row, r.getLayoutFor("header_arquivo"))

	for _, linha := range linhas {
		header[linha.Name] = r.decoder.Decode(linha.Block, linha)
	}

	return header
}

func (r *retornoCNAB240File) decodeFileTrailer(row string) map[string]interface{} {
	trailer := map[string]interface{}{}

	linhas := r.decoder.Parse("trailer_arquivo", row, r.getLayoutFor("trailer_arquivo"))

	for _, linha := range linhas {
		trailer[linha.Name] = r.decoder.Decode(linha.Block, linha)
	}

	return trailer
}

func (r *retornoCNAB240File) getLayoutFor(name string) map[interface{}]interface{} {
	config := r.layout.GetRetornoLayout()
	return config[name].(map[interface{}]interface{})
}
