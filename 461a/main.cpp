#include <iostream>
#include <algorithm>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
using namespace std;

int group[300005];

int cmp(const void *a , const void *b ) {
    return *(int *)a - *(int *)b;
}

int main() {
    int n;
    cin >> n;

    for (int i = 0; i < n; i++) {
        cin >> group[i];
    }

    if (n == 1) {
        cout << group[0] << endl;
        return 0;
    }

    qsort(group, n, sizeof(group[0]), cmp);

    long long score = 0;
    for (int i = 0; i < n - 2; i++) {
        score += (long long)(i + 2) * group[i];
    }

    score += (long long)(n) * (group[n-2] + group[n-1]);
    cout << score << endl;
}
