check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger:
	@echo Ensure you have the swagger CLI or this command will fail.
	@echo You can install the swagger CLI with: go get -u github.com/go-swagger/go-swagger/cmd/swagger
	@echo ....

	swagger generate spec -o ./swagger.yaml --scan-models

generate_client:
	swagger generate client -f ./swagger.yaml -A product_api