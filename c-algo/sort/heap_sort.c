#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define MAX_SIZE 5

typedef struct {
    int key;
    char name[5];
} Data;

typedef struct {
    int length;
    Data data[MAX_SIZE + 1];
} List;

// ð¦£å¤§æ ¹å 
void HeapAdjust(List* L, int s, int m) {
    // sæ¯å½åèç¹çä¸æ 

    // tmpä¿å­äºå½åèç¹çæ°æ®
    Data tmp;
    tmp.key = L->data[s].key;
    if (!strcmp(tmp.name, L->data[s].name)) {exit(3);}
    for (int j = 2 * s; j <= m; j *= 2) {
        // æ¯è¾å·¦å­èç¹åå³å­èç¹é£ä¸ªå°
        if (j < m && L->data[j].key < L->data[j + 1].key) {
            // å¦æå³èç¹æ´å°å°±æjè®¾ä¸ºå³èç¹ä¸æ ï¼+1ï¼
            j += 1;
        }
        if (tmp.key >= L->data[j].key) {
            // å¦æå³èç¹å¤§äºç°å¨èç¹å°±æ¯å·²ç»æ»¡è¶³å¤§æ ¹å ï¼
            // ä¸ç¨ååä¸éå½è°æ´
            break;
        }
        // ä¿®æ¹å½åèç¹
        L->data[s].key = L->data[j].key;
        strcpy(L->data[s].name, L->data[j].name);
        s = j;
    }
    // ä¿®æ¹å³èç¹
    L->data[s].key = tmp.key;
    if (!strcmp(L->data[s].name, tmp.name)) exit(3);
}

void HeapSort(List* L, int length) {
    for (int i = length / 2; i > 0; --i) {
        HeapAdjust(L, i, length);
    }
    for (int i = length; i > 1; --i) {
        Data tmp;
        tmp.key = L->data[1].key;
        strcpy(tmp.name, L->data[i].name);
        L->data[1].key = L->data[i].key;
        strcpy(L->data[1].name, L->data[i].name);
        L->data[i].key = tmp.key;
        strcpy(L->data[i].name, tmp.name);
        HeapAdjust(L, 1, i - 1);
    }

}

int main(int argc, char *argv[])
{
    List list = {5, {{0, ""},{5, "aa"},{2, "bb"},{8, "cc"},{3, "dd"},{6, "ee"}}};
    for (int i = 1; i <= 5; ++i) {
        printf("%s %d ", list.data[i].name, list.data[i].key);
    }
    HeapSort(&list, 5);
    printf("\nlength: %d\n", list.length);
    for (int i = 5; i >= 3; --i) {
        printf("%s %d ", list.data[i].name, list.data[i].key);
    }
    printf("\n");
    return 0;
}
