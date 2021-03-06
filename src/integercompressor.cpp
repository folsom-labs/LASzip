#include "integercompressor.hpp"

#include <stdlib.h>
#include <assert.h>

IntegerCompressor::IntegerCompressor(ArithmeticDecoder* dec, U32 bits, U32 contexts)
{
  crashif(!dec);
  crashif(bits == 0);

  this->dec = dec;
  this->bits = bits;
  this->contexts = contexts;
  this->bits_high = 8;
  this->range = 0;

  if (bits < 32)
  {
    corr_bits = bits;
    corr_range = 1u << bits;
		// the corrector must fall into this interval
    corr_min = -((I32)(corr_range/2));
  	corr_max = corr_min + corr_range - 1;
  }
	else
	{
    corr_bits = 32;
		corr_range = 0;
		// the corrector must fall into this interval
    corr_min = I32_MIN;
    corr_max = I32_MAX;
	}

  k = 0;

  mBits = nullptr;
  mCorrector = nullptr;
  mCorrector0 = nullptr;
}

IntegerCompressor::~IntegerCompressor()
{
  U32 i;
  if (mBits)
  {
    for (i = 0; i < contexts; i++)
    {
      dec->destroySymbolModel(mBits[i]);
    }
    delete [] mBits;
  }

  if (mCorrector)
  {
    crashif(!mCorrector0);
    dec->destroyBitModel(mCorrector0);
    for (i = 1; i <= corr_bits; i++)
    {
      dec->destroySymbolModel(mCorrector[i]);
    }
    delete [] mCorrector;
  }
}

void IntegerCompressor::initDecompressor()
{
  U32 i;

  assert(dec);

  // maybe create the models
  if (mBits == 0)
  {
    mBits = new ArithmeticModel*[contexts];
    for (i = 0; i < contexts; i++)
    {
      mBits[i] = dec->createSymbolModel(corr_bits+1);
    }

    mCorrector = new ArithmeticModel*[corr_bits+1];
    mCorrector0 = dec->createBitModel();
    for (i = 1; i <= corr_bits; i++)
    {
      if (i <= bits_high)
      {
        mCorrector[i] = dec->createSymbolModel(1<<i);
      }
      else
      {
        mCorrector[i] = dec->createSymbolModel(1<<bits_high);
      }
    }
  }

  // certainly init the models
  for (i = 0; i < contexts; i++)
  {
    dec->initSymbolModel(mBits[i]);
  }

  dec->initBitModel(mCorrector0);
  for (i = 1; i <= corr_bits; i++)
  {
    dec->initSymbolModel(mCorrector[i]);
  }
}

I32 IntegerCompressor::decompress(I32 pred, U32 context)
{
  assert(dec);
  I32 real = pred + readCorrector(mBits[context]);
  if (real < 0) real += corr_range;
  else if ((U32)(real) >= corr_range) real -= corr_range;
  return real;
}

I32 IntegerCompressor::readCorrector(ArithmeticModel* mBits)
{
  I32 c;

  // decode within which interval the corrector is falling

  k = dec->decodeSymbol(mBits);

  // decode the exact location of the corrector within the interval

  if (k) // then c is either smaller than 0 or bigger than 1
  {
    if (k < 32)
    {
      if (k <= bits_high) // for small k we can do this in one step
      {
        // decompress c with the range coder
        c = dec->decodeSymbol(mCorrector[k]);
      }
      else
      {
        // for larger k we need to do this in two steps
        int k1 = k-bits_high;
        // decompress higher bits with table
        c = dec->decodeSymbol(mCorrector[k]);
        // read lower bits raw
        int c1 = dec->readBits(k1);
        // put the corrector back together
        c = (c << k1) | c1;
      }
      // translate c back into its correct interval
      if (c >= (1<<(k-1))) // if c is in the interval [ 2^(k-1)  ...  + 2^k - 1 ]
      {
        // so we translate c back into the interval [ 2^(k-1) + 1  ...  2^k ] by adding 1
        c += 1;
      }
      else // otherwise c is in the interval [ 0 ...  + 2^(k-1) - 1 ]
      {
        // so we translate c back into the interval [ - (2^k - 1)  ...  - (2^(k-1)) ] by subtracting (2^k - 1)
        c -= ((1<<k) - 1);
      }
    }
    else
    {
      c = corr_min;
    }
  }
  else // then c is either 0 or 1
  {
    c = dec->decodeBit(mCorrector0);
  }

  return c;
}
