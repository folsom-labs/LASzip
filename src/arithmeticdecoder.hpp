#ifndef ARITHMETIC_DECODER_HPP
#define ARITHMETIC_DECODER_HPP

#include "mydefs.hpp"
#include "bytestreamin.hpp"

class ArithmeticModel;
class ArithmeticBitModel;

class ArithmeticDecoder
{
public:

/* Constructor & Destructor                                  */
  ArithmeticDecoder();
  ~ArithmeticDecoder();

/* Manage decoding                                           */
  BOOL init(ByteStreamIn* instream);
  void done();

/* Manage an entropy model for a single bit                  */
  ArithmeticBitModel* createBitModel();
  void initBitModel(ArithmeticBitModel* model);
  void destroyBitModel(ArithmeticBitModel* model);

/* Manage an entropy model for n symbols (table optional)    */
  ArithmeticModel* createSymbolModel(U32 n);
  void initSymbolModel(ArithmeticModel* model, U32* table=0);
  void destroySymbolModel(ArithmeticModel* model);

/* Decode a bit with modelling                               */
  U32 decodeBit(ArithmeticBitModel* model);

/* Decode a symbol with modelling                            */
  U32 decodeSymbol(ArithmeticModel* model);

/* Decode a bit without modelling                            */
  U32 readBit();

/* Decode bits without modelling                             */
  U32 readBits(U32 bits);

/* Decode an unsigned char without modelling                 */
  U8 readByte();

/* Decode an unsigned short without modelling                */
  U16 readShort();

/* Decode an unsigned int without modelling                  */
  U32 readInt();

/* Decode an unsigned 64 bit int without modelling           */
  U64 readInt64();

private:

  ByteStreamIn* instream;

  void renorm_dec_interval();
  U32 value, length;
};

#endif
