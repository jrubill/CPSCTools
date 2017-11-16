# Makefile for C++

CC = g++

CFLAGS = -Wall -std=c++11

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
