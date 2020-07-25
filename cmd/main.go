package main

import (
	"strings"

	"github.com/helloticket/superfile"
	"github.com/helloticket/superfile/layout/mte"
)

func main() {
	source := strings.NewReader(mte.ACJEF)
	layout, err := superfile.NewLayout(source)
	if err != nil {
		panic(err)
	}

	remessa := superfile.NewRemessa(layout)
	debug := superfile.NewRemessaDebug(remessa)
	debug.Write()
}
