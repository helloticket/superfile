package cnab

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCriarLayoutInvalido(t *testing.T) {
	layout, err := NewLayout("240", strings.NewReader(""))

	assert.NotNil(t, err)
	assert.Nil(t, layout)
}

func TestCriarRemessa(t *testing.T) {
	remessa := NewRemessa(&LayoutFake{})

	assert.NotNil(t, remessa)
}
