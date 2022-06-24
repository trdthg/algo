#include <pthread.h>
#include <semaphore.h>
#include <stdbool.h>
#include <stdio.h>
sem_t empty;
sem_t mutex;
sem_t stuffs;
sem_t reader;

int n = 0;
int m = 0;
int t = 0;
int N = 100000;

void *produce(void *p) {
  while (n < N) {
    sem_wait(&empty); // 缓冲区有空的
    sem_wait(&mutex);
    if (n >= N) {
      sem_post(&mutex);
      sem_post(&stuffs); // 缓冲区装一个
      break;
    }
    n++;
    t++;
    sem_post(&mutex);
    sem_post(&stuffs); // 缓冲区装一个
  }
  printf("Producer: [%d %d]\n", n, m);
  return 0;
}
void *consume(void *p) {
  while (true) {
    // 究级大锁
    sem_wait(&reader);

    if (m >= N) {
      sem_post(&reader);
      break;
    }

    sem_wait(&stuffs); // 有东西
    sem_wait(&mutex);  // 锁缓冲区
    m++;
    t--;
    sem_post(&mutex); // 那缓冲区
    sem_post(&empty); // 减一个东西

    sem_post(&reader); // 大锁结束
  }
  printf("Consumer: [%d %d %d]\n", n, m, t);
  return 0;
}

int main() {
  sem_init(&mutex, 0, 1);
  sem_init(&reader, 0, 1);
  sem_init(&empty, 0, 2);
  sem_init(&stuffs, 0, 0);
  pthread_t t1;
  pthread_t t2;
  pthread_t t3;
  pthread_t t4;
  pthread_create(&t1, NULL, produce, NULL);
  pthread_create(&t3, NULL, produce, NULL);
  pthread_create(&t2, NULL, consume, NULL);
  pthread_create(&t4, NULL, consume, NULL);

  int N = 3;
  pthread_t arr[N];
  for (int i = 0; i < N; i++) {
    pthread_create(&arr[i], NULL, consume, NULL);
  }

  pthread_exit(0);
  return 0;
}