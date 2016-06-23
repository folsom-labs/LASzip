#ifndef INTEGER_COMPRESSOR_HPP
#define INTEGER_COMPRESSOR_HPP

#include "arithmeticdecoder.hpp"

class IntegerCompressor
{
public:

  IntegerCompressor(ArithmeticDecoder* dec, U32 bits, U32 contexts);
  ~IntegerCompressor();

  void initDecompressor();
  I32 decompress(I32 iPred, U32 context=0);

  // Get the k corrector bits from the last compress/decompress call
  U32 getK() const {return k;};

private:
  void writeCorrector(I32 c, ArithmeticModel* model);
  I32 readCorrector(ArithmeticModel* model);

  U32 k;

  U32 contexts;
  U32 bits_high;

  U32 bits;
  U32 range;

  U32 corr_bits;
  U32 corr_range;
  I32 corr_min;
  I32 corr_max;

  ArithmeticDecoder* dec;

  ArithmeticModel** mBits;

  ArithmeticModel** mCorrector;
};

#endif
