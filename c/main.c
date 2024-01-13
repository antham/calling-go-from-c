#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "libhttp.h"

int main(int argc, char *argv[]) {
  GoString s = Get(100);

  char* status = malloc(s.n + 1);
  if (!status) {
    printf("A failure occurred");
    return 1;
  }
  memcpy(status, s.p, s.n);
  status[s.n] = '\0';
  printf("\nStatus: %s",status);
  free(status);
  return 0;
}
