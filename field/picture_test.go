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

func TestDecodeString(t *testing.T) {
	pic := picture{}

	assert.Equal(t, "Jose da Si", noError(pic.Decode("Jose da Si", "X(13)", Empty)))
	assert.Equal(t, nil, noError(pic.Decode("", "X(13)", Empty)))
}

func TestDecodeDataHora(t *testing.T) {
	pic := picture{}

	assert.Equal(t, time.Date(2017, 5, 10, 0, 0, 0, 0, time.UTC), noError(pic.Decode("10052017", "9(8)", Options{"data_format": "ddmmyyyy"})))
	assert.Equal(t, time.Date(2017, 5, 10, 0, 0, 0, 0, time.UTC), noError(pic.Decode("100517", "9(6)", Options{"data_format": "dmy"})))
	assert.Equal(t, time.Date(2017, 5, 10, 0, 0, 0, 0, time.UTC), noError(pic.Decode("10052017", "9(8)", Options{"data_format": "dmY"})))
	assert.Equal(t, time.Date(0, 1, 1, 12, 56, 10, 0, time.UTC), noError(pic.Decode("125610", "9(8)", Options{"data_format": "hhmmss"})))
	assert.Equal(t, time.Date(0, 1, 1, 12, 56, 0, 0, time.UTC), noError(pic.Decode("1256", "9(8)", Options{"data_format": "hhmm"})))
	assert.Equal(t, nil, noError(pic.Decode("", "9(8)", Options{"data_format": "hhmm"})))
}

func TestDecodeNumericoSimples(t *testing.T) {
	pic := picture{}

	assert.Equal(t, int64(100010), noError(pic.Decode("100010", "9(15)", Empty)))
	assert.Equal(t, int64(1000000005689), noError(pic.Decode("1000000005689", "9(32)", Empty)))
	assert.Equal(t, int64(1), noError(pic.Decode("1", "9(1)", Empty)))
	assert.Equal(t, int64(10), noError(pic.Decode("10", "9(2)", Empty)))
	assert.Equal(t, nil, noError(pic.Decode("", "9(1)", Empty)))
	assert.Equal(t, int64(10), noError(pic.Decode("10", "9(2)", Empty)))
}

func TestDecodeNumericoComVirgula(t *testing.T) {
	pic := picture{}

	assert.Equal(t, float64(52501.56), noError(pic.Decode("5250156", "9(5)V9(2)", Empty)))
	assert.Equal(t, float64(10.201), noError(pic.Decode("10201", "9(2)V9(3)", Empty)))
	assert.Equal(t, float64(0.001), noError(pic.Decode("1", "9(5)V9(3)", Empty)))
	assert.Equal(t, float64(0), noError(pic.Decode("0", "9(5)V9(3)", Empty)))
	assert.Equal(t, nil, noError(pic.Decode("", "9(5)V9(3)", Empty)))
}
