#include <iostream>
#include <string.h>
using namespace std;

int pos[1000005];

int main() {
    int n, v;
    int acc = 1;
    cin >> n;
    for (int i = 1; i <= n; i++) {
        cin >> v;
        for (int j = 0; j < v; j++) {
            pos[acc+j] = i;
        }
        acc += v;
    }

    int m;
    cin >> m;
    for (int i = 0; i < m; i++) {
        cin >> v;
        cout << pos[v] << endl;
    }
}
