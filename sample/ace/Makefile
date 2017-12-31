.PHONY: usage invoice

# Usage samples:
#
# make n=5 date=2017-12-01 d=20
# make h=160
# make d=20 h=4
# make w="Sa|Su"
#
# where
# n - invoice number
# date - invoice date
# d - working days (per 8 hours)
# h - working hours
# fa - fixed amount
# w - weekends (regexp of names like in `cal`)

usage:
	@echo -e "\n \
		Usage:\n \
		\n \
		make invoice [n=N] [date=YYYY-MM-DD] [d=N] [h=N] [w=Sa] [fa=N]\n \
		\n \
		where: \n \
		n - invoice number (default calculated by perviouse invoice file)\n \
		date - invoice date (default is first day of the next month)\n \
		d - working days (per 8 hours)\n \
		h - working hours\n \
		w - weekends (regexp of names like in 'cal')\n \
		fa - fixed amount\n \
		\n \
		Billed hours are calculated as (d*8 + h).\n \
		Default for d, h and fa is 0.\n \
		If w is provided, then d is calculated as number of wokring days in the month (preceeding invoice date).\n \
		Working days calculated as number of days in month minus nuber of weekend days.\n \
		\n \
		Example:\n \
		\n \
		make invoice n=7 w=\"Sa|Su\"\n"

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

# Check if weekends are provided and calculate working days if so.
# Note: working days are calculated for the month before invoice date.
# For example, if invoice is prepared as of 01 Dec, then Nov is used.
ifdef w
override d := $(shell m=$$(date -d "-1 day ${date}" +"%m %Y"); expr $$(cal $$m | awk 'NF {DAYS = $$NF}; END {print DAYS}') - $$(ncal -M $$m | awk '/$w/ {n +=NF-1}; END {print n}'))
# To use the same month as in the invoice:
#override d := $(shell m=$$(date -d "${date}" +"%m %Y"); expr $$(cal $$m | awk 'NF {DAYS = $$NF}; END {print DAYS}') - $$(ncal -M $$m | awk '/$w/ {n +=NF-1}; END {print n}'))
endif


# Check if number of working days or working hours provided
# Convert working days to working hours if necessary
h ?= 0
d ?= 0
fa ?= 0
override h := $(shell expr $d \* 8 + $h)

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
