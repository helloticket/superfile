package file

import (
	"bufio"
	"fmt"
	"io"

	"github.com/helloticket/superfile/helper"
	"github.com/helloticket/superfile/model"
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
	var pos int64 = 0

	for reader.Scan() {
		pos++
		linha := reader.Text()
		tipoRegistro := helper.ToSafeInt(linha[7:8])

		if registroHeaderArquivo == tipoRegistro {
			retorno.Header = r.decodeFileHeader(pos, retorno, linha)
		}

		if registroHeaderLote == tipoRegistro {
			numeroLote++
			loteCorrente = retorno.NovoLote(numeroLote)
			detalheCorrente = loteCorrente.NovoDetalhe()
			loteCorrente.Header = r.decodeLoteHeader(pos, retorno, linha)
		}

		if registroDetalhes == tipoRegistro {
			numeroSegmentos++
			segmento := r.decodeSegmento(pos, retorno, linha)
			detalheCorrente[fmt.Sprintf("%d.%s", numeroSegmentos, segmento.Nome)] = segmento.Valores
		}

		if registroTrailerLote == tipoRegistro {
			loteCorrente.Trailer = r.decodeLoteTrailer(pos, retorno, linha)
			loteCorrente.InserirDetalhe(detalheCorrente)
			retorno.InserirLote(loteCorrente)
			loteCorrente = nil
			detalheCorrente = nil
		}

		if registroTrailerArquivo == tipoRegistro {
			retorno.Trailer = r.decodeFileTrailer(pos, retorno, linha)
		}
	}

	return retorno
}

func (r *retornoCNAB240File) decodeSegmento(pos int64, retorno *model.Retorno, row string) model.Segmento {
	decode := DecodeRetorno{decoder: r.decoder, layout: r.layout}

	return decode.Segmento(pos, 13, 14, row, retorno)
}

func (r *retornoCNAB240File) decodeLoteTrailer(pos int64, retorno *model.Retorno, row string) map[string]interface{} {
	decode := DecodeRetorno{decoder: r.decoder, layout: r.layout}

	return decode.TrailerLote(pos, row, retorno)
}

func (r *retornoCNAB240File) decodeLoteHeader(pos int64, retorno *model.Retorno, row string) map[string]interface{} {
	decode := DecodeRetorno{decoder: r.decoder, layout: r.layout}

	return decode.HeaderLote(pos, row, retorno)
}

func (r *retornoCNAB240File) decodeFileHeader(pos int64, retorno *model.Retorno, row string) map[string]interface{} {
	decode := DecodeRetorno{decoder: r.decoder, layout: r.layout}

	return decode.HeaderArquivo(pos, row, retorno)
}

func (r *retornoCNAB240File) decodeFileTrailer(pos int64, retorno *model.Retorno, row string) map[string]interface{} {
	decode := DecodeRetorno{decoder: r.decoder, layout: r.layout}

	return decode.TrailerArquivo(pos, row, retorno)
}
