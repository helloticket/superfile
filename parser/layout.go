package parser

import (
	"errors"
	"strings"

	"github.com/helderfarias/superfile/helper"
	"github.com/helderfarias/superfile/model"
)

type ModeloLayout struct {
	config      model.FileConfigMap
	definitions model.RecordDetailMap
}

func NewModeloLayout(config model.FileConfigMap) (*ModeloLayout, error) {
	if len(config) == 0 {
		return nil, errors.New("Arquivo de definição de layout vazio")
	}

	layout := helper.ToString(config["layout"])
	if model.LayoutCNAB400 != layout &&
		model.LayoutCNAB240 != layout &&
		model.LayoutCCSITEF != layout {
		return nil, errors.New("Arquivo de definição de layout vazio")
	}

	cache := model.RecordDetailMap{}
	if value, ok := config["remessa"]; ok {
		cache = carregarDetalhes(value.(map[interface{}]interface{}))
	}

	return &ModeloLayout{
		config:      config,
		definitions: cache,
	}, nil
}

func (l *ModeloLayout) Validate() error {
	if len(l.config) == 0 {
		return errors.New("Arquivo de definição de layout vazio")
	}

	layout := helper.ToString(l.config["layout"])
	if model.LayoutCNAB400 != layout &&
		model.LayoutCNAB240 != layout &&
		model.LayoutCCSITEF != layout {
		return errors.New("Arquivo de definição de layout vazio")
	}

	return nil
}

func (l *ModeloLayout) GetSegmentoDefinitions() model.RecordDetailMap {
	return l.definitions
}

func (l *ModeloLayout) GetRemessaLayout() model.FileConfigMap {
	if len(l.config) == 0 {
		panic("Falta seção 'remessa' no arquivo de layout")
	}

	return l.config["remessa"].(map[interface{}]interface{})
}

func (l *ModeloLayout) GetRetornoLayout() model.FileConfigMap {
	if len(l.config) == 0 {
		panic("Falta seção 'retorno' no arquivo de layout")
	}

	return l.config["retorno"].(map[interface{}]interface{})
}

func (l *ModeloLayout) GetVersao() string {
	if value, ok := l.config["versao"].(string); ok {
		return strings.ToLower(value)
	}
	return ""
}

func (l *ModeloLayout) GetServico() string {
	if value, ok := l.config["servico"].(string); ok {
		return value
	}
	return ""
}

func (l *ModeloLayout) GetLayout() string {
	if value, ok := l.config["layout"].(string); ok {
		return strings.ToLower(value)
	}
	return ""
}

func (l *ModeloLayout) GetTamanhoRegistro() int64 {
	layout, ok := l.config["layout"].(string)
	if !ok {
		panic("Atributo 'layout' não definido no modelo")
	}

	if model.LayoutCNAB240 == strings.ToLower(layout) {
		return 240
	}

	if model.LayoutCNAB400 == strings.ToLower(layout) {
		return 400
	}

	return -1
}

func carregarDetalhes(remessa model.FileConfigMap) model.RecordDetailMap {
	definitions := model.RecordDetailMap{}

	if value, ok := remessa["detalhes"].(map[interface{}]interface{}); ok {
		for key := range value {
			definitions[key.(string)] = model.RecordMap{}
		}
	}

	return definitions
}
