use std::fmt::Debug;

use crate::array::{Array, F32Array, I32Array, PrimitiveArray, StringArray};

pub trait Scalar: Debug + Clone + Send + Sync + 'static {
    type ArrayType: Array<OwnedItem = Self>;
    type RefType<'a>: ScalarRef<'a, ScalarType = Self>;
    fn as_scalar_ref(&self) -> Self::RefType<'_>;
}

pub trait ScalarRef<'a>: Debug + Clone + Copy + Send + 'a {
    type ArrayType: Array<RefItem<'a> = Self>;
    type ScalarType: Scalar<RefType<'a> = Self>;
    fn to_owned_scalar(&self) -> Self::ScalarType;
}
pub enum ScalarImpl {
    Int32(i32),
    Float32(f32),
    String(String),
}

pub enum ScalarRefImpl<'a> {
    Int32(i32),
    Float32(f32),
    String(&'a str),
}

impl Scalar for i32 {
    type ArrayType = I32Array;

    type RefType<'a> = i32;

    fn as_scalar_ref(&self) -> Self::RefType<'_> {
        *self
    }
}

impl Scalar for f32 {
    type ArrayType = F32Array;

    type RefType<'a> = f32;

    fn as_scalar_ref(&self) -> Self::RefType<'_> {
        *self
    }
}

impl<'a> ScalarRef<'a> for i32 {
    type ArrayType = I32Array;

    type ScalarType = i32;

    fn to_owned_scalar(&self) -> Self::ScalarType {
        *self
    }
}

impl<'a> ScalarRef<'a> for f32 {
    type ArrayType = F32Array;

    type ScalarType = f32;

    fn to_owned_scalar(&self) -> Self::ScalarType {
        *self
    }
}

impl Scalar for String {
    type ArrayType = StringArray;

    type RefType<'a> = &'a str;

    fn as_scalar_ref(&self) -> Self::RefType<'_> {
        self.as_ref()
    }
}

impl<'a> ScalarRef<'a> for &'a str {
    type ArrayType = StringArray;

    type ScalarType = String;

    fn to_owned_scalar(&self) -> Self::ScalarType {
        self.to_string()
    }
}
