package parser

import (
	"strings"

	"github.com/helderfarias/cnab-go/model"
)

type ModeloLayout struct {
	config      model.FileConfigMap
	definitions model.RecordDetailMap
}

func NewModeloLayout(config model.FileConfigMap) *ModeloLayout {
	cache := model.RecordDetailMap{}

	if value, ok := config["remessa"]; ok {
		cache = carregarDetalhes(value.(map[interface{}]interface{}))
	}

	return &ModeloLayout{
		config:      config,
		definitions: cache,
	}
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
	return l.config["versao"].(string)
}

func (l *ModeloLayout) GetServico() string {
	return l.config["servico"].(string)
}

func (l *ModeloLayout) GetLayout() string {
	return l.config["layout"].(string)
}

func (l *ModeloLayout) GetTamanhoRegistro() int64 {
	layout, ok := l.config["layout"].(string)
	if !ok {
		panic("Atributo 'layout' não definido no modelo")
	}

	if "cnab240" == strings.ToLower(layout) {
		return 240
	}

	if "cnab400" == strings.ToLower(layout) {
		return 400
	}

	return 240
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
