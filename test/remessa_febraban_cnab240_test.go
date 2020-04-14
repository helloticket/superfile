package test

import (
	"strings"
	"testing"
	"time"

	"github.com/helderfarias/cnab-go"
	"github.com/helderfarias/cnab-go/file"
	"github.com/helderfarias/cnab-go/layout/febraban"
	"github.com/stretchr/testify/assert"
)

func TestRemessaCnab240Pagamentos(t *testing.T) {
	source := strings.NewReader(febraban.CNAB240Pagamentos)
	layout, err := cnab.NewLayout("240", source)
	remessa := cnab.NewRemessa(layout)

	remessa.Header["codigo_banco"] = 341
	remessa.Header["tipo_inscricao_empresa"] = 1
	remessa.Header["numero_inscricao_empresa"] = 12234567000186
	remessa.Header["codigo_convenio_banco"] = "12234567000186"
	remessa.Header["agencia_mantenedora_conta"] = 341
	remessa.Header["digito_verificador_agencia"] = "12"
	remessa.Header["numero_conta_corrente"] = 122345670001
	remessa.Header["digito_verificador_conta"] = "23"
	remessa.Header["digito_verificador_agencia_conta"] = "24"
	remessa.Header["nome_empresa"] = "NOME DA EMPRESA"
	remessa.Header["nome_banco"] = "BANCO BRASIL"
	remessa.Header["codigo_remessa_retorno"] = "1"
	remessa.Header["data_geracao_arquivo"] = time.Date(2017, 5, 10, 10, 2, 1, 4, time.UTC)
	remessa.Header["hora_geracao_arquivo"] = time.Date(2017, 5, 10, 10, 2, 1, 4, time.UTC)
	remessa.Header["numero_sequencial_arquivo"] = 1
	remessa.Header["densidade_gravacao_arquivo"] = 12345
	remessa.Header["reservado_empresa_01"] = "PRODUCAO"

	lote := remessa.NovoLote()

	lote.Header["codigo_banco"] = 341
	lote.Header["lote_servico"] = lote.Sequencial
	lote.Header["tipo_servico"] = "01"
	lote.Header["tipo_inscricao_empresa"] = 1
	lote.Header["versao_layout_lote"] = "030"
	lote.Header["forma_lancamento"] = 12
	lote.Header["tipo_inscricao"] = 2
	lote.Header["inscricao_empresa"] = 12234567000186
	lote.Header["numero_inscricao_empresa"] = 12234567000186
	lote.Header["agencia"] = 2932
	lote.Header["conta"] = "24992"
	lote.Header["dac"] = 9
	lote.Header["numero"] = 12344
	lote.Header["nome_empresa"] = "NOME EMPRESA LOTE"
	lote.Header["numero_sequencial_arquivo_retorno"] = 1
	lote.Header["data_gravacao"] = time.Date(2017, 5, 10, 0, 0, 0, 0, time.UTC)
	lote.Header["data_credito"] = time.Date(2017, 5, 10, 0, 0, 0, 0, time.UTC)
	lote.Header["agencia_mantenedora_conta"] = 12345
	lote.Header["numero_conta_corrente"] = 123456789122
	lote.Header["cep"] = 58040
	lote.Header["indicativo_forma_pagamento_servico"] = 12

	detalhe := lote.NovoDetalhe()

	detalhe["segmento_a"]["lote_servico"] = 1
	detalhe["segmento_a"]["codigo_banco"] = 341
	detalhe["segmento_a"]["tipo_registro"] = 3
	detalhe["segmento_a"]["numero_sequencial_registro_lote"] = 1
	detalhe["segmento_a"]["codigo_segmento_registro_detalhe"] = "A"
	detalhe["segmento_a"]["tipo_movimento"] = 0
	detalhe["segmento_a"]["codigo_instrucao_movimento"] = "00"
	detalhe["segmento_a"]["codigo_camara_centralizadora"] = "700"
	detalhe["segmento_a"]["codigo_banco_favorecido"] = 341
	detalhe["segmento_a"]["agencia_mantenedora_conta_favorecido"] = 3158
	detalhe["segmento_a"]["digito_verificador_agencia"] = ""
	detalhe["segmento_a"]["numero_conta_corrente"] = 38094
	detalhe["segmento_a"]["digito_verificador_conta"] = "3"
	detalhe["segmento_a"]["digito_verificador_agencia_conta"] = ""
	detalhe["segmento_a"]["nome_favorecido"] = "GLAUBER PORTELLA"
	detalhe["segmento_a"]["numero_documento_atribuido_empresa"] = "12345"
	detalhe["segmento_a"]["data_pagamento"] = time.Date(2017, 5, 10, 0, 0, 0, 0, time.UTC)
	detalhe["segmento_a"]["tipo_moeda"] = "BRL"
	detalhe["segmento_a"]["quantidade_moeda"] = 1
	detalhe["segmento_a"]["valor_pagamento"] = "15000"
	detalhe["segmento_a"]["numero_documento_atribuido_banco"] = "123456"
	detalhe["segmento_a"]["data_real_efetivacao_pagamento"] = time.Date(2017, 5, 10, 0, 0, 0, 0, time.UTC)
	detalhe["segmento_a"]["valor_real_efetivacao_pagamento"] = "15000"
	detalhe["segmento_a"]["outras_informacoes"] = ""
	detalhe["segmento_a"]["complemento_tipo_servico"] = "06"
	detalhe["segmento_a"]["codigo_finalidade_ted"] = "123456"
	detalhe["segmento_a"]["complemento_finalidade_pagamento"] = "0"
	detalhe["segmento_a"]["aviso_favorecido"] = "0"
	detalhe["segmento_a"]["codigos_ocorrencias_retorno"] = "00"

	detalhe["segmento_b"]["codigo_banco"] = 341
	detalhe["segmento_b"]["lote_servico"] = 1
	detalhe["segmento_b"]["tipo_registro"] = 3
	detalhe["segmento_b"]["numero_sequencial_registro_lote"] = 1
	detalhe["segmento_b"]["codigo_segmento_registro_detalhe"] = "B"
	detalhe["segmento_b"]["tipo_inscricao_favorecido"] = 1
	detalhe["segmento_b"]["numero_inscricao_favorecido"] = "05771095613"
	detalhe["segmento_b"]["logradouro"] = "RUA ALVARENGA"
	detalhe["segmento_b"]["numero"] = 40
	detalhe["segmento_b"]["complemento"] = ""
	detalhe["segmento_b"]["bairro"] = "GUARANI"
	detalhe["segmento_b"]["cidade"] = "BELO HORIZONTE"
	detalhe["segmento_b"]["cep"] = "31814"
	detalhe["segmento_b"]["complemento_cep"] = "500"
	detalhe["segmento_b"]["estado"] = "MG"
	detalhe["segmento_b"]["data_vencimento_nominal"] = time.Date(2017, 5, 10, 0, 0, 0, 0, time.UTC)
	detalhe["segmento_b"]["valor_documento_nominal"] = "1500"
	detalhe["segmento_b"]["valor_abatimento"] = "0"
	detalhe["segmento_b"]["valor_desconto"] = "0"
	detalhe["segmento_b"]["valor_mora"] = "0"
	detalhe["segmento_b"]["valor_multa"] = "0"
	detalhe["segmento_b"]["codigo_documento_favorecido"] = "05771095613"
	detalhe["segmento_b"]["aviso_favorecido"] = 0
	detalhe["segmento_b"]["exclusivo_siape_01"] = 0
	detalhe["segmento_b"]["codigo_ispb"] = 60701190

	detalhe["segmento_c"]["codigo_banco"] = 341
	detalhe["segmento_c"]["lote_servico"] = 1
	detalhe["segmento_c"]["tipo_registro"] = 3
	detalhe["segmento_c"]["numero_sequencial_registro_lote"] = 1
	detalhe["segmento_c"]["codigo_segmento_registro_detalhe"] = "C"
	detalhe["segmento_c"]["valor_ir"] = "0"
	detalhe["segmento_c"]["valor_iss"] = "0"
	detalhe["segmento_c"]["valor_iof"] = "0"
	detalhe["segmento_c"]["valor_outras_deducoes"] = "0"
	detalhe["segmento_c"]["valor_outros_acrescimos"] = "0"
	detalhe["segmento_c"]["agencia_favorecido"] = 3158
	detalhe["segmento_c"]["digito_verificador_agencia"] = ""
	detalhe["segmento_c"]["numero_conta_corrente"] = 38094
	detalhe["segmento_c"]["digito_verificador_conta"] = "3"
	detalhe["segmento_c"]["digito_verificador_agencia_conta"] = "3"
	detalhe["segmento_c"]["valor_inss"] = "0"

	lote.InserirDetalhe(detalhe)

	lote.Trailer["codigo_banco"] = 341
	lote.Trailer["lote_servico"] = lote.Sequencial
	lote.Trailer["exclusivo_febraban_01"] = ""
	lote.Trailer["quantidade_registros_lote"] = 1
	lote.Trailer["somatoria_valores"] = "10000"
	lote.Trailer["somatoria_quantidade_moedas"] = "1"
	lote.Trailer["numero_aviso_debito"] = "0"

	remessa.InserirLote(lote)

	remessa.Trailer["codigo_banco"] = 341
	remessa.Trailer["lote_servico"] = 9999
	remessa.Trailer["tipo_registro"] = 9
	remessa.Trailer["quantidade_lotes_arquivo"] = remessa.TotalLotes()
	remessa.Trailer["quantidade_registros_arquivo"] = remessa.TotalRegistros()
	remessa.Trailer["quantidade_contas_conciliacao_lotes"] = 1

	remessaFile := file.NewRemessaFile(remessa, "febraban-pagamento-cnab240.rem")

	arquivo := remessaFile.Write()

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, remessa)
	assert.NotNil(t, arquivo)
	assertFile(t, 7, arquivo.Name(), true)
}
