#include <iostream>
#include <stdlib.h>
using namespace std;

struct laptop {
    int price;
    int quality;
};

int cmp(const void * a, const void * b) {
    return ((laptop*)a)->price - ((laptop*)b)->price; 
}

laptop lap[100005];

int main() {
    int n;
    cin >> n;

    for (int i = 0; i < n; i++) {
        cin >> lap[i].price >> lap[i].quality;
    }

    qsort(lap, n, sizeof(laptop), cmp);
    for (int i = 0; i < n - 1; i++) {
        if (lap[i].quality > lap[i+1].quality) {
            cout << "Happy Alex" << endl;
            return 0;
        }
    }

    cout << "Poor Alex" << endl;
    return 0;
}
