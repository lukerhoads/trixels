# Need to generate
# - Go bindings
# - Typescript bindings
# - ABI

gen_go_bindings:
	abigen --pkg abigen --sol ./contracts/*.sol --out ./abigen/*.go