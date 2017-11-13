package funcs

import (
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
	Pkg = use.FuncMap{
		"spellMoney": spell.Money,
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
		use.Pkg{Prefix: "fmt_", Funcs: Pkg},
	)
)

// Get returns initialized map of funcs.
func Get() use.FuncMap {
	return funcs
}
