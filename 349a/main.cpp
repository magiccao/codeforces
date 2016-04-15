#include <iostream>
using namespace std;

int main() {
    int kinds[3] = {0};
    int n, v;
    cin >> n;

    bool ok = true;
    for (int i = 0; i < n; i++) {
        cin >> v;
        if (ok) {
            if (i == 0 && v != 25) {
                ok = false;
                continue;
            }

            if (v == 25) {
                kinds[1] += 1;
            } else if (v == 50) {
                if (kinds[1] == 0) {
                    ok = false;
                } else {
                    kinds[1] -= 1;
                    kinds[2] += 1;
                }
            } else {
                if (kinds[1] > 0) {
                    if (kinds[2] > 0) {
                        kinds[1] -= 1;
                        kinds[2] -= 1;
                    } else if (kinds[1] >= 3) {
                        kinds[1] -= 3;
                    } else {
                        ok = false;
                    }
                } else {
                    ok = false;
                }
            }
        }
    }

    if (ok) {
        cout << "YES" << endl;
    } else {
        cout << "NO" << endl;
    }
}
