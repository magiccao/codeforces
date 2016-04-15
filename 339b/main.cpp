#include <stdlib.h>
#include <stdio.h>
#include <iostream>
using namespace std;

int main() {
    int n, m, v;
    scanf("%d %d", &n, &m);
    
    int prev = 1;
    long long step = 0;
    int i = 0;
    for (; i < m; i++) {
        scanf("%d", &v);
        if (v >= prev) {
            step += (v - prev);
        } else {
            step += (n + v - prev);
        }
        prev = v;
    }

    cout << step;
}
