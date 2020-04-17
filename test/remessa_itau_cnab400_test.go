package test

import (
	"strings"
	"testing"

	"github.com/helderfarias/cnab-go"
	"github.com/helderfarias/cnab-go/file"
	"github.com/helderfarias/cnab-go/layout/itau"
	"github.com/stretchr/testify/assert"
)

func TestRemessaItauCnab400Cobranca(t *testing.T) {
	source := strings.NewReader(itau.CNAB400Cobranca)
	layout, err := cnab.NewLayout("400", source)

	remessa := cnab.NewRemessa(layout)
	remessa.Header["agencia"] = 0000
	remessa.Header["conta"] = 00000
	remessa.Header["dac"] = 0
	remessa.Header["nome_empresa"] = "0"
	remessa.Header["data_geracao"] = 000000

	lote := remessa.NovoLote()
	remessa.InserirLote(lote)

	detalhe := lote.NovoDetalhe()
	lote.InserirDetalhe(detalhe)
	detalhe["segmento_1"]["codigo_inscricao"] = 00
	detalhe["segmento_1"]["numero_inscricao"] = 00000000000000
	detalhe["segmento_1"]["agencia"] = 0000
	detalhe["segmento_1"]["conta"] = 00000
	detalhe["segmento_1"]["dac"] = 0
	detalhe["segmento_1"]["instrucao_alegacao"] = 0000
	detalhe["segmento_1"]["uso_empresa"] = "0"
	detalhe["segmento_1"]["nosso_numero"] = 00000000
	detalhe["segmento_1"]["quantidade_moeda"] = 0000000000000
	detalhe["segmento_1"]["numero_carteira"] = 000
	detalhe["segmento_1"]["uso_banco"] = "0"
	detalhe["segmento_1"]["carteira"] = "0"
	detalhe["segmento_1"]["codigo_ocorrencia"] = 00
	detalhe["segmento_1"]["numero_documento"] = "0"
	detalhe["segmento_1"]["vencimento"] = 000000
	detalhe["segmento_1"]["valor"] = 0000000000000
	detalhe["segmento_1"]["agencia_cobradora"] = 00000
	detalhe["segmento_1"]["especie"] = "0"
	detalhe["segmento_1"]["data_emissao"] = 000000
	detalhe["segmento_1"]["instrucao_1"] = "0"
	detalhe["segmento_1"]["instrucao_2"] = "0"
	detalhe["segmento_1"]["juros_1_dia"] = 0000000000000
	detalhe["segmento_1"]["desconto_ate"] = 000000
	detalhe["segmento_1"]["valor_desconto"] = 0000000000000
	detalhe["segmento_1"]["valor_iof"] = 0000000000000
	detalhe["segmento_1"]["abatimento"] = 0000000000000
	detalhe["segmento_1"]["codigo_inscricao_pagador"] = 00
	detalhe["segmento_1"]["numero_inscricao_pagador"] = 00000000000000
	detalhe["segmento_1"]["nome_pagador"] = "0"
	detalhe["segmento_1"]["logradouro"] = "0"
	detalhe["segmento_1"]["bairro"] = "0"
	detalhe["segmento_1"]["cep"] = 00000000
	detalhe["segmento_1"]["cidade"] = "0"
	detalhe["segmento_1"]["estado"] = "0"
	detalhe["segmento_1"]["sacador_avalista"] = "0"
	detalhe["segmento_1"]["data_mora"] = 000000
	detalhe["segmento_1"]["prazo"] = 00
	detalhe["segmento_1"]["numero_sequencial_registro"] = 000000

	remessa.Trailer["numero_sequencial_registro"] = 000000

	remessaFile := file.NewRemessaFile(remessa, "itau-cobranca-cnab400.rem")

	arquivo := remessaFile.Write()

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, remessa)
	assert.NotNil(t, arquivo)
	assertFile(t, 3, arquivo.Name(), false)
}
