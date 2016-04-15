#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <iostream>
#include <algorithm>
using namespace std;

const int N = 1e5+5;
int num[N];

long long points() {
    // p means max scores when do not select a[k-1];
    // q means max scores when do select a[k-1]
    long long p = 0, q = 0;
    for (int i = 1; i < N; i++) {
        long long tmp = max(p, q);    
        q = p + (long long)i * num[i];
        p = tmp;
    }

    return max(p, q);
}

int main() {
    int n, v;
    cin >> n;
    memset(num, 0, sizeof(num));

    for (int i = 1; i <= n; i++) {
        cin >> v;
        num[v] += 1;
    }

    cout << points();
}
