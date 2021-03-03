package mte

const AEJ = `
# FORMATO: AEJ
# OBJETIVO DO ARQUIVO: Arquivo Eletrônico de Jornada – AEJ
#
# TAMANHO DO REGISTRO
# O Tamanho do Registro é de 232 bytes.
#
# ALINHAMENTO DE CAMPOS
# - Campos Numéricos (9) = Sempre à direita e preenchidos com zeros à esquerda.
# - Campos Alfanuméricos (X) = Sempre à esquerda e preenchidos com brancos à direita. 
#
servico: 'relatorio_aej'
versao: '1.0.0'
layout: 'aej'

global:
  ordenar_escrita_usando_campo: tipo_registro
  adicao_caracter_a_direita: "|"

remessa:
  header_arquivo:
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
    caepf_empregador:
      pos: [26,39]
      picture: '9(14)'
    cno_empregador:
      pos: [40,51]
      picture: '9(12)'
    razao_social_empregador:
      pos: [52,201]
      picture: 'X(150)'
    data_inicio:
      pos: [202,209]
      picture: '9(8)'
      data_format: "ddmmyyyy"
    data_fim:
      pos: [210,217]
      picture: '9(8)'
      data_format: "ddmmyyyy"
    data_geracao:
      pos: [218,225]
      picture: '9(8)'
      data_format: "ddmmyyyy"
    hora_geracao:
      pos: [226,229]
      picture: '9(4)'
      data_format: "hhmm"
    versao_aej:
      pos: [230,232]
      picture: 'X(3)'
      default: "001"
  
  trailer_arquivo:
    tipo_registro:
      pos: [10,10]
      picture: '9(1)'
      default: 9
    quantidade_registro_tipo_1:
      pos: [11,19]
      picture: '9(9)'
    quantidade_registro_tipo_2:
      pos: [20,28]
      picture: '9(9)'
    quantidade_registro_tipo_3:
      pos: [29,37]
      picture: '9(9)'
    quantidade_registro_tipo_4:
      pos: [38,46]
      picture: '9(9)'
    quantidade_registro_tipo_5:
      pos: [47,55]
      picture: '9(9)'
    quantidade_registro_tipo_6:
      pos: [56,64]
      picture: '9(9)'
    
  
  detalhes:
    segmento_2:
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 2
      identificador_vinculo:
        pos: [11,19]
        picture: '9(9)'
      cpf_empregado:
        pos: [20,30]
        picture: '9(11)'
      nome_empregado:
        pos: [31,180]
        picture: 'X(150)'
    segmento_3:
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 3
      codigo_hor_contratual:
        pos: [11,40]
        picture: 'X(30)'
      duracao_jornada:
        pos: [41,52]
        picture: '9(12)'
      hora_entrada_1:
        pos: [53,56]
        picture: '9(4)'
        data_format: "hhmm"
      hora_saida_1:
        pos: [57,60]
        picture: '9(4)'
        data_format: "hhmm"
      hora_entrada_2:
        pos: [61,64]
        picture: '9(4)'
        data_format: "hhmm"
      hora_saida_2:
        pos: [65,68]
        picture: '9(4)'
        data_format: "hhmm"
    segmento_4:
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 4
      identificador_vinculo:
        pos: [11,19]
        picture: '9(9)'
      data_marcacao:
        pos: [20,27]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      hora_marcacao:
        pos: [28,31]
        picture: '9(4)'
        data_format: "hhmm"
      num_rep:
        pos: [32,48]
        picture: '9(17)'
        default: 99999999999999999
      tipo_marcacao:
        pos: [49,49]
        picture: 'X(1)'
      seq_entrada_saida:
        pos: [50,52]
        picture: '9(3)'
      fonte_marcacao:
        pos: [53,53]
        picture: 'X(1)'
      codigo_hor_contratual:
        pos: [54,83]
        picture: 'X(30)'
      motivo:
        pos: [84,233]
        picture: 'X(150)'
    segmento_5:
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 5
      identificador_vinculo:
        pos: [11,19]
        picture: '9(9)'
      mat_esocial:
        pos: [20,49]
        picture: 'X(30)'
    segmento_6:
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 6
      nome_dev:
        pos: [11,160]
        picture: 'X(150)'
      versao_ptrp:
        pos: [161,168]
        picture: 'X(8)'
      tipo_id_dev:
        pos: [169,169]
        picture: '9(1)'
        default: 1
      id_dev:
        pos: [170,183]
        picture: '9(14)'
      razao_social_dev:
        pos: [184,333]
        picture: 'X(150)'
      email_dev:
        pos: [334,383]
        picture: 'X(50)'

retorno:
  header_arquivo:
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
    caepf_empregador:
      pos: [26,39]
      picture: '9(14)'
    cno_empregador:
      pos: [40,51]
      picture: '9(12)'
    razao_social_empregador:
      pos: [52,201]
      picture: 'X(150)'
    data_inicio:
      pos: [202,209]
      picture: '9(8)'
      data_format: "ddmmyyyy"
    data_fim:
      pos: [210,217]
      picture: '9(8)'
      data_format: "ddmmyyyy"
    data_geracao:
      pos: [218,225]
      picture: '9(8)'
      data_format: "ddmmyyyy"
    hora_geracao:
      pos: [226,229]
      picture: '9(4)'
      data_format: "hhmm"
    versao_aej:
      pos: [230,232]
      picture: 'X(3)'
      default: "001"
  
  trailer_arquivo:
    tipo_registro:
      pos: [10,10]
      picture: '9(1)'
      default: 9
    quantidade_registro_tipo_1:
      pos: [11,19]
      picture: '9(9)'
    quantidade_registro_tipo_2:
      pos: [20,28]
      picture: '9(9)'
    quantidade_registro_tipo_3:
      pos: [29,37]
      picture: '9(9)'
    quantidade_registro_tipo_4:
      pos: [38,46]
      picture: '9(9)'
    quantidade_registro_tipo_5:
      pos: [47,55]
      picture: '9(9)'
    quantidade_registro_tipo_6:
      pos: [56,64]
      picture: '9(9)'
    
  
  detalhes:
    segmento_2:
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 2
      identificador_vinculo:
        pos: [11,19]
        picture: '9(9)'
      cpf_empregado:
        pos: [20,30]
        picture: '9(11)'
      nome_empregado:
        pos: [31,180]
        picture: 'X(150)'
    segmento_3:
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 3
      codigo_hor_contratual:
        pos: [11,40]
        picture: 'X(30)'
      duracao_jornada:
        pos: [41,52]
        picture: '9(12)'
      hora_entrada_1:
        pos: [53,56]
        picture: '9(4)'
        data_format: "hhmm"
      hora_saida_1:
        pos: [57,60]
        picture: '9(4)'
        data_format: "hhmm"
      hora_entrada_2:
        pos: [61,64]
        picture: '9(4)'
        data_format: "hhmm"
      hora_saida_2:
        pos: [65,68]
        picture: '9(4)'
        data_format: "hhmm"
    segmento_4:
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 4
      identificador_vinculo:
        pos: [11,19]
        picture: '9(9)'
      data_marcacao:
        pos: [20,27]
        picture: '9(8)'
        data_format: "ddmmyyyy"
      hora_marcacao:
        pos: [28,31]
        picture: '9(4)'
        data_format: "hhmm"
      num_rep:
        pos: [32,48]
        picture: '9(17)'
        default: 99999999999999999
      tipo_marcacao:
        pos: [49,49]
        picture: 'X(1)'
      seq_entrada_saida:
        pos: [50,52]
        picture: '9(3)'
      fonte_marcacao:
        pos: [53,53]
        picture: 'X(1)'
      codigo_hor_contratual:
        pos: [54,83]
        picture: 'X(30)'
      motivo:
        pos: [84,233]
        picture: 'X(150)'
    segmento_5:
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 5
      identificador_vinculo:
        pos: [11,19]
        picture: '9(9)'
      mat_esocial:
        pos: [20,49]
        picture: 'X(30)'
    segmento_6:
      tipo_registro:
        pos: [10,10]
        picture: '9(1)'
        default: 6
      nome_dev:
        pos: [11,160]
        picture: 'X(150)'
      versao_ptrp:
        pos: [161,168]
        picture: 'X(8)'
      tipo_id_dev:
        pos: [169,169]
        picture: '9(1)'
        default: 1
      id_dev:
        pos: [170,183]
        picture: 'X(14)'
      razao_social_dev:
        pos: [184,333]
        picture: '9(150)'
      email_dev:
        pos: [334,383]
        picture: 'X(50)'
`
