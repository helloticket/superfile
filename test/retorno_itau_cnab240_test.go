package test

import (
	"os"
	"strings"
	"testing"

	"github.com/helderfarias/cnab-go"
	"github.com/helderfarias/cnab-go/file"
	"github.com/helderfarias/cnab-go/layout/itau"
	"github.com/stretchr/testify/assert"
)

func TestRetornoItauCnab240Cobranca(t *testing.T) {
	source := strings.NewReader(itau.CNAB240Cobranca)
	layout, err := cnab.NewLayout("240", source)

	f, _ := os.Open("cobranca_itau_cnab240.ret")
	defer f.Close()
	arquivo, err := file.NewRetornoFile(layout, f)
	remessa := arquivo.Read()

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, arquivo)
	assert.NotNil(t, remessa)
	assert.Equal(t, int64(1), remessa.TotalLotes())
}
