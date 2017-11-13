// See
// https://github.com/uber-go/icu4go
// http://userguide.icu-project.org/formatparse/numbers
// http://icu-project.org/apiref/icu4c/unum_8h.html

package spell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpellNumberUS(t *testing.T) {
	assert := assert.New(t)
	actual, err := Number("en-US", 142)
	assert.NoError(err)
	assert.Equal("one hundred forty-two", actual)
}

func TestSpellNumberRU(t *testing.T) {
	assert := assert.New(t)
	actual, err := Number("ru-RU", 142)
	assert.NoError(err)
	assert.Equal("сто сорок два", actual)
}

func TestSpellMoneyUS(t *testing.T) {
	assert := assert.New(t)
	actual, err := Money("en-US", 1977.22, RUB)
	assert.NoError(err)
	assert.Equal("One thousand nine hundred seventy-seven Russian Rubles 22 kopecks", actual)
	actual, err = Money("en-US", 1977.22, USD)
	assert.NoError(err)
	assert.Equal("One thousand nine hundred seventy-seven US Dollars 22 cents", actual)
}

func TestSpellMoneyRU(t *testing.T) {
	assert := assert.New(t)
	actual, err := Money("ru-RU", 1977.22, RUB)
	assert.NoError(err)
	assert.Equal("Одна тысяча девятьсот семьдесят семь рублей 22 копейки", actual)

	actual, err = Money("ru-RU", 1977.22, USD)
	assert.NoError(err)
	assert.Equal("Одна тысяча девятьсот семьдесят семь долларов США 22 цента", actual)

	actual, err = Money("ru-RU", 42.01, USD)
	assert.NoError(err)
	assert.Equal("Сорок два доллара США 01 цент", actual)

	actual, err = Money("ru-RU", 42.01, RUB)
	assert.NoError(err)
	assert.Equal("Сорок два рубля 01 копейка", actual)
}
