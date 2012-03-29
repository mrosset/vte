CFLAGS	+= -g $(shell pkg-config --cflags vte)
LDFLAGS += -g $(shell pkg-config --libs vte)

all: clean vte
	./vte

vte:

clean:
	rm vte
