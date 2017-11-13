[![Build Status](https://travis-ci.org/nikolay-turpitko/x-go-templator.svg?branch=master)](https://travis-ci.org/nikolay-turpitko/x-go-templator)

# x-go-templator

## What?

**x-go-templator** is a tool to generate simple documents using templates.

I wrote it to automate creation of invoices. But it is able to create another
types of simple one or several page reports as well.

**TL;DR** It's an experimental project (so, "x-" prefix in the name).
I'm doing it for my own needs and fun. PRs are welcome, but I have no time
(and intention) for providing good docs and support.

Think about it like about Go `text/template` engine plus YAML file parser
wrapped as a CLI tool. Not much more than that.

You give it a template file, a file with less volatile data in the YAML and
more volatile data via CLI args and ENV variables.

Template accesses both CLI args and YAML content via embedded scripting
language and uses them to fill report.

Report goes to stdout, you redirect it whenever you like.

There are some predefined functions, available within templates. You may use
them to manipulate available data.

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

    go get github.com/nikolay-turpitko/x-go-templator

Refer .travis.yml for build requrements and steps.

**Note:** to install ICU lib, you may use script `.build/install-icu4c.sh`.  But I
had to remove existing icu-devtools installation, because of conflict (`sudo
apt purge icu-devtools`).

I borrowed this script and ICU lib binding approach from
https://github.com/uber-go/icu4go. Refer it for details.

**Note:** pandoc, texlive and wkhtmltopdf are not build requrements, these
tools used by the script as illustration of possible workflow.

TODO:
- check if it goes to go/bin
- probably, add a link to binary, built by Travis?

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

### How to find which cyrillic (T2A) fonts do you have in your local TeX

    # find installed fonts
    updmap-sys --listmaps | cut -f 2 | cut -f 1 -d . | sort

    # find t2a fonts
    find /usr/share/texlive/ -path "**/t2a*"

    # TODO: don't know how automate. Look at them and try to guess...

    # use name from first output to check pdf doc and find T2A in it
    texdoc cantarell


