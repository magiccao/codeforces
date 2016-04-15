#include <iostream>

using namespace std;

int flowers[200005];

int main() {
    int n;
    cin >> n;

    int min = 1000000005;
    int max = 0;
    long long minn = 0, maxn = 0;
    for (int i = 0; i < n; i++) {
        cin >> flowers[i];
        if (flowers[i] > max) {
            max = flowers[i];
        }
        if (flowers[i] < min) {
            min = flowers[i];
        }
    }

    int diff = max - min;
    for (int i = 0; i < n; i++) {
        if (flowers[i] == min) {
            minn += 1;
        } else if (flowers[i] == max) {
            maxn += 1;
        }
    }

    if (minn == n) {
        cout << 0 << " " << (long long)n * (n - 1) / 2;
    } else {
        cout << diff << " " << minn * maxn << endl;
    }
}
