package test

import (
	"os"
	"strings"
	"testing"

	"github.com/helloticket/superfile"
	"github.com/helloticket/superfile/file"
	"github.com/helloticket/superfile/layout/bb"
	"github.com/stretchr/testify/assert"
)

func TestRetornosBBExtratoContaCorrenteCnab240Cobranca(t *testing.T) {
	source := strings.NewReader(bb.CNAB240ExtratoContaCorrente)
	layout, err1 := superfile.NewLayout(source)

	f, _ := os.Open("fixtures/bb_extrato_cc.ret")
	defer f.Close()
	arquivo, err2 := superfile.NewRetornoFile(layout, f)
	retorno := arquivo.Read()

	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.NotNil(t, layout)
	assert.NotNil(t, arquivo)
	assert.NotNil(t, retorno)
	assert.Equal(t, int64(1), retorno.TotalLotes())

	lote := retorno.Lotes[0]
	assert.Equal(t, int64(124), lote.Header["lote_servico"])
	assert.Equal(t, int64(12234567000186), lote.Header["numero_inscricao_empresa"])
	assert.Equal(t, "NOME EMPRESA LOTE", lote.Header["nome_empresa"])

	segmentos := retorno.Segmentos()
	assert.Len(t, segmentos, 25)
}

func TestRemessaBBExtratoContaCorrenteCnab240Cobranca(t *testing.T) {
	source := strings.NewReader(bb.CNAB240ExtratoContaCorrente)
	layout, err := superfile.NewLayout(source)

	remessa := superfile.NewRemessa(layout)

	lote := remessa.NovoLote()
	lote.Header["codigo_banco"] = 1
	lote.Header["lote_servico"] = 124
	lote.Header["tipo_inscricao_empresa"] = 0
	lote.Header["numero_inscricao_empresa"] = 12234567000186
	lote.Header["codigo_convenio_banco"] = "0"
	lote.Header["agencia_mantenedora_conta"] = 00000
	lote.Header["digito_verificador_agencia"] = "0"
	lote.Header["numero_conta_corrente"] = "0"
	lote.Header["digito_verificador_conta_corrente"] = "0"
	lote.Header["digito_verificador_agencia_conta_corrente"] = "0"
	lote.Header["nome_empresa"] = "NOME EMPRESA LOTE"
	lote.Header["data_saldo_inicial"] = 00000000
	lote.Header["valor_saldo_inicial"] = 000000000000000000
	lote.Header["situacao_saldo_inicial"] = "0"
	lote.Header["posicao_saldo_inicial"] = "0"
	lote.Header["moeda_referenciada_extrato"] = "0"
	lote.Header["numero_sequencia_extrato"] = 00000
	remessa.InserirLote(lote)

	for i := 1; i <= 25; i++ {
		detalhe := lote.NovoDetalhe()
		detalhe["segmento_e"]["codigo_banco"] = 000
		detalhe["segmento_e"]["lote_servico"] = 10 + i
		detalhe["segmento_e"]["tipo_inscricao_empresa"] = 0
		detalhe["segmento_e"]["numero_inscricao_empresa"] = 00000000000000
		detalhe["segmento_e"]["codigo_convenio_banco"] = "0"
		detalhe["segmento_e"]["agencia_mantenedora_conta"] = 00000
		detalhe["segmento_e"]["digito_verificador_agencia"] = "0"
		detalhe["segmento_e"]["numero_conta_corrente"] = 000000000000
		detalhe["segmento_e"]["digito_verificador_conta_corrente"] = "0"
		detalhe["segmento_e"]["digito_verificador_agencia_conta_corrente"] = "0"
		detalhe["segmento_e"]["nome_empresa"] = "NOME EMPRESA LOTE"
		detalhe["segmento_e"]["natureza_lancamento"] = "0"
		detalhe["segmento_e"]["tipo_complemento_lancamento"] = 00
		detalhe["segmento_e"]["complemento_lancamento"] = "0"
		detalhe["segmento_e"]["identificacao_isencao_cpmf"] = "0"
		detalhe["segmento_e"]["data_contabil"] = 00000000
		detalhe["segmento_e"]["data_lancamento"] = 00000000
		detalhe["segmento_e"]["valor_lancamento"] = 000000000000000000
		detalhe["segmento_e"]["tipo_lancamento"] = "0"
		detalhe["segmento_e"]["categoria_lancamento"] = 000
		detalhe["segmento_e"]["codigo_historico_banco"] = "0"
		detalhe["segmento_e"]["descricao_historico_lancamento_banco"] = "0"
		detalhe["segmento_e"]["numero_documento_complemento"] = "0"
		lote.InserirDetalhe(detalhe)
	}

	lote.Trailer["codigo_banco"] = 000
	lote.Trailer["lote_servico"] = 0000
	lote.Trailer["tipo_inscricao_empresa"] = 0
	lote.Trailer["numero_inscricao_empresa"] = 00000000000000
	lote.Trailer["codigo_convenio_banco"] = "0"
	lote.Trailer["agencia_mantenedora_conta"] = 00000
	lote.Trailer["digito_verificador_agencia"] = "0"
	lote.Trailer["numero_conta_corrente"] = 000000000000
	lote.Trailer["digito_verificador_conta_corrente"] = "0"
	lote.Trailer["digito_verificador_agencia_conta"] = "0"
	lote.Trailer["saldo_bloqueado_acima_24horas"] = 000000000000000000
	lote.Trailer["limite_conta"] = 000000000000000000
	lote.Trailer["saldo_bloqueado_ate_24horas"] = 000000000000000000
	lote.Trailer["data_saldo_final"] = 00000000
	lote.Trailer["valor_saldo_final"] = 000000000000000000
	lote.Trailer["situacao_saldo_final"] = "0"
	lote.Trailer["posicao_saldo_final"] = "0"
	lote.Trailer["quantidade_registros_lote"] = 000000
	lote.Trailer["somatorio_valores_debito"] = 000000000000000000
	lote.Trailer["somatorio_valores_credito"] = 000000000000000000

	remessaFile := file.NewRemessaFile(remessa, "bb_extrato_cc.rem")

	arquivo := remessaFile.Write()

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, remessa)
	assert.NotNil(t, arquivo)
	assertFile(t, 27, arquivo.Name(), true)
}
