package itau

const CNAB400Cobranca = `
# FORMATO: ITAU - CNAB400
# OBJETIVO DO ARQUIVO: COBRANÇA BANCÁRIA
#
# TAMANHO DO REGISTRO
# O Tamanho do Registro é de 400 bytes.
#
# ALINHAMENTO DE CAMPOS
# - Campos Numéricos (9) = Sempre à direita e preenchidos com zeros à esquerda.
# - Campos Alfanuméricos (X) = Sempre à esquerda e preenchidos com brancos à direita. 
#
servico: 'cobranca'
versao: 'fev2016'
layout: 'cnab400'

remessa:
  header_arquivo:
    tipo_registro:
      pos: [1,1]
      picture: '9(1)'
      default: 0
    operacao:
      pos: [2,2]
      picture: '9(1)'
      default: 1
    literal_remessa:
      pos: [3,9]
      picture: 'X(7)'
      default: 'REMESSA'
    codigo_servico:
      pos: [10,11]
      picture: '9(2)'
      default: '01'
    literal_servico:
      pos: [12,26]
      picture: 'X(15)'
      default: 'COBRANCA'
    agencia:
      pos: [27,30]
      picture: '9(4)'
    zeros_01:
      pos: [31,32]
      picture: '9(2)'
      default: 0
    conta:
      pos: [33,37]
      picture: '9(5)'
    dac:
      pos: [38,38]
      picture: '9(1)'
    brancos_01:
      pos: [39,46]
      picture: 'X(8)'
      default: ''
    nome_empresa:
      pos: [47,76]
      picture: 'X(30)'
    codigo_banco:
      pos: [77,79]
      picture: '9(3)'
      default: 341
    nome_banco:
      pos: [80,94]
      picture: 'X(15)'
      default: 'BANCO ITAU SA'
    data_geracao:
      pos: [95,100]
      picture: '9(6)'
    brancos_02:
      pos: [101,394]
      picture: 'X(294)'
      default: ''
    numero_sequencial_registro:
      pos: [395,400]
      picture: '9(6)'
      default: '000001'
  
  trailer_arquivo:
    tipo_registro:
      pos: [1,1]
      picture: '9(1)'
      default: 9
    brancos:
      pos: [2,394]
      picture: 'X(393)'
      default: ''
    numero_sequencial_registro:
      pos: [395,400]
      picture: '9(6)'
  
  detalhes:
    segmento_1:
      tipo_registro:
        pos: [1,1]
        picture: '9(1)'
        default: 1
      codigo_inscricao:
        pos: [2,3]
        picture: '9(2)'
      numero_inscricao:
        pos: [4,17]
        picture: '9(14)'
      agencia:
        pos: [18,21]
        picture: '9(4)'
      zeros_01:
        pos: [22,23]
        picture: '9(2)'
        default: 0
      conta:
        pos: [24,28]
        picture: '9(5)'
      dac:
        pos: [29,29]
        picture: '9(1)'
      brancos_01:
        pos: [30,33]
        picture: 'X(4)'
        default: ''
      instrucao_alegacao:
        pos: [34,37]
        picture: '9(4)'
      uso_empresa:
        pos: [38,62]
        picture: 'X(25)'
      nosso_numero:
        pos: [63,70]
        picture: '9(8)'
      quantidade_moeda:
        pos: [71,83]
        picture: '9(8)V9(5)'
      numero_carteira:
        pos: [84,86]
        picture: '9(3)'
      uso_banco:
        pos: [87,107]
        picture: 'X(21)'
      carteira:
        pos: [108,108]
        picture: 'X(1)'
      codigo_ocorrencia:
        pos: [109,110]
        picture: '9(2)'
      numero_documento:
        pos: [111,120]
        picture: 'X(10)'
      vencimento:
        pos: [121,126]
        picture: '9(6)'
      valor:
        pos: [127,139]
        picture: '9(11)V9(2)'
      codigo_banco:
        pos: [140,142]
        picture: '9(3)'
        default: 341
      agencia_cobradora:
        pos: [143,147]
        picture: '9(5)'
      especie:
        pos: [148,149]
        picture: 'X(2)'
      aceite:
        pos: [150,150]
        picture: 'X(1)'
        default: 'N'
      data_emissao:
        pos: [151,156]
        picture: '9(6)'
      instrucao_1:
        pos: [157,158]
        picture: 'X(2)'
      instrucao_2:
        pos: [159,160]
        picture: 'X(2)'
      juros_1_dia:
        pos: [161,173]
        picture: '9(11)V9(2)'
      desconto_ate:
        pos: [174,179]
        picture: '9(6)'
      valor_desconto:
        pos: [180,192]
        picture: '9(11)V9(2)'
      valor_iof:
        pos: [193,205]
        picture: '9(11)V9(2)'
      abatimento:
        pos: [206,218]
        picture: '9(11)V9(2)'
      codigo_inscricao_pagador:
        pos: [219,220]
        picture: '9(2)'
      numero_inscricao_pagador:
        pos: [221,234]
        picture: '9(14)'
      nome_pagador:
        pos: [235,264]
        picture: 'X(30)'
      brancos_02:
        pos: [265,274]
        picture: 'X(10)'
        default: ''
      logradouro:
        pos: [275,314]
        picture: 'X(40)'
      bairro:
        pos: [315,326]
        picture: 'X(12)'
      cep:
        pos: [327,334]
        picture: '9(8)'
      cidade:
        pos: [335,349]
        picture: 'X(15)'
      estado:
        pos: [350,351]
        picture: 'X(2)'
      sacador_avalista:
        pos: [352,381]
        picture: 'X(30)'
      brancos_03:
        pos: [382,385]
        picture: 'X(4)'
        default: ''
      data_mora:
        pos: [386,391]
        picture: '9(6)'
      prazo:
        pos: [392,393]
        picture: '9(2)'
      brancos_04:
        pos: [394,394]
        picture: 'X(1)'
        default: ''
      numero_sequencial_registro:
        pos: [395,400]
        picture: '9(6)'

    # opcional - complemento detalhe - multa
    segmento_2:
      tipo_registro:
        pos: [1,1]
        picture: '9(1)'
        default: 2
      codigo_multa:
        pos: [2,2]
        picture: 'X(1)'
      data_multa:
        pos: [3,10]
        picture: '9(8)'
      multa:
        pos: [11,23]
        picture: '9(13)'
      brancos:
        pos: [24,394]
        picture: 'X(370)'
        default: ''
      numero_sequencial_registro:
        pos: [395,400]
        picture: '9(6)'
  
    # opcional - rateio de credito
    segmento_4:
      tipo_registro:
        pos: [1,1]
        picture: '9(1)'
        default: 4
      codigo_inscricao:
        pos: [2,3]
        picture: '9(2)'
      numero_inscricao:
        pos: [4,17]
        picture: '9(14)'
      agencia:
        pos: [18,21]
        picture: '9(4)'
      zeros:
        pos: [22,23]
        picture: '9(2)'
        default: 0
      conta:
        pos: [24,28]
        picture: '9(5)'
      dac:
        pos: [29,29]
        picture: '9(1)'
      numero_carteira:
        pos: [30,32]
        picture: '9(3)'
      nosso_numero:
        pos: [33,40]
        picture: '9(8)'
      dac_nosso_numero:
        pos: [41,41]
        picture: '9(1)'
      sequencia:
        pos: [42,43]
        picture: '9(2)'
      agencia_01:
        pos: [44,47]
        picture: '9(4)'
      conta_01:
        pos: [48,54]
        picture: '9(7)'
      dac_01:
        pos: [55,55]
        picture: '9(1)'
      valor_01:
        pos: [56,68]
        picture: '9(11)V9(2)'
      agencia_02:
        pos: [69,72]
        picture: '9(4)'
      conta_02:
        pos: [73,79]
        picture: '9(7)'
      dac_02:
        pos: [80,80]
        picture: '9(1)'
      valor_02:
        pos: [81,93]
        picture: '9(11)V9(2)'
      agencia_03:
        pos: [94,97]
        picture: '9(4)'
      conta_03:
        pos: [98,104]
        picture: '9(7)'
      dac_03:
        pos: [105,105]
        picture: '9(1)'
      valor_03:
        pos: [106,118]
        picture: '9(11)V9(2)'
      agencia_04:
        pos: [119,122]
        picture: '9(4)'
      conta_04:
        pos: [123,129]
        picture: '9(7)'
      dac_04:
        pos: [130,130]
        picture: '9(1)'
      valor_04:
        pos: [131,143]
        picture: '9(11)V9(2)'
      agencia_05:
        pos: [144,147]
        picture: '9(4)'
      conta_05:
        pos: [148,154]
        picture: '9(7)'
      dac_05:
        pos: [155,155]
        picture: '9(1)'
      valor_05:
        pos: [156,168]
        picture: '9(11)V9(2)'
      agencia_06:
        pos: [169,172]
        picture: '9(4)'
      conta_06:
        pos: [173,179]
        picture: '9(7)'
      dac_06:
        pos: [180,180]
        picture: '9(1)'
      valor_06:
        pos: [181,193]
        picture: '9(11)V9(2)'
      agencia_07:
        pos: [194,197]
        picture: '9(4)'
      conta_07:
        pos: [198,204]
        picture: '9(7)'
      dac_07:
        pos: [205,205]
        picture: '9(1)'
      valor_07:
        pos: [206,218]
        picture: '9(11)V9(2)'
      agencia_08:
        pos: [219,222]
        picture: '9(4)'
      conta_08:
        pos: [223,229]
        picture: '9(7)'
      dac_08:
        pos: [230,230]
        picture: '9(1)'
      valor_08:
        pos: [231,243]
        picture: '9(11)V9(2)'
      agencia_09:
        pos: [244,247]
        picture: '9(4)'
      conta_09:
        pos: [248,254]
        picture: '9(7)'
      dac_09:
        pos: [255,255]
        picture: '9(1)'
      valor_09:
        pos: [256,268]
        picture: '9(11)V9(2)'
      agencia_10:
        pos: [269,272]
        picture: '9(4)'
      conta_10:
        pos: [273,279]
        picture: '9(7)'
      dac_10:
        pos: [280,280]
        picture: '9(1)'
      valor_10:
        pos: [281,293]
        picture: '9(11)V9(2)'
      agencia_11:
        pos: [294,297]
        picture: '9(4)'
      conta_11:
        pos: [298,304]
        picture: '9(7)'
      dac_11:
        pos: [305,305]
        picture: '9(1)'
      valor_11:
        pos: [306,318]
        picture: '9(11)V9(2)'
      agencia_12:
        pos: [319,322]
        picture: '9(4)'
      conta_12:
        pos: [323,329]
        picture: '9(7)'
      dac_12:
        pos: [330,330]
        picture: '9(1)'
      valor_12:
        pos: [331,343]
        picture: '9(11)V9(2)'
      agencia_13:
        pos: [344,347]
        picture: '9(4)'
      conta_13:
        pos: [348,354]
        picture: '9(7)'
      dac_13:
        pos: [355,355]
        picture: '9(1)'
      valor_13:
        pos: [356,368]
        picture: '9(11)V9(2)'
      agencia_14:
        pos: [369,372]
        picture: '9(4)'
      conta_14:
        pos: [373,379]
        picture: '9(7)'
      dac_14:
        pos: [380,380]
        picture: '9(1)'
      valor_14:
        pos: [381,393]
        picture: '9(11)V9(2)'
      tipo_valor:
        pos: [394,394]
        picture: '9(1)'
      numero_sequencial_registro:
        pos: [395,400]
        picture: '9(6)'

    # opcional - cobrança e-mail e/ou dados do sacador avalista
    segmento_5:
      tipo_registro:
        pos: [1,1]
        picture: '9(1)'
        default: 5
      endereco_email:
        pos: [2,121]
        picture: 'X(120)'
      codigo_inscricao:
        pos: [122,123]
        picture: '9(2)'
      numero_inscricao:
        pos: [124,137]
        picture: '9(14)'
      logradouro:
        pos: [138,177]
        picture: 'X(40)'
      bairro:
        pos: [178,189]
        picture: 'X(12)'
      cep:
        pos: [190,197]
        picture: '9(8)'
      cidade:
        pos: [198,212]
        picture: 'X(15)'
      estado:
        pos: [213,214]
        picture: 'X(2)'
      brancos:
        pos: [215,394]
        picture: 'X(180)'
        default: ''
      numero_sequencial_registro:
        pos: [395,400]
        picture: '9(6)'

retorno:
  header_arquivo:
    tipo_registro:
      pos: [1,1]
      picture: '9(1)'
      default: 0
    codigo_retorno:
      pos: [2,2]
      picture: '9(1)'
      default: 2
    literal_retorno:
      pos: [3,9]
      picture: 'X(7)'
      default: 'RETORNO'
    codigo_servico:
      pos: [10,11]
      picture: '9(2)'
      default: '01'
    literal_servico:
      pos: [12,26]
      picture: 'X(15)'
      default: 'COBRANCA'
    agencia:
      pos: [27,30]
      picture: '9(4)'
    zeros:
      pos: [31,32]
      picture: '9(2)'
      default: 0
    conta:
      pos: [33,37]
      picture: '9(5)'
    dac:
      pos: [38,38]
      picture: '9(1)'
    brancos_01:
      pos: [39,46]
      picture: 'X(8)'
      default: ''
    nome_empresa:
      pos: [47,76]
      picture: 'X(30)'
    codigo_banco:
      pos: [77,79]
      picture: '9(3)'
      default: 341
    nome_banco:
      pos: [80,94]
      picture: 'X(15)'
      default: 'BANCO ITAU SA'
    data_geracao:
      pos: [95,100]
      picture: '9(6)'
    densidade:
      pos: [101,105]
      picture: '9(5)'
    unidade_densidade:
      pos: [106,108]
      picture: 'X(3)'
      default: 'BPI'
    numero_sequencial_arquivo:
      pos: [109,113]
      picture: '9(5)'
    data_credito:
      pos: [114,119]
      picture: '9(6)'
    brancos_02:
      pos: [120,394]
      picture: 'X(275)'
      default: ''
    numero_sequencial_registro:
      pos: [395,400]
      picture: '9(6)'
      default: '000001'

  trailer_arquivo:
    tipo_registro:
      pos: [1,1]
      picture: '9(1)'
      default: 9
    codigo_retorno:
      pos: [2,2]
      picture: '9(1)'
      default: 2
    codigo_servico:
      pos: [3,4]
      picture: '9(2)'
      default: '01'
    codigo_banco:
      pos: [5,7]
      picture: '9(3)'
      default: 341
    brancos_01:
      pos: [8,17]
      picture: 'X(10)'
      default: ''
    quantidade_titulos_simples:
      pos: [18,25]
      picture: '9(8)'
    valor_total_simples:
      pos: [26,39]
      picture: '9(12)V9(2)'
    aviso_bancario_01:
      pos: [40,47]
      picture: 'X(8)'
    brancos_02:
      pos: [48,57]
      picture: 'X(10)'
      default: ''
    quantidade_titulos_vinculada:
      pos: [58,65]
      picture: '9(8)'
    valor_total_vinculada:
      pos: [66,79]
      picture: '9(12)V9(2)'
    aviso_bancario_02:
      pos: [80,87]
      picture: 'X(8)'
    brancos_03:
      pos: [88,177]
      picture: 'X(90)'
      default: ''
    quantidade_titulos_direta:
      pos: [178,185]
      picture: '9(8)'
    valor_total_direta:
      pos: [186,199]
      picture: '9(12)V9(2)'
    aviso_bancario_03:
      pos: [200,207]
      picture: 'X(8)'
    controle_arquivo:
      pos: [208,212]
      picture: '9(5)'
    quantidade_detalhes:
      pos: [213,220]
      picture: '9(8)'
    valor_total_informado:
      pos: [221,234]
      picture: '9(12)V9(2)'
    brancos_04:
      pos: [235,394]
      picture: 'X(160)'
      default: ''
    numero_sequencial_registro:
      pos: [395,400]
      picture: '9(6)'

  detalhes:
    segmento_1:
      tipo_registro:
        pos: [1,1]
        picture: '9(1)'
        default: 1
      codigo_inscricao:
        pos: [2,3]
        picture: '9(2)'
      numero_inscricao:
        pos: [4,17]
        picture: '9(14)'
      agencia:
        pos: [18,21]
        picture: '9(4)'
      zeros_01:
        pos: [22,23]
        picture: '9(2)'
        default: '01'
      conta:
        pos: [24,28]
        picture: '9(5)'
      dac:
        pos: [29,29]
        picture: '9(1)'
      brancos_01:
        pos: [30,37]
        picture: 'X(8)'
        default: ''
      uso_empresa:
        pos: [38,62]
        picture: 'X(25)'
      nosso_numero_01:
        pos: [63,70]
        picture: '9(8)'
      brancos_02:
        pos: [71,82]
        picture: 'X(12)'
        default: ''
      numero_carteira:
        pos: [83,85]
        picture: '9(3)'
      nosso_numero_02:
        pos: [86,93]
        picture: '9(8)'
      dac_nosso_numero:
        pos: [94,94]
        picture: '9(1)'
      brancos_03:
        pos: [95,107]
        picture: 'X(13)'
        default: ''
      codigo_carteira:
        pos: [108,108]
        picture: 'X(1)'
      codigo_ocorrencia:
        pos: [109,110]
        picture: '9(2)'
      data_ocorrencia:
        pos: [111,116]
        picture: '9(6)'
      numero_documento:
        pos: [117,126]
        picture: 'X(10)'
      nosso_numero_confirmacao:
        pos: [127,134]
        picture: '9(8)'
      brancos_04:
        pos: [135,146]
        picture: 'X(12)'
        default: ''
      vencimento:
        pos: [147,152]
        picture: '9(6)'
      valor_titulo:
        pos: [153,165]
        picture: '9(11)V9(2)'
      codigo_banco:
        pos: [166,168]
        picture: '9(3)'
        default: 341
      agencia_cobradora:
        pos: [169,172]
        picture: '9(4)'
      dac_agencia_cobradora:
        pos: [173,173]
        picture: '9(1)'
      especie:
        pos: [174,175]
        picture: '9(2)'
      tarifa_cobranca:
        pos: [176,188]
        picture: '9(11)V9(2)'
      brancos_05:
        pos: [189,214]
        picture: 'X(26)'
        default: ''
      valor_iof:
        pos: [215,227]
        picture: '9(11)V9(2)'
      valor_abatimento:
        pos: [228,240]
        picture: '9(11)V9(2)'
      descontos:
        pos: [241,253]
        picture: '9(11)V9(2)'
      valor_principal:
        pos: [254,266]
        picture: '9(11)V9(2)'
      juros_mora_multa:
        pos: [267,279]
        picture: '9(11)V9(2)'
      outros_creditos:
        pos: [280,292]
        picture: '9(11)V9(2)'
      boleto_dda:
        pos: [293,293]
        picture: 'X(1)'
      brancos_06:
        pos: [294,295]
        picture: 'X(2)'
        default: ''
      data_credito:
        pos: [296,301]
        picture: 'X(6)'
      instrucao_cancelada:
        pos: [302,305]
        picture: '9(4)'
      brancos_07:
        pos: [306,311]
        picture: 'X(6)'
        default: ''
      zeros_02:
        pos: [312,324]
        picture: '9(13)'
        default: 0
      nome_pagador:
        pos: [325,354]
        picture: 'X(30)'
      brancos_08:
        pos: [355,377]
        picture: 'X(23)'
        default: ''
      erros_mensagem_informativa:
        pos: [378,385]
        picture: 'X(8)'
      brancos_09:
        pos: [386,392]
        picture: 'X(7)'
        default: ''
      codigo_liquidacao:
        pos: [393,394]
        picture: 'X(2)'
      numero_sequencial_registro:
        pos: [395,400]
        picture: '9(6)'

    segmento_4:
      tipo_registro:
        pos: [1,1]
        picture: '9(1)'
        default: 4
      codigo_inscricao:
        pos: [2,3]
        picture: '9(2)'
      numero_inscricao:
        pos: [4,17]
        picture: '9(14)'
      agencia:
        pos: [18,21]
        picture: '9(4)'
      zeros:
        pos: [22,23]
        picture: '9(2)'
        default: 0
      conta:
        pos: [24,28]
        picture: '9(5)'
      dac:
        pos: [29,29]
        picture: '9(1)'
      brancos_01:
        pos: [30,37]
        picture: 'X(8)'
        default: ''
      uso_empresa:
        pos: [38,62]
        picture: 'X(25)'
      nosso_numero_01:
        pos: [63,70]
        picture: '9(8)'
      brancos_02:
        pos: [71,82]
        picture: 'X(12)'
        default: ''
      numero_carteira:
        pos: [83,85]
        picture: '9(3)'
      nosso_numero_02:
        pos: [86,93]
        picture: '9(8)'
      dac_nosso_numero:
        pos: [94,94]
        picture: '9(1)'
      brancos_03:
        pos: [95,107]
        picture: 'X(13)'
        default: ''
      codigo_carteira:
        pos: [108,108]
        picture: 'X(1)'
      codigo_ocorrencia:
        pos: [109,110]
        picture: '9(2)'
      sequencia:
        pos: [111,112]
        picture: '9(2)'
      valor_titulo:
        pos: [113,125]
        picture: '9(11)V9(2)'
      agencia_01:
        pos: [126,129]
        picture: '9(4)'
      conta_01:
        pos: [130,136]
        picture: '9(7)'
      dac_01:
        pos: [137,137]
        picture: '9(1)'
      valor_01:
        pos: [138,150]
        picture: '9(11)V9(2)'
      valor_encargos_01:
        pos: [151,160]
        picture: '9(8)V9(2)'
      agencia_02:
        pos: [161,164]
        picture: '9(4)'
      conta_02:
        pos: [165,171]
        picture: '9(7)'
      dac_02:
        pos: [172,172]
        picture: '9(1)'
      valor_02:
        pos: [173,185]
        picture: '9(11)V9(2)'
      valor_encargos_02:
        pos: [186,195]
        picture: '9(8)V9(2)'
      agencia_03:
        pos: [196,199]
        picture: '9(4)'
      conta_03:
        pos: [200,206]
        picture: '9(7)'
      dac_03:
        pos: [207,207]
        picture: '9(1)'
      valor_03:
        pos: [208,220]
        picture: '9(11)V9(2)'
      valor_encargos_03:
        pos: [221,230]
        picture: '9(8)V9(2)'
      agencia_04:
        pos: [231,234]
        picture: '9(4)'
      conta_04:
        pos: [235,241]
        picture: '9(7)'
      dac_04:
        pos: [242,242]
        picture: '9(1)'
      valor_04:
        pos: [243,255]
        picture: '9(11)V9(2)'
      valor_encargos_04:
        pos: [256,265]
        picture: '9(8)V9(2)'
      agencia_05:
        pos: [266,269]
        picture: '9(4)'
      conta_05:
        pos: [270,276]
        picture: '9(7)'
      dac_05:
        pos: [277,277]
        picture: '9(1)'
      valor_05:
        pos: [278,290]
        picture: '9(11)V9(2)'
      valor_encargos_05:
        pos: [291,300]
        picture: '9(8)V9(2)'
      agencia_06:
        pos: [301,304]
        picture: '9(4)'
      conta_06:
        pos: [305,311]
        picture: '9(7)'
      dac_06:
        pos: [312,312]
        picture: '9(1)'
      valor_06:
        pos: [313,325]
        picture: '9(11)V9(2)'
      valor_encargos_06:
        pos: [326,335]
        picture: '9(8)V9(2)'
      agencia_07:
        pos: [336,339]
        picture: '9(4)'
      conta_07:
        pos: [340,346]
        picture: '9(7)'
      dac_07:
        pos: [347,347]
        picture: '9(1)'
      valor_07:
        pos: [348,360]
        picture: '9(11)V9(2)'
      valor_encargos_07:
        pos: [361,370]
        picture: '9(8)V9(2)'
      brancos_04:
        pos: [371,393]
        picture: 'X(23)'
        default: ''
      tipo_valor:
        pos: [394,394]
        picture: 'X(1)'
      numero_sequencial_registro:
        pos: [395,400]
        picture: '9(6)'
`
