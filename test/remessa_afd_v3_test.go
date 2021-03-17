package test

import (
	"strings"
	"testing"
	"time"

	"github.com/helloticket/superfile"
	"github.com/helloticket/superfile/file"
	"github.com/helloticket/superfile/layout/mte"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestRemessaAFDV3(t *testing.T) {
	source := strings.NewReader(mte.AFD_V3)
	layout, err := superfile.NewLayout(source)

	remessa := superfile.NewRemessa(layout)
	remessa.Header["codigo_registro"] = 1
	remessa.Header["tipo_registro"] = 1
	remessa.Header["identificador_empregador"] = 1
	remessa.Header["cpf_cnpj_empregador"] = "94281944095"
	remessa.Header["cno_caepf_empregador"] = "94281944095124"
	remessa.Header["razao_social_empregador"] = "HELOO SOLUTIONS"
	remessa.Header["inpi_empregador"] = "10000000000000000"
	remessa.Header["data_inicio"] = time.Now()
	remessa.Header["data_fim"] = time.Now()
	remessa.Header["data_geracao"] = time.Now()
	remessa.Header["hora_geracao"] = time.Now()
	remessa.Header["crc_registro"] = "1000"

	lote := remessa.NovoLote()
	remessa.InserirLote(lote)

	for i := 0; i < 10; i++ {
		detalhe := lote.NovoDetalhe()
		detalhe["segmento_2"]["codigo_registro"] = 2
		detalhe["segmento_2"]["data_gravacao"] = time.Now()
		detalhe["segmento_2"]["hora_gravacao"] = time.Now()
		detalhe["segmento_2"]["cpf_responsavel"] = "564281944095"
		detalhe["segmento_2"]["identificador_empregador"] = 2
		detalhe["segmento_2"]["cpf_cnpj_empregador"] = "564281944095"
		detalhe["segmento_2"]["cno_caepf_empregador"] = "10000000000000"
		detalhe["segmento_2"]["razao_social_ou_nome_empregador"] = "JOSE DA SILVA"
		detalhe["segmento_2"]["local_prestacao_servico"] = "RUA DOS SONHOS"
		detalhe["segmento_2"]["crc_registro"] = "1000"

		detalhe["segmento_3"]["codigo_registro"] = 3
		detalhe["segmento_3"]["data_gravacao"] = time.Now()
		detalhe["segmento_3"]["hora_gravacao"] = time.Now()
		detalhe["segmento_3"]["cpf_empregado"] = "70746928042"
		detalhe["segmento_3"]["crc_registro"] = "1000"

		detalhe["segmento_4"]["codigo_registro"] = 4
		detalhe["segmento_4"]["data_ajuste"] = time.Now()
		detalhe["segmento_4"]["hora_ajuste"] = time.Now()
		detalhe["segmento_4"]["data_ajustada"] = time.Now()
		detalhe["segmento_4"]["hora_ajustada"] = time.Now()
		detalhe["segmento_4"]["cpf_responsavel"] = "564281944095"
		detalhe["segmento_4"]["crc_registro"] = "1000"

		detalhe["segmento_5"]["codigo_registro"] = 5
		detalhe["segmento_5"]["data_gravacao"] = time.Now()
		detalhe["segmento_5"]["hora_gravacao"] = time.Now()
		detalhe["segmento_5"]["tipo_operacao"] = "I"
		detalhe["segmento_5"]["cpf_empregado"] = "564281944095"
		detalhe["segmento_5"]["nome_empregado"] = "HELOO SOLUTIONS"
		detalhe["segmento_5"]["demais_dados"] = "2000"
		detalhe["segmento_5"]["cpf_responsavel"] = "564281944095"
		detalhe["segmento_5"]["crc_registro"] = "1000"

		detalhe["segmento_6"]["codigo_registro"] = 5
		detalhe["segmento_6"]["data_gravacao"] = time.Now()
		detalhe["segmento_6"]["hora_gravacao"] = time.Now()
		detalhe["segmento_6"]["tipo_evento"] = 01

		detalhe["segmento_7"]["codigo_registro"] = 5
		detalhe["segmento_7"]["data_marcacao"] = time.Now()
		detalhe["segmento_7"]["hora_marcacao"] = time.Now()
		detalhe["segmento_7"]["cpf_empregado"] = "564281944095"
		detalhe["segmento_7"]["data_gravacao"] = time.Now()
		detalhe["segmento_7"]["hora_gravacao"] = time.Now()
		detalhe["segmento_7"]["identificador_marcacao"] = 02
		detalhe["segmento_7"]["codigo_hash"] = "shhashhashhash"
		lote.InserirDetalhe(detalhe)
	}

	remessa.Trailer["quantidade_registro_tipo_2"] = "1"
	remessa.Trailer["quantidade_registro_tipo_3"] = "1"
	remessa.Trailer["quantidade_registro_tipo_4"] = "1"
	remessa.Trailer["quantidade_registro_tipo_5"] = "1"
	remessa.Trailer["quantidade_registro_tipo_6"] = "1"
	remessa.Trailer["quantidade_registro_tipo_7"] = "1"
	remessaFile := file.NewRemessaFile(remessa, "afdv3.rem")

	arquivo := remessaFile.Write()

	teste := arquivo.Name()

	logrus.Info(teste)

	assert.Nil(t, err)
	assert.NotNil(t, layout)
	assert.NotNil(t, remessa)
	assert.NotNil(t, arquivo)
	assertFile(t, 62, arquivo.Name(), true)
}
