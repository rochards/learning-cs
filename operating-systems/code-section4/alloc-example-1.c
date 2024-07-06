#include <stdlib.h>

typedef struct
{
    char source_ip[16];
    char destination_ip[16];
    int source_port;
    int destination_port;
    int packet_length;
} Packet;

#define PACKET_ARRAY_LENGTH 1000000

int main() {
    
    Packet *packets[PACKET_ARRAY_LENGTH];
    for (int i = 0; i < PACKET_ARRAY_LENGTH; i++) {
        packets[i] = (Packet *) malloc(sizeof(Packet)); // 1,000,000 of system calls involved
    }

    for (int i = 0; i < PACKET_ARRAY_LENGTH; i++) {
        free(packets[i]); // 1,000,000 of system calls involved
    }

    return 0;
}


