package test

import (
	"strings"
	"testing"

	"github.com/helderfarias/cnab-go"
	"github.com/helderfarias/cnab-go/file"
	"github.com/helderfarias/cnab-go/layout/itau"
	"github.com/stretchr/testify/assert"
)

func TestRemessaItauCnab240Pagamento(t *testing.T) {
	source := strings.NewReader(itau.CNAB240Pagamentos)
	layout, err := cnab.NewLayout("240", source)
	remessa := cnab.NewRemessa(layout)

	remessa.Header["tipo_inscricao"] = 0
	remessa.Header["inscricao_numero"] = 00000000000000
	remessa.Header["agencia_debito"] = 00000
	remessa.Header["conta_debito"] = 000000000000
	remessa.Header["dac_debito"] = 0
	remessa.Header["nome_empresa"] = "0"
	remessa.Header["nome_banco"] = "0"
	remessa.Header["data_geracao"] = 00000000
	remessa.Header["hora_geracao"] = 000000

	lote := remessa.NovoLote()

	lote.Header["codigo_lote"] = 0000
	lote.Header["tipo_pagamento"] = 00
	lote.Header["forma_pagamento"] = 00
	lote.Header["tipo_inscricao_debito"] = 0
	lote.Header["inscricao_numero"] = 00000000000000
	lote.Header["identificacao_lancamento"] = "0"
	lote.Header["agencia_debito"] = 00000
	lote.Header["conta_debito"] = 000000000000
	lote.Header["dac_debito"] = 0
	lote.Header["nome_empresa"] = "0"
	lote.Header["finalidade_lote"] = "0"
	lote.Header["historico_cc_debito"] = "0"
	lote.Header["endereco_empresa"] = "0"
	lote.Header["numero"] = 00000
	lote.Header["cidade"] = "0"
	lote.Header["cep"] = 00000000
	lote.Header["estado"] = "0"
	lote.Header["codigo_ocorrencias"] = "0"

	detalhe := lote.NovoDetalhe()

	detalhe["segmento_a"]["codigo_lote"] = 0000
	detalhe["segmento_a"]["numero_registro"] = 00000
	detalhe["segmento_a"]["tipo_movimento"] = 000
	detalhe["segmento_a"]["codigo_camara_centralizadora"] = 000
	detalhe["segmento_a"]["codigo_banco_favorecido"] = 000
	detalhe["segmento_a"]["agencia_favorecido"] = "0"
	detalhe["segmento_a"]["nome_favorecido"] = "0"
	detalhe["segmento_a"]["numero_doc"] = "0"
	detalhe["segmento_a"]["data_pagto"] = 00000000
	detalhe["segmento_a"]["codigo_ispb"] = 00000000
	detalhe["segmento_a"]["valor_pagto"] = 000000000000000
	detalhe["segmento_a"]["nosso_numero"] = "0"
	detalhe["segmento_a"]["data_efetiva"] = 00000000
	detalhe["segmento_a"]["valor_efetivo"] = 000000000000000
	detalhe["segmento_a"]["finalidade"] = "0"
	detalhe["segmento_a"]["num_documento"] = 000000
	detalhe["segmento_a"]["num_inscricao_favorecido"] = 00000000000000
	detalhe["segmento_a"]["finalidade_doc_status_funcionario"] = "0"
	detalhe["segmento_a"]["finalidade_ted"] = "0"
	detalhe["segmento_a"]["aviso"] = "0"
	detalhe["segmento_a"]["codigo_ocorrencias"] = "0"

	detalhe["segmento_b"]["codigo_lote"] = 0000
	detalhe["segmento_b"]["numero_registro"] = 00000
	detalhe["segmento_b"]["tipo_inscricao_favorecido"] = 0
	detalhe["segmento_b"]["inscricao_favorecido"] = 00000000000000
	detalhe["segmento_b"]["endereco"] = "0"
	detalhe["segmento_b"]["numero"] = 00000
	detalhe["segmento_b"]["bairro"] = "0"
	detalhe["segmento_b"]["cidade"] = "0"
	detalhe["segmento_b"]["cep"] = 00000000
	detalhe["segmento_b"]["estado"] = "0"
	detalhe["segmento_b"]["email"] = "0"
	detalhe["segmento_b"]["codigo_ocorrencias"] = "0"

	detalhe["segmento_c"]["codigo_lote"] = 0000
	detalhe["segmento_c"]["numero_registro"] = 00000
	detalhe["segmento_c"]["valor_contrib_lucro_liquido"] = 000000000000000
	detalhe["segmento_c"]["vencimento"] = "0"
	detalhe["segmento_c"]["valor_documento"] = 000000000000000
	detalhe["segmento_c"]["valor_pis"] = 000000000000000
	detalhe["segmento_c"]["valor_ir"] = 000000000000000
	detalhe["segmento_c"]["valor_iss"] = 000000000000000
	detalhe["segmento_c"]["valor_cofins"] = 000000000000000
	detalhe["segmento_c"]["desconto"] = 000000000000000
	detalhe["segmento_c"]["abatimento"] = 000000000000000
	detalhe["segmento_c"]["outras_deducoes"] = 000000000000000
	detalhe["segmento_c"]["mora"] = 000000000000000
	detalhe["segmento_c"]["multa"] = 000000000000000
	detalhe["segmento_c"]["outros_acrescimos"] = 000000000000000
	detalhe["segmento_c"]["numero_fatura_documento"] = "0"

	detalhe["segmento_d"]["codigo_lote"] = 0000
	detalhe["segmento_d"]["numero_registro"] = 00000
	detalhe["segmento_d"]["mes_ano_pagamento"] = 000000
	detalhe["segmento_d"]["centro_custo"] = "0"
	detalhe["segmento_d"]["codigo_funcionario"] = "0"
	detalhe["segmento_d"]["cargo"] = "0"
	detalhe["segmento_d"]["ferias_de"] = 00000000
	detalhe["segmento_d"]["ferias_ate"] = 00000000
	detalhe["segmento_d"]["dependentes_ir"] = 00
	detalhe["segmento_d"]["dependentes_sf"] = 00
	detalhe["segmento_d"]["horas_semanais"] = 00
	detalhe["segmento_d"]["salario_contribuicao"] = 000000000000000
	detalhe["segmento_d"]["fgts"] = 000000000000000
	detalhe["segmento_d"]["valor_creditos"] = 000000000000000
	detalhe["segmento_d"]["valor_debito"] = 000000000000000
	detalhe["segmento_d"]["valor_liquido"] = 000000000000000
	detalhe["segmento_d"]["valor_fixo_base"] = 000000000000000
	detalhe["segmento_d"]["base_calculo_irrf"] = 000000000000000
	detalhe["segmento_d"]["base_calculo_fgts"] = 000000000000000
	detalhe["segmento_d"]["disponibilizacao"] = "0"
	detalhe["segmento_d"]["codigo_ocorrencia"] = "0"

	detalhe["segmento_e"]["codigo_lote"] = 0000
	detalhe["segmento_e"]["numero_registro"] = 00000
	detalhe["segmento_e"]["tipo_movimento"] = "0"
	detalhe["segmento_e"]["informacoes_complementares"] = "0"
	detalhe["segmento_e"]["codigo_ocorrencia"] = "0"

	detalhe["segmento_f"]["codigo_lote"] = 0000
	detalhe["segmento_f"]["numero_registro"] = 00000
	detalhe["segmento_f"]["mensagem"] = "0"
	detalhe["segmento_f"]["codigo_ocorrencia"] = "0"

	detalhe["segmento_z"]["codigo_lote"] = 0000
	detalhe["segmento_z"]["numero_registro"] = 00000
	detalhe["segmento_z"]["autenticacao"] = "0"
	detalhe["segmento_z"]["numero_doc"] = "0"
	detalhe["segmento_z"]["nosso_numero"] = "0"

	lote.InserirDetalhe(detalhe)

	lote.Trailer["codigo_lote"] = 0000
	lote.Trailer["total_registros_lote"] = 000000
	lote.Trailer["total_valor_pagtos"] = 0
	lote.Trailer["codigos_ocorrencias"] = "0"

	remessa.InserirLote(lote)

	remessa.Trailer["total_lotes_arquivo"] = remessa.TotalLotes()
	remessa.Trailer["total_registros"] = remessa.TotalRegistros()

	remessaFile := file.NewRemessaFile(remessa, "itau-pagamentos-cnab240.rem")

	arquivo := remessaFile.Write()

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, remessa)
	assert.NotNil(t, arquivo)
	assertFile(t, 11, arquivo.Name(), false)
}

