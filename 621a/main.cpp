#include <iostream>
#include <algorithm>
using namespace std;

int main() {
    int n, v;
    long long sum = 0;
    int minOdd = 1000000005;
    cin >> n;
    for (int i = 0; i < n; i++) {
        cin >> v;
        sum += v;
        if (v % 2 == 1) {
            if (minOdd > v) {
                minOdd = v;
            }
        }
    }

    if (sum % 2 == 0) {
        cout << sum << endl;
    } else {
        cout << sum - minOdd << endl;
    }
}


