
#include <iostream>
using namespace std;
int n, step, pos;
char c[1005];

void print()
{
    cout << "Step " << step++ << ":";
    for (int i = 1; i <= 2 * n + 2; i++)
        cout << c[i];
    cout << endl;
}

void init(int n)
{
    step = 0;
    pos = 2 * n + 1; //空位的开始下标
    for (int i = 1; i <= n; i++)
        c[i] = '1'; //白子
    for (int i = n + 1; i <= 2 * n; i++)
        c[i] = '2';                    //黑子
    c[2 * n + 1] = c[2 * n + 2] = '0'; //空位
    print();
}

void move(int k)
{
    for (int j = 0; j <= 1; j++)
    {
        c[pos + j] = c[k + j];
        c[k + j] = '0';
    }
    pos = k; //更新空位
    print();
}

// 1 1 1 1 2 2 2 2|0 0 移动中间
// 1 1 1 0 0 2 2 2|1 2 固定
// 1 1 1 2 2 2|0 0|1 2 还原，移动中间
// 1 1 0 0 2 2|1 2|1 2 固定
void mv(int n)
{
    if (n == 4)
    {
        move(4);
        move(8);
        move(2);
        move(7);
        move(1);
    }
    else
    {
        move(n);
        move(2 * n - 1);
        mv(n - 1);
    }
}

int main()
{
    cin >> n;
    init(n);
    mv(n);
    cout << step - 1 << endl; //总数
    return 0;
}
