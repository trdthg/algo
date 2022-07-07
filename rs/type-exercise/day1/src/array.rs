use std::fmt::Debug;

mod iterator;
mod primitive_array;
mod string_array;

pub use iterator::*;
pub use primitive_array::*;
pub use string_array::*;

use crate::scalar::{Scalar, ScalarRefImpl};

pub trait Array: Send + Sync + Sized + 'static + TryFrom<ArrayImpl> + Into<ArrayImpl>
where
    for<'a> Self::OwnedItem: Scalar<RefType<'a> = Self::RefItem<'a>>,
{
    type Builder: ArrayBuilder<Array = Self>;

    type OwnedItem: Scalar<ArrayType = Self>;

    type RefItem<'a>: Copy + Debug;

    /// 返回一个引用
    fn get(&self, idx: usize) -> Option<Self::RefItem<'_>>;

    /// 总数量。
    fn len(&self) -> usize;

    /// 是否为空
    fn is_empty(&self) -> bool {
        self.len() == 0
    }

    fn iter(&self) -> ArrayIterator<Self>;
}

pub trait ArrayBuilder {
    type Array: Array<Builder = Self>;

    fn with_capacity(capacity: usize) -> Self;

    fn push(&mut self, value: Option<<Self::Array as Array>::RefItem<'_>>);

    fn finish(self) -> Self::Array;
}

fn eval_binary<I1: Array, I2: Array>(i1: &ArrayImpl, i2: &ArrayImpl) -> Result<ArrayImpl, ()>
where
    for<'a> &'a I1: TryFrom<&'a ArrayImpl, Error = ()>,
    for<'a> &'a I2: TryFrom<&'a ArrayImpl, Error = ()>,
{
    let i1: &I1 = i1.try_into()?;
    let i2: &I2 = i2.try_into()?;
    // some black magic
    todo!()
}

pub enum ArrayImpl {
    Int32(I32Array),
    Float32(F32Array),
    String(StringArray),
}

impl TryFrom<ArrayImpl> for I32Array {
    type Error = ();

    fn try_from(value: ArrayImpl) -> Result<Self, Self::Error> {
        match value {
            ArrayImpl::Int32(value) => Ok(value),
            _ => Err(()),
        }
    }
}

impl From<I32Array> for ArrayImpl {
    fn from(v: I32Array) -> Self {
        ArrayImpl::Int32(v)
    }
}

impl TryFrom<ArrayImpl> for F32Array {
    type Error = ();

    fn try_from(value: ArrayImpl) -> Result<Self, Self::Error> {
        match value {
            ArrayImpl::Float32(value) => Ok(value),
            _ => Err(()),
        }
    }
}

impl From<F32Array> for ArrayImpl {
    fn from(v: F32Array) -> Self {
        ArrayImpl::Float32(v)
    }
}

impl TryFrom<ArrayImpl> for StringArray {
    type Error = ();

    fn try_from(value: ArrayImpl) -> Result<Self, Self::Error> {
        match value {
            ArrayImpl::String(value) => Ok(value),
            _ => Err(()),
        }
    }
}

impl From<StringArray> for ArrayImpl {
    fn from(v: StringArray) -> Self {
        ArrayImpl::String(v)
    }
}

impl ArrayImpl {
    fn get(&self, idx: usize) -> Option<ScalarRefImpl<'_>> {
        todo!()
    }
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

    #[test]
    fn test_build_int32_array() {
        let data = vec![Some(1), Some(2), Some(3), None, Some(5)];
        let array = build_array_from_vec::<I32Array>(&data[..]);
        check_array_eq(&array, &data[..]);
    }

    #[test]
    fn test_build_string_array() {
        let data = vec![Some("1"), Some("2"), Some("3"), None, Some("5"), Some("")];
        let array = build_array_from_vec::<StringArray>(&data[..]);
        check_array_eq(&array, &data[..]);
    }

    #[test]
    fn test() {}
}
