# Need to generate
# - Go bindings
# - Typescript bindings
# - ABI

ABI_DIR=abigenBindings

setup_contracts:
	yarn types 

generate_go_bindings:
	sh scripts/generate_go_bindings.sh $(ABI_DIR) abigen

build_server:
	go build main.go

dev:
	yarn dev:deploy 
	go run main.go