package fortes

const ImportacaoMovimento = `
# FORMATO: PS
# OBJETIVO DO ARQUIVO: Fortes Pessoal (Importação de Movimentos)
#
# TAMANHO DO REGISTRO
# O Tamanho do Registro é de 37 bytes.
#
# ALINHAMENTO DE CAMPOS
# - Campos Numéricos (9) = Sempre à direita e preenchidos com zeros à esquerda.
# - Campos Alfanuméricos (X) = Sempre à esquerda e preenchidos com brancos à direita. 
#
servico: 'fortes'
versao: '3.0.0'
layout: 'ps'

global:
  alinhamento_alfanumerico: 'X'

remessa:
  header_arquivo:
    identificacao_registro:
      pos: [1,1]
      picture: '9(1)'
      default: 0
    versao_layout:
      pos: [2,2]
      picture: '9(1)'
      default: 3
    sistema_destino:
      pos: [3,12]
      picture: 'X(10)'
      default: 'AC'
    sistema_origem:
      pos: [13,22]
      picture: 'X(10)'
    codigo_empresa_sistema_destino:
      pos: [23,26]
      picture: 'X(4)'
    comentario:
      pos: [27,66]
      picture: 'X(40)'

  trailer_arquivo:
    identificacao_registro:
      pos: [1,1]
      picture: 'X(1)'
      default: 'Z'
    quantidade_movimentos:
      pos: [2,6]
      picture: '9(5)'
    total_valor_ref_eventos:
      pos: [7,21]
      picture: '9(12)V9(2)'
      decimal_separator: '.'
    total_valor_eventos:
      pos: [22,36]
      picture: '9(12)V9(2)'
      decimal_separator: '.'
    total_valores_situacao_empregado:
      pos: [37,51]
      picture: '9(12)V9(2)'
      decimal_separator: '.'
  
  detalhes:
    # registro detalhe folha
    segmento_1:
      identificacao_registro:
        pos: [1,1]
        picture: '9(1)'
        default: '1'
      codigo_empregado:
        pos: [2,7]
        picture: '9(6)'
      codigo_evento:
        pos: [8,10]
        picture: '9(3)'
      valor_referencia_calculo_evento:
        pos: [11,20]
        picture: '9(7)V9(2)'
        decimal_separator: '.'
      valor_evento:
        pos: [21,35]
        picture: '9(12)V9(2)'
        decimal_separator: '.'

    # registro detalhe eventos por data
    segmento_2:
      identificacao_registro:
        pos: [1,1]
        picture: '9(1)'
        default: '2'
      codigo_empregado:
        pos: [2,7]
        picture: '9(6)'
      codigo_evento:
        pos: [8,10]
        picture: '9(3)'
      valor_evento:
        pos: [11,25]
        picture: '9(12)V9(2)'
        decimal_separator: '.'
      data:
        pos: [26,33]
        picture: '9(8)'
        data_format: 'ddmmyyyy'

    # registro detalhe ocorrência por data
    segmento_3:
      identificacao_registro:
        pos: [1,1]
        picture: '9(1)'
        default: '3'
      codigo_empregado:
        pos: [2,7]
        picture: '9(6)'
      codigo_ocorrencia:
        pos: [8,10]
        picture: '9(3)'
      data:
        pos: [11,18]
        picture: '9(8)'
        data_format: 'ddmmyyyy'
      referencia:
        pos: [19,22]
        picture: '9(4)'

    # registro detalhe situação do empregado
    segmento_4:
      identificacao_registro:
        pos: [1,1]
        picture: '9(1)'
        default: '4'
      codigo_empregado:
        pos: [2,7]
        picture: '9(6)'
      valor_novo_salario:
        pos: [8,22]
        picture: '9(12)V9(2)'
        decimal_separator: '.'
      data_alteracao:
        pos: [23,30]
        picture: '9(8)'
        data_format: 'ddmmyyyy'

    # registro detalhe alteração de cargo
    segmento_5:
      identificacao_registro:
        pos: [1,1]
        picture: '9(1)'
        default: '5'
      codigo_empregado:
        pos: [2,7]
        picture: '9(6)'
      codigo_cargo:
        pos: [8,10]
        picture: '9(3)'
      data_alteracao:
        pos: [11,18]
        picture: '9(8)'
        data_format: 'ddmmyyyy'

retorno:
  header_arquivo:
    identificacao_registro:
      pos: [1,1]
      picture: '9(1)'
      default: 0
    versao_layout:
      pos: [2,2]
      picture: '9(1)'
      default: 3
    sistema_destino:
      pos: [3,12]
      picture: 'X(10)'
      default: 'AC'
    sistema_origem:
      pos: [13,22]
      picture: 'X(10)'
    codigo_empresa_sistema_destino:
      pos: [23,26]
      picture: 'X(4)'
    comentario:
      pos: [27,66]
      picture: 'X(40)'

  trailer_arquivo:
    identificacao_registro:
      pos: [1,1]
      picture: '9(1)'
      default: 'Z'
    quantidade_movimentos:
      pos: [2,6]
      picture: 'X(5)'
    total_valor_ref_eventos:
      pos: [7,21]
      picture: '9(12)V9(2)'
      decimal_separator: '.'
    total_valor_eventos:
      pos: [22,36]
      picture: '9(12)V9(2)'
      decimal_separator: '.'
    total_valores_situacao_empregado:
      pos: [37,51]
      picture: '9(12)V9(2)'
      decimal_separator: '.'
  
  detalhes:
    # registro detalhe folha
    segmento_1:
      identificacao_registro:
        pos: [1,1]
        picture: '9(1)'
        default: '1'
      codigo_empregado:
        pos: [2,7]
        picture: '9(6)'
      codigo_evento:
        pos: [8,10]
        picture: '9(3)'
      valor_referencia_calculo_evento:
        pos: [11,20]
        picture: '9(7)V9(2)'
        decimal_separator: '.'
      valor_evento:
        pos: [21,35]
        picture: '9(12)V9(2)'
        decimal_separator: '.'

    # registro detalhe eventos por data
    segmento_2:
      identificacao_registro:
        pos: [1,1]
        picture: '9(1)'
        default: '2'
      codigo_empregado:
        pos: [2,7]
        picture: '9(6)'
      codigo_evento:
        pos: [8,10]
        picture: '9(3)'
      valor_evento:
        pos: [11,25]
        picture: '9(12)V9(2)'
        decimal_separator: '.'
      data:
        pos: [26,33]
        picture: '9(8)'
        data_format: 'ddmmyyyy'

    # registro detalhe ocorrência por data
    segmento_3:
      identificacao_registro:
        pos: [1,1]
        picture: '9(1)'
        default: '3'
      codigo_empregado:
        pos: [2,7]
        picture: '9(6)'
      codigo_ocorrencia:
        pos: [8,10]
        picture: '9(3)'
      data:
        pos: [11,18]
        picture: '9(8)'
        data_format: 'ddmmyyyy'
      referencia:
        pos: [19,22]
        picture: '9(4)'

    # registro detalhe situação do empregado
    segmento_4:
      identificacao_registro:
        pos: [1,1]
        picture: '9(1)'
        default: '4'
      codigo_empregado:
        pos: [2,7]
        picture: '9(6)'
      valor_novo_salario:
        pos: [8,22]
        picture: '9(12)V9(2)'
        decimal_separator: '.'
      data_alteracao:
        pos: [23,30]
        picture: '9(8)'
        data_format: 'ddmmyyyy'

    # registro detalhe alteração de cargo
    segmento_5:
      identificacao_registro:
        pos: [1,1]
        picture: '9(1)'
        default: '5'
      codigo_empregado:
        pos: [2,7]
        picture: '9(6)'
      codigo_cargo:
        pos: [8,10]
        picture: '9(3)'
      data_alteracao:
        pos: [11,18]
        picture: '9(8)'
        data_format: 'ddmmyyyy'
`
