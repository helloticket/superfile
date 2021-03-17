package mte

const AFD_V3 = `
# FORMATO: AFDV3
# OBJETIVO DO ARQUIVO: Arquivo-Fonte de Dados – AFDV3
#
# TAMANHO DO REGISTRO
# O Tamanho do Registro é de 232 bytes.
#
# ALINHAMENTO DE CAMPOS
# - Campos Numéricos (9) = Sempre à direita e preenchidos com zeros à esquerda.
# - Campos Alfanuméricos (X) = Sempre à esquerda e preenchidos com brancos à direita. 
#
servico: 'relatorio_afd'
versao: '3.0.0'
layout: 'afd'

global:
  ordenar_escrita_por: campo_segmento
  ordenar_escrita_usando_campo: codigo_registro

remessa:
  header_arquivo:
    codigo_registro:
      pos: [1,9]
      picture: '9(9)'
      default: 0
    tipo_registro:
      pos: [10,10]
      picture: '9(1)'
      default: 1
    identificador_empregador:
      pos: [11,11]
      picture: '9(1)'
    cpf_cnpj_empregador:
      pos: [12,25]
      picture: 'X(14)'
    cno_caepf_empregador:
      pos: [26,39]
      picture: '9(14)'
    razao_social_empregador:
      pos: [40,189]
      picture: 'X(150)'
    inpi_empregador:
      pos: [190,206]
      picture: '9(17)'
      default: 99999999999999999
    data_inicio:
      pos: [207,214]
      picture: '9(8)'
      data_format: "ddmmyyyy"
    data_fim:
      pos: [215,222]
      picture: '9(8)'
      data_format: "ddmmyyyy"
    data_geracao:
      pos: [223,230]
      picture: '9(8)'
      data_format: "ddmmyyyy"
    hora_geracao:
      pos: [231,234]
      picture: '9(4)'
      data_format: "hhmm"
    layout_versao:
      pos: [235,237]
      picture: '9(3)'
      default: 3
    crc_registro:
      pos: [238,241]
      picture: 'X(4)'
  
  trailer_arquivo:
    codigo_registro:
      pos: [1,9]
      picture: '9(9)'
      default: 999999999
    quantidade_registro_tipo_2:
      pos: [10,18]
      picture: 'X(9)'
    quantidade_registro_tipo_3:
      pos: [19,27]
      picture: 'X(9)'
    quantidade_registro_tipo_4:
      pos: [28,36]
      picture: 'X(9)'
    quantidade_registro_tipo_5:
      pos: [37,45]
      picture: 'X(9)'
    quantidade_registro_tipo_6:
      pos: [46,54]
      picture: 'X(9)'
    quantidade_registro_tipo_7:
      pos: [55,63]
      picture: 'X(9)'
    tipo_registro:
      pos: [64,64]
      picture: '9(1)'
      default: 9
  
  detalhes:
    segmento_2:
      codigo_registro:
        pos: [1,9]
        picture: '9(9)'
        default: '0'
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 2
      data_gravacao:
        pos: [11,18]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      hora_gravacao:
        pos: [19,22]
        picture: '9(4)'
        data_format: "hhmm"
      cpf_responsavel:
        pos: [23,36]
        picture: 'X(14)'
      identificador_empregador:
        pos: [37,37]
        picture: '9(1)'
      cpf_cnpj_empregador:
        pos: [38,51]
        picture: 'X(14)'
      cno_caepf_empregador:
        pos: [52,65]
        picture: 'X(14)'
      razao_social_ou_nome_empregador:
        pos: [66,215]
        picture: 'X(150)'
      local_prestacao_servico:
        pos: [216,315]
        picture: 'X(100)'
      crc_registro:
        pos: [316,319]
        picture: 'X(4)'
    segmento_3:
      codigo_registro:
        pos: [1,9]
        picture: '9(9)'
        default: 0
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 3
      data_marcacao:
        pos: [11,18]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      hora_marcacao:
        pos: [19,22]
        picture: '9(4)'
        data_format: "hhmm"
      cpf_empregado:
        pos: [23,34]
        picture: 'X(12)'
      crc_registro:
        pos: [35,38]
        picture: 'X(4)'
    segmento_4:
      codigo_registro:
        pos: [1,9]
        picture: '9(9)'
        default: 0
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 4
      data_ajuste:
        pos: [11,18]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      hora_ajuste:
        pos: [19,22]
        picture: '9(4)'
        data_format: "hhmm"
      data_ajustada:
        pos: [23,30]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      hora_ajustada:
        pos: [31,34]
        picture: '9(4)'
        data_format: "hhmm"
      cpf_responsavel:
        pos: [35,45]
        picture: 'X(11)'
      crc_registro:
        pos: [46,49]
        picture: 'X(4)'
    segmento_5:
      codigo_registro:
        pos: [1,9]
        picture: '9(9)'
        default: 0
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 5
      data_gravacao:
        pos: [11,18]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      hora_gravacao:
        pos: [19,22]
        picture: '9(4)'
        data_format: "hhmm"
      tipo_operacao:
        pos: [23,23]
        picture: 'X(1)'
      cpf_empregado:
        pos: [24,35]
        picture: 'X(12)'
      nome_empregado:
        pos: [36,87]
        picture: 'X(52)'
      demais_dados:
        pos: [88,91]
        picture: 'X(4)'
      cpf_responsavel:
        pos: [92,102]
        picture: 'X(11)'
      crc_registro:
        pos: [103,106]
        picture: 'X(4)'
    segmento_6:
      codigo_registro:
        pos: [1,9]
        picture: '9(9)'
        default: 0
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 6
      data_gravacao:
        pos: [11,18]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      hora_gravacao:
        pos: [19,22]
        picture: '9(4)'
        data_format: "hhmm"
      tipo_evento:
        pos: [23,24]
        picture: '9(2)'
    segmento_7:
      codigo_registro:
        pos: [1,9]
        picture: '9(9)'
        default: 0
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 7
      data_marcacao:
        pos: [11,18]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      hora_marcacao:
        pos: [19,22]
        picture: '9(4)'
        data_format: "hhmm"
      cpf_empregado:
        pos: [23,34]
        picture: 'X(12)'
      data_gravacao:
        pos: [35,42]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      hora_gravacao:
        pos: [43,46]
        picture: '9(4)'
        data_format: "hhmm"
      identificador_marcacao:
        pos: [47,48]
        picture: '9(2)'
      codigo_hash:
        pos: [49,112]
        picture: 'X(64)'

retorno:
  header_arquivo:
    codigo_registro:
      pos: [1,9]
      picture: '9(9)'
      default: 0
    tipo_registro:
      pos: [10,10]
      picture: '9(1)'
      default: 1
    identificador_empregador:
      pos: [11,11]
      picture: '9(1)'
    cpf_cnpj_empregador:
      pos: [12,25]
      picture: 'X(14)'
    cno_caepf_empregador:
      pos: [26,39]
      picture: 'X(14)'
    razao_social_empregador:
      pos: [40,189]
      picture: 'X(150)'
    inpi_empregador:
      pos: [190,206]
      picture: 'X(17)'
      default: 99999999999999999
    data_inicio:
      pos: [207,214]
      picture: '9(8)'
      data_format: "ddmmyyyy"
    data_fim:
      pos: [215,222]
      picture: '9(8)'
      data_format: "ddmmyyyy"
    data_geracao:
      pos: [223,230]
      picture: '9(8)'
      data_format: "ddmmyyyy"
    hora_geracao:
      pos: [231,234]
      picture: '9(4)'
      data_format: "hhmm"
    layout_versao:
      pos: [235,237]
      picture: '9(3)'
      default: 3
    crc_registro:
      pos: [238,241]
      picture: 'X(4)'

  trailer_arquivo:
    codigo_registro:
      pos: [1,9]
      picture: '9(9)'
      default: 999999999
    quantidade_registro_tipo_2:
      pos: [10,18]
      picture: 'X(9)'
    quantidade_registro_tipo_3:
      pos: [19,27]
      picture: 'X(9)'
    quantidade_registro_tipo_4:
      pos: [28,36]
      picture: 'X(9)'
    quantidade_registro_tipo_5:
      pos: [37,45]
      picture: 'X(9)'
    quantidade_registro_tipo_6:
      pos: [46,54]
      picture: 'X(9)'
    quantidade_registro_tipo_7:
      pos: [55,63]
      picture: 'X(9)'
    tipo_registro:
      pos: [64,64]
      picture: '9(1)'
      default: 9

  detalhes:
    segmento_2:
      codigo_registro:
        pos: [1,9]
        picture: '9(9)'
        default: '0'
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 2
      data_gravacao:
        pos: [11,18]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      hora_gravacao:
        pos: [19,22]
        picture: '9(4)'
        data_format: "hhmm"
      cpf_responsavel:
        pos: [23,36]
        picture: 'X(14)'
      identificador_empregador:
        pos: [37,37]
        picture: '9(1)'
      cpf_cnpj_empregador:
        pos: [38,51]
        picture: 'X(14)'
      cno_caepf_empregador:
        pos: [52,65]
        picture: 'X(14)'
      razao_social_ou_nome_empregador:
        pos: [66,215]
        picture: 'X(150)'
      local_prestacao_servico:
        pos: [216,315]
        picture: 'X(100)'
      crc_registro:
        pos: [316,319]
        picture: 'X(4)'
    segmento_3:
      codigo_registro:
        pos: [1,9]
        picture: '9(9)'
        default: 0
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 3
      data_marcacao:
        pos: [11,18]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      hora_marcacao:
        pos: [19,22]
        picture: '9(4)'
        data_format: "hhmm"
      cpf_empregado:
        pos: [23,34]
        picture: 'X(12)'
      crc_registro:
        pos: [35,38]
        picture: 'X(4)'
    segmento_4:
      codigo_registro:
        pos: [1,9]
        picture: '9(9)'
        default: 0
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 4
      data_ajuste:
        pos: [11,18]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      hora_ajuste:
        pos: [19,22]
        picture: '9(4)'
        data_format: "hhmm"
      data_ajustada:
        pos: [23,30]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      hora_ajustada:
        pos: [31,34]
        picture: '9(4)'
        data_format: "hhmm"
      cpf_responsavel:
        pos: [35,45]
        picture: 'X(11)'
      crc_registro:
        pos: [46,49]
        picture: 'X(4)'
    segmento_5:
      codigo_registro:
        pos: [1,9]
        picture: '9(9)'
        default: 0
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 5
      data_gravacao:
        pos: [11,18]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      hora_gravacao:
        pos: [19,22]
        picture: '9(4)'
        data_format: "hhmm"
      tipo_operacao:
        pos: [23,23]
        picture: 'X(1)'
      cpf_empregado:
        pos: [24,35]
        picture: 'X(12)'
      nome_empregado:
        pos: [36,87]
        picture: 'X(52)'
      demais_dados:
        pos: [88,91]
        picture: 'X(4)'
      cpf_responsavel:
        pos: [92,102]
        picture: 'X(11)'
      crc_registro:
        pos: [103,106]
        picture: 'X(4)'
    segmento_6:
      codigo_registro:
        pos: [1,9]
        picture: '9(9)'
        default: 0
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 6
      data_gravacao:
        pos: [11,18]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      hora_gravacao:
        pos: [19,22]
        picture: '9(4)'
        data_format: "hhmm"
      tipo_evento:
        pos: [23,24]
        picture: '9(2)'
    segmento_7:
      codigo_registro:
        pos: [1,9]
        picture: '9(9)'
        default: 0
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 7
      data_marcacao:
        pos: [11,18]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      hora_marcacao:
        pos: [19,22]
        picture: '9(4)'
        data_format: "hhmm"
      cpf_empregado:
        pos: [23,34]
        picture: 'X(12)'
      data_gravacao:
        pos: [35,42]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      hora_gravacao:
        pos: [43,46]
        picture: '9(4)'
        data_format: "hhmm"
      identificador_marcacao:
        pos: [47,48]
        picture: '9(2)'
      codigo_hash:
        pos: [49,112]
        picture: 'X(64)'
`
