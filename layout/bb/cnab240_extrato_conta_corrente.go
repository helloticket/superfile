package bb

const CNAB240ExtratoContaCorrente = `
# FORMATO: BANCO DO BRASIL - FEBRABAN CNAB240
# OBJETIVO DO ARQUIVO: Extrato de Conta Corrente para Conciliação Bancária
#
# TAMANHO DO REGISTRO
# O Tamanho do Registro é de 240 bytes.
#
# ALINHAMENTO DE CAMPOS
# - Campos Numéricos (9) = Sempre à direita e preenchidos com zeros à esquerda.
# - Campos Alfanuméricos (X) = Sempre à esquerda e preenchidos com brancos à direita. 
#
servico: 'extrato_conta_corrente'
versao: '032'
layout: 'cnab240'

remessa:
  header_lote:
    codigo_banco:
      pos: [1,3]
      picture: '9(3)'
    lote_servico:
      pos: [4,7]
      picture: '9(4)'
    tipo_registro:
      pos: [8,8]
      picture: '9(1)'
      default: '1'
    tipo_operacao:
      pos: [9,9]
      picture: 'X(1)'
      default: 'E'
    tipo_servico:
      pos: [10,11]
      picture: '9(2)'
      default: '04'
    forma_lancamento:
      pos: [12,13]
      picture: '9(2)'
      default: '40'
    versao_layout_lote:
      pos: [14,16]
      picture: '9(3)'
      default: '032'
    exclusivo_febraban_01:
      pos: [17,17]
      picture: 'X(1)'
      default: ''
    tipo_inscricao_empresa:
      pos: [18,18]
      picture: '9(1)'
    numero_inscricao_empresa:
      pos: [19,32]
      picture: '9(14)'
    codigo_convenio_banco:
      pos: [33,52]
      picture: 'X(20)'
    agencia_mantenedora_conta:
      pos: [53,57]
      picture: '9(5)'
    digito_verificador_agencia:
      pos: [58,58]
      picture: 'X(1)'
    numero_conta_corrente:
      pos: [59,70]
      picture: 'X(12)'
    digito_verificador_conta_corrente:
      pos: [71,71]
      picture: 'X(1)'
    digito_verificador_agencia_conta_corrente:
      pos: [72,72]
      picture: 'X(1)'
    nome_empresa:
      pos: [73,102]
      picture: 'X(30)'
    exclusivo_febraban_02:
      pos: [103,142]
      picture: 'X(40)'
      default: ''
    data_saldo_inicial:
      pos: [143,150]
      picture: '9(8)'
    valor_saldo_inicial:
      pos: [151,168]
      picture: '9(16)V9(2)'
    situacao_saldo_inicial:
      pos: [169,169]
      picture: 'X(1)'
    posicao_saldo_inicial:
      pos: [170,170]
      picture: 'X(1)'
    moeda_referenciada_extrato:
      pos: [171,173]
      picture: 'X(3)'
    numero_sequencia_extrato:
      pos: [174,178]
      picture: '9(5)'
    exclusivo_febraban_03:
      pos: [179,240]
      picture: 'X(62)'
      default: ''

  trailer_lote:
    codigo_banco:
      pos: [1,3]
      picture: '9(3)'
    lote_servico:
      pos: [4,7]
      picture: '9(4)'
    tipo_registro:
      pos: [8,8]
      picture: '9(1)'
      default: 5
    exclusivo_febraban_01:
      pos: [9,17]
      picture: 'X(9)'
      default: ''
    tipo_inscricao_empresa:
      pos: [18,18]
      picture: '9(1)'
    numero_inscricao_empresa:
      pos: [19,32]
      picture: '9(14)'
    codigo_convenio_banco:
      pos: [33,52]
      picture: 'X(20)'
    agencia_mantenedora_conta:
      pos: [53,57]
      picture: '9(5)'
    digito_verificador_agencia:
      pos: [58,58]
      picture: 'X(1)'
    numero_conta_corrente:
      pos: [59,70]
      picture: '9(12)'
    digito_verificador_conta_corrente:
      pos: [71,71]
      picture: 'X(1)'
    digito_verificador_agencia_conta:
      pos: [72,72]
      picture: 'X(1)'
    exclusivo_febraban_02:
      pos: [73,88]
      picture: 'X(16)'
      default: ''
    saldo_bloqueado_acima_24horas:
      pos: [89,106]
      picture: '9(16)V9(2)'
    limite_conta:
      pos: [107,124]
      picture: '9(16)V9(2)'
    saldo_bloqueado_ate_24horas:
      pos: [125,142]
      picture: '9(16)V9(2)'
    data_saldo_final:
      pos: [143,150]
      picture: '9(8)'
    valor_saldo_final:
      pos: [151,168]
      picture: '9(16)V9(2)'
    situacao_saldo_final:
      pos: [169,169]
      picture: 'X(1)'
    posicao_saldo_final:
      pos: [170,170]
      picture: 'X(1)'
    quantidade_registros_lote:
      pos: [171,176]
      picture: '9(6)'
    somatorio_valores_debito:
      pos: [177,194]
      picture: '9(16)V9(2)'
    somatorio_valores_credito:
      pos: [195,212]
      picture: '9(16)V9(2)'
    exclusivo_febraban_03:
      pos: [213,240]
      picture: 'X(28)'
      default: ''

  detalhes:
    segmento_e:
      codigo_banco:
        pos: [1,3]
        picture: '9(3)'
      lote_servico:
        pos: [4,7]
        picture: '9(4)'
      tipo_registro:
        pos: [8,8]
        picture: '9(1)'
        default: '3'
      sequencial_registro_lote:
        pos: [9,13]
        picture: '9(5)'
        default: '1'
      codigo_segmento_registro_detalhe:
        pos: [14,14]
        picture: 'X(1)'
        default: 'E'
      exclusivo_febraban_01:
        pos: [15,17]
        picture: 'X(3)'
        default: ''
      tipo_inscricao_empresa:
        pos: [18,18]
        picture: '9(1)'
      numero_inscricao_empresa:
        pos: [19,32]
        picture: '9(14)'
      codigo_convenio_banco:
        pos: [33,52]
        picture: 'X(20)'
      agencia_mantenedora_conta:
        pos: [53,57]
        picture: '9(5)'
      digito_verificador_agencia:
        pos: [58,58]
        picture: 'X(1)'
      numero_conta_corrente:
        pos: [59,70]
        picture: '9(12)'
      digito_verificador_conta_corrente:
        pos: [71,71]
        picture: 'X(1)'
      digito_verificador_agencia_conta_corrente:
        pos: [72,72]
        picture: 'X(1)'
      nome_empresa:
        pos: [73,102]
        picture: 'X(30)'
      exclusivo_febraban_02:
        pos: [103,108]
        picture: 'X(6)'
        default: ''
      natureza_lancamento:
        pos: [109,111]
        picture: 'X(3)'
      tipo_complemento_lancamento:
        pos: [112,113]
        picture: '9(2)'
      complemento_lancamento:
        pos: [114,133]
        picture: 'X(20)'
      identificacao_isencao_cpmf:
        pos: [134,134]
        picture: 'X(1)'
      data_contabil:
        pos: [135,142]
        picture: '9(8)'
      data_lancamento:
        pos: [143,150]
        picture: '9(8)'
      valor_lancamento:
        pos: [151,168]
        picture: '9(16)V9(2)'
      tipo_lancamento:
        pos: [169,169]
        picture: 'X(1)'
      categoria_lancamento:
        pos: [170,172]
        picture: '9(3)'
      codigo_historico_banco:
        pos: [173,176]
        picture: 'X(4)'
      descricao_historico_lancamento_banco:
        pos: [177,201]
        picture: 'X(25)'
      numero_documento_complemento:
        pos: [202,240]
        picture: 'X(39)' 

retorno:
  header_lote:
    codigo_banco:
      pos: [1,3]
      picture: '9(3)'
    lote_servico:
      pos: [4,7]
      picture: '9(4)'
    tipo_registro:
      pos: [8,8]
      picture: '9(1)'
      default: '1'
    tipo_operacao:
      pos: [9,9]
      picture: 'X(1)'
      default: 'E'
    tipo_servico:
      pos: [10,11]
      picture: '9(2)'
      default: '04'
    forma_lancamento:
      pos: [12,13]
      picture: '9(2)'
      default: '40'
    versao_layout_lote:
      pos: [14,16]
      picture: '9(3)'
      default: '032'
    exclusivo_febraban_01:
      pos: [17,17]
      picture: 'X(1)'
      default: ''
    tipo_inscricao_empresa:
      pos: [18,18]
      picture: '9(1)'
    numero_inscricao_empresa:
      pos: [19,32]
      picture: '9(14)'
    codigo_convenio_banco:
      pos: [33,52]
      picture: 'X(20)'
    agencia_mantenedora_conta:
      pos: [53,57]
      picture: '9(5)'
    digito_verificador_agencia:
      pos: [58,58]
      picture: 'X(1)'
    numero_conta_corrente:
      pos: [59,70]
      picture: 'X(12)'
    digito_verificador_conta_corrente:
      pos: [71,71]
      picture: 'X(1)'
    digito_verificador_agencia_conta_corrente:
      pos: [72,72]
      picture: 'X(1)'
    nome_empresa:
      pos: [73,102]
      picture: 'X(30)'
    exclusivo_febraban_02:
      pos: [103,142]
      picture: 'X(40)'
      default: ''
    data_saldo_inicial:
      pos: [143,150]
      picture: '9(8)'
    valor_saldo_inicial:
      pos: [151,168]
      picture: '9(16)V9(2)'
    situacao_saldo_inicial:
      pos: [169,169]
      picture: 'X(1)'
    posicao_saldo_inicial:
      pos: [170,170]
      picture: 'X(1)'
    moeda_referenciada_extrato:
      pos: [171,173]
      picture: 'X(3)'
    numero_sequencia_extrato:
      pos: [174,178]
      picture: '9(5)'
    exclusivo_febraban_03:
      pos: [179,240]
      picture: 'X(62)'
      default: ''

  trailer_lote:
    codigo_banco:
      pos: [1,3]
      picture: '9(3)'
    lote_servico:
      pos: [4,7]
      picture: '9(4)'
    tipo_registro:
      pos: [8,8]
      picture: '9(1)'
      default: 5
    exclusivo_febraban_01:
      pos: [9,17]
      picture: 'X(9)'
      default: ''
    tipo_inscricao_empresa:
      pos: [18,18]
      picture: '9(1)'
    numero_inscricao_empresa:
      pos: [19,32]
      picture: '9(14)'
    codigo_convenio_banco:
      pos: [33,52]
      picture: 'X(20)'
    agencia_mantenedora_conta:
      pos: [53,57]
      picture: '9(5)'
    digito_verificador_agencia:
      pos: [58,58]
      picture: 'X(1)'
    numero_conta_corrente:
      pos: [59,70]
      picture: '9(12)'
    digito_verificador_conta_corrente:
      pos: [71,71]
      picture: 'X(1)'
    digito_verificador_agencia_conta:
      pos: [72,72]
      picture: 'X(1)'
    exclusivo_febraban_02:
      pos: [73,88]
      picture: 'X(16)'
      default: ''
    saldo_bloqueado_acima_24horas:
      pos: [89,106]
      picture: '9(16)V9(2)'
    limite_conta:
      pos: [107,124]
      picture: '9(16)V9(2)'
    saldo_bloqueado_ate_24horas:
      pos: [125,142]
      picture: '9(16)V9(2)'
    data_saldo_final:
      pos: [143,150]
      picture: '9(8)'
    valor_saldo_final:
      pos: [151,168]
      picture: '9(16)V9(2)'
    situacao_saldo_final:
      pos: [169,169]
      picture: 'X(1)'
    posicao_saldo_final:
      pos: [170,170]
      picture: 'X(1)'
    quantidade_registros_lote:
      pos: [171,176]
      picture: '9(6)'
    somatorio_valores_debito:
      pos: [177,194]
      picture: '9(16)V9(2)'
    somatorio_valores_credito:
      pos: [195,212]
      picture: '9(16)V9(2)'
    exclusivo_febraban_03:
      pos: [213,240]
      picture: 'X(28)'
      default: ''

  detalhes:
    segmento_e:
      codigo_banco:
        pos: [1,3]
        picture: '9(3)'
      lote_servico:
        pos: [4,7]
        picture: '9(4)'
      tipo_registro:
        pos: [8,8]
        picture: '9(1)'
        default: '3'
      sequencial_registro_lote:
        pos: [9,13]
        picture: '9(5)'
        default: '1'
      codigo_segmento_registro_detalhe:
        pos: [14,14]
        picture: 'X(1)'
        default: 'E'
      exclusivo_febraban_01:
        pos: [15,17]
        picture: 'X(3)'
        default: ''
      tipo_inscricao_empresa:
        pos: [18,18]
        picture: '9(1)'
      numero_inscricao_empresa:
        pos: [19,32]
        picture: '9(14)'
      codigo_convenio_banco:
        pos: [33,52]
        picture: 'X(20)'
      agencia_mantenedora_conta:
        pos: [53,57]
        picture: '9(5)'
      digito_verificador_agencia:
        pos: [58,58]
        picture: 'X(1)'
      numero_conta_corrente:
        pos: [59,70]
        picture: '9(12)'
      digito_verificador_conta_corrente:
        pos: [71,71]
        picture: 'X(1)'
      digito_verificador_agencia_conta_corrente:
        pos: [72,72]
        picture: 'X(1)'
      nome_empresa:
        pos: [73,102]
        picture: 'X(30)'
      exclusivo_febraban_02:
        pos: [103,108]
        picture: 'X(6)'
        default: ''
      natureza_lancamento:
        pos: [109,111]
        picture: 'X(3)'
      tipo_complemento_lancamento:
        pos: [112,113]
        picture: '9(2)'
      complemento_lancamento:
        pos: [114,133]
        picture: 'X(20)'
      identificacao_isencao_cpmf:
        pos: [134,134]
        picture: 'X(1)'
      data_contabil:
        pos: [135,142]
        picture: '9(8)'
      data_lancamento:
        pos: [143,150]
        picture: '9(8)'
      valor_lancamento:
        pos: [151,168]
        picture: '9(16)V9(2)'
      tipo_lancamento:
        pos: [169,169]
        picture: 'X(1)'
      categoria_lancamento:
        pos: [170,172]
        picture: '9(3)'
      codigo_historico_banco:
        pos: [173,176]
        picture: 'X(4)'
      descricao_historico_lancamento_banco:
        pos: [177,201]
        picture: 'X(25)'
      numero_documento_complemento:
        pos: [202,240]
        picture: 'X(39)'
`
