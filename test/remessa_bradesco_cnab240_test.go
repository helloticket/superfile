package test

import (
	"strings"
	"testing"

	"github.com/helloticket/superfile"
	"github.com/helloticket/superfile/file"
	"github.com/helloticket/superfile/layout/bradesco"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestRemessaBradescoCnab240Pagamento(t *testing.T) {
	source := strings.NewReader(bradesco.CNAB240Pagamentos)
	layout, err := superfile.NewLayout(source)
	remessa := superfile.NewRemessa(layout)

	remessa.Header["codigo_banco"] = 237
	remessa.Header["tipo_inscricao_empresa"] = 0
	remessa.Header["numero_inscricao_empresa"] = 00000000000000
	remessa.Header["codigo_convenio_banco"] = "00000000000000000000"
	remessa.Header["agencia_mantenedora_conta"] = 00000
	remessa.Header["digito_Verificador_agencia"] = "0"
	remessa.Header["numero_conta_corrente"] = 000000000000
	remessa.Header["digito_verificador_conta"] = "0"
	remessa.Header["digito_verificador_agencia_conta"] = "0"
	remessa.Header["nome_empresa"] = "Empresa"
	remessa.Header["nome_banco"] = "Bradesco"
	remessa.Header["exclusivo_febraban_cnab_2"] = "00000000000"
	remessa.Header["codigo_remessa_retorno"] = 0

	lote := remessa.NovoLote()

	lote.Header["codigo_banco"] = 237
	lote.Header["lote_Servico"] = 0000
	lote.Header["tipo_servico"] = 00
	lote.Header["forma_lancamento"] = 00
	lote.Header["numero_versao_layout_lote"] = 000
	lote.Header["tipo_inscricao_empresa"] = 0
	lote.Header["numero_inscricao_empresa"] = 00000000000000
	lote.Header["agencia_mantenedora_conta"] = 00000
	lote.Header["digito_verificador_agencia"] = "0"
	lote.Header["numero_conta_corrente"] = 0000000000000
	lote.Header["digito_verificador_conta"] = "0"
	lote.Header["digito_verificador_agencia_conta"] = "0"

	detalhe := lote.NovoDetalhe()

	detalhe["segmento_a"]["codigo_banco"] = 237
	detalhe["segmento_a"]["numero_sequencial_registro"] = lote.Sequencial
	detalhe["segmento_a"]["tipo_movimento"] = 0
	detalhe["segmento_a"]["codigo_instrucao_motimento"] = 00
	detalhe["segmento_a"]["codigo_camara_centralizadora"] = 000
	detalhe["segmento_a"]["codigo_banco_favorecido"] = 000
	detalhe["segmento_a"]["agencia_favorecido"] = 00000
	detalhe["segmento_a"]["digito_verificador_agencia"] = "0"
	detalhe["segmento_a"]["numero_conta_corrente"] = 000000000000
	detalhe["segmento_a"]["digito_verificador_conta"] = "0"
	detalhe["segmento_a"]["digito_verificador_agencia_conta"] = "0"
	detalhe["segmento_a"]["nome_favorecido"] = "Eugenio Costa"
	detalhe["segmento_a"]["numero_doc"] = "000000000000000000000"

	detalhe["segmento_b"]["codigo_banco"] = 237
	detalhe["segmento_b"]["lote_servico"] = 0000
	detalhe["segmento_b"]["numero_sequencial_registro"] = lote.Sequencial
	detalhe["segmento_b"]["tipo_inscricao_favorecido"] = 0
	detalhe["segmento_b"]["numero_inscricao_favorecido"] = 00000000000000
	detalhe["segmento_b"]["endereco"] = "0000000000"
	detalhe["segmento_b"]["numero"] = 00000
	detalhe["segmento_b"]["bairro"] = "0"
	detalhe["segmento_b"]["cidade"] = "0"
	detalhe["segmento_b"]["cep"] = 00000
	detalhe["segmento_b"]["complemento_cep"] = "-000"
	detalhe["segmento_b"]["estado"] = "00"
	detalhe["segmento_b"]["data_vencimento"] = "00000000"
	detalhe["segmento_b"]["valor_documento"] = 0
	detalhe["segmento_b"]["valor_abatimento"] = 0
	detalhe["segmento_b"]["valor_desconto"] = 0
	detalhe["segmento_b"]["valor_mora"] = 0

	detalhe["segmento_c"]["codigo_banco"] = 237
	detalhe["segmento_c"]["lote_servico"] = 0000
	detalhe["segmento_c"]["numero_sequencial_registro"] = lote.Sequencial
	detalhe["segmento_c"]["codigo_segmento_registro_detalhe"] = "0"
	detalhe["segmento_c"]["valor_ir"] = "1000"
	detalhe["segmento_c"]["valor_iss"] = "1000"
	detalhe["segmento_c"]["valor_iof"] = "1000"
	detalhe["segmento_c"]["agencia_mantenedora_conta_favorecido"] = 12457
	detalhe["segmento_c"]["digito_verificador_agencia"] = "N"
	detalhe["segmento_c"]["numero_conta_corrente"] = 457894512457
	detalhe["segmento_c"]["digito_verificador_conta"] = "0"
	detalhe["segmento_c"]["digito_verificador_agencia_conta"] = "0"
	detalhe["segmento_c"]["valor_inss"] = "5000"

	detalhe["segmento_5"]["codigo_banco"] = 237
	detalhe["segmento_5"]["lote_servico"] = 0000
	detalhe["segmento_5"]["numero_sequencial_registro_lote"] = lote.Sequencial
	detalhe["segmento_5"]["codigo_segmento_registro_detalhe"] = ""
	detalhe["segmento_5"]["numero_lista_debito"] = 000000000
	detalhe["segmento_5"]["endereco"] = "0000000000"
	detalhe["segmento_5"]["numero"] = 00000
	detalhe["segmento_5"]["bairro"] = "0"
	detalhe["segmento_5"]["cidade"] = "0"
	detalhe["segmento_5"]["cep"] = 00000
	detalhe["segmento_5"]["complemento_cep"] = "-000"
	detalhe["segmento_5"]["estado"] = "00"
	detalhe["segmento_5"]["data_vencimento"] = "00000000"
	detalhe["segmento_5"]["valor_documento"] = 0
	detalhe["segmento_5"]["valor_abatimento"] = 0
	detalhe["segmento_5"]["valor_desconto"] = 0
	detalhe["segmento_5"]["valor_mora"] = 0

	detalhe["segmento_z"]["codigo_banco"] = 000
	detalhe["segmento_z"]["numero_sequencia_registro"] = 5

	lote.InserirDetalhe(detalhe)

	lote.Trailer["codigo_banco_cempensacao"] = 237
	lote.Trailer["quantidade_registros_lote"] = 000000

	remessa.InserirLote(lote)

	remessa.Trailer["quantidade_lotes_arquivo"] = remessa.TotalLotes()
	remessa.Trailer["quantidade_registros_arquivo"] = remessa.TotalRegistros()

	remessaFile := file.NewRemessaFile(remessa, "bradesco-pagamentos-cnab240.rem")

	arquivo := remessaFile.Write()

	logrus.Info("arquivo", arquivo.Name())
	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, remessa)
	assert.NotNil(t, arquivo)
	assertFile(t, 9, arquivo.Name(), true)
}
