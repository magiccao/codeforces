#include <iostream>
using namespace std;

int main() {
    int n1, n2, k, m, v, amax, bmin;
    cin >> n1 >> n2 >> k >> m;

    for (int i = 1; i <= n1; i++) {
        cin >> v;
        if (i == k) {
            amax = v;
        }
    }

    for (int i = 1; i <= n2; i++) {
        cin >> v;
        if (i == n2 - m + 1) {
            bmin = v;
        }
    }

    if (amax < bmin) {
        cout << "YES" << endl;
    } else {
        cout << "NO" << endl;
    }

    return 0;
}
