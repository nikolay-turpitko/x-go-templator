
all: md-pdf ace-pdf

md-pdf:
	@echo
	@echo "Generating pdf using markdown template..."
	@echo "Required tools: make, go, pandoc"
	@mkdir -p out
	@go run x-go-templator.go \
		-template sample/md/sample.tmpl \
		-data sample/md/sample.yml \
		Title="Markdown generation example" \
		| tee out/sample.md \
        | pandoc -s -S \
                -f markdown \
				--latex-engine=pdflatex \
				-V lang=ru \
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
	@go run x-go-templator.go \
		-template sample/ace/invoice.ace \
		-data sample/ace/invoice.yml \
		Number=5 \
		On=$$(date --rfc-3339=date) \
		BilledHours=$$(expr 8 \* 26) \
		AdditionalAmount=500 \
		> out/invoice.html
	@wkhtmltopdf out/invoice.html \
		out/invoice.pdf
