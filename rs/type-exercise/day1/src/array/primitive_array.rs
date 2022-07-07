use std::fmt::Debug;

use bitvec::vec::BitVec;

use super::iterator::ArrayIterator;
use super::{Array, ArrayBuilder};
use crate::scalar::{Scalar, ScalarRef};
// pub trait PrimitiveType: Debug + Copy + Default + Send + Sync + 'static {}ault + Debug + 'static
// {}
pub trait PrimitiveType: Copy + Send + Sync + Default + Debug + 'static {}

pub type I32Array = PrimitiveArray<i32>;
pub type F32Array = PrimitiveArray<f32>;

impl PrimitiveType for i32 {}
impl PrimitiveType for f32 {}
/// 存储顶长数据
///
/// data:    233abc
///          ^  ^  ^
///          |  |  |--|
/// offsets: 0, 3, 6, 6
/// bitmap: true, true, false
pub struct PrimitiveArray<T: PrimitiveType> {
    /// 实际数据
    data: Vec<T>,
    /// 位图
    bitmap: BitVec,
}

impl<T> Array for PrimitiveArray<T>
where
    T: PrimitiveType,
    T: Scalar<ArrayType = Self>,
    for<'a> T: ScalarRef<'a, ArrayType = Self, ScalarType = T>,
    for<'a> T: Scalar<RefType<'a> = T>,
{
    type Builder = PrimitiveArrayBuilder<T>;

    type RefItem<'a> = T;

    type OwnedItem = T;

    fn get(&self, idx: usize) -> Option<Self::RefItem<'_>> {
        if !self.bitmap[idx] {
            return None;
        }
        Some(self.data[idx])
    }
    fn len(&self) -> usize {
        self.data.len()
    }

    fn iter(&self) -> super::iterator::ArrayIterator<Self> {
        ArrayIterator::new(self)
    }
}

pub struct PrimitiveArrayBuilder<T: PrimitiveType> {
    data: Vec<T>,
    bitmap: BitVec,
}

impl<T> ArrayBuilder for PrimitiveArrayBuilder<T>
where
    T: PrimitiveType,
    T: Scalar<ArrayType = PrimitiveArray<T>>,
    for<'a> T: ScalarRef<'a, ArrayType = PrimitiveArray<T>, ScalarType = T>,
    for<'a> T: Scalar<RefType<'a> = T>,
{
    type Array = PrimitiveArray<T>;

    fn with_capacity(capacity: usize) -> Self {
        Self {
            data: Vec::with_capacity(capacity),
            bitmap: BitVec::with_capacity(capacity),
        }
    }

    fn push(&mut self, value: Option<T>) {
        match value {
            Some(v) => {
                self.data.push(v);
                self.bitmap.push(true);
            }
            None => {
                self.data.push(T::default());
                self.bitmap.push(false);
            }
        }
    }

    fn finish(self) -> Self::Array {
        Self::Array {
            data: self.data,
            bitmap: self.bitmap,
        }
    }
}