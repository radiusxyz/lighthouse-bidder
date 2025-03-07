build:
	go build -o ./bin/bidder

run: build
	./bin/bidder -bidder.address=b_1
.PHONY: run