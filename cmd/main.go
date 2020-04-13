package main

import (
	"strings"

	"github.com/helderfarias/cnab-go"
	"github.com/helderfarias/cnab-go/layout/bb"
	"github.com/helderfarias/cnab-go/output"
)

func main() {
	source := strings.NewReader(bb.CNAB240Pagamentos)
	layout, _ := cnab.NewLayout("240", source)
	remessa := cnab.NewRemessa(layout)
	debug := output.NewRemessaDebug(remessa)
	debug.Write()
}
