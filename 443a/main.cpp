#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <iostream>

using namespace std;

int main() {
	bool hit[26];
    char c;

	int cnt = 0;
    char str[4];

    ios::sync_with_stdio(false);
    while(gets(str)) {
        for (int i = 0; i < strlen(str); i++) {
            if (islower(str[i])) {
                c = str[i] - 'a';
		        if (!hit[c]) {
			        cnt += 1;
			        hit[c] = true;
		        }
            }
        }
    }

    printf("%d", cnt);
}
