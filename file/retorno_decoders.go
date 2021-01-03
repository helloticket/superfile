package file

import (
	"fmt"
	"strings"

	"github.com/helloticket/superfile/model"
)

type DecodeRetorno struct {
	layout  model.Layout
	decoder *Decoder
}

func (d *DecodeRetorno) Segmento(pos int64, start int64, end int64, row string, retorno *model.Retorno) model.Segmento {
	segmento := fmt.Sprintf("segmento_%s", strings.ToLower(row[start:end]))
	layout := d.getLayoutFor("detalhes")
	if layout[segmento] == nil {
		retorno.RegistrarFalha(pos, fmt.Errorf("Segmento %v não encontrado", segmento))
		return model.Segmento{}
	}

	layout = layout[segmento].(map[interface{}]interface{})
	block := fmt.Sprintf("detalhes.%s", segmento)
	linhas := d.parse(block, row, layout)

	valores := model.RecordMap{}

	for _, l := range linhas {
		val, err := d.decode(l.Block, l)
		valores[l.Name] = val
		valores[fmt.Sprintf("%s_error", l.Name)] = err
		retorno.RegistrarFalha(pos, err)
	}

	return model.Segmento{
		Nome:    segmento,
		Valores: valores,
	}
}

func (d *DecodeRetorno) HeaderArquivo(pos int64, row string, retorno *model.Retorno) map[string]interface{} {
	header := map[string]interface{}{}
	if d.getLayoutFor("header_arquivo") == nil {
		return header
	}

	linhas := d.parse("header_arquivo", row, d.getLayoutFor("header_arquivo"))

	for _, linha := range linhas {
		val, err := d.decode(linha.Block, linha)
		header[linha.Name] = val
		header[fmt.Sprintf("%s_error", linha.Name)] = err
		retorno.RegistrarFalha(pos, err)
	}

	return header
}

func (d *DecodeRetorno) TrailerArquivo(pos int64, row string, retorno *model.Retorno) map[string]interface{} {
	trailer := map[string]interface{}{}

	linhas := d.parse("trailer_arquivo", row, d.getLayoutFor("trailer_arquivo"))

	for _, linha := range linhas {
		val, err := d.decode(linha.Block, linha)
		trailer[linha.Name] = val
		trailer[fmt.Sprintf("%s_error", linha.Name)] = err
		retorno.RegistrarFalha(pos, err)
	}

	return trailer
}

func (d *DecodeRetorno) TrailerLote(pos int64, row string, retorno *model.Retorno) map[string]interface{} {
	trailer := map[string]interface{}{}
	if d.getLayoutFor("trailer_lote") == nil {
		return trailer
	}

	linhas := d.parse("trailer_lote", row, d.getLayoutFor("trailer_lote"))

	for _, linha := range linhas {
		val, err := d.decode(linha.Block, linha)
		trailer[linha.Name] = val
		trailer[fmt.Sprintf("%s_error", linha.Name)] = err
		retorno.RegistrarFalha(pos, err)
	}

	return trailer
}

func (d *DecodeRetorno) HeaderLote(pos int64, row string, retorno *model.Retorno) map[string]interface{} {
	header := map[string]interface{}{}
	if d.getLayoutFor("header_lote") == nil {
		return header
	}

	linhas := d.parse("header_lote", row, d.getLayoutFor("header_lote"))

	for _, linha := range linhas {
		val, err := d.decode(linha.Block, linha)
		header[linha.Name] = val
		header[fmt.Sprintf("%s_error", linha.Name)] = err
		retorno.RegistrarFalha(pos, err)
	}

	return header
}

func (d *DecodeRetorno) parse(block string, row string, layout model.FileConfigMap) []model.Linha {
	return d.decoder.Parse(block, row, layout)
}

func (d *DecodeRetorno) decode(block string, row model.Linha) (interface{}, error) {
	return d.decoder.Decode(block, row)
}

func (d *DecodeRetorno) getLayoutFor(name string) map[interface{}]interface{} {
	config := d.layout.GetRetornoLayout()
	if config[name] == nil {
		return nil
	}

	return config[name].(map[interface{}]interface{})
}
