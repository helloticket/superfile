package superfile

import (
	"errors"
	"io"
	"io/ioutil"
	"strings"

	"github.com/helloticket/superfile/file"
	"github.com/helloticket/superfile/model"
	"github.com/helloticket/superfile/parser"
	"gopkg.in/yaml.v2"
)

func NewLayout(modelo *strings.Reader) (model.Layout, error) {
	data, err := ioutil.ReadAll(modelo)
	if err != nil {
		return nil, err
	}

	config := map[interface{}]interface{}{}
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}

	if len(config) == 0 {
		return nil, errors.New("Arquivo de definição de layout vazio")
	}

	layout, err := parser.NewModeloLayout(config)
	if err != nil {
		return nil, err
	}

	if err := layout.Validate(); err != nil {
		return nil, err
	}

	return layout, err
}

func NewRemessa(data model.Layout) *model.Remessa {
	return model.NewRemessa(data)
}

func NewRetornoFile(layout model.Layout, content io.Reader) (file.RetornoFile, error) {
	return file.NewRetornoFile(layout, content)
}

func NewRemessaDebug(remessa *model.Remessa) *file.RemessaDebug {
	return file.NewRemessaDebug(remessa)
}
