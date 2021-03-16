package model

const LayoutCNAB240 = "cnab240"
const LayoutCNAB400 = "cnab400"
const LayoutCCSITEF = "ccsitef"
const LayoutAFD = "afd"
const LayoutAFDV3 = "afdv3"
const LayoutAFDT = "afdt"
const LayoutACJEF = "acjef"
const LayoutAEJ = "aej"
const LayoutFortesPS = "ps"

var RegisterLayouts = map[string]bool{
	LayoutCNAB240:  true,
	LayoutCNAB400:  true,
	LayoutCCSITEF:  true,
	LayoutAFD:      true,
	LayoutAFDV3:    true,
	LayoutAFDT:     true,
	LayoutACJEF:    true,
	LayoutAEJ:      true,
	LayoutFortesPS: true,
}

const GlobalSettingsOrdenarEscritaPorSufixo = "sufixo_segmento"
const GlobalSettingsOrdenarEscritaPorNomeCampo = "campo_segmento"
const GlobalSettingsOrdenarEscritaPorSufixoECampo = "sufixo_e_campo_segmento"
