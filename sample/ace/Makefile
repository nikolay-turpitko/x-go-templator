
# Usage samples:
#
# make n=5 date=2017-12-01 d=20
# make h=160
# make d=20 h=4
#
# where
# n - invoice number
# date - invoice date
# d - working days (per 8 hours)
# h - working hours
# fa - fixed amount

SHELL := /bin/bash

# Find a file with name like "Invoice-NN.pdf", where NN is like 01, 02, ...
# Extract number NN from it's name
# Increment NN
# Use it as a new invoice number
nn := $(shell v=$$(find . -name "Invoice-*.pdf" -printf "%f\n" | sort | tail -n 1); echo $${v//[^0-9]/})
nn := $(shell printf %02d $$(expr ${nn} + 1))
n ?= ${nn}

# Calculate the date of the first day of the next month
# Use it as a new invoice date, if date is not provided
date ?= $(shell date -d "+1 month -$$(($$(date +%d)-1)) days" --rfc-3339=date)

# Check if number of working days or working hours provided
# Convert working days to working hours if necessary
h ?= 0
d ?= 0
fa ?= 0
override h := $(shell echo $$(expr $d \* 8 + $h))

invoice:
	@echo "About to generate invoice #$n as of ${date}, billed $h hours, plus fixed ${fa} USD"
	@read -p "Press ENTER to continue or Ctrl+C to cancel..." discard
	@echo "Generating invoice..."
	@x-go-templator \
		-template ./invoice.ace \
		-data ./invoice.yml \
		Number=$n \
		On=${date} \
		BilledHours=$h \
		AdditionalAmount=${fa} \
		> Invoice-$n.html
	@wkhtmltopdf Invoice-$n.html \
		Invoice-$n.pdf
	@rm Invoice-$n.html
