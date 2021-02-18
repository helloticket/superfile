package file

import (
	"fmt"
	"io"
	"strings"

	"github.com/helloticket/superfile/helper"
	"github.com/helloticket/superfile/model"
)

type retornoAEJFile struct {
	content io.Reader
	layout  model.Layout
	decoder *Decoder
}

func (r *retornoAEJFile) Read() *model.Retorno {
	const registroHeaderArquivo = 1
	const registroTrailerArquivo = 999999999

	numeroSegmentos := int64(0)
	retorno := model.NewRetorno(r.layout)
	loteCorrente := retorno.NovoLote(1)
	detalheCorrente := loteCorrente.NovoDetalhe()
	loteCorrente.InserirDetalhe(detalheCorrente)
	retorno.InserirLote(loteCorrente)

	NewReader(r.content).Mapper(func(pos int64, source string) {
		linha := source
		if caracter := r.layout.GlobalSettings()["adicao_caracter_a_direita"]; caracter != "" {
			linha = strings.ReplaceAll(source, "|", "")
		}

		tipoRegistro := helper.ToSafeInt(linha[9:10])
		codigoRegistro := helper.ToSafeInt(linha[0:9])

		if registroHeaderArquivo == tipoRegistro {
			retorno.Header = r.decodeFileHeader(pos, retorno, linha)
		} else if codigoRegistro == registroTrailerArquivo && tipoRegistro == 9 {
			retorno.Trailer = r.decodeFileTrailer(pos, retorno, linha)
		} else {
			numeroSegmentos++
			segmento := r.decodeSegmento(pos, retorno, linha)
			detalheCorrente[fmt.Sprintf("%d.%s", numeroSegmentos, segmento.Nome)] = segmento.Valores
		}
	})

	return retorno
}

func (r *retornoAEJFile) decodeSegmento(pos int64, retorno *model.Retorno, row string) model.Segmento {
	decode := DecodeRetorno{decoder: r.decoder, layout: r.layout}

	return decode.Segmento(pos, 9, 10, row, retorno)
}

func (r *retornoAEJFile) decodeFileHeader(pos int64, retorno *model.Retorno, row string) map[string]interface{} {
	decode := DecodeRetorno{decoder: r.decoder, layout: r.layout}

	return decode.HeaderArquivo(pos, row, retorno)
}

func (r *retornoAEJFile) decodeFileTrailer(pos int64, retorno *model.Retorno, row string) map[string]interface{} {
	decode := DecodeRetorno{decoder: r.decoder, layout: r.layout}

	return decode.TrailerArquivo(pos, row, retorno)
}
