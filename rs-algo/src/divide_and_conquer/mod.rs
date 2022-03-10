/// 分治法
///
/// 步骤
/// - 分解(Devide)
/// - 解决(Conquer)
/// - 合并(Combine)
///
/// 使用问题的特征
/// - 缩小到一定规模容易解决
/// - 可以分解为若干个规模较小的性质相同的问题
/// - 子问题的解可以合并为该问题的解
/// - 个各自问题之间是相互独立的
///
/// 如果子问题不是互相独立的, 可以用动态规划
// fn divide_and_conquer<T>(problem: T) {
//     if (|problem| <= n0) {
//         adhoc(P)
//     }

//     let sub_problems: Vec<T> = devide(problem);
//     for sub_problem in sub_problems {
//         res_i = divide_and_conquer(sub_problem);
//     }
//     return merge(res_1, res_2, ..., res_n);
// }
mod tree;
