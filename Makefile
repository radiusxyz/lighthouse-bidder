CLUSTER1-1-1_ROLLUP_IDS ?= "cluster1-1"
CLUSTER1-1-1_ADDRESS ?= 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65
CLUSTER1-1-1_PRIVATE_KEY ?= 0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a

CLUSTER1-1-2_ROLLUP_IDS ?= "cluster1-1"
CLUSTER1-1-2_ADDRESS ?= 0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc
CLUSTER1-1-2_PRIVATE_KEY ?= 0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba

CLUSTER1-1-3_ROLLUP_IDS ?= "cluster1-1"
CLUSTER1-1-3_ADDRESS ?= 0x976EA74026E726554dB657fA54763abd0C3a0aa9
CLUSTER1-1-3_PRIVATE_KEY ?= 0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e

CLUSTER1-1-4_ROLLUP_IDS ?= "cluster1-1"
CLUSTER1-1-4_ADDRESS ?= 0x14dC79964da2C08b23698B3D3cc7Ca32193d9955
CLUSTER1-1-4_PRIVATE_KEY ?= 0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356

CLUSTER1-1-5_ROLLUP_IDS ?= "cluster1-1"
CLUSTER1-1-5_ADDRESS ?= 0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f
CLUSTER1-1-5_PRIVATE_KEY ?= 0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97

CLUSTER1-1-6_ROLLUP_IDS ?= "cluster1-1"
CLUSTER1-1-6_ADDRESS ?= 0xBcd4042DE499D14e55001CcbB24a551F3b954096
CLUSTER1-1-6_PRIVATE_KEY ?= 0xf214f2b2cd398c806f84e317254e0f0b801d0643303237d97a22a48e01628897

CLUSTER1-1-7_ROLLUP_IDS ?= "cluster1-1"
CLUSTER1-1-7_ADDRESS ?= 0x71bE63f3384f5fb98995898A86B02Fb2426c5788
CLUSTER1-1-7_PRIVATE_KEY ?= 0x701b615bbdfb9de65240bc28bd21bbc0d996645a3dd57e7b12bc2bdf6f192c82

CLUSTER1-1-8_ROLLUP_IDS ?= "cluster1-1"
CLUSTER1-1-8_ADDRESS ?= 0xFABB0ac9d68B0B445fB7357272Ff202C5651694a
CLUSTER1-1-8_PRIVATE_KEY ?= 0xa267530f49f8280200edf313ee7af6b827f2a8bce2897751d06a843f644967b1

CLUSTER1-1-9_ROLLUP_IDS ?= "cluster1-1"
CLUSTER1-1-9_ADDRESS ?= 0x1CBd3b2770909D4e10f157cABC84C7264073C9Ec
CLUSTER1-1-9_PRIVATE_KEY ?= 0x47c99abed3324a2707c28affff1267e45918ec8c3f20b8aa892e8b065d2942dd

CLUSTER1-1-10_ROLLUP_IDS ?= "cluster1-1"
CLUSTER1-1-10_ADDRESS ?= 0xdF3e18d64BC6A983f673Ab319CCaE4f1a57C7097
CLUSTER1-1-10_PRIVATE_KEY ?= 0xc526ee95bf44d8fc405a158bb884d9d1238d99f0612e9f33d006bb0789009aaa

CLUSTER1-1-11_ROLLUP_IDS ?= "cluster1-1"
CLUSTER1-1-11_ADDRESS ?= 0xcd3B766CCDd6AE721141F452C550Ca635964ce71
CLUSTER1-1-11_PRIVATE_KEY ?= 0x8166f546bab6da521a8369cab06c5d2b9e46670292d85c875ee9ec20e84ffb61

CLUSTER1-2-1_ROLLUP_IDS ?= "cluster1-2"
CLUSTER1-2-1_ADDRESS ?= 0x2546BcD3c84621e976D8185a91A922aE77ECEc30
CLUSTER1-2-1_PRIVATE_KEY ?= 0xea6c44ac03bff858b476bba40716402b03e41b8e97e276d1baec7c37d42484a0

build:
	go build -o ./bin/bidder

run: build
	./bin/bidder -bidder.address=$(CLUSTER1-1-1_ADDRESS) -bidder.private.key=$(CLUSTER1-1-1_PRIVATE_KEY) -rollup.ids=$(CLUSTER1-1-1_ROLLUP_IDS)
