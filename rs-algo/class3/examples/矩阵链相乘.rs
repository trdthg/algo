fn main() {
    //
    let mut p = [5, 7, 4, 3, 5];
    let res = solve(&mut p[..]);
}

// 输入原始纬度信息 -> 最小乘法次数，分割点位置
fn solve(p: &mut [i32]) -> (u32, &[&[usize]]) {
    let m = [[0; 100]; 100];
    
    (1, &[&[1, 2, 3]])
}

fn trace_back(i: usize, j: usize, split_points: &[&[usize]]) {
    if i == j {
        print!("A{}", i);
        return;
    }
    print!("(");
    trace_back(i, split_points[i][j], split_points);
    trace_back(split_points[i][j], j, split_points);
    print!(")");
}
