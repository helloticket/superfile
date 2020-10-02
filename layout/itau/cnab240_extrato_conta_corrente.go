package itau

const CNAB240ExtratoContaCorrente = `
# FORMATO: ITAU - CNAB240
# OBJETIVO DO ARQUIVO: EXTRATO DE CONTA CORRENTE + APLIC AUT
#
# TAMANHO DO REGISTRO
# O Tamanho do Registro é de 240 bytes.
#
# ALINHAMENTO DE CAMPOS
# - Campos Numéricos (9) = Sempre à direita e preenchidos com zeros à esquerda.
# - Campos Alfanuméricos (X) = Sempre à esquerda e preenchidos com brancos à direita.
#
servico: "extrato_conta_corrente"
versao: "5.0"
layout: "cnab240"

remessa:
  header_arquivo:
    codigo_banc:
      pos: [1, 3]
      picture: "9(3)"
      default: 341
    lote_servico:
      pos: [4, 7]
      picture: "9(4)"
      default: 0000
    registro_header:
      pos: [8, 8]
      picture: "X(1)"
      default: "0"
    complemento_header:
      pos: [9, 17]
      picture: "X(9)"
    tipo_inscricao_cliente:
      pos: [18, 18]
      picture: "X(1)"
    numero_cnpj_cpf_cliente:
      pos: [19, 32]
      picture: "9(14)"
    complemento_registro_01:
      pos: [33, 47]
      picture: "X(15)"
    codigo_empresa_banco:
      pos: [48, 52]
      picture: "X(5)"
    complemento_registro_02:
      pos: [53, 53]
      picture: "9(1)"
      default: 0
    numero_agencia_mantenedora_conta:
      pos: [54, 57]
      picture: "9(4)"
    digito_verificador_agencia:
      pos: [58, 58]
      picture: "X(1)"
    complemento_registro_03:
      pos: [59, 65]
      picture: "9(7)"
      default: 0000000
    numero_conta_corrente_cliente:
      pos: [66, 70]
      picture: "9(5)"
    complemento_registro_04:
      pos: [71, 71]
      picture: "X(1)"
    digito_verificador_agencia_conta:
      pos: [72, 72]
      picture: "X(1)"
    nome_empresa:
      pos: [73, 102]
      picture: "X(30)"
    nome_banco:
      pos: [103, 132]
      picture: "X(30)"
      default: "BANCO ITAU"
    complemento_registro_05:
      pos: [133, 142]
      picture: "X(10)"
    codigo_retorno:
      pos: [143, 143]
      picture: "X(1)"
      default: "2"
    data_geracao:
      pos: [144, 151]
      picture: "9(8)"
      data_format: "ddmmyyyy"
    hora_geracao:
      pos: [152, 157]
      picture: "9(6)"
      data_format: "hhmmss"
    numero_sequencial_arquivo:
      pos: [158, 163]
      picture: "9(6)"
    numero_versao_layout:
      pos: [164, 166]
      picture: "9(3)"
      default: 50
    complemento_registro_06:
      pos: [167, 171]
      picture: "9(5)"
      default: 00000
    uso_reservado_banco:
      pos: [172, 191]
      picture: "X(20)"
    complemento_registro_07:
      pos: [192, 240]
      picture: "X(49)"

  header_lote:
    codigo_banc:
      pos: [1, 3]
      picture: "9(3)"
      default: 341
    lote_servico:
      pos: [4, 7]
      picture: "9(4)"
    registro_header:
      pos: [8, 8]
      picture: "X(1)"
      default: "1"
    tipo_operacao:
      pos: [9, 9]
      picture: "X(1)"
      default: "E"
    tipo_servico:
      pos: [10, 11]
      picture: "9(2)"
      default: "4"
    forma_lancamento:
      pos: [12, 13]
      picture: "9(2)"
      default: "40"
    numero_versao_layout:
      pos: [14, 16]
      picture: "X(3)"
      default: "050"
    complemento_registro_01:
      pos: [17, 17]
      picture: "X(1)"
    tipo_inscricao_cliente:
      pos: [18, 18]
      picture: "X(1)"
    numero_inscricao_cliente:
      pos: [19, 32]
      picture: "9(14)"
    identificacao_tipo_conta:
      pos: [33, 36]
      picture: "X(4)"
    complemento_registro_02:
      pos: [37, 47]
      picture: "X(11)"
    codigo_empresa_banco:
      pos: [48, 52]
      picture: "X(5)"
    complemento_registro_03:
      pos: [53, 53]
      picture: "9(1)"
      default: 0
    numero_agencia_mantenedora_conta:
      pos: [54, 57]
      picture: "9(4)"
    digito_verificador_agencia:
      pos: [58, 58]
      picture: "X(1)"
    complemento_registro_04:
      pos: [59, 65]
      picture: "9(7)"
      default: 0000000
    numero_conta_corrente_cliente:
      pos: [66, 70]
      picture: "9(5)"
    complemento_registro_05:
      pos: [71, 71]
      picture: "X(1)"
    digito_verificador_agencia_conta:
      pos: [72, 72]
      picture: "X(1)"
    nome_empresa:
      pos: [73, 102]
      picture: "X(30)"
    complemento_registro_06:
      pos: [103, 142]
      picture: "X(40)"
    data_saldo_inicial:
      pos: [143, 150]
      picture: "9(8)"
      data_format: "ddmmyyyy"
    valor_saldo_inicial:
      pos: [151, 168]
      picture: '9(16)V9(2)'
    situacao_saldo_inicial:
      pos: [169, 169]
      picture: 'X(1)'
    posicao_saldo_inicial:
      pos: [170, 170]
      picture: 'X(1)'
    moeda_referenciada_extrato:
      pos: [171, 173]
      picture: 'X(3)'
      default: "BRL"
    numero_sequencia_extrato:
      pos: [174, 178]
      picture: '9(5)'
    complemento_registro_07:
      pos: [179, 240]
      picture: 'X(62)'

  trailer_lote:
    codigo_banc:
      pos: [1, 3]
      picture: "9(3)"
      default: 341
    lote_servico:
      pos: [4, 7]
      picture: "9(4)"
    tipo_registro:
      pos: [8, 8]
      picture: "9(1)"
      default: 5
    complemento_registro_01:
      pos: [9, 17]
      picture: 'X(9)'
    tipo_inscricao_cliente:
      pos: [18, 18]
      picture: 'X(1)'
    numero_inscricao_cliente:
      pos: [19, 32]
      picture: '9(14)'
    complemento_registro_02:
      pos: [33, 47]
      picture: 'X(15)'
    codigo_empresa_banco:
      pos: [48, 52]
      picture: 'X(5)'
    complemento_registro_03:
      pos: [53, 53]
      picture: '9(1)'
      default: 0
    numero_agencia_mantenedora_conta:
      pos: [54, 57]
      picture: '9(4)'
    digito_verificador_agencia:
      pos: [58, 58]
      picture: 'X(1)'
    complemento_registro_04:
      pos: [59, 65]
      picture: '9(7)'
      default: 0000000
    numero_conta_corrente:
      pos: [66, 70]
      picture: '9(5)'
    complemento_registro_05:
      pos: [71, 71]
      picture: 'X(1)'
    digito_verificador_agencia_conta:
      pos: [72, 72]
      picture: 'X(1)'
    complemento_registro_06:
      pos: [73, 88]
      picture: 'X(16)'
    saldo_bloqueado_acima_24_horas:
      pos: [89, 106]
      picture: '9(16)V9(2)'
    limite_conta:
      pos: [107, 124]
      picture: '9(16)V9(2)'
    saldo_bloqueado_ate_24_horas:
      pos: [125, 142]
      picture: '9(16)V9(2)'
    data_saldo_final:
      pos: [143, 150]
      picture: "9(8)"
      data_format: "ddmmyyyy"
    valor_saldo_inicial:
      pos: [151, 168]
      picture: '9(16)V9(2)'
    situacao_saldo_inicial:
      pos: [169, 169]
      picture: 'X(1)'
    posicao_saldo_inicial:
      pos: [170, 170]
      picture: 'X(1)'
    quantidade_registros_lote:
      pos: [171, 176]
      picture: '9(6)'
    somatorio_valores_a_debito:
      pos: [177, 194]
      picture: '9(16)V9(2)'
    somatorio_valores_a_credito:
      pos: [195, 212]
      picture: '9(16)V9(2)'
    somatorio_lancamentos_dia_lancamentos_futuros:
      pos: [213, 230]
      picture: '9(16)V9(2)'
    complemento_registro_07:
      pos: [231, 240]
      picture: 'X(10)'

  trailer_arquivo:
    codigo_banc:
      pos: [1, 3]
      picture: "9(3)"
      default: 341
    lote_servico:
      pos: [4, 7]
      picture: "9(4)"
      default: 9999
    tipo_registro:
      pos: [8, 8]
      picture: "9(1)"
      default: 9
    complemento_registro_01:
      pos: [9, 17]
      picture: "X(9)"
    quantidade_lotes:
      pos: [18, 23]
      picture: "9(6)"
    quantidade_registros:
      pos: [24, 29]
      picture: "9(6)"
    quantidade_contas_conciliacao:
      pos: [30, 35]
      picture: "9(6)"
    complemento_registro_02:
      pos: [36, 240]
      picture: "X(205)"

  detalhes:
    segmento_e:
      codigo_banco:
        pos: [1,3]
        picture: "9(3)"
        default: "341"
      lote_servico:
        pos: [4,7]
        picture: '9(4)'
      registro_detalhe_lote:
        pos: [8,8]
        picture: 'X(1)'
        default: '3'
      numero_sequencial_registro_lote:
        pos: [9,13]
        picture: '9(5)'
      codigo_segmento_registro_detalhe:
        pos: [14,14]
        picture: 'X(1)'
        default: 'E'
      identificacao_tipo_lancamento:
        pos: [15,15]
        picture: '9(1)'
      complemento_registro_01:
        pos: [16,17]
        picture: 'X(2)'
        default: ""
      tipo_inscricao_cliente:
        pos: [18,18]
        picture: 'X(1)'
      numero_inscricao_cliente:
        pos: [19,32]
        picture: '9(14)'
      complemento_registro_02:
        pos: [33,47]
        picture: 'X(15)'
        default: ""
      codigo_empresa_banco:
        pos: [48,52]
        picture: 'X(5)'
      complemento_registro_03:
        pos: [53,53]
        picture: '9(1)'
        default: 0
      numero_agencia_mantenedora_conta:
        pos: [54,57]
        picture: '9(4)'
      digito_verificador_agencia:
        pos: [58,58]
        picture: 'X(1)'
      complemento_registro_04:
        pos: [59,65]
        picture: '9(7)'
        default: 0000000
      numero_conta_corrente:
        pos: [66,70]
        picture: '9(5)'
      complemento_registro_05:
        pos: [71,71]
        picture: 'X(1)'
        default: ""
      digito_verificador_agencia_conta:
        pos: [72,72]
        picture: 'X(1)'
      nome_empresa:
        pos: [73,102]
        picture: 'X(30)'
      uso_banco:
        pos: [103,108]
        picture: 'X(6)'
      natureza_lancamento:
        pos: [109,111]
        picture: 'X(3)'
      tipo_complemento_lancamento:
        pos: [112,113]
        picture: '9(2)'
      banco_origem_lancamento:
        pos: [114,116]
        picture: '9(3)'
      agencia_origem_lancamento_01:
        pos: [117,121]
        picture: '9(5)'
      agencia_conta_origem_lancamento:
        pos: [122,133]
        picture: '9(12)'
      idenficacao_isencao_cpmf:
        pos: [134,134]
        picture: 'X(1)'
      data_contabil:
        pos: [135,142]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      data_lancamento:
        pos: [143,150]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      valor_lancamento:
        pos: [151,168]
        picture: '9(16)V9(2)'
      tipo_lancamento:
        pos: [169,169]
        picture: 'X(1)'
      categoria_lancamento:
        pos: [170,172]
        picture: '9(3)'
      identificacao_codigo_fluxo_caixa:
        pos: [173,176]
        picture: 'X(4)'
      descricao_historico_lacto_banco:
        pos: [177,201]
        picture: 'X(25)'
      agencia_origem_lancamento_02:
        pos: [202,205]
        picture: '9(4)'
      complemento_registro_06:
        pos: [206,207]
        picture: '9(2)'
        default: 0
      conta_origem_lancamento:
        pos: [208,212]
        picture: '9(5)'
      dac_agencia_conta_origem:
        pos: [213,213]
        picture: '9(1)'
      tipo_inscricao_emitente:
        pos: [214,214]
        picture: 'X(1)'
      numero_inscricao_emitente:
        pos: [215,228]
        picture: 'X(14)'
      complemento_registro_07:
        pos: [229,234]
        picture: 'X(6)'
        default: ""
      numero_documento_complemento:
        pos: [235,240]
        picture: 'X(6)'

retorno:
  header_arquivo:
    codigo_banco:
      pos: [1, 3]
      picture: "9(3)"
      default: 341
    lote_servico:
      pos: [4, 7]
      picture: "9(4)"
      default: 0000
    registro_header:
      pos: [8, 8]
      picture: "X(1)"
      default: "0"
    complemento_header:
      pos: [9, 17]
      picture: "X(9)"
    tipo_inscricao_cliente:
      pos: [18, 18]
      picture: "X(1)"
    numero_cnpj_cpf_cliente:
      pos: [19, 32]
      picture: "9(14)"
    complemento_registro_01:
      pos: [33, 47]
      picture: "X(15)"
    codigo_empresa_banco:
      pos: [48, 52]
      picture: "X(5)"
    complemento_registro_02:
      pos: [53, 53]
      picture: "9(1)"
      default: 0
    numero_agencia_mantenedora_conta:
      pos: [54, 57]
      picture: "9(4)"
    digito_verificador_agencia:
      pos: [58, 58]
      picture: "X(1)"
    complemento_registro_03:
      pos: [59, 65]
      picture: "9(7)"
      default: 0000000
    numero_conta_corrente_cliente:
      pos: [66, 70]
      picture: "9(5)"
    complemento_registro_04:
      pos: [71, 71]
      picture: "X(1)"
    digito_verificador_agencia_conta:
      pos: [72, 72]
      picture: "X(1)"
    nome_empresa:
      pos: [73, 102]
      picture: "X(30)"
    nome_banco:
      pos: [103, 132]
      picture: "X(30)"
      default: "BANCO ITAU"
    complemento_registro_05:
      pos: [133, 142]
      picture: "X(10)"
    codigo_retorno:
      pos: [143, 143]
      picture: "X(1)"
      default: "2"
    data_geracao:
      pos: [144, 151]
      picture: "9(8)"
      data_format: "ddmmyyyy"
    hora_geracao:
      pos: [152, 157]
      picture: "9(6)"
      data_format: "hhmmss"
    numero_sequencial_arquivo:
      pos: [158, 163]
      picture: "9(6)"
    numero_versao_layout:
      pos: [164, 166]
      picture: "9(3)"
      default: 50
    complemento_registro_06:
      pos: [167, 171]
      picture: "9(5)"
      default: 00000
    uso_reservado_banco:
      pos: [172, 191]
      picture: "X(20)"
    complemento_registro_07:
      pos: [192, 240]
      picture: "X(49)"

  header_lote:
    codigo_banco:
      pos: [1, 3]
      picture: "9(3)"
      default: 341
    lote_servico:
      pos: [4, 7]
      picture: "9(4)"
    registro_header:
      pos: [8, 8]
      picture: "X(1)"
      default: "1"
    tipo_operacao:
      pos: [9, 9]
      picture: "X(1)"
      default: "E"
    tipo_servico:
      pos: [10, 11]
      picture: "9(2)"
      default: "4"
    forma_lancamento:
      pos: [12, 13]
      picture: "9(2)"
      default: "40"
    numero_versao_layout:
      pos: [14, 16]
      picture: "X(3)"
      default: "050"
    complemento_registro_01:
      pos: [17, 17]
      picture: "X(1)"
    tipo_inscricao_cliente:
      pos: [18, 18]
      picture: "X(1)"
    numero_inscricao_cliente:
      pos: [19, 32]
      picture: "9(14)"
    identificacao_tipo_conta:
      pos: [33, 36]
      picture: "X(4)"
    complemento_registro_02:
      pos: [37, 47]
      picture: "X(11)"
    codigo_empresa_banco:
      pos: [48, 52]
      picture: "X(5)"
    complemento_registro_03:
      pos: [53, 53]
      picture: "9(1)"
      default: 0
    numero_agencia_mantenedora_conta:
      pos: [54, 57]
      picture: "9(4)"
    digito_verificador_agencia:
      pos: [58, 58]
      picture: "X(1)"
    complemento_registro_04:
      pos: [59, 65]
      picture: "9(7)"
      default: 0000000
    numero_conta_corrente_cliente:
      pos: [66, 70]
      picture: "9(5)"
    complemento_registro_05:
      pos: [71, 71]
      picture: "X(1)"
    digito_verificador_agencia_conta:
      pos: [72, 72]
      picture: "X(1)"
    nome_empresa:
      pos: [73, 102]
      picture: "X(30)"
    complemento_registro_06:
      pos: [103, 142]
      picture: "X(40)"
    data_saldo_inicial:
      pos: [143, 150]
      picture: "9(8)"
      data_format: "ddmmyyyy"
    valor_saldo_inicial:
      pos: [151, 168]
      picture: '9(16)V9(2)'
    situacao_saldo_inicial:
      pos: [169, 169]
      picture: 'X(1)'
    posicao_saldo_inicial:
      pos: [170, 170]
      picture: 'X(1)'
    moeda_referenciada_extrato:
      pos: [171, 173]
      picture: 'X(3)'
      default: "BRL"
    numero_sequencia_extrato:
      pos: [174, 178]
      picture: '9(5)'
    complemento_registro_07:
      pos: [179, 240]
      picture: 'X(62)'

  trailer_lote:
    codigo_banco:
      pos: [1, 3]
      picture: "9(3)"
      default: 341
    lote_servico:
      pos: [4, 7]
      picture: "9(4)"
    tipo_registro:
      pos: [8, 8]
      picture: "9(1)"
      default: 5
    complemento_registro_01:
      pos: [9, 17]
      picture: 'X(9)'
    tipo_inscricao_cliente:
      pos: [18, 18]
      picture: 'X(1)'
    numero_inscricao_cliente:
      pos: [19, 32]
      picture: '9(14)'
    complemento_registro_02:
      pos: [33, 47]
      picture: 'X(15)'
    codigo_empresa_banco:
      pos: [48, 52]
      picture: 'X(5)'
    complemento_registro_03:
      pos: [53, 53]
      picture: '9(1)'
      default: 0
    numero_agencia_mantenedora_conta:
      pos: [54, 57]
      picture: '9(4)'
    digito_verificador_agencia:
      pos: [58, 58]
      picture: 'X(1)'
    complemento_registro_04:
      pos: [59, 65]
      picture: '9(7)'
      default: 0000000
    numero_conta_corrente:
      pos: [66, 70]
      picture: '9(5)'
    complemento_registro_05:
      pos: [71, 71]
      picture: 'X(1)'
    digito_verificador_agencia_conta:
      pos: [72, 72]
      picture: 'X(1)'
    complemento_registro_06:
      pos: [73, 88]
      picture: 'X(16)'
    saldo_bloqueado_acima_24_horas:
      pos: [89, 106]
      picture: '9(16)V9(2)'
    limite_conta:
      pos: [107, 124]
      picture: '9(16)V9(2)'
    saldo_bloqueado_ate_24_horas:
      pos: [125, 142]
      picture: '9(16)V9(2)'
    data_saldo_final:
      pos: [143, 150]
      picture: "9(8)"
      data_format: "ddmmyyyy"
    valor_saldo_inicial:
      pos: [151, 168]
      picture: '9(16)V9(2)'
    situacao_saldo_inicial:
      pos: [169, 169]
      picture: 'X(1)'
    posicao_saldo_inicial:
      pos: [170, 170]
      picture: 'X(1)'
    quantidade_registros_lote:
      pos: [171, 176]
      picture: '9(6)'
    somatorio_valores_a_debito:
      pos: [177, 194]
      picture: '9(16)V9(2)'
    somatorio_valores_a_credito:
      pos: [195, 212]
      picture: '9(16)V9(2)'
    somatorio_lancamentos_dia_lancamentos_futuros:
      pos: [213, 230]
      picture: '9(16)V9(2)'
    complemento_registro_07:
      pos: [231, 240]
      picture: 'X(10)'

  trailer_arquivo:
    codigo_banco:
      pos: [1, 3]
      picture: "9(3)"
      default: 341
    lote_servico:
      pos: [4, 7]
      picture: "9(4)"
      default: 9999
    tipo_registro:
      pos: [8, 8]
      picture: "9(1)"
      default: 9
    complemento_registro_01:
      pos: [9, 17]
      picture: "X(9)"
    quantidade_lotes:
      pos: [18, 23]
      picture: "9(6)"
    quantidade_registros:
      pos: [24, 29]
      picture: "9(6)"
    quantidade_contas_conciliacao:
      pos: [30, 35]
      picture: "9(6)"
    complemento_registro_02:
      pos: [36, 240]
      picture: "X(205)"

  detalhes:
    segmento_e:
      codigo_banco:
        pos: [1,3]
        picture: "9(3)"
        default: "341"
      lote_servico:
        pos: [4,7]
        picture: '9(4)'
      registro_detalhe_lote:
        pos: [8,8]
        picture: 'X(1)'
        default: '3'
      numero_sequencial_registro_lote:
        pos: [9,13]
        picture: '9(5)'
      codigo_segmento_registro_detalhe:
        pos: [14,14]
        picture: 'X(1)'
        default: 'E'
      identificacao_tipo_lancamento:
        pos: [15,15]
        picture: '9(1)'
      complemento_registro_01:
        pos: [16,17]
        picture: 'X(2)'
        default: ""
      tipo_inscricao_cliente:
        pos: [18,18]
        picture: 'X(1)'
      numero_inscricao_cliente:
        pos: [19,32]
        picture: '9(14)'
      complemento_registro_02:
        pos: [33,47]
        picture: 'X(15)'
        default: ""
      codigo_empresa_banco:
        pos: [48,52]
        picture: 'X(5)'
      complemento_registro_03:
        pos: [53,53]
        picture: '9(1)'
        default: 0
      numero_agencia_mantenedora_conta:
        pos: [54,57]
        picture: '9(4)'
      digito_verificador_agencia:
        pos: [58,58]
        picture: 'X(1)'
      complemento_registro_04:
        pos: [59,65]
        picture: '9(7)'
        default: 0000000
      numero_conta_corrente:
        pos: [66,70]
        picture: '9(5)'
      complemento_registro_05:
        pos: [71,71]
        picture: 'X(1)'
        default: ""
      digito_verificador_agencia_conta:
        pos: [72,72]
        picture: 'X(1)'
      nome_empresa:
        pos: [73,102]
        picture: 'X(30)'
      uso_banco:
        pos: [103,108]
        picture: 'X(6)'
      natureza_lancamento:
        pos: [109,111]
        picture: 'X(3)'
      tipo_complemento_lancamento:
        pos: [112,113]
        picture: '9(2)'
      banco_origem_lancamento:
        pos: [114,116]
        picture: '9(3)'
      agencia_origem_lancamento_01:
        pos: [117,121]
        picture: '9(5)'
      agencia_conta_origem_lancamento:
        pos: [122,133]
        picture: '9(12)'
      idenficacao_isencao_cpmf:
        pos: [134,134]
        picture: 'X(1)'
      data_contabil:
        pos: [135,142]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      data_lancamento:
        pos: [143,150]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      valor_lancamento:
        pos: [151,168]
        picture: '9(16)V9(2)'
      tipo_lancamento:
        pos: [169,169]
        picture: 'X(1)'
      categoria_lancamento:
        pos: [170,172]
        picture: '9(3)'
      identificacao_codigo_fluxo_caixa:
        pos: [173,176]
        picture: 'X(4)'
      descricao_historico_lacto_banco:
        pos: [177,201]
        picture: 'X(25)'
      agencia_origem_lancamento_02:
        pos: [202,205]
        picture: '9(4)'
      complemento_registro_06:
        pos: [206,207]
        picture: '9(2)'
        default: 0
      conta_origem_lancamento:
        pos: [208,212]
        picture: '9(5)'
      dac_agencia_conta_origem:
        pos: [213,213]
        picture: '9(1)'
      tipo_inscricao_emitente:
        pos: [214,214]
        picture: 'X(1)'
      numero_inscricao_emitente:
        pos: [215,228]
        picture: 'X(14)'
      complemento_registro_07:
        pos: [229,234]
        picture: 'X(6)'
        default: ""
      numero_documento_complemento:
        pos: [235,240]
        picture: 'X(6)'
`
