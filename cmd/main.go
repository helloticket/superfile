package main

import (
	"strings"

	"github.com/helderfarias/superfile"
	"github.com/helderfarias/superfile/layout/itau"
)

func main() {
	source := strings.NewReader(itau.CNAB400Cobranca)
	layout, err := superfile.NewLayout(source)
	if err != nil {
		panic(err)
	}

	remessa := superfile.NewRemessa(layout)
	debug := superfile.NewRemessaDebug(remessa)
	debug.Write()
}
