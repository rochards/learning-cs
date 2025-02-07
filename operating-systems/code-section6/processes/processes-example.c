#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/wait.h>
#include <time.h>

void process_function() {
    pid_t pid = getpid();
    pid_t parent_pid = getppid();

    srand(time(NULL) ^ pid);
    int sleep_time = rand() % 10 + 1;

    printf("Process started: PID = %d, Parent PID = %d\n", pid, parent_pid);
    printf("Process (PID = %d) sleeping for %d seconds...\n", pid, sleep_time);
    
    sleep(sleep_time);

    printf("Process (PID = %d) woke up!\n", pid);
}

int main() {
    pid_t pid1, pid2;

    pid1 = fork();
    if (pid1 < 0) {
        perror("Failed to create process 1");
        return 1;
    } else if (pid1 == 0) {
        process_function();
        exit(0);
    }

    pid2 = fork();
    if (pid2 < 0) {
        perror("Failed to create process 2");
        return 1;
    } else if (pid2 == 0) {
        process_function();
        exit(0);
    }

    waitpid(pid1, NULL, 0);
    waitpid(pid2, NULL, 0);

    printf("Both processes have finished execution.\n");
    return 0;
}
