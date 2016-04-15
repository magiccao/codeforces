#include <stdlib.h>
#include <stdio.h>

int main() {
   char op[2][3];
   int n;
   scanf("%d", &n);

   int cnt = 0;
   int id = 0;
   int i = 0;

   for (; i < n; i++, id = 1-id) {
       scanf("%s", op[id]);
       if (i == 0) {
            cnt += 1;
            continue;
       }

       if (op[id][0] - op[1-id][1] == 0) {
            cnt += 1;
       }
   }

   printf("%d", cnt);
}
