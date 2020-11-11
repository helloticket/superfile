package test

import (
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"github.com/helloticket/superfile"
	"github.com/helloticket/superfile/file"
	"github.com/helloticket/superfile/layout/ccsitef"
	"github.com/stretchr/testify/assert"
)

func TestRemessaCCSitefExtratoBoaVista(t *testing.T) {
	source := strings.NewReader(ccsitef.ExtratoEletronico)
	layout, err := superfile.NewLayout(source)
	remessa := superfile.NewRemessa(layout)

	remessa.Header["data_geracao"] = 20200522
	remessa.Header["hora_geracao"] = 104002
	remessa.Header["id_movimento"] = 000000
	remessa.Header["nome_administradora"] = "Hello Ticket"
	remessa.Header["identificacao_remetente"] = 0001
	remessa.Header["identificacao_destinatario"] = 000001
	remessa.Header["nseq_registro"] = 000001

	lote := remessa.NovoLote()

	lote.Header["lote_servico"] = lote.Sequencial
	lote.Header["data_movimento"] = time.Date(2020, 5, 21, 0, 0, 0, 0, time.UTC)
	lote.Header["identificacao_moeda"] = "RE"
	lote.Header["nseq"] = 1

	detalhe := lote.NovoDetalhe()

	detalhe["segmento_cv"]["identificacao_loja"] = "12345678000195"
	detalhe["segmento_cv"]["nsu_host_transacao"] = 192168001017
	detalhe["segmento_cv"]["data_transacao"] = 20200521
	detalhe["segmento_cv"]["hora_transacao"] = 132418
	detalhe["segmento_cv"]["tipo_lancamento"] = 1
	detalhe["segmento_cv"]["data_lancamento"] = 20200617
	detalhe["segmento_cv"]["tipo_produto"] = "D"
	detalhe["segmento_cv"]["meio_captura"] = 5
	detalhe["segmento_cv"]["valor_bruno_venda"] = 10103
	detalhe["segmento_cv"]["valor_desconto"] = 505
	detalhe["segmento_cv"]["valor_liquida_venda"] = 9598
	detalhe["segmento_cv"]["numero_cartao"] = "0"
	detalhe["segmento_cv"]["numero_parcela"] = 00
	detalhe["segmento_cv"]["numero_total_parcelas"] = 00
	detalhe["segmento_cv"]["nsu_host_parcela"] = "192168001017"
	detalhe["segmento_cv"]["valor_bruto_parcela"] = 00000000000
	detalhe["segmento_cv"]["valor_desconto_parcela"] = 00000000000
	detalhe["segmento_cv"]["valor_liquido_parcela"] = 00000000000
	detalhe["segmento_cv"]["banco"] = 001
	detalhe["segmento_cv"]["agencia"] = 12345
	detalhe["segmento_cv"]["conta"] = "9876"
	detalhe["segmento_cv"]["codigo_autorizacao"] = 000000000000
	detalhe["segmento_cv"]["nseq"] = 2
	lote.InserirDetalhe(detalhe)

	detalhe1 := lote.NovoDetalhe()
	detalhe1["segmento_cv"]["identificacao_loja"] = "12345678000195"
	detalhe1["segmento_cv"]["nsu_host_transacao"] = 192168001017
	detalhe1["segmento_cv"]["data_transacao"] = 20200521
	detalhe1["segmento_cv"]["hora_transacao"] = 141403
	detalhe1["segmento_cv"]["tipo_lancamento"] = 1
	detalhe1["segmento_cv"]["data_lancamento"] = 20200617
	detalhe1["segmento_cv"]["tipo_produto"] = "D"
	detalhe1["segmento_cv"]["meio_captura"] = 5
	detalhe1["segmento_cv"]["valor_bruno_venda"] = 3000
	detalhe1["segmento_cv"]["valor_desconto"] = 150
	detalhe1["segmento_cv"]["valor_liquida_venda"] = 2850
	detalhe1["segmento_cv"]["numero_cartao"] = "0"
	detalhe1["segmento_cv"]["numero_parcela"] = 00
	detalhe1["segmento_cv"]["numero_total_parcelas"] = 00
	detalhe1["segmento_cv"]["nsu_host_parcela"] = "192168001017"
	detalhe1["segmento_cv"]["valor_bruto_parcela"] = 00000000000
	detalhe1["segmento_cv"]["valor_desconto_parcela"] = 00000000000
	detalhe1["segmento_cv"]["valor_liquido_parcela"] = 00000000000
	detalhe1["segmento_cv"]["banco"] = 001
	detalhe1["segmento_cv"]["agencia"] = 12345
	detalhe1["segmento_cv"]["conta"] = "9876"
	detalhe1["segmento_cv"]["codigo_autorizacao"] = 000000000000
	detalhe1["segmento_cv"]["nseq"] = 3
	lote.InserirDetalhe(detalhe1)

	detalhe2 := lote.NovoDetalhe()
	detalhe2["segmento_cv"]["identificacao_loja"] = "12345678000195"
	detalhe2["segmento_cv"]["nsu_host_transacao"] = 192168001017
	detalhe2["segmento_cv"]["data_transacao"] = 20200521
	detalhe2["segmento_cv"]["hora_transacao"] = 173524
	detalhe2["segmento_cv"]["tipo_lancamento"] = 1
	detalhe2["segmento_cv"]["data_lancamento"] = 20200617
	detalhe2["segmento_cv"]["tipo_produto"] = "D"
	detalhe2["segmento_cv"]["meio_captura"] = 5
	detalhe2["segmento_cv"]["valor_bruno_venda"] = 2000
	detalhe2["segmento_cv"]["valor_desconto"] = 100
	detalhe2["segmento_cv"]["valor_liquida_venda"] = 1900
	detalhe2["segmento_cv"]["numero_cartao"] = "0"
	detalhe2["segmento_cv"]["numero_parcela"] = 00
	detalhe2["segmento_cv"]["numero_total_parcelas"] = 00
	detalhe2["segmento_cv"]["nsu_host_parcela"] = "192168001017"
	detalhe2["segmento_cv"]["valor_bruto_parcela"] = 00000000000
	detalhe2["segmento_cv"]["valor_desconto_parcela"] = 00000000000
	detalhe2["segmento_cv"]["valor_liquido_parcela"] = 00000000000
	detalhe2["segmento_cv"]["banco"] = 001
	detalhe2["segmento_cv"]["agencia"] = 12345
	detalhe2["segmento_cv"]["conta"] = "9876"
	detalhe2["segmento_cv"]["codigo_autorizacao"] = 000000000000
	detalhe2["segmento_cv"]["nseq"] = 4
	lote.InserirDetalhe(detalhe2)

	detalhe3 := lote.NovoDetalhe()
	detalhe3["segmento_cv"]["identificacao_loja"] = "12345678000195"
	detalhe3["segmento_cv"]["nsu_host_transacao"] = 192168001017
	detalhe3["segmento_cv"]["data_transacao"] = 20200521
	detalhe3["segmento_cv"]["hora_transacao"] = 184319
	detalhe3["segmento_cv"]["tipo_lancamento"] = 1
	detalhe3["segmento_cv"]["data_lancamento"] = 20200617
	detalhe3["segmento_cv"]["tipo_produto"] = "D"
	detalhe3["segmento_cv"]["meio_captura"] = 5
	detalhe3["segmento_cv"]["valor_bruno_venda"] = 6322
	detalhe3["segmento_cv"]["valor_desconto"] = 316
	detalhe3["segmento_cv"]["valor_liquida_venda"] = 6006
	detalhe3["segmento_cv"]["numero_cartao"] = "0"
	detalhe3["segmento_cv"]["numero_parcela"] = 00
	detalhe3["segmento_cv"]["numero_total_parcelas"] = 00
	detalhe3["segmento_cv"]["nsu_host_parcela"] = "192168001017"
	detalhe3["segmento_cv"]["valor_bruto_parcela"] = 00000000000
	detalhe3["segmento_cv"]["valor_desconto_parcela"] = 00000000000
	detalhe3["segmento_cv"]["valor_liquido_parcela"] = 00000000000
	detalhe3["segmento_cv"]["banco"] = 001
	detalhe3["segmento_cv"]["agencia"] = 12345
	detalhe3["segmento_cv"]["conta"] = "9876"
	detalhe3["segmento_cv"]["codigo_autorizacao"] = 000000000000
	detalhe3["segmento_cv"]["nseq"] = 5

	lote.InserirDetalhe(detalhe3)

	detalhe = lote.NovoDetalhe()
	detalhe["segmento_aj"]["identificacao_loja"] = "12345678000195"
	detalhe["segmento_aj"]["nsu_host_transacao_original"] = 192168001017
	detalhe["segmento_aj"]["data_transacao_original"] = 20200520
	detalhe["segmento_aj"]["numero_parcela"] = 00
	detalhe["segmento_aj"]["nsu_host_transacao"] = 192168001017
	detalhe["segmento_aj"]["data_transacao"] = 20200521
	detalhe["segmento_aj"]["hora_transacao"] = 074533
	detalhe["segmento_aj"]["tipo_lancamento"] = 1
	detalhe["segmento_aj"]["data_lancamento"] = 20200521
	detalhe["segmento_aj"]["meio_captura"] = 5
	detalhe["segmento_aj"]["tipo_ajuste"] = 1
	detalhe["segmento_aj"]["codigo_ajuste"] = 000
	detalhe["segmento_aj"]["descricao_motivo_ajuste"] = "0"
	detalhe["segmento_aj"]["valor_bruto_parcela"] = 00000000000
	detalhe["segmento_aj"]["valor_desconto_comissao"] = 00000000000
	detalhe["segmento_aj"]["valor_liquido"] = 3000
	detalhe["segmento_aj"]["banco"] = 001
	detalhe["segmento_aj"]["agencia"] = 12345
	detalhe["segmento_aj"]["conta"] = "9876"
	detalhe["segmento_aj"]["nseq"] = 6
	lote.InserirDetalhe(detalhe)

	detalhe = lote.NovoDetalhe()
	detalhe["segmento_cc"]["identificacao_loja"] = "12345678000195"
	detalhe["segmento_cc"]["nsu_host_transacao_original"] = 192168001017
	detalhe["segmento_cc"]["data_transacao_original"] = 20200520
	detalhe["segmento_cc"]["numero_parcela"] = 00
	detalhe["segmento_cc"]["nsu_host_transacao"] = 192168001017
	detalhe["segmento_cc"]["data_transacao"] = 20200521
	detalhe["segmento_cc"]["hora_transacao"] = 074533
	detalhe["segmento_cc"]["meio_captura"] = 5
	detalhe["segmento_cc"]["nseq"] = 7
	lote.InserirDetalhe(detalhe)

	lote.Trailer["lote_servico"] = lote.Sequencial
	lote.Trailer["quantidade_registros_lote"] = 6

	remessa.InserirLote(lote)
	remessa.Trailer["total_registros_transacao"] = 4
	remessa.Trailer["total_valores_creditos"] = 21425
	remessa.Trailer["nseq"] = 7

	remessaFile := file.NewRemessaFile(remessa, "ccsitef-vendas.rem")

	arquivo := remessaFile.Write()

	b, err := ioutil.ReadFile(arquivo.Name())
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("/Users/vinicioalves/teste/ccsitef-vendas-homologacao_2.rem", b, 0644)

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, remessa)
	assert.NotNil(t, arquivo)
	// assertFile(t, 7, arquivo.Name(), true)
}
