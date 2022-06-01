# Need to generate
# - Go bindings
# - Typescript bindings
# - ABI

ABI_DIR=abigenBindings

generate_go_bindings:
	sh scripts/generate_go_bindings.sh $(ABI_DIR) abigen

build_server: