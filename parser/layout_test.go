package parser

import (
	"testing"

	"github.com/helderfarias/cnab-go/model"
	"github.com/stretchr/testify/assert"
)

func TestCriarLayoutValido(t *testing.T) {
	l, err := NewModeloLayout(model.FileConfigMap{
		"layout": "cnab240",
	})

	assert.NotNil(t, l)
	assert.Nil(t, err)
}

func TestErroCriarLayoutVazio(t *testing.T) {
	l, err := NewModeloLayout(model.FileConfigMap{})

	assert.Nil(t, l)
	assert.EqualError(t, err, "Arquivo de definição de layout vazio")
}

func TestErroCriarLayoutComVersaoInvalida(t *testing.T) {
	l, err := NewModeloLayout(model.FileConfigMap{
		"layout": "cnabinvalido",
	})

	assert.Nil(t, l)
	assert.EqualError(t, err, "Arquivo de definição de layout vazio")
}

func TestCriarLayoutVazio(t *testing.T) {
	l := ModeloLayout{config: model.FileConfigMap{}}

	assert.NotNil(t, l)
}

func TestLancarPanicAoObterRemessaNaoDefinida(t *testing.T) {
	l := ModeloLayout{config: model.FileConfigMap{}}

	assert.PanicsWithValue(t, "Falta seção 'remessa' no arquivo de layout", func() {
		l.GetRemessaLayout()
	})
}

func TestLancarPanicAoObteRetornoNaoDefinido(t *testing.T) {
	l := ModeloLayout{config: model.FileConfigMap{}}

	assert.PanicsWithValue(t, "Falta seção 'retorno' no arquivo de layout", func() {
		l.GetRetornoLayout()
	})
}

func TestObterVersao(t *testing.T) {
	l := ModeloLayout{config: model.FileConfigMap{}}

	assert.Empty(t, l.GetVersao())
}

func TestObterServico(t *testing.T) {
	l := ModeloLayout{config: model.FileConfigMap{}}

	assert.Empty(t, l.GetServico())
}

func TestObterModeloLayout(t *testing.T) {
	l := ModeloLayout{config: model.FileConfigMap{}}

	assert.Empty(t, l.GetLayout())
}

func TestTamanhoRegistroPorLayout(t *testing.T) {
	cnab240 := ModeloLayout{config: model.FileConfigMap{
		"layout": "cnab240",
	}}

	cnab400 := ModeloLayout{config: model.FileConfigMap{
		"layout": "cnab400",
	}}

	assert.Equal(t, int64(240), cnab240.GetTamanhoRegistro())
	assert.Equal(t, int64(400), cnab400.GetTamanhoRegistro())
}

func TestErrorTamanhoRegistroPorLayout(t *testing.T) {
	l := ModeloLayout{config: model.FileConfigMap{}}

	assert.PanicsWithValue(t, "Atributo 'layout' não definido no modelo", func() {
		l.GetTamanhoRegistro()
	})
}
