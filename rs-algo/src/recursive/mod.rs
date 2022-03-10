mod equation;
mod permute;
// 阶乘
fn factorial(n: u32) -> u32 {
    match n {
        1 => n,
        n => n * factorial(n - 1),
    }
}

// fig
fn fib(n: u32) -> u32 {
    match n {
        0 | 1 => 1,
        n => fib(n - 1) + fib(n - 2),
    }
}

// ackerman
fn ackerman(n: u32, m: u32) -> u32 {
    match (n, m) {
        (1, 0) => 2,
        (0, m) if m >= 0 => 1,
        (n, 0) if n >= 2 => n + 2,
        (n, m) if n >= 1 && m >= 1 => ackerman(ackerman(n - 1, m), m - 1),
        _ => panic!("unsupported!"),
    }
}

// 全排列
fn permute(nums: Vec<u32>) -> Vec<Vec<u32>> {
    let mut res = vec![];
    return res;
}

// 汉诺它

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn a() {
        let a = 5;
        let res = factorial(a);
        println!("{res}")
    }

    #[test]
    fn fib_test() {
        println!("{}", fib(5))
    }

    #[test]
    fn ackerman_test() {
        println!("{}", ackerman(2, 3))
    }
}
