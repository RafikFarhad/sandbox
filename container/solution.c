#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main()
{
    int i;
    printf("Hello from code file\n");
    for(i=0; i< 3; i++) {
        printf("%d ...\n", i);
        sleep(1);
    }
    system("reboot");
    printf("Code file execution ended.\n");
    return 0;
}
