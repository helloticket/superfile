package cnab

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCriarLayout(t *testing.T) {
	layout, err := NewLayout("240", strings.NewReader(""))

	assert.Nil(t, err)
	assert.NotNil(t, layout)
}

func TestCriarRemessa(t *testing.T) {
	remessa := NewRemessa(&LayoutFake{})

	assert.NotNil(t, remessa)
}
