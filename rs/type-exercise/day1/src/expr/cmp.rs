use std::cmp::Ordering;

use crate::array::Array;

pub fn cmp_le<'a, I1: Array, I2: Array, C: Array>(i1: I1::RefItem<'a>, i2: I2::RefItem<'a>) -> bool
where
    I1::RefItem<'a>: Into<C::RefItem<'a>>,
    I2::RefItem<'a>: Into<C::RefItem<'a>>,
    C::RefItem<'a>: PartialOrd,
{
    i1.into().partial_cmp(&i2.into()).unwrap() == Ordering::Less
}

pub fn cmp_ge<'a, I1: Array, I2: Array, C: Array>(i1: I1::RefItem<'a>, i2: I2::RefItem<'a>) -> bool
where
    I1::RefItem<'a>: Into<C::RefItem<'a>>,
    I2::RefItem<'a>: Into<C::RefItem<'a>>,
    C::RefItem<'a>: PartialOrd,
{
    i1.into().partial_cmp(&i2.into()).unwrap() == Ordering::Greater
}
