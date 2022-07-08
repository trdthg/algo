use std::fmt::Debug;

mod impls;
pub use impls::*;

use crate::array::{Array, F32Array, I32Array, StringArray};

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
    Int16(i16),
    Int32(i32),
    Int64(i64),
    Float32(f32),
    Float64(f64),
    String(String),
    Bool(bool),
}

pub enum ScalarRefImpl<'a> {
    Int16(i16),
    Int32(i32),
    Int64(i64),
    Float32(f32),
    Float64(f64),
    String(&'a str),
    Bool(bool),
}
