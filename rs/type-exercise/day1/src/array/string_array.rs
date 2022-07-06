use bitvec::prelude::BitVec;

use super::{Array, ArrayBuilder};

/// 存储变长数据
///
/// Vec<Option<String>> 需要创建一个 String 对象
/// 连续读取 String 对于缓存不友好，String 是一个指针，数据则散落在内存各处，属于随机读
pub struct StringArray {
    /// 实际数据
    data: Vec<u8>,
    /// 每个字符串的偏移量
    offsets: Vec<usize>,
    /// 位图
    bitmap: BitVec,
}

impl Array for StringArray {
    type RefItem<'a> = &'a str;
    type Builder = StringArrayBuilder;

    fn get(&self, idx: usize) -> Option<Self::RefItem<'_>> {
        if !self.bitmap[idx] {
            return None;
        }
        let range = self.offsets[idx]..self.offsets[idx + 1];
        let res = unsafe { std::str::from_utf8_unchecked(&self.data[range]) };
        Some(res)
    }
    fn len(&self) -> usize {
        self.bitmap.len()
    }
}

pub struct StringArrayBuilder {
    data: Vec<u8>,
    bitmap: BitVec,
    offsets: Vec<usize>,
}
impl ArrayBuilder for StringArrayBuilder {
    type Array = StringArray;

    fn with_capacity(capacity: usize) -> Self {
        let mut offsets = Vec::with_capacity(capacity + 1);
        offsets.push(0);
        Self {
            data: Vec::with_capacity(capacity),
            bitmap: BitVec::with_capacity(capacity),
            offsets,
        }
    }

    fn push(&mut self, value: Option<<Self::Array as Array>::RefItem<'_>>) {
        match value {
            Some(v) => {
                self.data.extend(v.as_bytes());
                self.bitmap.push(true);
            }
            None => {
                self.bitmap.push(false);
            }
        }
        self.offsets.push(self.data.len());
    }

    fn finish(self) -> Self::Array {
        Self::Array {
            data: self.data,
            bitmap: self.bitmap,
            offsets: self.offsets,
        }
    }
}

fn eval_binary<I: Array, O: Array>(i1: I, i2: I) -> O {
    assert_eq!(i1.len(), i2.len(), "size mismatch");
    let mut builder = O::Builder::with_capacity(i1.len());
    for (i1, i2) in i1.iter().zip(i2.iter()) {
        //   builder.push(sql_func(i1, i2));
    }
    builder.finish()
}

#[cfg(test)]
mod tests {
    use super::*;

    // These are two examples of using generics over array.
    //
    // These functions work for all kinds of array, no matter fixed-length arrays like `I32Array`,
    // or variable-length ones like `StringArray`.

    /// Build an array from a vector of data
    fn build_array_from_vec<A: Array>(items: &[Option<A::RefItem<'_>>]) -> A {
        let mut builder = A::Builder::with_capacity(items.len());
        for item in items {
            builder.push(*item);
        }
        builder.finish()
    }

    /// Test if an array has the same content as a vector
    fn check_array_eq<'a, A: Array>(array: &'a A, vec: &[Option<A::RefItem<'a>>])
    where
        A::RefItem<'a>: PartialEq,
    {
        for (a, b) in array.iter().zip(vec.iter()) {
            assert_eq!(&a, b);
        }
    }

    // #[test]
    // fn test_build_int32_array() {
    //     let data = vec![Some(1), Some(2), Some(3), None, Some(5)];
    //     let array = build_array_from_vec::<I32Array>(&data[..]);
    //     check_array_eq(&array, &data[..]);
    // }

    #[test]
    fn test_build_string_array() {
        let data = vec![Some("1"), Some("2"), Some("3"), None, Some("5"), Some("")];
        let array = build_array_from_vec::<StringArray>(&data[..]);
        check_array_eq(&array, &data[..]);
    }
}
