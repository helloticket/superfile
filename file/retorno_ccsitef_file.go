package file

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/helloticket/superfile/model"
)

type retornoCCSITEFFile struct {
	content io.Reader
	layout  model.Layout
	decoder *Decoder
}

func (r *retornoCCSITEFFile) Read() *model.Retorno {
	const registroHeaderArquivo = "A0"
	const registroHeaderLote = "L0"
	const registroTrailerLote = "L9"
	const registroTrailerArquivo = "A9"

	numeroLote := int64(0)
	numeroSegmentos := int64(0)
	retorno := model.NewRetorno(r.layout)
	loteCorrente := retorno.NovoLote(numeroLote)
	detalheCorrente := loteCorrente.NovoDetalhe()

	reader := bufio.NewScanner(r.content)

	for reader.Scan() {
		linha := reader.Text()
		tipoRegistro := linha[0:2]

		if registroHeaderArquivo == tipoRegistro {
			retorno.Header = r.decodeFileHeader(linha)
		}

		if registroHeaderLote == tipoRegistro {
			numeroLote++
			loteCorrente = retorno.NovoLote(numeroLote)
			detalheCorrente = loteCorrente.NovoDetalhe()
			loteCorrente.Header = r.decodeLoteHeader(linha)
		}

		if tipoRegistro == "CV" ||
			tipoRegistro == "AJ" ||
			tipoRegistro == "CC" {
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

func (r *retornoCCSITEFFile) decodeSegmento(row string) model.Segmento {
	segmento := fmt.Sprintf("segmento_%s", strings.ToLower(row[0:2]))
	layout := r.getLayoutFor("detalhes")
	layout = layout[segmento].(map[interface{}]interface{})
	block := fmt.Sprintf("detalhes.%s", segmento)
	linhas := r.decoder.Parse(block, row, layout)

	valores := model.RecordMap{}

	for _, l := range linhas {
		valores[l.Name] = r.decoder.Decode(l.Block, l)
	}

	return model.Segmento{
		Nome:    segmento,
		Valores: valores,
	}
}

func (r *retornoCCSITEFFile) decodeLoteTrailer(row string) map[string]interface{} {
	trailer := map[string]interface{}{}

	linhas := r.decoder.Parse("trailer_lote", row, r.getLayoutFor("trailer_lote"))

	for _, linha := range linhas {
		trailer[linha.Name] = r.decoder.Decode(linha.Block, linha)
	}

	return trailer
}

func (r *retornoCCSITEFFile) decodeLoteHeader(row string) map[string]interface{} {
	header := map[string]interface{}{}

	linhas := r.decoder.Parse("header_lote", row, r.getLayoutFor("header_lote"))

	for _, linha := range linhas {
		header[linha.Name] = r.decoder.Decode(linha.Block, linha)
	}

	return header
}

func (r *retornoCCSITEFFile) decodeFileHeader(row string) map[string]interface{} {
	header := map[string]interface{}{}

	linhas := r.decoder.Parse("header_arquivo", row, r.getLayoutFor("header_arquivo"))

	for _, linha := range linhas {
		header[linha.Name] = r.decoder.Decode(linha.Block, linha)
	}

	return header
}

func (r *retornoCCSITEFFile) decodeFileTrailer(row string) map[string]interface{} {
	trailer := map[string]interface{}{}

	linhas := r.decoder.Parse("trailer_arquivo", row, r.getLayoutFor("trailer_arquivo"))

	for _, linha := range linhas {
		trailer[linha.Name] = r.decoder.Decode(linha.Block, linha)
	}

	return trailer
}

func (r *retornoCCSITEFFile) getLayoutFor(name string) map[interface{}]interface{} {
	config := r.layout.GetRetornoLayout()
	return config[name].(map[interface{}]interface{})
}
