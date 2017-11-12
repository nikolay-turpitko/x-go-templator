// See
// https://github.com/uber-go/icu4go
// http://userguide.icu-project.org/formatparse/numbers
// http://icu-project.org/apiref/icu4c/unum_8h.html

package spell

// #include "unicode/utypes.h"
// #include "bridge.h"
import "C"
import (
	"fmt"
	"math"
	"unicode"
	"unicode/utf8"
	"unsafe"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/ru"
	"github.com/go-playground/universal-translator"
)

// Number is function to spell a number for a locale
func Number(l string, n float64) (string, error) {
	const bufSize = 512
	a := C.double(n)
	locale := C.CString(l)
	defer func() { C.free(unsafe.Pointer(locale)) }()

	resSize := C.size_t(bufSize * C.sizeof_char)
	res := (*C.char)(C.malloc(resSize))
	defer func() { C.free(unsafe.Pointer(res)) }()

	err := C.spellNumber(a, locale, res, resSize)
	if err > 0 {
		return "", fmt.Errorf("error (%d) spelling a number %f", err, n)
	}
	return C.GoString(res), nil
}

// CurrencyCode is a type representing currencies, known by the package.
type CurrencyCode string

const (
	USD CurrencyCode = "USD"
	RUB CurrencyCode = "RUB"
)

var universalTraslator *ut.UniversalTranslator

func init() {
	e := en.New()
	universalTraslator = ut.New(e, e, ru.New())

	t, _ := universalTraslator.GetTranslator("en")
	t.AddCardinal(USD, "{0} US Dollar", locales.PluralRuleOne, false)
	t.AddCardinal(USD, "{0} US Dollars", locales.PluralRuleOther, false)
	t.AddCardinal(RUB, "{0} Russian Ruble", locales.PluralRuleOne, false)
	t.AddCardinal(RUB, "{0} Russian Rubles", locales.PluralRuleOther, false)
	t.AddCardinal(USD+"/100", "{0} cent", locales.PluralRuleOne, false)
	t.AddCardinal(USD+"/100", "{0} cents", locales.PluralRuleOther, false)
	t.AddCardinal(RUB+"/100", "{0} kopeck", locales.PluralRuleOne, false)
	t.AddCardinal(RUB+"/100", "{0} kopecks", locales.PluralRuleOther, false)

	t, _ = universalTraslator.GetTranslator("ru")
	t.AddCardinal(USD, "{0} доллар США", locales.PluralRuleOne, false)
	t.AddCardinal(USD, "{0} доллара США", locales.PluralRuleFew, false)
	t.AddCardinal(USD, "{0} долларов США", locales.PluralRuleMany, false)
	t.AddCardinal(USD, "{0} долларов США", locales.PluralRuleOther, false)
	t.AddCardinal(RUB, "{0} рубль", locales.PluralRuleOne, false)
	t.AddCardinal(RUB, "{0} рубля", locales.PluralRuleFew, false)
	t.AddCardinal(RUB, "{0} рублей", locales.PluralRuleMany, false)
	t.AddCardinal(RUB, "{0} рублей", locales.PluralRuleOther, false)
	t.AddCardinal(USD+"/100", "{0} цент", locales.PluralRuleOne, false)
	t.AddCardinal(USD+"/100", "{0} цента", locales.PluralRuleFew, false)
	t.AddCardinal(USD+"/100", "{0} центов", locales.PluralRuleMany, false)
	t.AddCardinal(USD+"/100", "{0} центов", locales.PluralRuleOther, false)
	t.AddCardinal(RUB+"/100", "{0} копейка", locales.PluralRuleOne, false)
	t.AddCardinal(RUB+"/100", "{0} копейки", locales.PluralRuleFew, false)
	t.AddCardinal(RUB+"/100", "{0} копеек", locales.PluralRuleMany, false)
	t.AddCardinal(RUB+"/100", "{0} копеек", locales.PluralRuleOther, false)
}

// Currency returns spelling of the amount in selected locale and currency.
func Currency(
	l string,
	n float64,
	curr CurrencyCode) (translation string, err error) {

	// Implementation note: Probably, I should have used ICU here as well,
	// instead of "universal-translator", but I was not able to find
	// documentation quick enough and experimenting with Go code is a lot
	// simpler.

	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	i, f := math.Modf(n)
	s, err := Number(l, i)
	if err != nil {
		return "", err
	}
	t, _ := universalTraslator.GetTranslator(l[:2])
	major, err := t.C(curr, i, 0, s)
	if err != nil {
		return "", err
	}
	ff := math.Trunc(100*f + 0.5)
	minor, err := t.C(curr+"/100", ff, 0, fmt.Sprintf("%02.0f", ff))
	return upperFirst(fmt.Sprintf("%s %s", major, minor)), nil
}

func upperFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}
