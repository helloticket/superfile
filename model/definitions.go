package model

type RecordMap map[string]interface{}

type RecordDetailMap map[string]RecordMap

type FileConfigMap map[interface{}]interface{}

type Layout interface {
	GetSegmentoDefinitions() RecordDetailMap

	GetRemessaLayout() FileConfigMap

	GetRetornoLayout() FileConfigMap

	GetVersao() string

	GetServico() string

	GetLayout() string

	GetTamanhoRegistro() int64
}
