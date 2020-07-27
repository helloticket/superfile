package file

import (
	"errors"
	"io"
	"os"

	"github.com/helloticket/superfile/model"
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

	if remessa.GetLayout() == model.LayoutCCSITEF {
		return &remessaCCSITEFFile{
			model:    remessa,
			fileName: fileName,
			encoder:  NewEncoder(remessa),
		}
	}

	if remessa.GetLayout() == model.LayoutAFD {
		return &remessaAFDFile{
			model:    remessa,
			fileName: fileName,
			encoder:  NewEncoder(remessa),
		}
	}

	if remessa.GetLayout() == model.LayoutAFDT {
		return &remessaAFDTFile{
			model:    remessa,
			fileName: fileName,
			encoder:  NewEncoder(remessa),
		}
	}

	if remessa.GetLayout() == model.LayoutACJEF {
		return &remessaACJEFFile{
			model:    remessa,
			fileName: fileName,
			encoder:  NewEncoder(remessa),
		}
	}

	if remessa.GetLayout() == model.LayoutFortesPS {
		return &remessaFortesPSFile{
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
			decoder: NewDecoder(layout),
		}, nil
	}

	if layout.GetLayout() == model.LayoutCCSITEF {
		return &retornoCCSITEFFile{
			layout:  layout,
			content: content,
			decoder: NewDecoder(layout),
		}, nil
	}

	if layout.GetLayout() == model.LayoutAFD {
		return &retornoAFDFile{
			layout:  layout,
			content: content,
			decoder: NewDecoder(layout),
		}, nil
	}

	if layout.GetLayout() == model.LayoutAFDT {
		return &retornoAFDTFile{
			layout:  layout,
			content: content,
			decoder: NewDecoder(layout),
		}, nil
	}

	if layout.GetLayout() == model.LayoutACJEF {
		return &retornoAFDTFile{
			layout:  layout,
			content: content,
			decoder: NewDecoder(layout),
		}, nil
	}

	if layout.GetLayout() == model.LayoutFortesPS {
		return &retornoFortesPSFile{
			layout:  layout,
			content: content,
			decoder: NewDecoder(layout),
		}, nil
	}

	return &retornoCNAB240File{
		layout:  layout,
		content: content,
		decoder: NewDecoder(layout),
	}, nil
}
