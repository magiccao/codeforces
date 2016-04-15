#include <stdlib.h>
#include <stdio.h>
#include <iostream>
using namespace std;

int main() {
    int n;
    cin>>n;

    int i = 1;
    int flag = 0;
    int max = 1;
    int cur;
    int prev;
    cin >> prev;

    for (; i < n; i++) {
        cin >> cur;
        if (cur < prev) {
            if (max < i - flag) {
                max = i - flag;
            }
            flag = i;
        }
        prev = cur;
    }

    if (max < n-flag) {
        max = n - flag;
    }
    cout << max;
}
