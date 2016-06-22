#ifndef MYDEFS_HPP
#define MYDEFS_HPP

typedef char               CHAR;

typedef int                I32;
typedef short              I16;
typedef char               I8;

typedef unsigned int       U32;
typedef unsigned short     U16;
typedef unsigned char      U8;

#if defined(_WIN32) && ! defined (__MINGW32__) // 64 byte integer under Windows
typedef unsigned __int64   U64;
typedef __int64            I64;
#else                                          // 64 byte integer elsewhere ...
typedef unsigned long long U64;
typedef long long          I64;
#endif

typedef float              F32;
typedef double             F64;

#if defined(_MSC_VER) && (_MSC_VER < 1300) || defined (__MINGW32__)
typedef int                BOOL;
#else
typedef bool               BOOL;
#endif

typedef union U32I32F32 { U32 u32; I32 i32; F32 f32; } U32I32F32;
typedef union U64I64F64 { U64 u64; I64 i64; F64 f64; } U64I64F64;

#define F32_MAX            +2.0e+37f
#define F32_MIN            -2.0e+37f

#define F64_MAX            +2.0e+307
#define F64_MIN            -2.0e+307

#define U8_MIN             ((U8)0x0)  // 0
#define U8_MAX             ((U8)0xFF) // 255
#define U8_MAX_MINUS_ONE   ((U8)0xFE) // 254
#define U8_MAX_PLUS_ONE    0x0100     // 256

#define U16_MIN            ((U16)0x0)    // 0
#define U16_MAX            ((U16)0xFFFF) // 65535
#define U16_MAX_MINUS_ONE  ((U16)0xFFFE) // 65534
#define U16_MAX_PLUS_ONE   0x00010000    // 65536

#define U32_MIN            ((U32)0x0)            // 0
#define U32_MAX            ((U32)0xFFFFFFFF)     // 4294967295
#define U32_MAX_MINUS_ONE  ((U32)0xFFFFFFFE)     // 4294967294
#if defined(WIN32)            // 64 byte unsigned int constant under Windows
#define U32_MAX_PLUS_ONE   0x0000000100000000    // 4294967296
#else                         // 64 byte unsigned int constant elsewhere ...
#define U32_MAX_PLUS_ONE   0x0000000100000000ull // 4294967296
#endif

#define I8_MIN             ((I8)0x80) // -128
#define I8_MAX             ((I8)0x7F) // 127

#define I16_MIN            ((I16)0x8000) // -32768
#define I16_MAX            ((I16)0x7FFF) // 32767

#define I32_MIN            ((I32)0x80000000) // -2147483648
#define I32_MAX            ((I32)0x7FFFFFFF) //  2147483647

#define I64_MIN            ((I64)0x8000000000000000)
#define I64_MAX            ((I64)0x7FFFFFFFFFFFFFFF)

#define U8_FOLD(n)      (((n) < U8_MIN) ? (n+U8_MAX_PLUS_ONE) : (((n) > U8_MAX) ? (n-U8_MAX_PLUS_ONE) : (n)))

#define I8_CLAMP(n)     (((n) <= I8_MIN) ? I8_MIN : (((n) >= I8_MAX) ? I8_MAX : ((I8)(n))))
#define U8_CLAMP(n)     (((n) <= U8_MIN) ? U8_MIN : (((n) >= U8_MAX) ? U8_MAX : ((U8)(n))))

#define I16_CLAMP(n)    (((n) <= I16_MIN) ? I16_MIN : (((n) >= I16_MAX) ? I16_MAX : ((I16)(n))))
#define U16_CLAMP(n)    (((n) <= U16_MIN) ? U16_MIN : (((n) >= U16_MAX) ? U16_MAX : ((U16)(n))))

#define I32_CLAMP(n)    (((n) <= I32_MIN) ? I32_MIN : (((n) >= I32_MAX) ? I32_MAX : ((I32)(n))))
#define U32_CLAMP(n)    (((n) <= U32_MIN) ? U32_MIN : (((n) >= U32_MAX) ? U32_MAX : ((U32)(n))))

#define I8_QUANTIZE(n) (((n) >= 0) ? (I8)((n)+0.5f) : (I8)((n)-0.5f))
#define U8_QUANTIZE(n) (((n) >= 0) ? (U8)((n)+0.5f) : (U8)(0))

#define I16_QUANTIZE(n) (((n) >= 0) ? (I16)((n)+0.5f) : (I16)((n)-0.5f))
#define U16_QUANTIZE(n) (((n) >= 0) ? (U16)((n)+0.5f) : (U16)(0))

#define I32_QUANTIZE(n) (((n) >= 0) ? (I32)((n)+0.5f) : (I32)((n)-0.5f))
#define U32_QUANTIZE(n) (((n) >= 0) ? (U32)((n)+0.5f) : (U32)(0))

#define I64_QUANTIZE(n) (((n) >= 0) ? (I64)((n)+0.5f) : (I64)((n)-0.5f))
#define U64_QUANTIZE(n) (((n) >= 0) ? (U64)((n)+0.5f) : (U64)(0))

#define I16_FLOOR(n) ((((I16)(n)) > (n)) ? (((I16)(n))-1) : ((I16)(n)))
#define I32_FLOOR(n) ((((I32)(n)) > (n)) ? (((I32)(n))-1) : ((I32)(n)))
#define I64_FLOOR(n) ((((I64)(n)) > (n)) ? (((I64)(n))-1) : ((I64)(n)))

#define I16_CEIL(n) ((((I16)(n)) < (n)) ? (((I16)(n))+1) : ((I16)(n)))
#define I32_CEIL(n) ((((I32)(n)) < (n)) ? (((I32)(n))+1) : ((I32)(n)))
#define I64_CEIL(n) ((((I64)(n)) < (n)) ? (((I64)(n))+1) : ((I64)(n)))

#define I8_FITS_IN_RANGE(n) (((n) >= I8_MIN) && ((n) <= I8_MAX) ? TRUE : FALSE)
#define U8_FITS_IN_RANGE(n) (((n) >= U8_MIN) && ((n) <= U8_MAX) ? TRUE : FALSE)
#define I16_FITS_IN_RANGE(n) (((n) >= I16_MIN) && ((n) <= I16_MAX) ? TRUE : FALSE)
#define U16_FITS_IN_RANGE(n) (((n) >= U16_MIN) && ((n) <= U16_MAX) ? TRUE : FALSE)

#define F32_IS_FINITE(n) ((F32_MIN < (n)) && ((n) < F32_MAX))
#define F64_IS_FINITE(n) ((F64_MIN < (n)) && ((n) < F64_MAX))

#define U32_ZERO_BIT_0(n) (((n)&(U32)0xFFFFFFFE))

#ifndef FALSE
#define FALSE   0
#endif

#ifndef TRUE
#define TRUE    1
#endif

#endif
