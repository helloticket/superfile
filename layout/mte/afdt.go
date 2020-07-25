package mte

const AFDT = `
# FORMATO: AFDT
# OBJETIVO DO ARQUIVO: Arquivo-Fonte de Dados Tratado – AFDT
#
# TAMANHO DO REGISTRO
# O Tamanho do Registro é de 215 bytes.
#
# ALINHAMENTO DE CAMPOS
# - Campos Numéricos (9) = Sempre à direita e preenchidos com zeros à esquerda.
# - Campos Alfanuméricos (X) = Sempre à esquerda e preenchidos com brancos à direita. 
#
servico: 'relatorio_afdt'
versao: '1.0.0'
layout: 'afdt'

remessa:
  header_arquivo:
    seq_registro:
      pos: [1,9]
      picture: '9(9)'
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
    data_inicio:
      pos: [188,195]
      picture: '9(8)'
      data_format: "ddmmyyyy"
    data_fim:
      pos: [196,203]
      picture: '9(8)'
      data_format: "ddmmyyyy"
    data_geracao:
      pos: [204,211]
      picture: '9(8)'
      data_format: "ddmmyyyy"
    hora_geracao:
      pos: [212,215]
      picture: '9(4)'
      data_format: "hhmm"

  trailer_arquivo:
    seq_registro:
      pos: [1,9]
      picture: '9(9)'
    tipo_registro:
      pos: [10,10]
      picture: '9(1)'
      default: 9
  
  detalhes:
    segmento_2:
      seq_registro:
        pos: [1,9]
        picture: 'X(9)'
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 2
      data_marcacao_ponto:
        pos: [11,18]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      hora_marcacao_ponto:
        pos: [19,22]
        picture: '9(4)'
        data_format: "hhmm"
      numero_pis:
        pos: [23,34]
        picture: '9(12)'
      numero_fabricacao_rep:
        pos: [35,51]
        picture: '9(17)'
      tipo_marcacao:
        pos: [52,52]
        picture: 'X(1)'
      seq_empregado_jornada_entrada_saida:
        pos: [53,54]
        picture: '9(2)'
      tipo_registro_eletro_digi_pre_assina:
        pos: [55,55]
        picture: 'X(1)'
      motivo:
        pos: [56,155]
        picture: 'X(100)'
   
retorno:
  header_arquivo:
    seq_registro:
      pos: [1,9]
      picture: '9(9)'
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
    data_inicio:
      pos: [188,195]
      picture: '9(8)'
      data_format: "ddmmyyyy"
    data_fim:
      pos: [196,203]
      picture: '9(8)'
      data_format: "ddmmyyyy"
    data_geracao:
      pos: [204,211]
      picture: '9(8)'
      data_format: "ddmmyyyy"
    hora_geracao:
      pos: [212,215]
      picture: '9(4)'
      data_format: "hhmm"

  trailer_arquivo:
    seq_registro:
      pos: [1,9]
      picture: '9(9)'
    tipo_registro:
      pos: [10,10]
      picture: '9(1)'
      default: 9
  
  detalhes:
    segmento_2:
      seq_registro:
        pos: [1,9]
        picture: 'X(9)'
        default: NST
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 2
      data_marcacao_ponto:
        pos: [11,18]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      hora_marcacao_ponto:
        pos: [19,22]
        picture: '9(4)'
        data_format: "hhmm"
      numero_pis:
        pos: [23,34]
        picture: '9(12)'
      numero_fabricacao_rep:
        pos: [35,51]
        picture: '9(17)'
      tipo_marcacao:
        pos: [52,52]
        picture: 'X(1)'
      seq_empregado_jornada_entrada_saida:
        pos: [53,54]
        picture: '9(2)'
      tipo_registro_eletro_digi_pre_assina:
        pos: [55,55]
        picture: 'X(1)'
      motivo:
        pos: [56,155]
        picture: 'X(100)'
`
