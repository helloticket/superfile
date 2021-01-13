package file

import (
	"fmt"
	"io"

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

	NewReader(r.content).Mapper(func(pos int64, linha string) {
		tipoRegistro := linha[0:2]

		if registroHeaderArquivo == tipoRegistro {
			retorno.Header = r.decodeFileHeader(pos, retorno, linha)
		}

		if registroHeaderLote == tipoRegistro {
			numeroLote++
			loteCorrente = retorno.NovoLote(numeroLote)
			detalheCorrente = loteCorrente.NovoDetalhe()
			loteCorrente.Header = r.decodeLoteHeader(pos, retorno, linha)
		}

		if tipoRegistro == "CV" ||
			tipoRegistro == "AJ" ||
			tipoRegistro == "CC" {
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
	})

	return retorno
}

func (r *retornoCCSITEFFile) decodeSegmento(pos int64, retorno *model.Retorno, row string) model.Segmento {
	decode := DecodeRetorno{decoder: r.decoder, layout: r.layout}

	return decode.Segmento(pos, 0, 2, row, retorno)
}

func (r *retornoCCSITEFFile) decodeLoteTrailer(pos int64, retorno *model.Retorno, row string) map[string]interface{} {
	decode := DecodeRetorno{decoder: r.decoder, layout: r.layout}

	return decode.TrailerLote(pos, row, retorno)
}

func (r *retornoCCSITEFFile) decodeLoteHeader(pos int64, retorno *model.Retorno, row string) map[string]interface{} {
	decode := DecodeRetorno{decoder: r.decoder, layout: r.layout}

	return decode.HeaderLote(pos, row, retorno)
}

func (r *retornoCCSITEFFile) decodeFileHeader(pos int64, retorno *model.Retorno, row string) map[string]interface{} {
	decode := DecodeRetorno{decoder: r.decoder, layout: r.layout}

	return decode.HeaderArquivo(pos, row, retorno)
}

func (r *retornoCCSITEFFile) decodeFileTrailer(pos int64, retorno *model.Retorno, row string) map[string]interface{} {
	decode := DecodeRetorno{decoder: r.decoder, layout: r.layout}

	return decode.TrailerArquivo(pos, row, retorno)
}
