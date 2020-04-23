package file

import (
	"errors"
	"io"
	"os"

	"github.com/helderfarias/superfile/model"
)

type RetornoFile interface {
	Read() *model.Retorno
}

type RemessaFile interface {
	Write() *os.File
}

func NewRemessaFile(remessa *model.Remessa, fileName string) RemessaFile {
	if remessa.GetLayout() == model.LayoutCNAB400 {
		return &remessaCNAB400File{
			model:    remessa,
			fileName: fileName,
			encoder:  NewEncoder(remessa),
		}
	}

	return &remessaCNAB240File{
		model:    remessa,
		fileName: fileName,
		encoder:  NewEncoder(remessa),
	}
}

func NewRetornoFile(layout model.Layout, content io.Reader) (RetornoFile, error) {
	if layout == nil {
		return nil, errors.New("Layout não definido tratar retorno de arquivo")
	}

	if err := layout.Validate(); err != nil {
		return nil, err
	}

	if layout.GetLayout() == model.LayoutCNAB400 {
		return &retornoCNAB400File{
			layout:  layout,
			content: content,
			decoder: NewDecoder(),
		}, nil
	}

	return &retornoCNAB240File{
		layout:  layout,
		content: content,
		decoder: NewDecoder(),
	}, nil
}
