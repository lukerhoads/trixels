gen_abi:
	abigen --pkg abigen --sol ./contracts/*.sol --out ./abigen/*.go

gen_abi_2:
	abigen --abi=abigenBindings/abi/*.abi --pkg abigen --out *.go