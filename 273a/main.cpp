# include <iostream>
using namespace std;

int main() {
    int n, p1, p2, v1, v2;
    cin >> n;
    int cashes = 1;
    int max = 0;

    cin >> p1 >> p2;

    for (int i = 1; i < n; i++) {
        cin >> v1 >> v2;
        if (v1 == p1 && v2 == p2) {
            cashes += 1;
        } else {
            if (cashes > max) {
                max = cashes;
            }
            p1 = v1;
            p2 = v2;
            cashes = 1;
        }
    }

    if (cashes > max) {
        max = cashes;
    }

    cout << max << endl;
}
