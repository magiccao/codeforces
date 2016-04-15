#include <iostream>
using namespace std;

int main() {
    int n;
    cin >> n;

    long long v;
    long long sum[3] = {0};
    for (int i = 0; i < n; i++) {
        cin >> v;
        sum[0] += v;
    }
    for (int i = 0; i < n - 1; i++) {
        cin >> v;
        sum[1] += v;
    }
    for (int i = 0; i < n - 2; i++) {
        cin >> v;
        sum[2] += v;
    }

    cout << sum[0] - sum[1] << endl 
         << sum[1] - sum[2] << endl;
}
