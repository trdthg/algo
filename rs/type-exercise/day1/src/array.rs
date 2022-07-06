use std::fmt::Debug;

use self::iterator::ArrayIterator;

mod iterator;
mod primitive_array;
mod string_array;

pub trait Array {
    type RefItem<'a>: Copy + Debug
    where
        Self: 'a;

    type Builder: ArrayBuilder<Array = Self>;

    /// 返回一个引用
    fn get(&self, idx: usize) -> Option<Self::RefItem<'_>>;

    /// 总数量。
    fn len(&self) -> usize;

    /// 是否为空
    fn is_empty(&self) -> bool {
        self.len() == 0
    }

    fn iter(&self) -> ArrayIterator<Self>
    where
        Self: Sized,
    {
        ArrayIterator::new(self)
    }
}

pub trait ArrayBuilder {
    type Array: Array<Builder = Self>;

    fn with_capacity(capacity: usize) -> Self;

    fn push(&mut self, value: Option<<Self::Array as Array>::RefItem<'_>>);

    fn finish(self) -> Self::Array;
}
