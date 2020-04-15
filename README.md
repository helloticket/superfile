# cnab-go

Parser de arquivos de configuração CNAB (layouts) baseado no projeto CNAB Layouts (http://glauberportella.github.io/cnab-layouts)

## Instalação

```bash
go get github.com/helderfarias/cnab-go
```

## Debug de layout

```go
  go run cmd/main.go
```

## Arquivo de remessa

```go
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
  log.Println(arquivo)
```

## Arquivo de retorno

```go
  source := strings.NewReader(itau.CNAB240Cobranca)
	layout, err := cnab.NewLayout("240", source)

	f, _ := os.Open("cobranca_itau_cnab240.ret")
	defer f.Close()
	arquivo, err := file.NewRetornoFile(layout, f)
  retorno := arquivo.Read()
  log.Println(retorno)
```
