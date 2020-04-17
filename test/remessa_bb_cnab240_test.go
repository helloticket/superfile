package test

import (
	"strings"
	"testing"
	"time"

	"github.com/helderfarias/cnab-go"
	"github.com/helderfarias/cnab-go/file"
	"github.com/helderfarias/cnab-go/layout/bb"
	"github.com/stretchr/testify/assert"
)

func TestRemessaBBCnab240Cobranca(t *testing.T) {
	source := strings.NewReader(bb.CNAB240Cobranca)
	layout, err := cnab.NewLayout(source)
	remessa := cnab.NewRemessa(layout)

	remessa.Header["tipo_inscricao_empresa"] = 0
	remessa.Header["numero_inscricao_empresa"] = 00000000000000
	remessa.Header["numero_convenio"] = 000000000
	remessa.Header["cobranca_cedente"] = 0000
	remessa.Header["numero_carteira_cobranca"] = 00
	remessa.Header["variacao_carteira_cobranca"] = 000
	remessa.Header["campo_reservado_bb"] = "0"
	remessa.Header["agencia_mantenedora_conta"] = 00000
	remessa.Header["digito_verificador_agencia"] = "0"
	remessa.Header["numero_conta_corrente"] = 000000000000
	remessa.Header["digito_verificador_conta"] = "0"
	remessa.Header["digito_verificador_agencia_conta"] = "0"
	remessa.Header["nome_empresa"] = "0"
	remessa.Header["codigo_remessa_retorno"] = 0
	remessa.Header["data_geracao_arquivo"] = time.Date(2017, 5, 10, 0, 0, 0, 0, time.UTC)
	remessa.Header["hora_geracao_arquivo"] = time.Date(2017, 5, 10, 0, 0, 0, 0, time.UTC)
	remessa.Header["numero_sequencial_arquivo"] = 000000

	lote := remessa.NovoLote()

	lote.Header["lote_servico"] = lote.Sequencial
	lote.Header["tipo_inscricao_empresa"] = 2
	lote.Header["numero_inscricao_empresa"] = 122345670001867
	lote.Header["numero_convenio"] = 2932
	lote.Header["cobranca_cedente"] = 123456789
	lote.Header["numero_carteira_cobranca"] = 12
	lote.Header["variacao_carteira_cobranca"] = 123
	lote.Header["agencia_mantenedora_conta"] = 12345
	lote.Header["digito_verificador_agencia"] = "1"
	lote.Header["numero_conta_corrente"] = 123456789101
	lote.Header["digito_verificador_conta"] = "1"
	lote.Header["digito_verificador_agencia_conta"] = "1"
	lote.Header["numero_remessa_retorno"] = 12345678
	lote.Header["nome_empresa"] = "NOME EMPRESA LOTE"
	lote.Header["data_gravacao"] = time.Date(2017, 5, 10, 0, 0, 0, 0, time.UTC)
	lote.Header["data_credito"] = time.Date(2017, 5, 10, 0, 0, 0, 0, time.UTC)

	detalhe := lote.NovoDetalhe()

	detalhe["segmento_p"]["lote_servico"] = 0000
	detalhe["segmento_p"]["numero_registro"] = 00000
	detalhe["segmento_p"]["agencia_conta"] = 00000
	detalhe["segmento_p"]["verificador_agencia"] = "0"
	detalhe["segmento_p"]["numero_conta"] = 000000000000
	detalhe["segmento_p"]["verificador_conta"] = "0"
	detalhe["segmento_p"]["identificacao_titulo"] = "0"
	detalhe["segmento_p"]["codigo_carteira"] = 0
	detalhe["segmento_p"]["numero_documento"] = "0"
	detalhe["segmento_p"]["data_vencimento"] = 00000000
	detalhe["segmento_p"]["valor_nominal"] = 000000000000000
	detalhe["segmento_p"]["data_emissao"] = 00000000
	detalhe["segmento_p"]["codigo_juros_mora"] = 0
	detalhe["segmento_p"]["data_juros_mora"] = 00000000
	detalhe["segmento_p"]["juros_mora_dia"] = 000000000000000

	detalhe["segmento_q"]["lote_servico"] = 0000
	detalhe["segmento_q"]["numero_registro"] = 00000
	detalhe["segmento_q"]["codigo_movimento_remessa"] = 00
	detalhe["segmento_q"]["tipo_inscricao"] = 0
	detalhe["segmento_q"]["numero_inscricao"] = 000000000000000
	detalhe["segmento_q"]["nome"] = "0"
	detalhe["segmento_q"]["endereco"] = "0"
	detalhe["segmento_q"]["bairro"] = "0"
	detalhe["segmento_q"]["cep"] = 00000
	detalhe["segmento_q"]["sufixo_cep"] = 000
	detalhe["segmento_q"]["cidade"] = "0"
	detalhe["segmento_q"]["uf"] = "0"

	detalhe["segmento_r"]["lote_servico"] = 0000
	detalhe["segmento_r"]["numero_registro"] = 00000
	detalhe["segmento_r"]["codigo_movimento_remessa"] = 00

	detalhe["segmento_s"]["lote_servico"] = 0000
	detalhe["segmento_s"]["numero_registro"] = 00000
	detalhe["segmento_s"]["codigo_movimento_remessa"] = 00

	detalhe["segmento_y01"]["lote_servico"] = 00000
	detalhe["segmento_y01"]["numero_registro"] = 00000
	detalhe["segmento_y01"]["codigo_movimento_remessa"] = 00
	detalhe["segmento_y01"]["tipo_inscricao"] = 0
	detalhe["segmento_y01"]["numero_inscricao"] = 000000000000000
	detalhe["segmento_y04"]["lote_servico"] = 0000
	detalhe["segmento_y04"]["numero_registro"] = 00000
	detalhe["segmento_y04"]["codigo_movimento"] = 00
	detalhe["segmento_y04"]["email"] = "0"
	detalhe["segmento_y04"]["codigo_ddd"] = 00
	detalhe["segmento_y04"]["numero_celular"] = 00000000

	detalhe["segmento_y05"]["lote_servico"] = 0000
	detalhe["segmento_y05"]["numero_registro"] = 00000
	detalhe["segmento_y05"]["codigo_movimento"] = 00
	detalhe["segmento_y05"]["identificacao_cheque_01"] = "0"

	detalhe["segmento_y50"]["lote_servico"] = 0000
	detalhe["segmento_y50"]["numero_registro"] = 00000
	detalhe["segmento_y50"]["codigo_movimento"] = 00
	detalhe["segmento_y50"]["agencia_mantenedora"] = 00000
	detalhe["segmento_y50"]["verificador_agencia_mantenedora"] = "0"
	detalhe["segmento_y50"]["numero_conta"] = 000000000000
	detalhe["segmento_y50"]["verificador_conta"] = "0"
	detalhe["segmento_y50"]["verificador_agencia_conta"] = "0"
	detalhe["segmento_y50"]["numero_identificacao_titulo"] = "0"
	detalhe["segmento_y50"]["codigo_calculo_rateio"] = 0
	detalhe["segmento_y50"]["tipo_valor"] = 0
	detalhe["segmento_y50"]["valor"] = 000000000000000
	detalhe["segmento_y50"]["codigo_banco_credito_benef"] = 000
	detalhe["segmento_y50"]["codigo_agencia_credito_benef"] = 00000
	detalhe["segmento_y50"]["digito_agencia_credito_benef"] = "0"
	detalhe["segmento_y50"]["conta_credito_beneficiario"] = 000000000000
	detalhe["segmento_y50"]["digito_conta_credito_benef"] = "0"
	detalhe["segmento_y50"]["digito_agencia_conta_benef"] = "0"
	detalhe["segmento_y50"]["nome_beneficiario"] = "0"
	detalhe["segmento_y50"]["identificacao_parcela_rateio"] = "0"
	detalhe["segmento_y50"]["qtde_dias_credito_benef"] = 000
	detalhe["segmento_y50"]["data_credito_beneficiario"] = 00000000
	detalhe["segmento_y50"]["identificacao_rejeicoes"] = 0000000000

	detalhe["segmento_y51"]["lote_servico"] = 0000
	detalhe["segmento_y51"]["numero_registro"] = 00000
	detalhe["segmento_y51"]["codigo_movimento"] = 00
	detalhe["segmento_y51"]["numero_nf_01"] = "0"
	detalhe["segmento_y51"]["valor_nf_01"] = 000000000000000
	detalhe["segmento_y51"]["data_emissao_nf_01"] = 00000000

	lote.InserirDetalhe(detalhe)

	lote.Trailer["lote_servico"] = lote.Sequencial
	lote.Trailer["quantidade_registros_lote"] = 1

	remessa.InserirLote(lote)

	remessa.Trailer["lote_servico"] = 9999
	remessa.Trailer["tipo_registro"] = 9
	remessa.Trailer["quantidade_lotes_arquivo"] = remessa.TotalLotes()
	remessa.Trailer["quantidade_registros_arquivo"] = remessa.TotalRegistros()
	remessa.Trailer["quantidade_contas_conciliacao_lotes"] = 1

	remessaFile := file.NewRemessaFile(remessa, "bb-cobranca-cnab240.rem")

	arquivo := remessaFile.Write()

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, remessa)
	assert.NotNil(t, arquivo)
	assertFile(t, 13, arquivo.Name(), true)
}

