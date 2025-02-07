CC = gcc

CFLAGS = -Wall -Wextra

TARGET = threads-example

SRC = threads-example.c

all: $(TARGET)

$(TARGET): $(SRC)
	$(CC) $(SRC) -o $(TARGET) $(CFLAGS)

clean:
	rm -f $(TARGET)

.PHONY: all clean
