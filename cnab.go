package cnab

import (
	"io/ioutil"
	"strings"

	"github.com/helderfarias/cnab-go/model"
	"github.com/helderfarias/cnab-go/parser"
	"gopkg.in/yaml.v2"
)

func NewLayout(versao string, modelo *strings.Reader) (model.Layout, error) {
	data, err := ioutil.ReadAll(modelo)
	if err != nil {
		return nil, err
	}

	config := map[interface{}]interface{}{}
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}

	return parser.NewModeloLayout(config), nil
}

func NewRemessa(data model.Layout) *model.Remessa {
	return model.NewRemessa(data)
}
