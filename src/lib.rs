// def a(n = 8) {
//     while n >= 1 do {
//         j = 1
//         for
//         for j = 1 to n; do {
//             k = k + 1
//             n = n / 2
//             return k
//         }
//     }
// }
fn a(n: u32) -> u32 {
    let mut n = n;
    let mut k = 0;
    while n >= 1 {
        println!("1: {n}");
        for j in 1..n {
            println!("2: {n}");
            k += 1;
            n /= 2;
        }
    }
    return k;
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn it_works() {
        let result = 2 + 2;
        assert_eq!(result, 4);
    }

    #[test]
    fn test_a() {
        assert_eq!(a(8), 7);
    }
}
