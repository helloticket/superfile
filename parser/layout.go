package parser

import (
	"errors"
	"strings"

	"github.com/helloticket/superfile/helper"
	"github.com/helloticket/superfile/model"
)

type ModeloLayout struct {
	config      model.FileConfigMap
	definitions model.RecordDetailMap
	global      map[string]string
}

func NewModeloLayout(config model.FileConfigMap) (*ModeloLayout, error) {
	if len(config) == 0 {
		return nil, errors.New("Arquivo de definição de layout vazio")
	}

	layout := helper.ToString(config["layout"])
	if !model.RegisterLayouts[layout] {
		return nil, errors.New("Arquivo de definição de layout vazio")
	}

	globalSettings := map[string]string{}

	if global, ok := config["global"].(map[interface{}]interface{}); ok && global != nil {
		if value, ok := global["alinhamento_alfanumerico"].(string); ok && value != "" {
			globalSettings["global_alinhamento_alfanumerico"] = value
		}
	}

	return &ModeloLayout{
		config: config,
		global: globalSettings,
	}, nil
}

func (l *ModeloLayout) GlobalSettings() map[string]string {
	return l.global
}

func (l *ModeloLayout) Validate() error {
	if len(l.config) == 0 {
		return errors.New("Arquivo de definição de layout vazio")
	}

	layout := helper.ToString(l.config["layout"])
	if !model.RegisterLayouts[layout] {
		return errors.New("Arquivo de definição de layout vazio")
	}

	return nil
}

func (l *ModeloLayout) GetSegmentoDefinitions() model.RecordDetailMap {
	cache := model.RecordDetailMap{}
	if value, ok := l.config["remessa"]; ok {
		cache = carregarDetalhes(value.(map[interface{}]interface{}))
	}
	return cache
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
