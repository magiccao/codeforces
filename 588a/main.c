#include <stdlib.h>
#include <stdio.h>

int main() {
    int n, q, p;
    scanf("%d", &n);

    int buy, price;
    scanf("%d %d", &buy, &price);

    int i = 1;
    long long cost = 0;
    for (; i < n; i++) {
        scanf("%d %d", &q, &p);
        if (p >= price) {
            buy += q;
        } else {
            cost += (long long)buy * price;
            buy = q;
            price = p;
        }
    }

    cost += (long long)buy * price;
    printf("%d", cost);
    

}
