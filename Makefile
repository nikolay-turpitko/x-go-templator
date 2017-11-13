# Below line is for Mac
# export ICU_LIB = /usr/local/opt/icu4c

# Below line is tested on Ubuntu 16.04
#export ICU_LIB = /usr
export CGO_CFLAGS += -I${ICU_LIB}/include
export CGO_LDFLAGS += -L${ICU_LIB}/lib -licui18n -licuuc -licudata

all: clean test build md-pdf ace-pdf

clean:
	@-rm ./x-go-templator 2> /dev/null || :
@-rm -r ./out 2> /dev/null || :

test:
	@go test -v $$(glide nv)

build:
	@go build

md-pdf:
	@echo
	@echo "Generating pdf using markdown template..."
	@echo "Required tools: make, go, pandoc"
	@mkdir -p out
	@./x-go-templator \
		-template sample/md/sample.tmpl \
		-data sample/md/sample.yml \
		Title="Markdown generation example" \
		| tee out/sample.md \
		| pandoc -s -S \
		-f markdown \
		--latex-engine=xelatex \
		-V lang="russian" \
		-V otherlangs="english" \
		-V mainfont="Linux Libertine O" \
		-V sansfont="Linux Libertine O" \
		-V monofont="Linux Libertine Mono O" \
		-V documentclass="paper" \
		-V papersize="a4" \
		-V geometry:margin=1.5cm \
		-V fontsize="10pt" \
		-V colorlinks=true \
		-V links-as-notes=true \
		-V title-meta="PDF generation sample (from markdown template)" \
		-V author-meta="Nikolay Turpitko" \
		-o out/sample-md.pdf
ace-pdf:
	@echo
	@echo "Generating pdf (via html) using ace template..."
	@echo "Required tools: make, go, wkhtmltopdf"
	@mkdir -p out
	@cp sample/ace/stylesheet.css out
	@./x-go-templator \
		-template sample/ace/invoice.ace \
		-data sample/ace/invoice.yml \
		Number=5 \
		On=$$(date --rfc-3339=date) \
		BilledHours=$$(expr 8 \* 26) \
		AdditionalAmount=500 \
		> out/invoice.html
	@wkhtmltopdf out/invoice.html \
		out/invoice.pdf
