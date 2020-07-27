package main

import (
	"strings"

	"github.com/helloticket/superfile"
	"github.com/helloticket/superfile/layout/fortes"
)

func main() {
	source := strings.NewReader(fortes.ImportacaoMovimento)
	layout, err := superfile.NewLayout(source)
	if err != nil {
		panic(err)
	}

	remessa := superfile.NewRemessa(layout)
	debug := superfile.NewRemessaDebug(remessa)
	debug.Write()
}
