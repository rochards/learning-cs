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
    
    Packet *packets = (Packet *)malloc(PACKET_ARRAY_LENGTH * sizeof(Packet)); // just one syscall involved
    Packet *p2 = (packets + 1); // the compiler will do the math based on the size of Packet to access packets like an array

    free(packets); // just on syscall involved

    return 0;
}


