build:
	go build -o ./bin/bidder

run: build
	./bin/bidder -bidder.address=b_1 -rollup.id=r_1
.PHONY: run

cluster1-1-1: build
	./bin/bidder -bidder.address=b_1 -rollup.id=r_1
.PHONY: cluster1-1-1

cluster1-1-2: build
	./bin/bidder -bidder.address=b_2 -rollup.id=r_1
.PHONY: cluster1-1-2

cluster1-2-1: build
	./bin/bidder -bidder.address=b_3 -rollup.id=r_2
.PHONY: cluster1-2-1