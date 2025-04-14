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

cluster1-1-3: build
	./bin/bidder -bidder.address=b_3 -rollup.id=r_1
.PHONY: cluster1-1-3

cluster1-1-4: build
	./bin/bidder -bidder.address=b_4 -rollup.id=r_1
.PHONY: cluster1-1-4

cluster1-1-5: build
	./bin/bidder -bidder.address=b_5 -rollup.id=r_1
.PHONY: cluster1-1-5

cluster1-1-6: build
	./bin/bidder -bidder.address=b_6 -rollup.id=r_1
.PHONY: cluster1-1-2

cluster1-1-7: build
	./bin/bidder -bidder.address=b_7 -rollup.id=r_1
.PHONY: cluster1-1-7

cluster1-1-8: build
	./bin/bidder -bidder.address=b_8 -rollup.id=r_1
.PHONY: cluster1-1-8

cluster1-1-9: build
	./bin/bidder -bidder.address=b_9 -rollup.id=r_1
.PHONY: cluster1-1-9

cluster1-1-10: build
	./bin/bidder -bidder.address=b_10 -rollup.id=r_1
.PHONY: cluster1-1-10

cluster1-1-11: build
	./bin/bidder -bidder.address=b_11 -rollup.id=r_1
.PHONY: cluster1-1-11

cluster1-2-1: build
	./bin/bidder -bidder.address=b_3 -rollup.id=r_2
.PHONY: cluster1-2-1