#ifndef LAS_READ_ITEM_H
#define LAS_READ_ITEM_H

#include "mydefs.hpp"

class ByteStreamIn;

class LASreadItem
{
public:
  virtual void read(U8* item)=0;

  virtual ~LASreadItem(){};
};

class LASreadItemRaw : public LASreadItem
{
public:
  LASreadItemRaw()
  {
    instream = 0;
  };
  BOOL init(ByteStreamIn* instream)
  {
    if (!instream) return FALSE;
    this->instream = instream;
    return TRUE;
  };
  virtual ~LASreadItemRaw(){};
protected:
  ByteStreamIn* instream;
};

class LASreadItemCompressed : public LASreadItem
{
public:
  virtual BOOL init(const U8* item)=0;

  virtual ~LASreadItemCompressed(){};
};

#endif
