#include <stdio.h>
#include <unistd.h>


int main()
{
    int i;
    printf("Hello from solution\n");
    for(i=0; i< 3; i++) {
        printf("%d ...\n", i);
        sleep(1);
    }
    printf("Solution ended.\n");
    return 0;
}