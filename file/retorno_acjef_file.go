package file

import (
	"fmt"
	"io"

	"github.com/helloticket/superfile/helper"
	"github.com/helloticket/superfile/model"
)

type retornoACJEFFile struct {
	content io.Reader
	layout  model.Layout
	decoder *Decoder
}

func (r *retornoACJEFFile) Read() *model.Retorno {
	const registroHeaderArquivo = 1
	const registroTrailerArquivo = 9

	numeroSegmentos := int64(0)
	retorno := model.NewRetorno(r.layout)
	loteCorrente := retorno.NovoLote(1)
	detalheCorrente := loteCorrente.NovoDetalhe()
	loteCorrente.InserirDetalhe(detalheCorrente)
	retorno.InserirLote(loteCorrente)

	NewReader(r.content).Mapper(func(pos int64, linha string) {
		tipoRegistro, err := helper.ToInt(linha[9:10])
		if err != nil {
			retorno.RegistrarFalha(pos, fmt.Errorf("Falha ao decodificar segmento na linha %v, bloco [9:10]. Erro %v", pos, err))
		}

		if registroHeaderArquivo == tipoRegistro {
			retorno.Header = r.decodeFileHeader(pos, retorno, linha)
		} else if tipoRegistro == registroTrailerArquivo {
			retorno.Trailer = r.decodeFileTrailer(pos, retorno, linha)
		} else {
			numeroSegmentos++
			segmento := r.decodeSegmento(pos, retorno, linha)
			detalheCorrente[fmt.Sprintf("%d.%s", numeroSegmentos, segmento.Nome)] = segmento.Valores
		}
	})

	return retorno
}

func (r *retornoACJEFFile) decodeSegmento(pos int64, retorno *model.Retorno, row string) model.Segmento {
	decode := DecodeRetorno{decoder: r.decoder, layout: r.layout}

	return decode.Segmento(pos, 9, 10, row, retorno)
}

func (r *retornoACJEFFile) decodeFileHeader(pos int64, retorno *model.Retorno, row string) map[string]interface{} {
	decode := DecodeRetorno{decoder: r.decoder, layout: r.layout}

	return decode.HeaderArquivo(pos, row, retorno)
}

func (r *retornoACJEFFile) decodeFileTrailer(pos int64, retorno *model.Retorno, row string) map[string]interface{} {
	decode := DecodeRetorno{decoder: r.decoder, layout: r.layout}

	return decode.TrailerArquivo(pos, row, retorno)
}
