[![Build Status](https://travis-ci.com/nikolay-turpitko/x-go-templator.svg?branch=master)](https://travis-ci.com/nikolay-turpitko/x-go-templator)

# x-go-templator

**TL;DR** It's an experimental project (so, "x-" prefix in the name).
I'm doing it for my own needs and fun. PRs are welcome, but I have no time
(and intention) for providing good docs and support.

**NOTE:** When I started it some time ago, there were no so much similar tools for the task,
otherwise I'd use one. If you do not need `ace` templates or spelling numbers
into words with `icu4c` or some other custom template functions, absent in other
projects, you may like more lightweight tool, or more feature full one, like
one of these:

- https://github.com/tsg/gotpl

- https://github.com/dreadatour/go-cli-template

- https://github.com/hairyhenderson/gomplate

This project is more havyweight due to usage of `icu4c`.

**NOTE:** Finally, I found a time to create a [lighter version of similar tool for the same needs]
(https://gitlab.com/nikolay-turpitko/gotmpl) and now can abandon this one in favor of a new one.

## What is it?

**x-go-templator** is a tool to generate simple documents using templates.

I wrote it to automate creation of invoices. But it is able to create another
types of simple one or several page reports as well.

Think about it like about Go `text/template` engine plus YAML file parser
wrapped as a CLI tool. Not much more than that. But it's a quite powerful
combination, capable to replace such tools as Hugo (static web site generator)
or Jasper Reports for simple use cases (and much simpler to install and use in
scripts). I especially like to use it with ACE templates, which give concise
and light Slim-like syntax. What is more important for me, is that I can easily
extend it however I need.

You work with it as follows.

You give it a template file, a file with less volatile data in the YAML and
more volatile data via CLI args and ENV variables.

Template accesses both CLI args and YAML content via embedded scripting
language and uses them to fill report. To be more precise, tool collect all
ENV and args to the key-value maps and puts them into template's script context.

You use construction such as `{{.Args.myarg}}`, `{{.Env.HOME}}`,
`{{.Data.somekey}}`, `{{.Vars.myvar}}` to access CLI arguments, environment
variables, values from YAML file and your own script variables respectively.

To pass CLI arguments to your template, you put them in form of `key=value` on
the command line after all the program's CLI flags.

You may access complex data fields with "dot" syntax like `v.k1.k2` or with
`index` function like `index v k1 k2` (see Go docs on `text/template` for more
details).

There are some predefined functions, available within templates. You may use
them to manipulate available data.

You may invoke pre-defined functions, passing them arguments, like `myfunc arg1
arg2`. You may use functions, defined by the `text/template` or `html/template`
engines (depending on which one you use), plus bunch of functions, provided by
the tool in a kind of "packages" (similar prefix). I'm a bit lazy to create
documentation about these functions. You may find examples of usage in my
samples and use a link to source code to find out what else is available.

Just a couple of interesting usages of predefined functions:

- `math_add x y`, `math_mul a b`, ... - allow to make simple math inside
  template;
- `fmt_spellMoney "en" .Vars.total "USD"` - allows to spell money amount in
  human-readable text;
- `set myvar 42` - allows to introduce custom variable and give it a value, you
  may use it later like `.Vars.myvar` (useful within range loops);
- `os_exec "./my-script" arg1 arg2 | set var` - allows to run arbitrary OS
  process (can be "/bin/bash", for example) and use it's output;
- `os_readTxtFile "./my-file"` - allows to process text files;

Report goes to stdout, you redirect it whenever you like. I like to redirect it
to the input of `pandoc` or `wkhtmltopdf` to create PDF documents from it.

And with this you may be as creative in your templates as you like.

Program use template file extension as a hint which engine to use:

- .ace  -> "https://github.com/yosssi/ace"
- .html -> "html/template"
- .\*   -> "html/text"

**BTW**, project illustrates:

- quite involved usage of template engines (with custom functions, scripting and access to ENV and CLI vars);
- usege of icu4c lib via cgo binding;
- spelling numbers to English and Russian words using icu4c;
- usage of Slim-like markup language for templates;
- generation of pdf docs from html templates using pandoc and wkhtmltopdf;
- usage of custom fonts in html (giving nice pdf);
- actual invoice template I use for Russian exchange currency control and bank;
- usage of glide and travis;

Though, it's not rocket science, and I brought all the pieces from elsewhere,
I just listed it here for myself to know where to find it when I need it again.

## How to use?

### Build & Install

    go get -d github.com/nikolay-turpitko/x-go-templator
    cd $GOPATH/src/github.com/nikolay-turpitko/x-go-templator
    glide i
    .build/install-icu4c.sh
    make clean test build
    # or, to install it into $GOPATH/bin
    make install

Refer .travis.yml for build requirements and steps.

**Note:** to install ICU lib, you may use script `.build/install-icu4c.sh`.
But I had to remove existing icu-devtools installation, because of conflict
(`sudo apt purge icu-devtools`).

I borrowed this script and ICU lib binding approach from
https://github.com/uber-go/icu4go. Refer it for details.

**Note:** pandoc, texlive and wkhtmltopdf are not build requirements, these
tools used by the script as illustration of possible workflow.

**Note:** I experimented with TeX sample (`make md-pdf`) on Travis for an hours
and was finally able to run it only with `lualatex` engine, though it works
perfectly well at my local Ubuntu 16.04 with either `pdflatex`, `xelatex` or
`lualatex`.

Issues I had with different engines on Travis:
- pdflatex: babel don't accept Russian language setting;
- xelatex: fails with some strange error (TeX capacity exceeded, sorry [input
  stack size=5000]).

Currently I'm not going to provide pre-build binary or package, it's too many
hustle for such a simple tool.

### Create a template file and (optionally) a data file

    vim my-template.ace my-data.yml

You may use plain text files, html or ace files as a template (see samples).
You may use some pre-defined functions in your templates (see samples).
You may use environment variables or variables, defined in the command line in
your templates (see samples).

### Use a tool to generate document based on your template and data file

    x-go-templator -h

    x-go-templator -template my-template.ace -data my-data.yml MY_VAR="some data" > my-doc.html

Use samples to get an idea.

BTW, Makefile in this project can be viewed not only as build tool, but also as
a tool usage example.  Take a look at it to have an idea how to generate pdf
files and use command line arguments in the templates.

## Additional documentation:

### About templates:

- https://golang.org/pkg/text/template/
- https://golang.org/pkg/html/template/
- https://github.com/yosssi/ace

### About YAML:

- http://www.yaml.org/start.html

### About additional functions in templates

I borrowed them from my other project. There is no much text, but functions are
listed along with the source code, and they are mostly just wrappers around Go
functions, so usage should be clear.  You may also use usage samples in that
project to have an idea how to use these functions in your templates.

- https://godoc.org/github.com/nikolay-turpitko/structor/funcs
- https://github.com/nikolay-turpitko/structor/blob/master/funcs_test.go

### How to use custom fonts for html templates with wkhtmltopd

- https://stackoverflow.com/questions/10611828/use-custom-fonts-with-wkhtmltopdf/16972315#16972315

Namely, I used https://www.fontsquirrel.com/tools/webfont-generator, uploaded
one of free fonts from Ubuntu's /usr/share/fonts/truetype dir, choose an expert
mode and removed all unused font types and char sets. Than I had to fix font
family names in downloaded css.

## Some random ideas (TODO):

- [ ] add template function to load data from csv file (+sample);
- [ ] add template function to load data from DB connection (+sample);
- [ ] add template function to read data from stdin (without `os_exec "cat"`);
- [ ] add flag to change template delims (to use tex template);
- [ ] try to generate tex file from template;

