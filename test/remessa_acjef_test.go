package test

import (
	"strings"
	"testing"
	"time"

	"github.com/helloticket/superfile"
	"github.com/helloticket/superfile/file"
	"github.com/helloticket/superfile/layout/mte"
	"github.com/stretchr/testify/assert"
)

func TestRemessaACJEF(t *testing.T) {
	source := strings.NewReader(mte.ACJEF)
	layout, err := superfile.NewLayout(source)

	remessa := superfile.NewRemessa(layout)
	remessa.Header["seq_registro"] = 1
	remessa.Header["identificador_empregador"] = 1
	remessa.Header["cpf_cnpj_empregador"] = 94281944095
	remessa.Header["cei_empregador"] = 200000000000
	remessa.Header["razao_social_empregador"] = "1"
	remessa.Header["data_inicio"] = time.Now()
	remessa.Header["data_fim"] = time.Now()
	remessa.Header["data_geracao"] = time.Now()
	remessa.Header["hora_geracao"] = time.Now()

	lote := remessa.NovoLote()
	remessa.InserirLote(lote)

	seq := 1

	for i := 1; i <= 10; i++ {
		seq++
		detalhe := lote.NovoDetalhe()
		detalhe["segmento_2"]["seq_registro"] = seq
		detalhe["segmento_2"]["codigo_horario"] = 12
		detalhe["segmento_2"]["entrada"] = time.Now()
		detalhe["segmento_2"]["inicio_intervalo"] = time.Now()
		detalhe["segmento_2"]["fim_intervalo"] = time.Now()
		detalhe["segmento_2"]["saida"] = time.Now()
		lote.InserirDetalhe(detalhe)
	}

	for i := 1; i <= 10; i++ {
		seq++
		detalhe := lote.NovoDetalhe()
		detalhe["segmento_3"]["seq_registro"] = seq
		detalhe["segmento_3"]["numero_pis"] = 200000000000
		detalhe["segmento_3"]["data_inicio_jornada"] = time.Now()
		detalhe["segmento_3"]["primeiro_horario_jornada"] = 12
		detalhe["segmento_3"]["codigo_horario_previsto_jornada"] = 1
		detalhe["segmento_3"]["horas_diurnas_nao_extraordinarias"] = 1
		detalhe["segmento_3"]["horas_noturnas_nao_extraordinarias"] = 2
		detalhe["segmento_3"]["horas_extras_1"] = time.Now()
		detalhe["segmento_3"]["percentual_adicional_horas_extras_1"] = 12
		detalhe["segmento_3"]["modalidades_hora_extra_1"] = "0"
		detalhe["segmento_3"]["horas_extras_2"] = time.Now()
		detalhe["segmento_3"]["percentual_adicional_horas_extras_2"] = 12
		detalhe["segmento_3"]["modalidades_hora_extra_2"] = "0"
		detalhe["segmento_3"]["horas_extra_3"] = time.Now()
		detalhe["segmento_3"]["percentual_adicional_horas_extras_3"] = 12
		detalhe["segmento_3"]["modalidades_hora_extra_3"] = "0"
		detalhe["segmento_3"]["horas_extra_4"] = time.Now()
		detalhe["segmento_3"]["percentual_adicional_horas_extras_4"] = 12
		detalhe["segmento_3"]["modalidades_hora_extra_4"] = "1"
		detalhe["segmento_3"]["horas_faltas_atrasos"] = time.Now()
		detalhe["segmento_3"]["sinal_horas_compensar"] = 1
		detalhe["segmento_3"]["saldo_horas_compensar"] = time.Now()
		lote.InserirDetalhe(detalhe)
	}

	remessa.Trailer["seq_registro"] = 1
	remessaFile := file.NewRemessaFile(remessa, "acjef.rem")

	arquivo := remessaFile.Write()

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, remessa)
	assert.NotNil(t, arquivo)
	assertFile(t, 22, arquivo.Name(), true)
}
