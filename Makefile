# Need to generate
# - Go bindings
# - Typescript bindings
# - ABI

gen_go_bindings:
	abigen --pkg abigen --sol ./contracts/*.sol --out ./abigen/*.go

gen_go_bindings_2:
	abigen --abi=abigenBindings/abi/*.abi --pkg abigen --out *


gen_ts_bindings:
	npx hardhat typechain

gen_abi_json:
