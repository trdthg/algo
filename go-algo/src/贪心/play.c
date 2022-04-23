#include <stdio.h>

int find_min_index(int arr[], int l, int len, int cut_point[], int *cut_n)
{
    int min = 1000000;
    int cnt = 0;
    for (int i = l; i < l + len; i++)
    {
        if (arr[i] < min)
        {
            cnt = 0;
            min = arr[i];
            cut_point[cnt] = i;
            cnt += 1;
        }
        else if (arr[i] == min)
        {
            cut_point[cnt] = i;
            cnt += 1;
        }
    }
    *cut_n = cnt;
    return min;
}
void solve(int m[], int l, int len, int *sum)
{
    if (len == 0)
    {
        return;
    }
    else if (len == 1)
    {
        // printf("good %d %d\n", m[l], *sum);
        *sum += m[l];
        return;
    }
    int cut_point[100000];
    int cut_n;
    int min = find_min_index(m, l, len, cut_point, &cut_n);
    // printf(">> %d %d %d\n", min, cut_n, *sum);
    *sum += min * cut_n;
    for (int i = l; i < l + len; i++)
    {
        m[i] -= min;
    }
    cut_point[cut_n] = len;
    for (int i = 0; i < cut_n + 1; i++)
    {
        // printf("|> %d\n", cut_point[i]);
    }
    for (int i = 0; i < cut_n + 1; i++)
    {
        int r = cut_point[i];
        solve(m, l, r - l, sum);
        l = r + 1;
    }
    return;
}

int main(int argc, char const *argv[])
{
    int n = 0;
    scanf("%d", &n);
    int m[100000] = {0};
    for (int i = 0; i < n; i++)
    {
        scanf("%d", &m[i]);
    }
    int sum = 0;
    solve(m, 0, n, &sum);
    printf("%d", sum);
    return 0;
}