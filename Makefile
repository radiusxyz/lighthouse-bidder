CLUSTER1-1-1_ROLLUP_IDS ?= "r_1"
CLUSTER1-1-1_ADDRESS ?= 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65

CLUSTER1-1-2_ROLLUP_IDS ?= "r_1"
CLUSTER1-1-2_ADDRESS ?= 0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc

CLUSTER1-1-3_ROLLUP_IDS ?= "r_1"
CLUSTER1-1-3_ADDRESS ?= 0x976EA74026E726554dB657fA54763abd0C3a0aa9

CLUSTER1-1-4_ROLLUP_IDS ?= "r_1"
CLUSTER1-1-4_ADDRESS ?= 0x14dC79964da2C08b23698B3D3cc7Ca32193d9955

CLUSTER1-1-5_ROLLUP_IDS ?= "r_1"
CLUSTER1-1-5_ADDRESS ?= 0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f

CLUSTER1-1-6_ROLLUP_IDS ?= "r_1"
CLUSTER1-1-6_ADDRESS ?= 0xBcd4042DE499D14e55001CcbB24a551F3b954096

CLUSTER1-1-7_ROLLUP_IDS ?= "r_1"
CLUSTER1-1-7_ADDRESS ?= 0x71bE63f3384f5fb98995898A86B02Fb2426c5788

CLUSTER1-1-8_ROLLUP_IDS ?= "r_1"
CLUSTER1-1-8_ADDRESS ?= 0xFABB0ac9d68B0B445fB7357272Ff202C5651694a

CLUSTER1-1-9_ROLLUP_IDS ?= "r_1"
CLUSTER1-1-9_ADDRESS ?= 0x1CBd3b2770909D4e10f157cABC84C7264073C9Ec

CLUSTER1-1-10_ROLLUP_IDS ?= "r_1"
CLUSTER1-1-10_ADDRESS ?= 0xdF3e18d64BC6A983f673Ab319CCaE4f1a57C7097

CLUSTER1-1-11_ROLLUP_IDS ?= "r_1"
CLUSTER1-1-11_ADDRESS ?= 0xcd3B766CCDd6AE721141F452C550Ca635964ce71

CLUSTER1-2-1_ROLLUP_IDS ?= "r_2"
CLUSTER1-2-1_ADDRESS ?= 0x2546BcD3c84621e976D8185a91A922aE77ECEc30

build:
	go build -o ./bin/bidder

run: build
	./bin/bidder -bidder.address=$(CLUSTER1-1-1_ADDRESS) -rollup.ids=$(CLUSTER1-1-1_ROLLUP_IDS)
.PHONY: run

cluster1-1-1: build
	./bin/bidder -bidder.address=$(CLUSTER1-1-1_ADDRESS) -rollup.ids=$(CLUSTER1-1-1_ROLLUP_IDS)
.PHONY: cluster1-1-1

cluster1-1-2: build
	./bin/bidder -bidder.address=$(CLUSTER1-1-2_ADDRESS) -rollup.ids=$(CLUSTER1-1-2_ROLLUP_IDS)
.PHONY: cluster1-1-2

cluster1-1-3: build
	./bin/bidder -bidder.address=$(CLUSTER1-1-3_ADDRESS) -rollup.ids=$(CLUSTER1-1-3_ROLLUP_IDS)
.PHONY: cluster1-1-3

cluster1-1-4: build
	./bin/bidder -bidder.address=$(CLUSTER1-1-4_ADDRESS) -rollup.ids=$(CLUSTER1-1-4_ROLLUP_IDS)
.PHONY: cluster1-1-4

cluster1-1-5: build
	./bin/bidder -bidder.address=$(CLUSTER1-1-5_ADDRESS) -rollup.ids=$(CLUSTER1-1-5_ROLLUP_IDS)
.PHONY: cluster1-1-5

cluster1-1-6: build
	./bin/bidder -bidder.address=$(CLUSTER1-1-6_ADDRESS) -rollup.ids=$(CLUSTER1-1-6_ROLLUP_IDS)
.PHONY: cluster1-1-2

cluster1-1-7: build
	./bin/bidder -bidder.address=$(CLUSTER1-1-7_ADDRESS) -rollup.ids=$(CLUSTER1-1-7_ROLLUP_IDS)
.PHONY: cluster1-1-7

cluster1-1-8: build
	./bin/bidder -bidder.address=$(CLUSTER1-1-8_ADDRESS) -rollup.ids=$(CLUSTER1-1-8_ROLLUP_IDS)
.PHONY: cluster1-1-8

cluster1-1-9: build
	./bin/bidder -bidder.address=$(CLUSTER1-1-9_ADDRESS) -rollup.ids=$(CLUSTER1-1-9_ROLLUP_IDS)
.PHONY: cluster1-1-9

cluster1-1-10: build
	./bin/bidder -bidder.address=$(CLUSTER1-1-10_ADDRESS) -rollup.ids=$(CLUSTER1-1-10_ROLLUP_IDS)
.PHONY: cluster1-1-10

cluster1-1-11: build
	./bin/bidder -bidder.address=$(CLUSTER1-1-11_ADDRESS) -rollup.ids=$(CLUSTER1-1-11_ROLLUP_IDS)
.PHONY: cluster1-1-11

cluster1-2-1: build
	./bin/bidder -bidder.address=$(CLUSTER1-2-1_ADDRESS) -rollup.ids=$(CLUSTER1-2-1_ROLLUP_IDS)
.PHONY: cluster1-2-1