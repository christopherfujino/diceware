#include <cstdio>
#include <cstdlib> // for std::srand
#include <cctype>
#include <ctime>   // for std::time

#include "wordlist.hpp"

const char *getWord() {
  const int wordIndex = std::rand() % wordSize;
  return words[wordIndex];
}

int main() {
  {
    timespec ts;
    clock_gettime(CLOCK_MONOTONIC, &ts);
    // XOR seconds & nanoseconds
    std::srand((unsigned int)(ts.tv_sec ^ ts.tv_nsec));
  }
  for (int i = 0; i < 4; i++) {
    auto word = getWord();
    if (i % 2 != 0) {
      for (size_t j = 0; j < wordSize; j++) {
        char *ch = (char *)word;
        while (*ch) {
          *ch = toupper(*ch);
          ch++;
        }
      }
    }
    printf("%s", word);
  }
  printf("\n");
}
