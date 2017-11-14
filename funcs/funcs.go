package funcs

import (
	"time"

	"github.com/nikolay-turpitko/structor/funcs/bytes"
	"github.com/nikolay-turpitko/structor/funcs/crypt"
	"github.com/nikolay-turpitko/structor/funcs/encoding"
	"github.com/nikolay-turpitko/structor/funcs/goquery"
	funcs_os "github.com/nikolay-turpitko/structor/funcs/os"
	"github.com/nikolay-turpitko/structor/funcs/regexp"
	funcs_str "github.com/nikolay-turpitko/structor/funcs/strings"
	"github.com/nikolay-turpitko/structor/funcs/use"
	"github.com/nikolay-turpitko/structor/funcs/xpath"

	"github.com/nikolay-turpitko/x-go-templator/funcs/math"
	"github.com/nikolay-turpitko/x-go-templator/spell"
)

var (
	fmtPkg = use.FuncMap{
		"spellMoney": spell.Money,
	}

	timePkg = use.FuncMap{
		"now":   time.Now,
		"parse": time.Parse,
		"const": func() map[string]interface{} {
			return map[string]interface{}{
				"ANSIC":       time.ANSIC,
				"UnixDate":    time.UnixDate,
				"RubyDate":    time.RubyDate,
				"RFC822":      time.RFC822,
				"RFC822Z":     time.RFC822Z,
				"RFC850":      time.RFC850,
				"RFC1123":     time.RFC1123,
				"RFC1123Z":    time.RFC1123Z,
				"RFC3339":     time.RFC3339,
				"RFC3339Nano": time.RFC3339Nano,
				"Kitchen":     time.Kitchen,
				"Stamp":       time.Stamp,
				"StampMilli":  time.StampMilli,
				"StampMicro":  time.StampMicro,
				"StampNano":   time.StampNano,

				"RFC3339Date": "2006-01-02",
				"DateUS":      "01/02/2006",
				"DateRU":      "02.01.2006",
			}
		},
	}

	funcs = use.Packages(
		use.Pkg{Prefix: "bytes_", Funcs: bytes.Pkg},
		use.Pkg{Prefix: "crypt_", Funcs: crypt.Pkg},
		use.Pkg{Prefix: "enc_", Funcs: encoding.Pkg},
		use.Pkg{Prefix: "gq_", Funcs: goquery.Pkg},
		use.Pkg{Prefix: "math_", Funcs: math.Pkg},
		use.Pkg{Prefix: "os_", Funcs: funcs_os.Pkg},
		use.Pkg{Prefix: "regex_", Funcs: regexp.Pkg},
		use.Pkg{Prefix: "str_", Funcs: funcs_str.Pkg},
		use.Pkg{Prefix: "xpath_", Funcs: xpath.Pkg},
		use.Pkg{Prefix: "fmt_", Funcs: fmtPkg},
		use.Pkg{Prefix: "time_", Funcs: timePkg},
	)
)

// Get returns initialized map of funcs.
func Get() use.FuncMap {
	return funcs
}
