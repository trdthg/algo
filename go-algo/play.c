#include <stdio.h>

int a[100000];
int b[100000];

void mergeInner(int begin, int mid, int end)
{
    int l = begin, r = mid;
    int cnt = begin;
    while (l < mid, r < end)
    {
        if (a[l] < a[r])
        {
            b[cnt] = a[l];
        }
        else
        {
            b[cnt] = a[r];
        }
        cnt++;
    }
    for (;;)
    {
        b[cnt] = a[l];
    }

    for (int i = begin; i < end; i++)
    {
        a[i] = b[i];
    }
}

void merge(int begin, int end)
{
    int step = 1;
    int mid, i, j, l, r;
    while (step < (end - begin) * 2)
    {
        for (int i = begin; i < step; i += step << 1)
        {
            l, mid, r = i, i + step, r > end ? r : end;
            if (mid > end)
            {
                continue;
            }
            mergeInner(l, mid, r);
        }
        step <<= 1;
    }

    // if (l >= r)
    // {
    //     return;
    // }
    // mid = (l + r) / 2;
    // merge(l, mid);
    // merge(mid, r);
}

int main(int argc, char const *argv[])
{
    int n, j, i;
    scanf("%d", &n);
    for (i = 0; i < n; i++)
    {
        scanf("%d", &a[i]);
    }

    for (i = 0; i < n; i++)
    {
        printf("%d", a[i]);
    }

    int cnt = 0;
    for (i = 0; i < n - 1; i++)
    {
        for (j = i + 1; j < n; j++)
        {
            if (a[i] > a[j])
            {
                cnt += 1;
            }
        }
    }
    printf("\n");
    merge(0, n);
    for (i = 0; i < n; i++)
    {
        printf("%d", a[i]);
    }
    return 0;
}
