#ifndef LAS_QUANTIZER_HPP
#define LAS_QUANTIZER_HPP

#include "mydefs.hpp"

class LASquantizer
{
public:
  F64 x_scale_factor;
  F64 y_scale_factor;
  F64 z_scale_factor;
  F64 x_offset;
  F64 y_offset;
  F64 z_offset;

  inline F64 get_x(const I32 X) const { return x_scale_factor*X+x_offset; };
  inline F64 get_y(const I32 Y) const { return y_scale_factor*Y+y_offset; };
  inline F64 get_z(const I32 Z) const { return z_scale_factor*Z+z_offset; };

  inline I32 get_X(const F64 x) const { if (x >= x_offset) return (I32)((x-x_offset)/x_scale_factor+0.5); else return (I32)((x-x_offset)/x_scale_factor-0.5); };
  inline I32 get_Y(const F64 y) const { if (y >= y_offset) return (I32)((y-y_offset)/y_scale_factor+0.5); else return (I32)((y-y_offset)/y_scale_factor-0.5); };
  inline I32 get_Z(const F64 z) const { if (z >= z_offset) return (I32)((z-z_offset)/z_scale_factor+0.5); else return (I32)((z-z_offset)/z_scale_factor-0.5); };

  LASquantizer()
  {
    x_scale_factor = 0.01;
    y_scale_factor = 0.01;
    z_scale_factor = 0.01;
    x_offset = 0.0;
    y_offset = 0.0;
    z_offset = 0.0;
  };

  LASquantizer & operator=(const LASquantizer & quantizer)
  {
    this->x_scale_factor = quantizer.x_scale_factor;
    this->y_scale_factor = quantizer.y_scale_factor;
    this->z_scale_factor = quantizer.z_scale_factor;
    this->x_offset = quantizer.x_offset;
    this->y_offset = quantizer.y_offset;
    this->z_offset = quantizer.z_offset;
    return *this;
  };
};

#endif
