# Should be equivalent to your list of C files, if you don't build selectively
SRC=$(wildcard *.go)
build: $(SRC)
	clear
	go build -o dominoes $(SRC)

clean:
	rm dominoes
