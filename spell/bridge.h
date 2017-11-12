// See
// https://github.com/uber-go/icu4go
// http://userguide.icu-project.org/formatparse/numbers
// http://icu-project.org/apiref/icu4c/unum_8h.html

#ifndef __C_NUMBER_BRIDGE_H__
#define __C_NUMBER_BRIDGE_H__

#include <stdlib.h>
#include <stdbool.h>
#include "unicode/utypes.h"

UErrorCode spellNumber(
        const double a,
        const char* locale,
        char* result,
        const size_t resultLength
        );

#endif //__C_NUMBER_BRIDGE_H__
