package field

import (
	"fmt"
	"math"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/helderfarias/superfile/helper"
)

type Picture interface {
	Encode(target interface{}, format string, opts Options) (string, error)

	Decode(target string, format string, opts Options) (interface{}, error)
}

type picture struct {
}

type Options map[string]string

var Empty = Options{}

var dateFormat = map[string]string{
	"ddmmyy":   "020106",
	"ddmmyyyy": "02012006",
	"dmY":      "02012006",
	"dmy":      "020106",
	"hhmmss":   "150405",
	"hhmm":     "1504",
	"yyyymmdd": "20060201",
}

func NewPicture() Picture {
	return &picture{}
}

func (p *picture) Decode(target string, pic string, opts Options) (interface{}, error) {
	match, err := p.createPattern(pic)
	if err != nil {
		panic(err)
	}

	if strings.TrimSpace(target) == "" {
		return nil, nil
	}

	if match["tipo1"] == "X" && match["tipo2"] == "" {
		return strings.TrimSpace(target), nil
	}

	if match["tipo1"] == "9" && match["tipo2"] == "" {
		if opts["data_format"] != "" {
			date, err := time.Parse(dateFormat[opts["data_format"]], strings.TrimSpace(target))
			if err != nil {
				return nil, err
			}
			return date, nil
		}

		size, err := helper.ToInt(match["tamanho1"])
		if err != nil {
			return "", err
		}

		i := big.NewInt(int64(size))
		i.SetString(target, 10)
		return i.Int64(), nil
	}

	if match["tipo1"] == "9" && match["tipo2"] == "V9" {
		leftSize, err := helper.ToInt(match["tamanho1"])
		if err != nil {
			return nil, err
		}

		rightSize, err := helper.ToInt(match["tamanho2"])
		if err != nil {
			return nil, err
		}

		value := target

		size := leftSize + rightSize
		if size > len(target) {
			value = helper.LeftPad(target, "0", size)
		}

		valorLeft := big.NewFloat(float64(leftSize))
		valorLeft.SetString(value[0:leftSize])

		valorRight := big.NewFloat(float64(rightSize))
		valorRight.SetString(value[leftSize:])

		if calc, _ := valorRight.Float64(); calc > 0 {
			i := new(big.Float)
			i.SetString(fmt.Sprintf("%s.%s", value[0:leftSize], value[leftSize:]))
			total, _ := i.Float64()
			return total, nil
		}

		total, _ := valorLeft.Float64()
		return total, nil
	}

	return nil, nil
}

func (p *picture) Encode(target interface{}, pic string, opts Options) (string, error) {
	match, err := p.createPattern(pic)
	if err != nil {
		panic(err)
	}

	if match["tipo1"] == "X" && match["tipo2"] == "" {
		if target == nil {
			size, err := helper.ToInt(match["tamanho1"])
			if err != nil {
				return "", err
			}

			return helper.RightPad("", " ", size), nil
		}

		value, ok := target.(string)
		if !ok {
			return "", fmt.Errorf("Format '%s' incorreto para valor '%v'", match["tipo1"], target)
		}

		size, err := helper.ToInt(match["tamanho1"])
		if err != nil {
			return "", err
		}

		value = value[0:helper.Min(size, len(value))]
		return helper.RightPad(value, " ", size), nil
	}

	if date, ok := target.(time.Time); ok && match["tipo1"] == "9" {
		if opts["data_format"] != "" {
			size, err := helper.ToInt(match["tamanho1"])
			if err != nil {
				return "", err
			}

			value := date.Format(dateFormat[opts["data_format"]])
			return helper.LeftPad(value, "0", size), nil
		}

		size, err := helper.ToInt(match["tamanho1"])
		if err != nil {
			return "", err
		}

		if size == 8 {
			value := date.Format(dateFormat["dmY"])
			return helper.LeftPad(value, "0", size), nil
		}

		if size == 6 {
			value := date.Format(dateFormat["dmy"])
			return helper.LeftPad(value, "0", size), nil
		}
	}

	if match["tipo1"] == "9" && match["tipo2"] == "" {
		value, err := helper.FormatNumber(target)
		if err != nil {
			return "", err
		}

		size, err := helper.ToInt(match["tamanho1"])
		if err != nil {
			return "", err
		}

		value = strings.Replace(value, ".", "", -1)
		return helper.LeftPad(value, "0", size), nil
	}

	if match["tipo1"] == "9" && match["tipo2"] == "V9" {
		value, err := helper.FormatNumber(target)
		if err != nil {
			return "", err
		}

		exp := strings.Split(value, ".")
		if len(exp) == 1 {
			exp = append(exp, "0")
		}

		leftSize, err := helper.ToInt(match["tamanho1"])
		if err != nil {
			return "", err
		}

		leftValue := helper.LeftPad(exp[0], "0", leftSize)

		rightSize, err := helper.ToInt(match["tamanho2"])
		if err != nil {
			return "", err
		}

		if len(exp[1]) > rightSize {
			extra := len(exp[1]) - rightSize
			extraPow := math.Pow10(extra)

			expPos, err := helper.ToInt(exp[1])
			if err != nil {
				return "", err
			}

			extraRounded := math.Round(float64(expPos) / float64(extraPow))
			exp[1] = strconv.FormatFloat(extraRounded, 'f', -1, 64)
		}
		rightValue := helper.RightPad(exp[1], "0", rightSize)

		return leftValue + rightValue, nil
	}

	return "", nil
}

func (p *picture) createPattern(format string) (map[string]string, error) {
	regexFormat := `(?P<tipo1>X|9)\((?P<tamanho1>[0-9]+)\)(?P<tipo2>(V9)?)\(?(?P<tamanho2>([0-9]+)?)\)?`

	pattern, err := regexp.Compile(regexFormat)
	if err != nil {
		return map[string]string{}, err
	}

	if !pattern.MatchString(format) {
		return map[string]string{}, fmt.Errorf("'%s' is not a valid format", format)
	}

	m := map[string]string{}

	tipo1 := pattern.SubexpNames()[1]
	if tipo1 != "" {
		m[tipo1] = pattern.ReplaceAllString(format, fmt.Sprintf("${%s}", tipo1))
	}

	tamanho1 := pattern.SubexpNames()[2]
	if tamanho1 != "" {
		m[tamanho1] = pattern.ReplaceAllString(format, fmt.Sprintf("${%s}", tamanho1))
	}

	tipo2 := pattern.SubexpNames()[3]
	if tipo2 != "" {
		m[tipo2] = pattern.ReplaceAllString(format, fmt.Sprintf("${%s}", tipo2))
	}

	tamanho2 := pattern.SubexpNames()[5]
	if tamanho2 != "" {
		m[tamanho2] = pattern.ReplaceAllString(format, fmt.Sprintf("${%s}", tamanho2))
	}

	return m, nil
}