func TestRemessaBBCnab240Pagamento(t *testing.T) {
	source := strings.NewReader(bb.CNAB240Pagamentos)
	layout, err := cnab.NewLayout(source)
	remessa := cnab.NewRemessa(layout)

	remessa.Header["codigo_banco"] = 000
	remessa.Header["tipo_inscricao_empresa"] = 0
	remessa.Header["numero_inscricao_empresa"] = 00000000000000
	remessa.Header["numero_convenio"] = 000000000
	remessa.Header["codigo_convenio"] = 0000
	remessa.Header["reservado_banco_01"] = "0"
	remessa.Header["agencia_mantenedora_conta"] = 00000
	remessa.Header["digito_verificador_agencia"] = "0"
	remessa.Header["numero_conta_corrente"] = 000000000000
	remessa.Header["digito_verificador_conta"] = "0"
	remessa.Header["digito_verificador_agencia_conta"] = "0"
	remessa.Header["nome_empresa"] = "0"
	remessa.Header["codigo_remessa_retorno"] = 0
	remessa.Header["data_geracao_arquivo"] = 00000000
	remessa.Header["hora_geracao_arquivo"] = 000000
	remessa.Header["numero_sequencial_arquivo"] = 000000

	lote := remessa.NovoLote()

	lote.Header["lote_servico"] = 0000
	lote.Header["tipo_servico"] = 00
	lote.Header["forma_lancamento"] = 00
	lote.Header["tipo_inscricao_empresa"] = 0
	lote.Header["numero_inscricao_empresa"] = 00000000000000
	lote.Header["numero_convenio"] = 000000000
	lote.Header["codigo_convenio_banco"] = "0"
	lote.Header["codigo_convenio"] = 0000
	lote.Header["agencia_mantenedora_conta"] = 00000
	lote.Header["digito_verificador_agencia"] = "0"
	lote.Header["numero_conta_corrente"] = 000000000000
	lote.Header["digito_verificador_conta"] = "0"
	lote.Header["digito_verificador_agencia_conta"] = "0"
	lote.Header["nome_empresa"] = "0"
	lote.Header["logradouro"] = "0"
	lote.Header["numero"] = 00000
	lote.Header["complemento"] = "0"
	lote.Header["cidade"] = "0"
	lote.Header["cep"] = 00000
	lote.Header["complemento_cep"] = "0"
	lote.Header["estado"] = "0"

	detalhe := lote.NovoDetalhe()
	detalhe["segmento_a"]["codigo_banco"] = 000
	detalhe["segmento_a"]["lote_servico"] = 0000
	detalhe["segmento_a"]["numero_sequencial_registro_lote"] = 00000
	detalhe["segmento_a"]["tipo_movimento"] = 0
	detalhe["segmento_a"]["codigo_instrucao_movimento"] = 00
	detalhe["segmento_a"]["codigo_camara_centralizadora"] = 000
	detalhe["segmento_a"]["codigo_banco_favorecido"] = 000
	detalhe["segmento_a"]["agencia_mantenedora_conta_favorecido"] = 00000
	detalhe["segmento_a"]["digito_verificador_agencia"] = "0"
	detalhe["segmento_a"]["numero_conta_corrente"] = 000000000000
	detalhe["segmento_a"]["digito_verificador_conta"] = "0"
	detalhe["segmento_a"]["digito_verificador_agencia_conta"] = "0"
	detalhe["segmento_a"]["nome_favorecido"] = "0"
	detalhe["segmento_a"]["numero_documento_atribuido_empresa"] = "0"
	detalhe["segmento_a"]["data_pagamento"] = 00000000
	detalhe["segmento_a"]["tipo_moeda"] = "0"
	detalhe["segmento_a"]["quantidade_moeda"] = 000000000000000
	detalhe["segmento_a"]["valor_pagamento"] = 000000000000000
	detalhe["segmento_a"]["numero_documento_atribuido_banco"] = "0"
	detalhe["segmento_a"]["data_real_efetivacao_pagamento"] = 00000000
	detalhe["segmento_a"]["valor_real_efetivacao_pagamento"] = 000000000000000
	detalhe["segmento_a"]["outras_informacoes"] = "0"
	detalhe["segmento_a"]["complemento_tipo_servico"] = "0"
	detalhe["segmento_a"]["codigo_finalidade_ted"] = "0"
	detalhe["segmento_a"]["complemento_finalidade_pagamento"] = "0"
	detalhe["segmento_a"]["aviso_favorecido"] = 0
	detalhe["segmento_a"]["codigos_ocorrencias_retorno"] = "0"

	detalhe["segmento_b"]["lote_servico"] = 0000
	detalhe["segmento_b"]["numero_sequencial_registro_lote"] = 00000
	detalhe["segmento_b"]["tipo_inscricao_favorecido"] = 0
	detalhe["segmento_b"]["numero_inscricao_favorecido"] = 00000000000000

	detalhe["segmento_c"]["lote_servico"] = 0000
	detalhe["segmento_c"]["numero_sequencial_registro_lote"] = 00000

	lote.InserirDetalhe(detalhe)

	lote.Trailer["lote_servico"] = 0000
	lote.Trailer["quantidade_registros_lote"] = 000000
	lote.Trailer["somatoria_valores"] = 000000000000000000
	lote.Trailer["codigos_ocorrencias_retorno"] = "0"

	remessa.InserirLote(lote)

	remessa.Trailer["quantidade_lotes_arquivo"] = 000000
	remessa.Trailer["quantidade_registros_arquivo"] = 000000

	remessaFile := file.NewRemessaFile(remessa, "bb-pagamentos-cnab240.rem")

	arquivo := remessaFile.Write()

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, remessa)
	assert.NotNil(t, arquivo)
	assertFile(t, 7, arquivo.Name(), true)
}