.PHONY: run

cluster1-1-1: build
	./bin/bidder -bidder.address=$(CLUSTER1-1-1_ADDRESS) -bidder.private.key=$(CLUSTER1-1-1_PRIVATE_KEY) -rollup.ids=$(CLUSTER1-1-1_ROLLUP_IDS)
.PHONY: cluster1-1-1

cluster1-1-2: build
	./bin/bidder -bidder.address=$(CLUSTER1-1-2_ADDRESS) -bidder.private.key=$(CLUSTER1-1-2_PRIVATE_KEY) -rollup.ids=$(CLUSTER1-1-2_ROLLUP_IDS)
.PHONY: cluster1-1-2

cluster1-1-3: build
	./bin/bidder -bidder.address=$(CLUSTER1-1-3_ADDRESS) -bidder.private.key=$(CLUSTER1-1-3_PRIVATE_KEY) -rollup.ids=$(CLUSTER1-1-3_ROLLUP_IDS)
.PHONY: cluster1-1-3

cluster1-1-4: build
	./bin/bidder -bidder.address=$(CLUSTER1-1-4_ADDRESS) -bidder.private.key=$(CLUSTER1-1-4_PRIVATE_KEY) -rollup.ids=$(CLUSTER1-1-4_ROLLUP_IDS)
.PHONY: cluster1-1-4

cluster1-1-5: build
	./bin/bidder -bidder.address=$(CLUSTER1-1-5_ADDRESS) -bidder.private.key=$(CLUSTER1-1-5_PRIVATE_KEY) -rollup.ids=$(CLUSTER1-1-5_ROLLUP_IDS)
.PHONY: cluster1-1-5

cluster1-1-6: build
	./bin/bidder -bidder.address=$(CLUSTER1-1-6_ADDRESS) -bidder.private.key=$(CLUSTER1-1-6_PRIVATE_KEY) -rollup.ids=$(CLUSTER1-1-6_ROLLUP_IDS)
.PHONY: cluster1-1-2

cluster1-1-7: build
	./bin/bidder -bidder.address=$(CLUSTER1-1-7_ADDRESS) -bidder.private.key=$(CLUSTER1-1-7_PRIVATE_KEY) -rollup.ids=$(CLUSTER1-1-7_ROLLUP_IDS)
.PHONY: cluster1-1-7

cluster1-1-8: build
	./bin/bidder -bidder.address=$(CLUSTER1-1-8_ADDRESS) -bidder.private.key=$(CLUSTER1-1-8_PRIVATE_KEY) -rollup.ids=$(CLUSTER1-1-8_ROLLUP_IDS)
.PHONY: cluster1-1-8

cluster1-1-9: build
	./bin/bidder -bidder.address=$(CLUSTER1-1-9_ADDRESS) -bidder.private.key=$(CLUSTER1-1-9_PRIVATE_KEY) -rollup.ids=$(CLUSTER1-1-9_ROLLUP_IDS)
.PHONY: cluster1-1-9

cluster1-1-10: build
	./bin/bidder -bidder.address=$(CLUSTER1-1-10_ADDRESS) -bidder.private.key=$(CLUSTER1-1-10_PRIVATE_KEY) -rollup.ids=$(CLUSTER1-1-10_ROLLUP_IDS)
.PHONY: cluster1-1-10

cluster1-1-11: build
	./bin/bidder -bidder.address=$(CLUSTER1-1-11_ADDRESS) -bidder.private.key=$(CLUSTER1-1-11_PRIVATE_KEY) -rollup.ids=$(CLUSTER1-1-11_ROLLUP_IDS)
.PHONY: cluster1-1-11

cluster1-2-1: build
	./bin/bidder -bidder.address=$(CLUSTER1-2-1_ADDRESS) -bidder.private.key=$(CLUSTER1-2-1_PRIVATE_KEY) -rollup.ids=$(CLUSTER1-2-1_ROLLUP_IDS)
.PHONY: cluster1-2-1

build-contract:
	solc --abi --bin contracts/ILighthouse.sol -o contracts/build --overwrite
	abigen --abi contracts/build/ILighthouse.abi --bin contracts/build/ILighthouse.bin --pkg bindings --out contracts/bindings/ILighthouse.go
.PHONY: build-contract