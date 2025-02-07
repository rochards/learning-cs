CC = gcc

CFLAGS = -Wall -Wextra

TARGET = processes-example

SRC = processes-example.c

all: $(TARGET)

$(TARGET): $(SRC)
	$(CC) $(SRC) -o $(TARGET) $(CFLAGS)

clean:
	rm -f $(TARGET)

.PHONY: all clean

