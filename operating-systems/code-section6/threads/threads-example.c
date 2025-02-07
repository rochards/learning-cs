#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <unistd.h>

void* thread_function() {
    int sleep_time = rand() % 10 + 1;
    pthread_t tid = pthread_self();
    pid_t pid = getpid();

    printf("Thread started: TID = %lu, PID = %d\n", tid, pid);
    printf("Thread (TID = %lu) sleeping for %d seconds...\n", tid, sleep_time);
    
    sleep(sleep_time);

    printf("Thread (TID = %lu, PID = %d) woke up!\n", tid, pid);
    pthread_exit(NULL);
}

int main() {
    srand(time(NULL));

    pthread_t thread1, thread2;

    if (pthread_create(&thread1, NULL, thread_function, NULL) != 0) {
        perror("Failed to create thread 1");
        return 1;
    }

    if (pthread_create(&thread2, NULL, thread_function, NULL) != 0) {
        perror("Failed to create thread 2");
        return 1;
    }

    pthread_join(thread1, NULL);
    pthread_join(thread2, NULL);

    printf("Both threads have finished execution.\n");
    return 0;
}
