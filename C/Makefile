# Makefile for C

CC = gcc

CFLAGS = -Wall

INCLUDE = *.h

SRC = *.c

TARGET = main

all:
	$(CC) $(CFLAGS) $(INCLUDE) $(SRC) -o $(TARGET)

clean:
	rm -f *.o
	rm main

run:
	./$(TARGET)
