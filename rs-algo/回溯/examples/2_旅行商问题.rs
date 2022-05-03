use std::collections::HashMap;

fn main() {
    let reader = std::io::stdin();
    let mut sc = Scanner::new(reader);
    let n = sc.scan::<usize>();
    let k = sc.scan::<usize>();
    let res = solve(n, k);
    println!("{:?}", res);
}

fn solve(n: usize, k: usize) -> Vec<Vec<usize>> {
    let mut other = Other {
        res: vec![],
        path: vec![],
    };
    backtracing(n, k, 1, &mut other);
    other.res
}

struct Other {
    res: Vec<Vec<usize>>,
    path: Vec<usize>,
}

// 1. 函数的参数和返回值
fn backtracing(n: usize, k: usize, start_index: usize, other: &mut Other) {
    if other.path.len() == k {
        other.res.push(other.path.clone());
    }
    if n + other.path.len() - start_index < k {
        return;
    }
    for i in start_index..=n {
        other.path.push(i);
        backtracing(n, k, i + 1, other);
        other.path.pop();
    }
}
struct Scanner<'a> {
    reader: std::io::BufReader<std::io::Stdin>,
    buffer: std::str::SplitAsciiWhitespace<'a>,
    buf: String,
}

impl<'a> Scanner<'a> {
    fn new(reader: std::io::Stdin) -> Self {
        Self {
            reader: std::io::BufReader::new(reader),
            buffer: "".split_ascii_whitespace(),
            buf: String::new(),
        }
    }

    fn scan<T: std::str::FromStr>(&mut self) -> T {
        loop {
            if let Some(s) = self.buffer.next() {
                return s.parse().ok().unwrap();
            } else {
                self.buf.clear();
                std::io::BufRead::read_line(&mut self.reader, &mut self.buf).unwrap();
                self.buffer = unsafe { std::mem::transmute(self.buf.split_ascii_whitespace()) };
            }
        }
    }
}
