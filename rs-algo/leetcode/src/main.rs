use std::collections::HashMap;

fn main() {
    println!("Hello, world!");
    assert_eq!(
        vec![4, 0, 1, 1, 3],
        smaller_numbers_than_current(vec![8, 1, 2, 2, 3])
    );
    assert_eq!(
        vec![4, 0, 1, 1, 3],
        smaller_numbers_than_current2(vec![8, 1, 2, 2, 3])
    );
}
pub fn array_pair_sum(nums: Vec<i32>) -> i32 {
    let mut nums = nums;
    nums.sort();
    let mut n = 0;
    for i in 0..nums.len() / 2 {
        n += nums[i * 2]
    }
    n
}
//给你一个下标从 0 开始的整数数组 nums 以及一个目标元素 target 。
// 目标下标 是一个满足 nums[i] == target 的下标 i 。
// 将 nums 按 非递减 顺序排序后，返回由 nums 中目标下标组成的列表。如果不存在目标下标，返回一个 空 列表。返回的列表必须按 递增 顺序排列。
pub fn target_indices(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut nums = nums;
    nums.sort();
    let mut arr = vec![];
    for (i, ele) in nums.iter().enumerate() {
        if *ele == target {
            arr.push(i as i32)
        } else if *ele > target {
            break;
        }
    }
    arr
}

// nums里最大的乘积之差
pub fn max_product_difference(nums: Vec<i32>) -> i32 {
    let max = i32::max(nums[0], nums[1]);
    let min = i32::min(nums[0], nums[1]);
    let mut max1 = max;
    let mut max2 = min;
    let mut min1 = min;
    let mut min2 = max;
    for i in 2..nums.len() {
        if max1 <= nums[i] {
            max2 = max1;
            max1 = nums[i];
        } else if max2 < nums[i] {
            max2 = nums[i];
        }

        if min1 >= nums[i] {
            min2 = min1;
            min1 = nums[i];
        } else if min2 >= nums[i] {
            min2 = nums[i];
        }
    }
    max1 * max2 - min1 * min2
}

// 最少的移动次数
pub fn min_moves_to_seat(seats: Vec<i32>, students: Vec<i32>) -> i32 {
    let mut seats = seats;
    let mut students = students;
    seats.sort_unstable();
    students.sort();
    let mut n = 0;
    for (s1, s2) in seats.iter().zip(students) {
        n += (s1 - s2).abs();
    }
    n
}

// 比它小的数字个数
pub fn smaller_numbers_than_current(nums: Vec<i32>) -> Vec<i32> {
    // 频次数组
    // [8, 1, 2, 2, 3]
    let mut map = vec![0; 100];
    // [0, 0, 0, 0, 0, 0, 0, 0, 0] 9个
    for (i, num) in nums.iter().enumerate() {
        map[*num as usize] += 1;
    }
    // [0, 1, 2, 1, 0, 0, 0, 0, 1] 9个
    for i in 1..map.len() {
        map[i] += map[i - 1];
    }
    // [0, 1, 3, 4, 4, 4, 4, 4, 5] 9个
    let mut arr = vec![];
    for i in 0..nums.len() {
        let res = if nums[i] == 0 {
            0
        } else {
            map[(nums[i] - 1) as usize]
        };
        arr.push(res);
    }
    arr
}

pub fn smaller_numbers_than_current2(nums: Vec<i32>) -> Vec<i32> {
    let mut map = vec![(0, 0); nums.len()];
    for (i, num) in nums.iter().enumerate() {
        map[i] = (*num, i);
    }
    map.sort_by(|a, b| a.0.cmp(&b.0));
    // [ (1,       2),     (1,   1), (10,  3)]
    //  last_num           num index
    // [0,           ]
    // last_i
    let mut arr = vec![0_i32; map.len()];
    let (mut last_num, mut last_i) = (0, -1);
    for i in 0..nums.len() {
        let index = map[i].1;
        let num = map[i].0;
        if last_i != -1 && num == last_num {
            arr[index] = arr[last_i as usize];
        } else {
            arr[index] = i as i32;
        }
        last_num = num;
        last_i = index as i32
    }
    arr
}
