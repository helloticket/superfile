package main

import (
	"strings"

	"github.com/helderfarias/cnab-go"
	"github.com/helderfarias/cnab-go/file"
	"github.com/helderfarias/cnab-go/layout/itau"
)

func main() {
	source := strings.NewReader(itau.CNAB240Cobranca)
	layout, _ := cnab.NewLayout("240", source)
	remessa := cnab.NewRemessa(layout)
	debug := file.NewRemessaDebug(remessa)
	debug.Write()
}
