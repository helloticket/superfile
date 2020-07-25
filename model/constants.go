package model

const LayoutCNAB240 = "cnab240"
const LayoutCNAB400 = "cnab400"
const LayoutCCSITEF = "ccsitef"
const LayoutAFD = "afd"
const LayoutAFDT = "afdt"
const LayoutACJEF = "acjef"

var RegisterLayouts = map[string]bool{
	LayoutCNAB240: true,
	LayoutCNAB400: true,
	LayoutCCSITEF: true,
	LayoutAFD:     true,
	LayoutAFDT:    true,
	LayoutACJEF:   true,
}
