package test

import (
	"strings"
	"testing"
	"time"

	"github.com/helloticket/superfile"
	"github.com/helloticket/superfile/file"
	"github.com/helloticket/superfile/layout/ccsitef"
	"github.com/stretchr/testify/assert"
)

func TestRemessaCCSitefExtrato(t *testing.T) {
	source := strings.NewReader(ccsitef.ExtratoEletronico)
	layout, err := superfile.NewLayout(source)
	remessa := superfile.NewRemessa(layout)

	remessa.Header["data_geracao"] = 00000000
	remessa.Header["hora_geracao"] = 000000
	remessa.Header["id_movimento"] = 000000
	remessa.Header["nome_administradora"] = "0"
	remessa.Header["identificacao_remetente"] = 0000
	remessa.Header["identificacao_destinatario"] = 000000
	remessa.Header["nseq_registro"] = 000000

	lote := remessa.NovoLote()

	lote.Header["lote_servico"] = lote.Sequencial
	lote.Header["data_movimento"] = time.Date(2017, 5, 10, 0, 0, 0, 0, time.UTC)
	lote.Header["identificacao_moeda"] = "RE"
	lote.Header["nseq"] = 1

	detalhe1 := lote.NovoDetalhe()
	detalhe1["segmento_cv"]["identificacao_loja"] = "1"
	detalhe1["segmento_cv"]["nsu_host_transacao"] = 100000000000
	detalhe1["segmento_cv"]["data_transacao"] = 00000000
	detalhe1["segmento_cv"]["hora_transacao"] = 000000
	detalhe1["segmento_cv"]["tipo_lancamento"] = 0
	detalhe1["segmento_cv"]["data_lancamento"] = 0000000
	detalhe1["segmento_cv"]["tipo_produto"] = "0"
	detalhe1["segmento_cv"]["meio_captura"] = 0
	detalhe1["segmento_cv"]["valor_bruno_venda"] = 00000000000
	detalhe1["segmento_cv"]["valor_desconto"] = 00000000000
	detalhe1["segmento_cv"]["valor_liquida_venda"] = 00000000000
	detalhe1["segmento_cv"]["numero_cartao"] = "0"
	detalhe1["segmento_cv"]["numero_parcela"] = 00
	detalhe1["segmento_cv"]["numero_total_parcelas"] = 00
	detalhe1["segmento_cv"]["nsu_host_parcela"] = "0"
	detalhe1["segmento_cv"]["valor_bruto_parcela"] = 00000000000
	detalhe1["segmento_cv"]["valor_desconto_parcela"] = 00000000000
	detalhe1["segmento_cv"]["valor_liquido_parcela"] = 00000000000
	detalhe1["segmento_cv"]["banco"] = 000
	detalhe1["segmento_cv"]["agencia"] = 000000
	detalhe1["segmento_cv"]["conta"] = "0"
	detalhe1["segmento_cv"]["codigo_autorizacao"] = 000000000000
	detalhe1["segmento_cv"]["nseq"] = 000000
	lote.InserirDetalhe(detalhe1)

	detalhe2 := lote.NovoDetalhe()
	detalhe2["segmento_cv"]["identificacao_loja"] = "2"
	detalhe2["segmento_cv"]["nsu_host_transacao"] = 200000000000
	detalhe2["segmento_cv"]["data_transacao"] = 00000000
	detalhe2["segmento_cv"]["hora_transacao"] = 000000
	detalhe2["segmento_cv"]["tipo_lancamento"] = 0
	detalhe2["segmento_cv"]["data_lancamento"] = 0000000
	detalhe2["segmento_cv"]["tipo_produto"] = "0"
	detalhe2["segmento_cv"]["meio_captura"] = 0
	detalhe2["segmento_cv"]["valor_bruno_venda"] = 00000000000
	detalhe2["segmento_cv"]["valor_desconto"] = 00000000000
	detalhe2["segmento_cv"]["valor_liquida_venda"] = 00000000000
	detalhe2["segmento_cv"]["numero_cartao"] = "0"
	detalhe2["segmento_cv"]["numero_parcela"] = 00
	detalhe2["segmento_cv"]["numero_total_parcelas"] = 00
	detalhe2["segmento_cv"]["nsu_host_parcela"] = "0"
	detalhe2["segmento_cv"]["valor_bruto_parcela"] = 00000000000
	detalhe2["segmento_cv"]["valor_desconto_parcela"] = 00000000000
	detalhe2["segmento_cv"]["valor_liquido_parcela"] = 00000000000
	detalhe2["segmento_cv"]["banco"] = 000
	detalhe2["segmento_cv"]["agencia"] = 000000
	detalhe2["segmento_cv"]["conta"] = "0"
	detalhe2["segmento_cv"]["codigo_autorizacao"] = 000000000000
	detalhe2["segmento_cv"]["nseq"] = 000000
	lote.InserirDetalhe(detalhe2)

	lote.Trailer["lote_servico"] = lote.Sequencial
	lote.Trailer["quantidade_registros_lote"] = 1

	remessa.InserirLote(lote)
	remessa.Trailer["total_registros_transacao"] = 000000
	remessa.Trailer["total_valores_creditos"] = 000000
	remessa.Trailer["nseq"] = 000000

	remessaFile := file.NewRemessaFile(remessa, "ccsitef-vendas.rem")

	arquivo := remessaFile.Write()

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, remessa)
	assert.NotNil(t, arquivo)
	assertFile(t, 6, arquivo.Name(), true)
}
