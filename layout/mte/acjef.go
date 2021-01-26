package mte

const ACJEF = `
# FORMATO: ACJEF
# OBJETIVO DO ARQUIVO: Arquivo de Controle de Jornada para Efeitos Fiscais– ACJEF
#
# TAMANHO DO REGISTRO
# O Tamanho do Registro é de 215 bytes.
#
# ALINHAMENTO DE CAMPOS
# - Campos Numéricos (9) = Sempre à direita e preenchidos com zeros à esquerda.
# - Campos Alfanuméricos (X) = Sempre à esquerda e preenchidos com brancos à direita. 
#
servico: 'relatorio_acjef'
versao: '1.0.0'
layout: 'acjef'

global:
  ordenar_escrita_por: campo_segmento
  ordenar_escrita_usando_campo: seq_registro

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
        picture: '9(9)'
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 2
      codigo_horario:
        pos: [11,14]
        picture: '9(4)'  
      entrada:
        pos: [15,18]
        picture: '9(4)'  
        data_format: "hhmm"
      inicio_intervalo:
        pos: [19,22]
        picture: '9(4)'
        data_format: "hhmm"
      fim_intervalo:
        pos: [23,26]
        picture: '9(4)'
        data_format: "hhmm"  
      saida:
        pos: [27,30]
        picture: '9(4)'
        data_format: "hhmm"  
    segmento_3:
      seq_registro:
        pos: [1,9]
        picture: '9(9)'
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 3
      numero_pis:
        pos: [11,22]
        picture: '9(12)'
      data_inicio_jornada:
        pos: [23,30]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      primeiro_horario_jornada:
        pos: [31,34]
        picture: '9(4)'
        data_format: "hhmm"
      codigo_horario_previsto_jornada:
        pos: [35,38]
        picture: '9(4)'
      horas_diurnas_nao_extraordinarias:
        pos: [39,42]
        picture: '9(4)'
        data_format: "hhmm"
      horas_noturnas_nao_extraordinarias:
        pos: [43,46]
        picture: '9(4)'
        data_format: "hhmm"
      horas_extras_1:
        pos: [47,50]
        picture: '9(4)'
        data_format: "hhmm"
      percentual_adicional_horas_extras_1:
        pos: [51,54]
        picture: '9(4)'
      modalidades_hora_extra_1:
        pos: [55,55]
        picture: 'X(1)'
      horas_extras_2:
        pos: [56,59]
        picture: '9(4)'
        data_format: "hhmm"
      percentual_adicional_horas_extras_2:
        pos: [60,63]
        picture: '9(4)'
      modalidades_hora_extra_2:
        pos: [64,64]
        picture: 'X(1)'
      horas_extra_3:
        pos: [65,68]
        picture: '9(4)'
        data_format: "hhmm"
      percentual_adicional_horas_extras_3:
        pos: [69,72]
        picture: '9(4)'
      modalidades_hora_extra_3:
        pos: [73,73]
        picture: 'X(1)'
      horas_extra_4:
        pos: [74,77]
        picture: '9(4)'
        data_format: "hhmm"
      percentual_adicional_horas_extras_4:
        pos: [78,81]
        picture: '9(4)'
      modalidades_hora_extra_4:
        pos: [82,82]
        picture: 'X(1)'
      horas_faltas_atrasos:
        pos: [83,86]
        picture: '9(4)'
        data_format: "hhmm"
      sinal_horas_compensar:
        pos: [87,87]
        picture: '9(1)'
      saldo_horas_compensar:
        pos: [88,91]
        picture: '9(4)'
        data_format: "hhmm"

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
        picture: '9(9)'
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 2
      codigo_horario:
        pos: [11,14]
        picture: '9(4)'  
      entrada:
        pos: [15,18]
        picture: '9(4)'  
      inicio_intervalo:
        pos: [19,22]
        picture: '9(4)'  
      fim_intervalo:
        pos: [23,26]
        picture: '9(4)'  
      saida:
        pos: [27,30]
        picture: '9(4)'  
    segmento_3:
      seq_registro:
        pos: [1,9]
        picture: '9(9)'
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 3
      numero_pis:
        pos: [11,22]
        picture: '9(12)'
      data_inicio_jornada:
        pos: [23,30]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      primeiro_horario_jornada:
        pos: [31,34]
        picture: '9(4)'
        data_format: "hhmm"
      codigo_horario_previsto_jornada:
        pos: [35,38]
        picture: '9(4)'
      horas_diurnas_nao_extraordinarias:
        pos: [39,42]
        picture: '9(4)'
        data_format: "hhmm"
      horas_noturnas_nao_extraordinarias:
        pos: [43,46]
        picture: '9(4)'
        data_format: "hhmm"
      horas_extras_1:
        pos: [47,50]
        picture: '9(4)'
        data_format: "hhmm"
      percentual_adicional_horas_extras_1:
        pos: [51,54]
        picture: '9(4)'
      modalidades_hora_extra_1:
        pos: [55,55]
        picture: 'X(1)'
      horas_extras_2:
        pos: [56,59]
        picture: '9(4)'
        data_format: "hhmm"
      percentual_adicional_horas_extras_2:
        pos: [60,63]
        picture: '9(4)'
      modalidades_hora_extra_2:
        pos: [64,64]
        picture: 'X(1)'
      horas_extra_3:
        pos: [65,68]
        picture: '9(4)'
        data_format: "hhmm"
      percentual_adicional_horas_extras_3:
        pos: [69,72]
        picture: '9(4)'
      modalidades_hora_extra_3:
        pos: [73,73]
        picture: 'X(1)'
      horas_extra_4:
        pos: [74,77]
        picture: '9(4)'
        data_format: "hhmm"
      percentual_adicional_horas_extras_4:
        pos: [78,81]
        picture: '9(4)'
      modalidades_hora_extra_4:
        pos: [82,82]
        picture: 'X(1)'
      horas_faltas_atrasos:
        pos: [83,86]
        picture: '9(4)'
        data_format: "hhmm"
      sinal_horas_compensar:
        pos: [87,87]
        picture: '9(1)'
      saldo_horas_compensar:
        pos: [88,91]
        picture: '9(4)'
        data_format: "hhmm"
`
