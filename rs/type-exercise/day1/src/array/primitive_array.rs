use std::fmt::Debug;

use bitvec::vec::BitVec;

use super::{Array, ArrayBuilder};
pub trait PrimitiveType: Debug + Copy + Default + 'static {}
/// 存储顶长数据
///
/// data:    233abc
///          ^  ^  ^
///          |  |  |--|
/// offsets: 0, 3, 6, 6
/// bitmap: true, true, false
#[derive()]
pub struct PrimitiveArray<T: PrimitiveType> {
    /// 实际数据
    data: Vec<T>,
    /// 位图
    bitmap: BitVec,
}

impl<T: PrimitiveType> Array for PrimitiveArray<T> {
    type RefItem<'a> = T;
    type Builder = PrimitiveArrayBuilder<T>;

    fn get(&self, idx: usize) -> Option<Self::RefItem<'_>> {
        if !self.bitmap[idx] {
            return None;
        }
        Some(self.data[idx])
    }
    fn len(&self) -> usize {
        self.data.len()
    }
}

pub struct PrimitiveArrayBuilder<T: PrimitiveType> {
    data: Vec<T>,
    bitmap: BitVec,
}
impl<T: PrimitiveType> ArrayBuilder for PrimitiveArrayBuilder<T> {
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
