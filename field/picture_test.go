package field

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEncodeString(t *testing.T) {
	pic := picture{}

	assert.NotNil(t, pic)
	assert.Equal(t, "Jose da Si   ", noError(pic.Encode("Jose da Si", "X(13)", Empty)))
	assert.Equal(t, "             ", noError(pic.Encode(nil, "X(13)", Empty)))
}

func TestEncodeDataHora(t *testing.T) {
	pic := picture{}

	data1 := time.Date(2017, 5, 10, 12, 56, 10, 0, time.UTC)

	assert.Equal(t, "10052017", noError(pic.Encode(data1, "9(8)", Options{"data_format": "ddmmyyyy"})))
	assert.Equal(t, "10052017", noError(pic.Encode(data1, "9(8)", Options{"data_format": "dmY"})))
	assert.Equal(t, "100517", noError(pic.Encode(data1, "9(6)", Options{"data_format": "dmy"})))
	assert.Equal(t, "10052017", noError(pic.Encode(data1, "9(8)", Empty)))
	assert.Equal(t, "100517", noError(pic.Encode(data1, "9(6)", Empty)))
	assert.Equal(t, "125610", noError(pic.Encode(data1, "9(6)", Options{"data_format": "hhmmss"})))
	assert.Equal(t, "1256", noError(pic.Encode(data1, "9(4)", Options{"data_format": "hhmm"})))
	assert.Equal(t, "0000", noError(pic.Encode(nil, "9(4)", Options{"data_format": "hhmm"})))
}

func TestEncodeNumericoSimples(t *testing.T) {
	pic := picture{}

	assert.Equal(t, "000000000010001", noError(pic.Encode(float64(1000.10), "9(15)", Empty)))
	assert.Equal(t, "00000000000000000001000000005689", noError(pic.Encode(float64(100000000.5689), "9(32)", Empty)))
	assert.Equal(t, "1", noError(pic.Encode(float64(1), "9(1)", Empty)))
	assert.Equal(t, "0", noError(pic.Encode(float64(10), "9(1)", Empty)))
	assert.Equal(t, "0", noError(pic.Encode(nil, "9(1)", Empty)))
	assert.Equal(t, "010", noError(pic.Encode("10", "9(3)", Empty)))
}

func TestEncodeNumericoComVirgula(t *testing.T) {
	pic := picture{}

	assert.Equal(t, "5250156", noError(pic.Encode(float64(52501.56), "9(5)V9(2)", Empty)))
	assert.Equal(t, "00010201", noError(pic.Encode(float64(10.201), "9(5)V9(3)", Empty)))
	assert.Equal(t, "00001000", noError(pic.Encode(float64(1), "9(5)V9(3)", Empty)))
	assert.Equal(t, "00000000", noError(pic.Encode(float64(0), "9(5)V9(3)", Empty)))
	assert.Equal(t, "00000000", noError(pic.Encode(nil, "9(5)V9(3)", Empty)))
}

func noError(val string, err error) string {
	if err != nil {
		return err.Error()
	}
	return val
}
