fn equation(n: u32, m: u32) -> u32 {
    match (n, m) {
        (_, 1) => 1,
        (1, _) => 1,
        // n <= m
        (n, m) if (n <= m) => 1 + equation(n, n - 1),
        // m > m
        (n, m) => equation(n - m, m) + equation(n, m - 1),
    }
}

// fn back_trace(list: Vec<u32>, k: )
// 正整数划分
fn split_num(n: u32) -> Vec<Vec<u32>> {
    let mut res = vec![];
    for i in (1..n + 1).rev() {
        let tmp = match i {
            // 6
            i if i == n => vec![vec![i]],
            // 5 + 1
            i if n - i == 1 => vec![vec![i, 1]],
            // 4 + 2 | 3 + 3
            i => {
                let mut head = vec![];
                let tail = split_num(n - i);
                head.extend(tail);
                head.iter_mut().for_each(|x| x.insert(0, i));
                head
            }
        };
        println!("{tmp:?}");
        res.extend(tmp);
    }
    res
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn split_num_test() {
        println!("{}", equation(4, 4));
        println!("{:?}", split_num(4));
    }
}
