package ccsitef

// ExtratoEletronico serviço/produto
const ExtratoEletronico = `
servico: "extrato"
versao: "1.6b–Agosto-2007"
layout: "ccsitef"

remessa:
  header_arquivo:
    codigo_registro:
      pos: [1, 2]
      picture: "X(2)"
      default: "A0"
    versao_layout:
      pos: [3, 8]
      picture: "X(6)"
      default: "001.6b"
    data_geracao:
      pos: [9, 16]
      picture: "9(8)"
      data_format: "yyyymmddd"
    hora_geracao:
      pos: [17, 22]
      picture: "9(6)"
      data_format: "hhmmss"
    id_movimento:
      pos: [23, 28]
      picture: "9(6)"
    nome_administradora:
      pos: [29, 58]
      picture: "X(30)"
    identificacao_remetente:
      pos: [59, 62]
      picture: "9(4)"
    identificacao_destinatario:
      pos: [63, 68]
      picture: "9(6)"
    nseq_registro:
      pos: [69, 74]
      picture: "9(6)"

  trailer_arquivo:
    codigo_registro:
      pos: [1, 2]
      picture: "X(2)"
      default: "A9"
    total_geral_registros:
      pos: [3, 8]
      picture: "9(6)"
    nseq_registro:
      pos: [9, 14]
      picture: "9(6)"

  header_lote:
    codigo_registro:
      pos: [1, 2]
      picture: "X(2)"
      default: "L0"
    data_movimento:
      pos: [3, 10]
      picture: "9(8)"
      data_format: "yyyymmdd"
    identificacao_moeda:
      pos: [11, 12]
      picture: "X(2)"
      data_format: "yyyymmdd"
    nseq:
      pos: [13, 18]
      picture: "9(6)"

  trailer_lote:
    codigo_registro:
      pos: [1, 2]
      picture: "X(2)"
      default: "L9"
    total_registros_transacao:
      pos: [3, 8]
      picture: "9(6)"
    total_valores_creditos:
      pos: [9, 22]
      picture: "9(12)V9(2)"
    nseq:
      pos: [23, 28]
      picture: "9(6)"

  detalhes:
    segmento_cv:
      codigo_banco:
        pos: [1, 2]
        picture: "X(2)"
        default: "CV"
      identificacao_loja:
        pos: [3, 17]
        picture: "X(15)"
      nsu_host_transacao:
        pos: [18, 29]
        picture: "9(12)"
      data_transacao:
        pos: [30, 37]
        picture: "9(8)"
        data_format: "yyyymmdd"
      hora_transacao:
        pos: [38, 43]
        picture: "9(6)"
        data_format: "hhmmss"
      tipo_lancamento:
        pos: [44, 44]
        picture: "9(1)"
      data_lancamento:
        pos: [45, 52]
        picture: "9(8)"
        data_format: "yyyymmdd"
      tipo_produto:
        pos: [53, 53]
        picture: "X(1)"
      meio_captura:
        pos: [54, 54]
        picture: "9(1)"
      valor_bruno_venda:
        pos: [55, 65]
        picture: "9(9)V9(2)"
      valor_desconto:
        pos: [66, 76]
        picture: "9(9)V9(2)"
      valor_liquida_venda:
        pos: [77, 87]
        picture: "9(9)V9(2)"
      numero_cartao:
        pos: [88, 106]
        picture: "X(19)"
      numero_parcela:
        pos: [107, 108]
        picture: "9(2)"
      numero_total_parcelas:
        pos: [109, 110]
        picture: "9(2)"
      nsu_host_parcela:
        pos: [111, 122]
        picture: "X(12)"
      valor_bruto_parcela:
        pos: [123, 133]
        picture: "9(9)V9(2)"
      valor_desconto_parcela:
        pos: [134, 144]
        picture: "9(9)V9(2)"
      valor_liquido_parcela:
        pos: [145, 155]
        picture: "9(9)V9(2)"
      banco:
        pos: [156, 158]
        picture: "9(3)"
      agencia:
        pos: [159, 164]
        picture: "9(6)"
      conta:
        pos: [165, 175]
        picture: "X(11)"
      codigo_autorizacao:
        pos: [176, 187]
        picture: "9(12)"
      nseq:
        pos: [188, 193]
        picture: "9(6)"

    segmento_aj:
      codigo_banco:
        pos: [1, 2]
        picture: "X(2)"
        default: "AJ"
      identificacao_loja:
        pos: [3, 17]
        picture: "X(15)"
      nsu_host_transacao_original:
        pos: [18, 29]
        picture: "9(12)"
      data_transacao_original:
        pos: [30, 37]
        picture: "9(8)"
        data_format: "yyyymmdd"
      numero_parcela:
        pos: [38, 39]
        picture: "9(2)"
      nsu_host_transacao:
        pos: [40, 51]
        picture: "9(12)"
      data_transacao:
        pos: [52, 59]
        picture: "9(8)"
        data_format: "yyyymmdd"
      hora_transacao:
        pos: [60, 65]
        picture: "9(6)"
        data_format: "hhmmss"
      tipo_lancamento:
        pos: [66, 66]
        picture: "9(1)"
      data_lancamento:
        pos: [67, 74]
        picture: "9(8)"
        data_format: "yyyymmdd"
      meio_captura:
        pos: [75, 75]
        picture: "9(1)"
      tipo_ajuste:
        pos: [76, 76]
        picture: "9(1)"
      codigo_ajuste:
        pos: [77, 79]
        picture: "9(3)"
      descricao_motivo_ajuste:
        pos: [80, 109]
        picture: "X(30)"
      valor_bruto_parcela:
        pos: [110, 120]
        picture: "9(9)V9(2)"
      valor_desconto_comissao:
        pos: [121, 131]
        picture: "9(9)V9(2)"
      valor_liquido:
        pos: [132, 142]
        picture: "9(9)V9(2)"
      banco:
        pos: [143, 145]
        picture: "9(3)"
      agencia:
        pos: [146, 151]
        picture: "9(6)"
      conta:
        pos: [152, 162]
        picture: "X(11)"
      nseq:
        pos: [163, 168]
        picture: "9(6)"

    segmento_cc:
      codigo_banco:
        pos: [1, 2]
        picture: "X(2)"
        default: "AJ"
      identificacao_loja:
        pos: [3, 17]
        picture: "X(15)"
      nsu_host_transacao_original:
        pos: [18, 29]
        picture: "9(12)"
      data_transacao_original:
        pos: [30, 37]
        picture: "9(8)"
        data_format: "yyyymmdd"
      numero_parcela:
        pos: [38, 39]
        picture: "9(2)"
      nsu_host_transacao:
        pos: [40, 51]
        picture: "9(12)"
      data_transacao:
        pos: [52, 59]
        picture: "9(8)"
        data_format: "yyyymmdd"
      hora_transacao:
        pos: [60, 65]
        picture: "9(6)"
        data_format: "hhmmss"
      meio_captura:
        pos: [66, 66]
        picture: "9(1)"
      nseq:
        pos: [67, 72]
        picture: "9(6)"

retorno:
  header_arquivo:
    codigo_registro:
      pos: [1, 2]
      picture: "X(2)"
      default: "A0"
    versao_layout:
      pos: [3, 8]
      picture: "X(6)"
      default: "001.6b"
    data_geracao:
      pos: [9, 16]
      picture: "9(8)"
      data_format: "yyyymmddd"
    hora_geracao:
      pos: [17, 22]
      picture: "9(6)"
      data_format: "hhmmss"
    id_movimento:
      pos: [23, 28]
      picture: "9(6)"
    nome_administradora:
      pos: [29, 58]
      picture: "X(30)"
    identificacao_remetente:
      pos: [59, 62]
      picture: "9(4)"
    identificacao_destinatario:
      pos: [63, 68]
      picture: "9(6)"
    nseq_registro:
      pos: [69, 74]
      picture: "9(6)"

  trailer_arquivo:
    codigo_registro:
      pos: [1, 2]
      picture: "X(2)"
      default: "A9"
    total_geral_registros:
      pos: [3, 8]
      picture: "9(6)"
    nseq_registro:
      pos: [9, 14]
      picture: "9(6)"

  header_lote:
    codigo_registro:
      pos: [1, 2]
      picture: "X(2)"
      default: "L0"
    data_movimento:
      pos: [3, 10]
      picture: "9(8)"
      data_format: "yyyymmdd"
    identificacao_moeda:
      pos: [11, 12]
      picture: "X(2)"
      data_format: "yyyymmdd"
    nseq:
      pos: [13, 18]
      picture: "9(6)"

  trailer_lote:
    codigo_registro:
      pos: [1, 2]
      picture: "X(2)"
      default: "L9"
    total_registros_transacao:
      pos: [3, 8]
      picture: "9(6)"
    total_valores_creditos:
      pos: [9, 22]
      picture: "9(12)V9(2)"
    nseq:
      pos: [23, 28]
      picture: "9(6)"

  detalhes:
    segmento_cv:
      codigo_banco:
        pos: [1, 2]
        picture: "X(2)"
        default: "CV"
      identificacao_loja:
        pos: [3, 17]
        picture: "X(15)"
      nsu_host_transacao:
        pos: [18, 29]
        picture: "9(12)"
      data_transacao:
        pos: [30, 37]
        picture: "9(8)"
        data_format: "yyyymmdd"
      hora_transacao:
        pos: [38, 43]
        picture: "9(6)"
        data_format: "hhmmss"
      tipo_lancamento:
        pos: [44, 44]
        picture: "9(1)"
      data_lancamento:
        pos: [45, 52]
        picture: "9(8)"
        data_format: "yyyymmdd"
      tipo_produto:
        pos: [53, 53]
        picture: "X(1)"
      meio_captura:
        pos: [54, 54]
        picture: "9(1)"
      valor_bruno_venda:
        pos: [55, 65]
        picture: "9(9)V9(2)"
      valor_desconto:
        pos: [66, 76]
        picture: "9(9)V9(2)"
      valor_liquida_venda:
        pos: [77, 87]
        picture: "9(9)V9(2)"
      numero_cartao:
        pos: [88, 106]
        picture: "X(19)"
      numero_parcela:
        pos: [107, 108]
        picture: "9(2)"
      numero_total_parcelas:
        pos: [109, 110]
        picture: "9(2)"
      nsu_host_parcela:
        pos: [111, 122]
        picture: "X(12)"
      valor_bruto_parcela:
        pos: [123, 133]
        picture: "9(9)V9(2)"
      valor_desconto_parcela:
        pos: [134, 144]
        picture: "9(9)V9(2)"
      valor_liquido_parcela:
        pos: [145, 155]
        picture: "9(9)V9(2)"
      banco:
        pos: [156, 158]
        picture: "9(3)"
      agencia:
        pos: [159, 164]
        picture: "9(6)"
      conta:
        pos: [165, 175]
        picture: "X(11)"
      codigo_autorizacao:
        pos: [176, 187]
        picture: "9(12)"
      nseq:
        pos: [188, 193]
        picture: "9(6)"

    segmento_aj:
      codigo_banco:
        pos: [1, 2]
        picture: "X(2)"
        default: "AJ"
      identificacao_loja:
        pos: [3, 17]
        picture: "X(15)"
      nsu_host_transacao_original:
        pos: [18, 29]
        picture: "9(12)"
      data_transacao_original:
        pos: [30, 37]
        picture: "9(8)"
        data_format: "yyyymmdd"
      numero_parcela:
        pos: [38, 39]
        picture: "9(2)"
      nsu_host_transacao:
        pos: [40, 51]
        picture: "9(12)"
      data_transacao:
        pos: [52, 59]
        picture: "9(8)"
        data_format: "yyyymmdd"
      hora_transacao:
        pos: [60, 65]
        picture: "9(6)"
        data_format: "hhmmss"
      tipo_lancamento:
        pos: [66, 66]
        picture: "9(1)"
      data_lancamento:
        pos: [67, 74]
        picture: "9(8)"
        data_format: "yyyymmdd"
      meio_captura:
        pos: [75, 75]
        picture: "9(1)"
      tipo_ajuste:
        pos: [76, 76]
        picture: "9(1)"
      codigo_ajuste:
        pos: [77, 79]
        picture: "9(3)"
      descricao_motivo_ajuste:
        pos: [80, 109]
        picture: "X(30)"
      valor_bruto_parcela:
        pos: [110, 120]
        picture: "9(9)V9(2)"
      valor_desconto_comissao:
        pos: [121, 131]
        picture: "9(9)V9(2)"
      valor_liquido:
        pos: [132, 142]
        picture: "9(9)V9(2)"
      banco:
        pos: [143, 145]
        picture: "9(3)"
      agencia:
        pos: [146, 151]
        picture: "9(6)"
      conta:
        pos: [152, 162]
        picture: "X(11)"
      nseq:
        pos: [163, 168]
        picture: "9(6)"

    segmento_cc:
      codigo_banco:
        pos: [1, 2]
        picture: "X(2)"
        default: "AJ"
      identificacao_loja:
        pos: [3, 17]
        picture: "X(15)"
      nsu_host_transacao_original:
        pos: [18, 29]
        picture: "9(12)"
      data_transacao_original:
        pos: [30, 37]
        picture: "9(8)"
        data_format: "yyyymmdd"
      numero_parcela:
        pos: [38, 39]
        picture: "9(2)"
      nsu_host_transacao:
        pos: [40, 51]
        picture: "9(12)"
      data_transacao:
        pos: [52, 59]
        picture: "9(8)"
        data_format: "yyyymmdd"
      hora_transacao:
        pos: [60, 65]
        picture: "9(6)"
        data_format: "hhmmss"
      meio_captura:
        pos: [66, 66]
        picture: "9(1)"
      nseq:
        pos: [67, 72]
        picture: "9(6)"
`
