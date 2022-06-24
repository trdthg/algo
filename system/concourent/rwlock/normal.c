#include "stdbool.h"
#include <pthread.h>
#include <semaphore.h>
#include <stdio.h>

sem_t write_lock;
sem_t can_read;
sem_t read_lock;
int reader_count = 0;

void *read(void *p) {
  sem_wait(&read_lock);
  if (reader_count == 0) {
    sem_wait(&write_lock);
  }
  reader_count += 1;
  sem_post(&read_lock);
  printf("reader(read) : %d\n", reader_count);
  sem_wait(&read_lock);
  reader_count -= 1;
  if (reader_count == 0) {
    sem_post(&write_lock);
  }
  sem_post(&read_lock);
  return 0;
}
// write 可能会阻塞
void *write(void *p) {
  sem_wait(&write_lock);
  printf("reader(write): %d\n", reader_count);
  sem_post(&write_lock);
  return 0;
}

int main() {
  sem_init(&write_lock, 0, 1);
  sem_init(&read_lock, 0, 1);
  // N 个 Reader
  int N = 3;
  pthread_t arr[N];
  for (int i = 0; i < N; i++) {
    pthread_create(&arr[i], NULL, read, NULL);
  }
  // 1 个 Writer
  pthread_t t;
  pthread_create(&t, NULL, write, NULL);
  pthread_exit(0);
  return 0;
}