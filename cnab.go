package cnab

import (
	"fmt"
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

	layout, err := parser.NewModeloLayout(config)
	if err != nil {
		return nil, err
	}

	esperado := strings.ToLower(versao)
	if !strings.HasPrefix(esperado, "cnab") {
		esperado = fmt.Sprintf("cnab%s", versao)
	}

	retornado := layout.GetLayout()
	if esperado != retornado {
		return nil, fmt.Errorf("Versão do layout não compativel. Esperado %s mas foi %s", versao, esperado)
	}

	return layout, err
}

func NewRemessa(data model.Layout) *model.Remessa {
	return model.NewRemessa(data)
}
