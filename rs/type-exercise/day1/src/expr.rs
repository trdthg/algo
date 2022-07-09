mod cmp;
mod string;
mod vectorize;

use std::marker::PhantomData;

use anyhow::Result;

use self::vectorize::{BinaryExprFunc, BinaryExpression};
use crate::array::*;
use crate::scalar::*;
use crate::TypeMismatch;

pub trait Expression {
    fn eval_expr(&self, data: &[&ArrayImpl]) -> Result<ArrayImpl>;
}

pub enum ExpressionFunc {
    CmpLe,
    CmpGe,
    CmpEq,
    CmpNe,
    StrContains,
}

pub fn build_binary_expression(f: ExpressionFunc) -> Box<dyn Expression> {
    use ExpressionFunc::*;

    use crate::expr::cmp::*;
    use crate::expr::string::*;
    use crate::expr::vectorize::*;

    match f {
        CmpLe => Box::new(BinaryExpression::<i32, i32, bool, _>::new(
            cmp_le::<i32, i32, i64>,
        )),
        CmpGe => Box::new(BinaryExpression::<i32, i32, bool, _>::new(
            cmp_ge::<i32, i32, i64>,
        )),
        CmpEq => Box::new(BinaryExpression::<i32, i32, bool, _>::new(
            cmp_eq::<i32, i32, i64>,
        )),
        CmpNe => Box::new(BinaryExpression::<i32, i32, bool, _>::new(
            cmp_ne::<i32, i32, i64>,
        )),
        StrContains => Box::new(BinaryExpression::<String, String, bool, _>::new(
            str_contains,
        )),
    }
}

impl<A: Scalar, B: Scalar, O: Scalar, F> Expression for BinaryExpression<A, B, O, F>
where
    for<'a> &'a A::ArrayType: TryFrom<&'a ArrayImpl, Error = TypeMismatch>,
    for<'a> &'a B::ArrayType: TryFrom<&'a ArrayImpl, Error = TypeMismatch>,
    F: BinaryExprFunc<A, B, O>,
{
    fn eval_expr(&self, data: &[&ArrayImpl]) -> Result<ArrayImpl> {
        self.eval_batch(data[0], data[1])
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::array::{Array, StringArray};
    use crate::scalar::ScalarRefImpl;

    #[test]
    fn test_build_str_contains() {
        let expr = build_binary_expression(ExpressionFunc::StrContains);

        for _ in 0..10 {
            let result = expr
                .eval_expr(&[
                    &StringArray::from_slice(&[Some("000"), Some("111"), None]).into(),
                    &StringArray::from_slice(&[Some("0"), Some("0"), None]).into(),
                ])
                .unwrap();
            assert_eq!(result.get(0).unwrap(), ScalarRefImpl::Bool(true));
            assert_eq!(result.get(1).unwrap(), ScalarRefImpl::Bool(false));
            assert!(result.get(2).is_none());
        }
    }
}
