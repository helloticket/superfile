package file

import (
	"fmt"
	"io"

	"github.com/helloticket/superfile/helper"
	"github.com/helloticket/superfile/model"
)

type retornoAFDV3File struct {
	content io.Reader
	layout  model.Layout
	decoder *Decoder
}

func (r *retornoAFDV3File) Read() *model.Retorno {
	const registroHeaderArquivo = 1
	const registroTrailerArquivo = 999999999

	numeroSegmentos := int64(0)
	retorno := model.NewRetorno(r.layout)
	loteCorrente := retorno.NovoLote(1)
	detalheCorrente := loteCorrente.NovoDetalhe()
	loteCorrente.InserirDetalhe(detalheCorrente)
	retorno.InserirLote(loteCorrente)

	NewReader(r.content).Mapper(func(pos int64, linha string) {
		tipoRegistro := helper.ToSafeInt(linha[9:10])
		codigoRegistro := helper.ToSafeInt(linha[0:9])

		if registroHeaderArquivo == tipoRegistro {
			retorno.Header = r.decodeFileHeader(pos, retorno, linha)
		} else if codigoRegistro == registroTrailerArquivo && tipoRegistro == 0 {
			retorno.Trailer = r.decodeFileTrailer(pos, retorno, linha)
		} else {
			numeroSegmentos++
			segmento := r.decodeSegmento(pos, retorno, linha)
			detalheCorrente[fmt.Sprintf("%d.%s", numeroSegmentos, segmento.Nome)] = segmento.Valores
		}
	})

	return retorno
}

func (r *retornoAFDV3File) decodeSegmento(pos int64, retorno *model.Retorno, row string) model.Segmento {
	decode := DecodeRetorno{decoder: r.decoder, layout: r.layout}

	return decode.Segmento(pos, 9, 10, row, retorno)
}

func (r *retornoAFDV3File) decodeFileHeader(pos int64, retorno *model.Retorno, row string) map[string]interface{} {
	decode := DecodeRetorno{decoder: r.decoder, layout: r.layout}

	return decode.HeaderArquivo(pos, row, retorno)
}

func (r *retornoAFDV3File) decodeFileTrailer(pos int64, retorno *model.Retorno, row string) map[string]interface{} {
	decode := DecodeRetorno{decoder: r.decoder, layout: r.layout}

	return decode.TrailerArquivo(pos, row, retorno)
}
