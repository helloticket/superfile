package test

import (
	"strings"
	"testing"

	"github.com/helderfarias/cnab-go"
	"github.com/helderfarias/cnab-go/layout/itau"
	"github.com/helderfarias/cnab-go/output"
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

	lote := remessa.NovoLote(1)

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

	remessaFile := output.NewRemessaFile(remessa, "itau-pagamentos-cnab240.rem")

	arquivo := remessaFile.Write()

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, remessa)
	assert.NotNil(t, arquivo)
	assertFile(t, 11, arquivo.Name(), true)
}
