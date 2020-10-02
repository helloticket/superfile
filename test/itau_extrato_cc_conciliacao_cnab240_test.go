package test

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/helloticket/superfile"
	"github.com/helloticket/superfile/file"
	"github.com/helloticket/superfile/layout/itau"
	"github.com/stretchr/testify/assert"
)

func TestRetornosItauExtratoContaCorrenteCnab240Cobranca(t *testing.T) {
	source := strings.NewReader(itau.CNAB240ExtratoContaCorrente)
	layout, err1 := superfile.NewLayout(source)

	f, _ := os.Open("fixtures/itau_extrato_cc.ret")
	defer f.Close()
	arquivo, err2 := superfile.NewRetornoFile(layout, f)
	retorno := arquivo.Read()

	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.NotNil(t, layout)
	assert.NotNil(t, arquivo)
	assert.NotNil(t, retorno)
	assert.Equal(t, int64(1), retorno.TotalLotes())
	assert.Equal(t, int64(00000000000000), retorno.Header["numero_cnpj_cpf_cliente"])

	lote := retorno.Lotes[0]
	assert.Equal(t, int64(0), lote.Header["lote_servico"])
	assert.Equal(t, int64(00000000000000), lote.Header["numero_conta_corrente_cliente"])
	assert.Equal(t, "0", lote.Header["nome_empresa"])

	segmentos := retorno.Segmentos()
	assert.Len(t, segmentos, 25)
}

