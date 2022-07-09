mod cmp;
mod string;
mod vectorize;

use std::marker::PhantomData;

use anyhow::Result;

use self::vectorize::BinaryExprFunc;
use crate::array::*;
use crate::scalar::*;
use crate::TypeMismatch;
