#include <stdlib.h>

int main() {
  int *ptr = malloc(sizeof(int)); // allocate memory

  *ptr = 10;
  *ptr += 1;

  free(ptr); // free the memory
}