func TestRemessaItauExtratoContaCorrenteCnab240Cobranca(t *testing.T) {
	source := strings.NewReader(itau.CNAB240ExtratoContaCorrente)
	layout, err := superfile.NewLayout(source)

	remessa := superfile.NewRemessa(layout)

	lote := remessa.NovoLote()
	remessa.Header["complemento_header"] = "0"
	remessa.Header["tipo_inscricao_cliente"] = "0"
	remessa.Header["numero_cnpj_cpf_cliente"] = 00000000000000
	remessa.Header["complemento_registro_01"] = "0"
	remessa.Header["codigo_empresa_banco"] = "0"
	remessa.Header["numero_agencia_mantenedora_conta"] = 0000
	remessa.Header["digito_verificador_agencia"] = "0"
	remessa.Header["numero_conta_corrente_cliente"] = 00000
	remessa.Header["complemento_registro_04"] = "0"
	remessa.Header["digito_verificador_agencia_conta"] = "0"
	remessa.Header["nome_empresa"] = "0"
	remessa.Header["complemento_registro_05"] = "0"
	remessa.Header["data_geracao"] = time.Now()
	remessa.Header["hora_geracao"] = time.Now()
	remessa.Header["numero_sequencial_arquivo"] = 000000
	remessa.Header["uso_reservado_banco"] = "0"
	remessa.Header["complemento_registro_07"] = "0"
	remessa.InserirLote(lote)

	lote.Header["lote_servico"] = 0000
	lote.Header["complemento_registro_01"] = "0"
	lote.Header["tipo_inscricao_cliente"] = "0"
	lote.Header["numero_inscricao_cliente"] = 00000000000000
	lote.Header["identificacao_tipo_conta"] = "0"
	lote.Header["complemento_registro_02"] = "0"
	lote.Header["codigo_empresa_banco"] = "0"
	lote.Header["numero_agencia_mantenedora_conta"] = 0000
	lote.Header["digito_verificador_agencia"] = "0"
	lote.Header["numero_conta_corrente_cliente"] = 00000
	lote.Header["complemento_registro_05"] = "0"
	lote.Header["digito_verificador_agencia_conta"] = "0"
	lote.Header["nome_empresa"] = "0"
	lote.Header["complemento_registro_06"] = "0"
	lote.Header["data_saldo_inicial"] = time.Now()
	lote.Header["valor_saldo_inicial"] = float64(10.30)
	lote.Header["situacao_saldo_inicial"] = "0"
	lote.Header["posicao_saldo_inicial"] = "0"
	lote.Header["numero_sequencia_extrato"] = 00000
	lote.Header["complemento_registro_07"] = "0"

	for i := 1; i <= 25; i++ {
		detalhe := lote.NovoDetalhe()
		detalhe["segmento_e"]["lote_servico"] = 0000
		detalhe["segmento_e"]["sequencial_registro_lote"] = 00000
		detalhe["segmento_e"]["identificacao_tipo_lancamento"] = 0
		detalhe["segmento_e"]["tipo_inscricao_cliente"] = "0"
		detalhe["segmento_e"]["numero_inscricao_cliente"] = 00000000000000
		detalhe["segmento_e"]["codigo_empresa_banco"] = "00000"
		detalhe["segmento_e"]["numero_agencia_mantenedora_conta"] = 0000
		detalhe["segmento_e"]["digito_verificador_agencia"] = "0"
		detalhe["segmento_e"]["numero_conta_corrente"] = 00000
		detalhe["segmento_e"]["digito_verificador_agencia_conta"] = "0"
		detalhe["segmento_e"]["nome_empresa"] = "NOMEEMPRESADETESTE444444444444"
		detalhe["segmento_e"]["uso_banco"] = "000000"
		detalhe["segmento_e"]["natureza_lancamento"] = "000"
		detalhe["segmento_e"]["tipo_complemento_lancamento"] = 00
		detalhe["segmento_e"]["banco_origem_lancamento"] = 000
		detalhe["segmento_e"]["agencia_conta_origem_lancamento"] = 000000000000
		detalhe["segmento_e"]["idenficacao_isencao_cpmf"] = "0"
		detalhe["segmento_e"]["data_contabil"] = time.Now()
		detalhe["segmento_e"]["data_lancamento"] = time.Now()
		detalhe["segmento_e"]["valor_lancamento"] = float64(10.30)
		detalhe["segmento_e"]["tipo_lancamento"] = "0"
		detalhe["segmento_e"]["categoria_lancamento"] = 000
		detalhe["segmento_e"]["identificacao_codigo_fluxo_caixa"] = "0000"
		detalhe["segmento_e"]["descricao_historico_lacto_banco"] = "0000000000000000000000000"
		detalhe["segmento_e"]["agencia_origem_lancamento"] = 0000
		detalhe["segmento_e"]["conta_origem_lancamento"] = 00000
		detalhe["segmento_e"]["dac_agencia_conta_origem"] = 0
		detalhe["segmento_e"]["tipo_inscricao_emitente"] = "0"
		detalhe["segmento_e"]["numero_inscricao_emitente"] = "00000000000000"
		detalhe["segmento_e"]["numero_documento_complemento"] = "888888"
		lote.InserirDetalhe(detalhe)
	}

	lote.Trailer["lote_servico"] = 0000
	lote.Trailer["complemento_registro_01"] = "0"
	lote.Trailer["tipo_inscricao_cliente"] = "0"
	lote.Trailer["numero_inscricao_cliente"] = 00000000000000
	lote.Trailer["complemento_registro_02"] = "0"
	lote.Trailer["codigo_empresa_banco"] = "0"
	lote.Trailer["numero_agencia_mantenedora_conta"] = 0000
	lote.Trailer["digito_verificador_agencia"] = "0"
	lote.Trailer["numero_conta_corrente"] = 00000
	lote.Trailer["complemento_registro_05"] = "0"
	lote.Trailer["digito_verificador_agencia_conta"] = "0"
	lote.Trailer["complemento_registro_06"] = "0"
	lote.Trailer["saldo_bloqueado_acima_24_horas"] = 000000000000000000
	lote.Trailer["limite_conta"] = 000000000000000000
	lote.Trailer["saldo_bloqueado_ate_24_horas"] = 000000000000000000
	lote.Trailer["data_saldo_final"] = time.Now()
	lote.Trailer["valor_saldo_inicial"] = float64(10.30)
	lote.Trailer["situacao_saldo_inicial"] = "0"
	lote.Trailer["posicao_saldo_inicial"] = "0"
	lote.Trailer["quantidade_registros_lote"] = 000000
	lote.Trailer["somatorio_valores_a_debito"] = 000000000000000000
	lote.Trailer["somatorio_valores_a_credito"] = 000000000000000000
	lote.Trailer["somatorio_lancamentos_dia_lancamentos_futuros"] = 000000000000000000

	remessa.Trailer["complemento_registro_01"] = "0"
	remessa.Trailer["quantidade_lotes"] = 000000
	remessa.Trailer["quantidade_registros"] = 000000
	remessa.Trailer["quantidade_contas_conciliacao"] = 000000
	remessa.Trailer["complemento_registro_02"] = "0"

	remessaFile := file.NewRemessaFile(remessa, "itau_extrato_cc.rem")

	arquivo := remessaFile.Write()

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, remessa)
	assert.NotNil(t, arquivo)
	assertFile(t, 29, arquivo.Name(), true)
}
