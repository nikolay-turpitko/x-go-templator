package funcs

import (
	"github.com/nikolay-turpitko/structor/funcs/bytes"
	"github.com/nikolay-turpitko/structor/funcs/crypt"
	"github.com/nikolay-turpitko/structor/funcs/encoding"
	"github.com/nikolay-turpitko/structor/funcs/goquery"
	"github.com/nikolay-turpitko/structor/funcs/math"
	funcs_os "github.com/nikolay-turpitko/structor/funcs/os"
	"github.com/nikolay-turpitko/structor/funcs/regexp"
	funcs_str "github.com/nikolay-turpitko/structor/funcs/strings"
	"github.com/nikolay-turpitko/structor/funcs/use"
	"github.com/nikolay-turpitko/structor/funcs/xpath"

	"github.com/nikolay-turpitko/x-go-templator/spell"
)

var (
	Pkg = use.FuncMap{
		"spellCurrency": spell.Currency,
	}

	funcs = use.Packages(
		use.Pkg{Prefix: "b_", Funcs: bytes.Pkg},
		use.Pkg{Prefix: "c_", Funcs: crypt.Pkg},
		use.Pkg{Prefix: "e_", Funcs: encoding.Pkg},
		use.Pkg{Prefix: "g_", Funcs: goquery.Pkg},
		use.Pkg{Prefix: "m_", Funcs: math.Pkg},
		use.Pkg{Prefix: "o_", Funcs: funcs_os.Pkg},
		use.Pkg{Prefix: "r_", Funcs: regexp.Pkg},
		use.Pkg{Prefix: "s_", Funcs: funcs_str.Pkg},
		use.Pkg{Prefix: "x_", Funcs: xpath.Pkg},
		use.Pkg{Funcs: Pkg}, // no prefix
	)
)

// Get returns initialized map of funcs.
func Get() use.FuncMap {
	return funcs
}
