package field

import (
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestEncodeString(t *testing.T) {
	pic := picture{}

	assert.NotNil(t, pic)
	assert.Equal(t, "Jose da Si   ", noError(pic.Encode("Jose da Si", "X(13)", Options{})))
	assert.Equal(t, "             ", noError(pic.Encode(nil, "X(13)", Options{})))
}

func TestEncodeDataHora(t *testing.T) {
	pic := picture{}

	data1 := time.Date(2017, 5, 10, 12, 56, 10, 0, time.UTC)

	assert.Equal(t, "10052017", noError(pic.Encode(data1, "9(8)", Options{"data_format": "ddmmyyyy"})))
	assert.Equal(t, "10052017", noError(pic.Encode(data1, "9(8)", Options{"data_format": "dmY"})))
	assert.Equal(t, "100517", noError(pic.Encode(data1, "9(6)", Options{"data_format": "dmy"})))
	assert.Equal(t, "10052017", noError(pic.Encode(data1, "9(8)", Options{})))
	assert.Equal(t, "100517", noError(pic.Encode(data1, "9(6)", Options{})))
	assert.Equal(t, "125610", noError(pic.Encode(data1, "9(6)", Options{"data_format": "hhmmss"})))
	assert.Equal(t, "1256", noError(pic.Encode(data1, "9(4)", Options{"data_format": "hhmm"})))
	assert.Equal(t, "0000", noError(pic.Encode(nil, "9(4)", Options{"data_format": "hhmm"})))
}

func TestEncodeNumericoSimples(t *testing.T) {
	pic := picture{}

	assert.Equal(t, "000000000010001", noError(pic.Encode(float64(1000.10), "9(15)", Options{})))
	assert.Equal(t, "00000000000000000001000000005689", noError(pic.Encode(float64(100000000.5689), "9(32)", Options{})))
	assert.Equal(t, "1", noError(pic.Encode(float64(1), "9(1)", Options{})))
	assert.Equal(t, "0", noError(pic.Encode(float64(10), "9(1)", Options{})))
	assert.Equal(t, "0", noError(pic.Encode(nil, "9(1)", Options{})))
	assert.Equal(t, "010", noError(pic.Encode("10", "9(3)", Options{})))
}

func TestEncodeNumericoComVirgula(t *testing.T) {
	pic := picture{}

	assert.Equal(t, "5250156", noError(pic.Encode(float64(52501.56), "9(5)V9(2)", Options{})))
	assert.Equal(t, "00010201", noError(pic.Encode(float64(10.201), "9(5)V9(3)", Options{})))
	assert.Equal(t, "00001000", noError(pic.Encode(float64(1), "9(5)V9(3)", Options{})))
	assert.Equal(t, "00000000", noError(pic.Encode(float64(0), "9(5)V9(3)", Options{})))
	assert.Equal(t, "00000000", noError(pic.Encode(nil, "9(5)V9(3)", Options{})))
	assert.Equal(t, "52501.56", noError(pic.Encode(float64(52501.56), "9(5)V9(2)", Options{"decimal_separator": "."})))
	assert.Equal(t, "52501,56", noError(pic.Encode(float64(52501.56), "9(5)V9(2)", Options{"decimal_separator": ","})))
	assert.Equal(t, "5250156", noError(pic.Encode(float64(52501.56), "9(5)V9(2)", Options{"decimal_separator": " "})))
	assert.Equal(t, "00000.00", noError(pic.Encode(float64(0), "9(5)V9(2)", Options{"decimal_separator": "."})))
}

func TestDecodeString(t *testing.T) {
	pic := picture{}

	assert.Equal(t, "Jose da Si", noError(pic.Decode("Jose da Si", "X(13)", Options{})))
	assert.Equal(t, nil, noError(pic.Decode("", "X(13)", Options{})))
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

	assert.Equal(t, int64(100010), noError(pic.Decode("100010", "9(15)", Options{})))
	assert.Equal(t, int64(1000000005689), noError(pic.Decode("1000000005689", "9(32)", Options{})))
	assert.Equal(t, int64(1), noError(pic.Decode("1", "9(1)", Options{})))
	assert.Equal(t, int64(10), noError(pic.Decode("10", "9(2)", Options{})))
	assert.Equal(t, nil, noError(pic.Decode("", "9(1)", Options{})))
	assert.Equal(t, int64(10), noError(pic.Decode("10", "9(2)", Options{})))
}

func TestDecodeNumericoComVirgula(t *testing.T) {
	pic := picture{}

	assert.Equal(t, float64(52501.56), noError(pic.Decode("5250156", "9(5)V9(2)", Options{})))
	assert.Equal(t, float64(10.201), noError(pic.Decode("10201", "9(2)V9(3)", Options{})))
	assert.Equal(t, float64(0.001), noError(pic.Decode("1", "9(5)V9(3)", Options{})))
	assert.Equal(t, float64(0), noError(pic.Decode("0", "9(5)V9(3)", Options{})))
	assert.Equal(t, nil, noError(pic.Decode("", "9(5)V9(3)", Options{})))
}

func TestEncodeStringGlobalSettings(t *testing.T) {
	viper.Set("global_alinhamento_alfanumerico", "X")
	defer viper.Reset()

	pic := picture{}

	assert.Equal(t, "0XXXXXXXXXXXX", noError(pic.Encode("0", "X(13)", Options{"global_alinhamento_alfanumerico": "X"})))
}

func TestDecodeStringGlobalSettings(t *testing.T) {
	viper.Set("global_alinhamento_alfanumerico", "X")
	defer viper.Reset()

	pic := picture{}

	assert.Equal(t, "1200", noError(pic.Decode("1200XXXX", "X(8)", Options{"global_alinhamento_alfanumerico": "X"})))
	assert.Nil(t, noError(pic.Decode("XXXXXXXX", "X(8)", Options{"global_alinhamento_alfanumerico": "X"})))
}
