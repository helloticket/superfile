package main

import (
	"strings"

	"github.com/helderfarias/cnab-go"
	"github.com/helderfarias/cnab-go/file"
	"github.com/helderfarias/cnab-go/layout/itau"
	"github.com/helderfarias/cnab-go/model"
)

func main() {
	source := strings.NewReader(itau.CNAB400Cobranca)

	layout, err := cnab.NewLayout(model.LayoutCNAB400, source)
	if err != nil {
		panic(err)
	}

	remessa := cnab.NewRemessa(layout)
	debug := file.NewRemessaDebug(remessa)
	debug.Write()
}
