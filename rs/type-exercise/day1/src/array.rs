use std::fmt::Debug;

mod iterator;
mod primitive_array;
mod string_array;

use iterator::*;
use primitive_array::*;
use string_array::*;

pub trait Array: Send + Sync + Sized + 'static {
    type RefItem<'a>: Copy + Debug;

    type OwnedItem: Debug;

    type Builder: ArrayBuilder<Array = Self>;

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
}
