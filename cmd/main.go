package main

import (
	"strings"

	"github.com/helderfarias/superfile"
	"github.com/helderfarias/superfile/layout/ccsitef"
)

func main() {
	source := strings.NewReader(ccsitef.ExtratoEletronico)
	layout, err := superfile.NewLayout(source)
	if err != nil {
		panic(err)
	}

	remessa := superfile.NewRemessa(layout)
	debug := superfile.NewRemessaDebug(remessa)
	debug.Write()
}
