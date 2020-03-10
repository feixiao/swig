#include <stdio.h>
#include "hello.h"

int hello(char *name, int age) {
    printf("Hello %s, your age is %d.\n", name, age);
    return age;
}
