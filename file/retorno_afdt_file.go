package file

import (
	"bufio"
	"fmt"
	"io"

	"github.com/helloticket/superfile/helper"
	"github.com/helloticket/superfile/model"
)

type retornoAFDTFile struct {
	content io.Reader
	layout  model.Layout
	decoder *Decoder
}

func (r *retornoAFDTFile) Read() *model.Retorno {
	const registroHeaderArquivo = 1
	const registroTrailerArquivo = 9

	numeroSegmentos := int64(0)
	retorno := model.NewRetorno(r.layout)
	loteCorrente := retorno.NovoLote(1)
	detalheCorrente := loteCorrente.NovoDetalhe()
	loteCorrente.InserirDetalhe(detalheCorrente)
	retorno.InserirLote(loteCorrente)

	reader := bufio.NewScanner(r.content)
	var pos int64 = 0

	for reader.Scan() {
		pos++
		linha := reader.Text()
		tipoRegistro := helper.ToSafeInt(linha[9:10])

		if registroHeaderArquivo == tipoRegistro {
			retorno.Header = r.decodeFileHeader(pos, retorno, linha)
		} else if tipoRegistro == registroTrailerArquivo {
			retorno.Trailer = r.decodeFileTrailer(pos, retorno, linha)
		} else {
			numeroSegmentos++
			segmento := r.decodeSegmento(pos, retorno, linha)
			detalheCorrente[fmt.Sprintf("%d.%s", numeroSegmentos, segmento.Nome)] = segmento.Valores
		}
	}

	return retorno
}

func (r *retornoAFDTFile) decodeSegmento(pos int64, retorno *model.Retorno, row string) model.Segmento {
	decode := DecodeRetorno{decoder: r.decoder, layout: r.layout}

	return decode.Segmento(pos, 9, 10, row, retorno)
}

func (r *retornoAFDTFile) decodeFileHeader(pos int64, retorno *model.Retorno, row string) map[string]interface{} {
	decode := DecodeRetorno{decoder: r.decoder, layout: r.layout}

	return decode.HeaderArquivo(pos, row, retorno)
}

func (r *retornoAFDTFile) decodeFileTrailer(pos int64, retorno *model.Retorno, row string) map[string]interface{} {
	decode := DecodeRetorno{decoder: r.decoder, layout: r.layout}

	return decode.TrailerArquivo(pos, row, retorno)
}
