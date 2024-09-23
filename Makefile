.PHONY: clean

clean:
	rm -f ./cdr

cdr: clean
	go build -o cdr ./cmd