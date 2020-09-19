package model

import "github.com/stretchr/testify/mock"

type mockLayout struct {
	mock.Mock
}

func (l *mockLayout) GetSegmentoDefinitions() RecordDetailMap {
	args := l.Called()
	return args.Get(0).(RecordDetailMap)
}

func (l *mockLayout) GetRemessaLayout() FileConfigMap {
	args := l.Called()
	return args.Get(0).(FileConfigMap)
}

func (l *mockLayout) GetRetornoLayout() FileConfigMap {
	args := l.Called()
	return args.Get(0).(FileConfigMap)
}

func (l *mockLayout) GetVersao() string {
	args := l.Called()
	return args.String()
}

func (l *mockLayout) GetServico() string {
	args := l.Called()
	return args.String()
}

func (l *mockLayout) GetLayout() string {
	args := l.Called()
	return args.String()
}

func (l *mockLayout) GetTamanhoRegistro() int64 {
	args := l.Called()
	return args.Get(0).(int64)
}

func (l *mockLayout) Validate() error {
	args := l.Called()
	return args.Error(0)
}

func (l *mockLayout) GlobalSettings() map[string]string {
	args := l.Called()
	return args.Get(0).(map[string]string)
}
