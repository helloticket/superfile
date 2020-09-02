package main

import (
	"strings"

	"github.com/helloticket/superfile"
	"github.com/helloticket/superfile/layout/bb"
)

func main() {
	source := strings.NewReader(bb.CNAB240ExtratoContaCorrente)
	layout, err := superfile.NewLayout(source)
	if err != nil {
		panic(err)
	}

	remessa := superfile.NewRemessa(layout)
	debug := superfile.NewRemessaDebug(remessa)
	debug.Write()
}
