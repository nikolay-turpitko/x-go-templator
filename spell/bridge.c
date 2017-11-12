// See
// https://github.com/uber-go/icu4go
// http://userguide.icu-project.org/formatparse/numbers
// http://icu-project.org/apiref/icu4c/unum_8h.html

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "unicode/unum.h"
#include "unicode/ucurr.h"
#include "unicode/ustring.h"

#include "bridge.h"

UErrorCode spellNumber(const double a, const char* locale, char* result, const size_t resultLength)
{
    UErrorCode status = U_ZERO_ERROR;
    const int32_t bufSize = resultLength / sizeof(UChar); // considering it's UNICODE.
    int32_t needed;

    UNumberFormat *fmt = unum_open(
            UNUM_SPELLOUT,     /* style         */
            0,                 /* pattern       */
            0,                 /* patternLength */
            locale,            /* locale        */
            0,                 /* parseErr      */
            &status);

    if (!U_FAILURE(status)) {
        /* Use the formatter to format the number as a string
           in the given locale. The return value is the buffer size needed.
           We pass in NULL for the UFieldPosition pointer because we don't
           care to receive that data. */
        UChar buf[bufSize];
        needed = unum_formatDouble(fmt, a, buf, bufSize, NULL, &status);
        /**
         * u_austrcpy docs from the header:
         *
         * Copy ustring to a byte string encoded in the default codepage.
         * Adds a null terminator.
         * Performs a UChar to host byte conversion
         */
        u_austrcpy(result, buf);

        /* Release the storage used by the formatter */
        unum_close(fmt);
    }

    return status;
}
