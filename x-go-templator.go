package main

import (
	"flag"
	"fmt"
	htmltemplate "html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	texttemplate "text/template"

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
	acetemplate "github.com/yosssi/ace"
	"gopkg.in/yaml.v2"
)

func main() {

	programName := filepath.Base(os.Args[0])
	flag.Usage = func() {
		fmt.Fprintf(
			os.Stderr,
			`
Usage of %[1]s:

%[1]s [FLAGS]

This tool generates document using template and data file.

Program uses file extansions as hint to decide which template engine to use and
how to parse input data:

.ace  -> "https://github.com/yosssi/ace"
.html -> "html/template"
.*    -> "text/template"

Currently, the only supported format for data files is YAML.
If "-" (dash simbol) in the "-data" flag means data is to be read from stdin.

Flags supported by the program:

`,
			programName)
		flag.PrintDefaults()
	}

	log.Println("Parsing flags...")

	args := struct {
		templateFileName string
		dataFileName     string
	}{}

	flag.StringVar(&args.templateFileName, "template", "template.html", "Template of the document")
	flag.StringVar(&args.dataFileName, "data", "-", "Data file")
	flag.Parse()

	var err error
	args.templateFileName, err = filepath.EvalSymlinks(args.templateFileName)
	if err != nil {
		log.Fatal("Cannot parse template file name. ", err)
	}
	args.dataFileName, err = filepath.EvalSymlinks(args.dataFileName)
	if err != nil {
		log.Fatal("Cannot parse data file name. ", err)
	}
	log.Printf("Arguments: %+v", args)
	log.Println("Parsing data...")

	data := make(map[string]interface{})
	if args.dataFileName != "" {
		var (
			b   []byte
			err error
		)
		if args.dataFileName == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b, err = ioutil.ReadFile(args.dataFileName)
		}
		if err != nil {
			log.Fatal("Cannot read data stream. ", err)
		}
		err = yaml.Unmarshal(b, &data)
		if err != nil {
			log.Fatal("Cannot parse data. ", err)
		}
	}

	ctx := struct {
		Args map[string]string
		Env  map[string]string
		Data interface{}            // Data from the file (yaml)
		Vars map[string]interface{} // Tmp vars for use in the template
	}{
		Args: parseProperties(flag.Args()),
		Env:  parseProperties(os.Environ()),
		Data: data,
		Vars: map[string]interface{}{},
	}

	log.Printf(
		"Parsed %d cmd arg(s), %d env var(s), %d data key(s)",
		len(ctx.Args),
		len(ctx.Env),
		len(data))
	log.Println("Generating output file...")

	funcs := use.Packages(
		use.Pkg{Prefix: "b_", Funcs: bytes.Pkg},
		use.Pkg{Prefix: "c_", Funcs: crypt.Pkg},
		use.Pkg{Prefix: "e_", Funcs: encoding.Pkg},
		use.Pkg{Prefix: "g_", Funcs: goquery.Pkg},
		use.Pkg{Prefix: "m_", Funcs: math.Pkg},
		use.Pkg{Prefix: "o_", Funcs: funcs_os.Pkg},
		use.Pkg{Prefix: "r_", Funcs: regexp.Pkg},
		use.Pkg{Prefix: "s_", Funcs: funcs_str.Pkg},
		use.Pkg{Prefix: "x_", Funcs: xpath.Pkg},
	)

	funcs["set"] = func(k string, v interface{}) string {
		ctx.Vars[k] = v
		return ""
	}

	type executor interface {
		Execute(wr io.Writer, data interface{}) error
	}

	var t executor

	ext := filepath.Ext(args.templateFileName)
	log.Printf("Extension = <%s>", filepath.Ext(args.templateFileName))
	switch ext {
	case ".ace":
		log.Println("Use ace engine.")
		base := strings.TrimSuffix(args.templateFileName, ext)
		log.Printf("Base name = <%s>", base)
		t, err = acetemplate.Load(base, "", &acetemplate.Options{
			FuncMap: htmltemplate.FuncMap(funcs),
		})
	case ".html":
		log.Println("Use html engine.")
		tpl, e := htmltemplate.ParseFiles(args.templateFileName)
		err = e
		if err == nil {
			t = tpl.Funcs(htmltemplate.FuncMap(funcs))
		}
	default:
		log.Println("Use html engine.")
		tpl, e := texttemplate.ParseFiles(args.templateFileName)
		err = e
		if err == nil {
			t = tpl.Funcs(texttemplate.FuncMap(funcs))
		}
	}
	if err != nil {
		log.Fatal("Cannot parse template. ", err)
	}
	err = t.Execute(os.Stdout, ctx)
	if err != nil {
		log.Fatal("Cannot execute template. ", err)
	}

	log.Println("Complete.")
}

func parseProperties(a []string) map[string]string {
	p := make(map[string]string, len(a))
	for _, s := range a {
		part := strings.Split(s, "=")
		if len(part) < 2 {
			p[s] = s
			continue
		}
		p[part[0]] = part[1]
	}
	return p
}
