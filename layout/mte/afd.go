package mte

const AFD = `
# FORMATO: AFD
# OBJETIVO DO ARQUIVO: Arquivo-Fonte de Dados – AFD
#
# TAMANHO DO REGISTRO
# O Tamanho do Registro é de 232 bytes.
#
# ALINHAMENTO DE CAMPOS
# - Campos Numéricos (9) = Sempre à direita e preenchidos com zeros à esquerda.
# - Campos Alfanuméricos (X) = Sempre à esquerda e preenchidos com brancos à direita. 
#
servico: 'relatorio_afd'
versao: '1.0.0'
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
      picture: '9(14)'
    cei_empregador:
      pos: [26,37]
      picture: '9(12)'
    razao_social_empregador:
      pos: [38,187]
      picture: 'X(150)'
    numero_rep:
      pos: [188,204]
      picture: '9(17)'
    data_inicio:
      pos: [205,212]
      picture: '9(8)'
      data_format: "ddmmyyyy"
    data_fim:
      pos: [213,220]
      picture: '9(8)'
      data_format: "ddmmyyyy"
    data_geracao:
      pos: [221,228]
      picture: '9(8)'
      data_format: "ddmmyyyy"
    hora_geracao:
      pos: [229,232]
      picture: '9(4)'
      data_format: "hhmm"
  
  trailer_arquivo:
    codigo_registro:
      pos: [1,9]
      picture: '9(9)'
      default: 999999999
    quantidade_registro_tipo_2:
      pos: [10,18]
      picture: '9(9)'
    quantidade_registro_tipo_3:
      pos: [19,27]
      picture: '9(9)'
    quantidade_registro_tipo_4:
      pos: [28,36]
      picture: '9(9)'
    quantidade_registro_tipo_5:
      pos: [37,45]
      picture: '9(9)'
    tipo_registro:
      pos: [46,46]
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
      identificador_empregador:
        pos: [23,23]
        picture: '9(1)'
      cpf_cnpj_empregador:
        pos: [24,37]
        picture: '9(14)'
      cei_empregador:
        pos: [38,49]
        picture: '9(12)'
      razao_social_ou_nome_empregador:
        pos: [50,199]
        picture: 'X(150)'
      local_prestacao_servico:
        pos: [200,299]
        picture: 'X(100)'
    segmento_3:
      codigo_registro:
        pos: [1,9]
        picture: '9(9)'
        default: 0
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 3
      data_gravacao:
        pos: [11,18]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      hora_gravacao:
        pos: [19,22]
        picture: '9(4)'
        data_format: "hhmm"
      numero_pis:
        pos: [23,34]
        picture: '9(12)'
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
      numero_pis:
        pos: [24,35]
        picture: '9(12)'
      nome_empregado:
        pos: [36,87]
        picture: 'X(52)'

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
      picture: '9(14)'
    cei_empregador:
      pos: [26,37]
      picture: '9(12)'
    razao_social_empregador:
      pos: [38,187]
      picture: 'X(150)'
    numero_rep:
      pos: [188,204]
      picture: '9(17)'
    data_inicio:
      pos: [205,212]
      picture: '9(8)'
      data_format: "ddmmyyyy"
    data_fim:
      pos: [213,220]
      picture: '9(8)'
      data_format: "ddmmyyyy"
    data_geracao:
      pos: [221,228]
      picture: '9(8)'
      data_format: "ddmmyyyy"
    hora_geracao:
      pos: [229,232]
      picture: '9(4)'
      data_format: "hhmm"
  
  trailer_arquivo:
    codigo_registro:
      pos: [1,9]
      picture: '9(9)'
      default: 999999999
    quantidade_registro_tipo_2:
      pos: [10,18]
      picture: '9(9)'
    quantidade_registro_tipo_3:
      pos: [19,27]
      picture: '9(9)'
    quantidade_registro_tipo_4:
      pos: [28,36]
      picture: '9(9)'
    quantidade_registro_tipo_5:
      pos: [37,45]
      picture: '9(9)'
    tipo_registro:
      pos: [46,46]
      picture: '9(1)'
      default: 9
  
  detalhes:
    segmento_2:
      codigo_registro:
        pos: [1,9]
        picture: 'X(9)'
        default: NST
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
      identificador_empregador:
        pos: [23,23]
        picture: '9(1)'
      cpf_cnpj_empregador:
        pos: [24,37]
        picture: '9(14)'
      cei_empregador:
        pos: [38,49]
        picture: '9(12)'
      razao_social_ou_nome_empregador:
        pos: [50,199]
        picture: 'X(150)'
      local_prestacao_servico:
        pos: [200,299]
        picture: 'X(100)'
    segmento_3:
      codigo_registro:
        pos: [1,9]
        picture: 'X(9)'
        default: NST
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 3
      data_gravacao:
        pos: [11,18]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      hora_gravacao:
        pos: [19,22]
        picture: '9(4)'
        data_format: "hhmm"
      numero_pis:
        pos: [23,34]
        picture: '9(12)'
    segmento_4:
      codigo_registro:
        pos: [1,9]
        picture: 'X(9)'
        default: NST
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
    segmento_5:
      codigo_registro:
        pos: [1,9]
        picture: 'X(9)'
        default: NST
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
      numero_pis:
        pos: [24,35]
        picture: '9(12)'
      nome_empregado:
        pos: [36,87]
        picture: 'X(52)'
`