func TestRemessaItauCnab240Cobranca(t *testing.T) {
	source := strings.NewReader(itau.CNAB240Cobranca)
	layout, err := cnab.NewLayout("240", source)
	remessa := cnab.NewRemessa(layout)

	remessa.Header["tipo_inscricao"] = 0
	remessa.Header["inscricao_numero"] = 00000000000000
	remessa.Header["agencia"] = 0000
	remessa.Header["conta"] = 00000
	remessa.Header["dac"] = 0
	remessa.Header["nome_empresa"] = "0"
	remessa.Header["data_geracao"] = 00000000
	remessa.Header["hora_geracao"] = 000000
	remessa.Header["numero_sequencial_arquivo_retorno"] = 000000

	lote := remessa.NovoLote()

	lote.Header["lote_servico"] = 0000
	lote.Header["tipo_inscricao"] = 0
	lote.Header["inscricao_empresa"] = 000000000000000
	lote.Header["agencia"] = 0000
	lote.Header["conta"] = 00000
	lote.Header["dac"] = 0
	lote.Header["nome_empresa"] = "0"
	lote.Header["numero_sequencial_arquivo_retorno"] = 00000000
	lote.Header["data_gravacao"] = 00000000
	lote.Header["data_credito"] = 00000000

	detalhe := lote.NovoDetalhe()

	detalhe["segmento_p"]["lote_servico"] = 0000
	detalhe["segmento_p"]["numero_sequencial_registro_lote"] = 00000
	detalhe["segmento_p"]["codigo_ocorrencia"] = 00
	detalhe["segmento_p"]["agencia"] = 0000
	detalhe["segmento_p"]["conta"] = 00000
	detalhe["segmento_p"]["dac"] = 0
	detalhe["segmento_p"]["carteira"] = 000
	detalhe["segmento_p"]["nosso_numero"] = 00000000
	detalhe["segmento_p"]["dac_nosso_numero"] = 0
	detalhe["segmento_p"]["numero_documento"] = "0"
	detalhe["segmento_p"]["vencimento"] = 00000000
	detalhe["segmento_p"]["valor_titulo"] = 000000000000000
	detalhe["segmento_p"]["dac_agencia_cobradora"] = 0
	detalhe["segmento_p"]["especie"] = 00
	detalhe["segmento_p"]["data_emissao"] = 00000000
	detalhe["segmento_p"]["data_juros_mora"] = 00000000
	detalhe["segmento_p"]["juros_1_dia"] = 000000000000000
	detalhe["segmento_p"]["data_1o_desconto"] = 00000000
	detalhe["segmento_p"]["valor_1o_desconto"] = 000000000000000
	detalhe["segmento_p"]["valor_iof"] = 000000000000000
	detalhe["segmento_p"]["valor_abatimento"] = 000000000000000
	detalhe["segmento_p"]["identificacao_titulo_empresa"] = "0"
	detalhe["segmento_p"]["codigo_negativacao_protesto"] = 0
	detalhe["segmento_p"]["prazo_negativacao_protesto"] = 00
	detalhe["segmento_p"]["codigo_baixa"] = 0
	detalhe["segmento_p"]["prazo_baixa"] = 00

	detalhe["segmento_q"]["lote_servico"] = 0000
	detalhe["segmento_q"]["numero_sequencial_registro_lote"] = 00000
	detalhe["segmento_q"]["codigo_ocorrencia"] = 00
	detalhe["segmento_q"]["tipo_inscricao"] = 0
	detalhe["segmento_q"]["inscricao_numero"] = 000000000000000
	detalhe["segmento_q"]["nome_pagador"] = "0"
	detalhe["segmento_q"]["logradouro"] = "0"
	detalhe["segmento_q"]["bairro"] = "0"
	detalhe["segmento_q"]["cep"] = 00000
	detalhe["segmento_q"]["sufixo_cep"] = 000
	detalhe["segmento_q"]["cidade"] = "0"
	detalhe["segmento_q"]["uf"] = "0"
	detalhe["segmento_q"]["tipo_inscricao_sacador"] = 0
	detalhe["segmento_q"]["inscricao_sacador"] = 000000000000000
	detalhe["segmento_q"]["nome_sacador"] = "0"

	detalhe["segmento_r"]["codigo_banco"] = 000
	detalhe["segmento_r"]["lote_servico"] = 0000
	detalhe["segmento_r"]["tipo_registro"] = 0
	detalhe["segmento_r"]["numero_sequencial_registro_lote"] = 00000
	detalhe["segmento_r"]["segmento"] = "0"
	detalhe["segmento_r"]["brancos_01"] = "0"
	detalhe["segmento_r"]["codigo_ocorrencia"] = 00
	detalhe["segmento_r"]["zeros_01"] = 0
	detalhe["segmento_r"]["valor_2o_desconto"] = 000000000000000
	detalhe["segmento_r"]["zeros_02"] = 0
	detalhe["segmento_r"]["data_3o_desconto"] = 00000000
	detalhe["segmento_r"]["valor_3o_desconto"] = 000000000000000
	detalhe["segmento_r"]["codigo_multa"] = 0
	detalhe["segmento_r"]["data_multa"] = "0"
	detalhe["segmento_r"]["multa"] = "0"
	detalhe["segmento_r"]["brancos_02"] = "0"
	detalhe["segmento_r"]["informacoes_pagador"] = "0"
	detalhe["segmento_r"]["brancos_03"] = "0"
	detalhe["segmento_r"]["codigo_ocorrencia_pagador"] = 00000000
	detalhe["segmento_r"]["zeros_03"] = 00000000
	detalhe["segmento_r"]["brancos_04"] = "0"
	detalhe["segmento_r"]["zeros_04"] = 000000000000
	detalhe["segmento_r"]["brancos_05"] = "0"
	detalhe["segmento_r"]["zeros_05"] = 0
	detalhe["segmento_r"]["brancos_06"] = "0"

	detalhe["segmento_y"]["lote_servico"] = 0000
	detalhe["segmento_y"]["numero_sequencial_registro_lote"] = 00000
	detalhe["segmento_y"]["tipo_inscricao"] = 0
	detalhe["segmento_y"]["inscricao_numero"] = 000000000000000
	detalhe["segmento_y"]["nome_sacador"] = "0"
	detalhe["segmento_y"]["endereco_sacador"] = "0"
	detalhe["segmento_y"]["bairro_sacador"] = "0"
	detalhe["segmento_y"]["cep_sacador"] = 00000000
	detalhe["segmento_y"]["cidade_sacador"] = "0"
	detalhe["segmento_y"]["uf_sacador"] = "0"

	lote.InserirDetalhe(detalhe)

	lote.Trailer["lote_servico"] = 0000
	lote.Trailer["quantidade_registros_lote"] = 000000
	lote.Trailer["quantidade_cobranca_simples"] = 000000
	lote.Trailer["valor_total_cobranca_simples"] = 00000000000000000
	lote.Trailer["quantidade_cobranca_vinculada"] = 000000
	lote.Trailer["valor_total_cobranca_vinculada"] = 00000000000000000

	remessa.InserirLote(lote)

	remessa.Trailer["total_lotes"] = remessa.TotalLotes()
	remessa.Trailer["total_registros"] = remessa.TotalRegistros()

	remessaFile := file.NewRemessaFile(remessa, "itau-cobranca-cnab240.rem")

	arquivo := remessaFile.Write()

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, remessa)
	assert.NotNil(t, arquivo)
	assertFile(t, 8, arquivo.Name(), true)
}
