# Need to generate
# - Go bindings
# - Typescript bindings
# - ABI

ABI_DIR=abigenBindings/abi

generate_go_bindings:
	sh scripts/generate_go_bindings.sh ABI_DIR

build_server: