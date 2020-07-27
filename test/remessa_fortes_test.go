package test

import (
	"strings"
	"testing"

	"github.com/helloticket/superfile"
	"github.com/helloticket/superfile/file"
	"github.com/helloticket/superfile/layout/fortes"
	"github.com/stretchr/testify/assert"
)

func TestRemessaFortes(t *testing.T) {
	source := strings.NewReader(fortes.ImportacaoMovimento)
	layout, err := superfile.NewLayout(source)

	remessa := superfile.NewRemessa(layout)
	remessa.Header["sistema_origem"] = "PONTO"
	remessa.Header["codigo_empresa_sistema_destino"] = "0"
	remessa.Header["comentario"] = "0"

	lote := remessa.NovoLote()
	remessa.InserirLote(lote)

	detalhe := lote.NovoDetalhe()
	detalhe["segmento_1"]["codigo_empregado"] = 255
	detalhe["segmento_1"]["codigo_evento"] = 102
	detalhe["segmento_1"]["valor_referencia_calculo_evento"] = 50.25
	detalhe["segmento_1"]["valor_evento"] = 250.00
	lote.InserirDetalhe(detalhe)

	detalhe = lote.NovoDetalhe()
	detalhe["segmento_1"]["codigo_empregado"] = 255
	detalhe["segmento_1"]["codigo_evento"] = 102
	detalhe["segmento_1"]["valor_referencia_calculo_evento"] = 25.25
	detalhe["segmento_1"]["valor_evento"] = 125.00
	lote.InserirDetalhe(detalhe)

	detalhe = lote.NovoDetalhe()
	detalhe["segmento_1"]["codigo_empregado"] = 255
	detalhe["segmento_1"]["codigo_evento"] = 102
	detalhe["segmento_1"]["valor_referencia_calculo_evento"] = 15.25
	detalhe["segmento_1"]["valor_evento"] = 75.00
	lote.InserirDetalhe(detalhe)

	detalhe = lote.NovoDetalhe()
	detalhe["segmento_2"]["codigo_empregado"] = 000000
	detalhe["segmento_2"]["codigo_evento"] = 000
	detalhe["segmento_2"]["valor_evento"] = 000000000000000
	detalhe["segmento_2"]["data"] = 00000000
	lote.InserirDetalhe(detalhe)

	detalhe = lote.NovoDetalhe()
	detalhe["segmento_3"]["codigo_empregado"] = 000000
	detalhe["segmento_3"]["codigo_ocorrencia"] = 000
	detalhe["segmento_3"]["data"] = 00000000
	detalhe["segmento_3"]["referencia"] = 0000
	lote.InserirDetalhe(detalhe)

	detalhe = lote.NovoDetalhe()
	detalhe["segmento_4"]["codigo_empregado"] = 000000
	detalhe["segmento_4"]["valor_novo_salario"] = 000000000000000
	detalhe["segmento_4"]["data_alteracao"] = 00000000
	lote.InserirDetalhe(detalhe)

	detalhe = lote.NovoDetalhe()
	detalhe["segmento_5"]["codigo_empregado"] = 000000
	detalhe["segmento_5"]["codigo_cargo"] = 000
	detalhe["segmento_5"]["data_alteracao"] = 00000000
	lote.InserirDetalhe(detalhe)

	remessa.Trailer["quantidade_movimentos"] = "0"
	remessa.Trailer["total_valor_ref_eventos"] = 90.75
	remessa.Trailer["total_valor_eventos"] = 450.00
	remessa.Trailer["total_valores_situacao_empregado"] = 14000.00
	remessaFile := file.NewRemessaFile(remessa, "fortes.ps")

	arquivo := remessaFile.Write()

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, remessa)
	assert.NotNil(t, arquivo)
	assertFile(t, 9, arquivo.Name(), true)
}
