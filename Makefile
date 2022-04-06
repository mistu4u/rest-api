check-swagger:
	./install.sh

swagger: check-swagger
	GO111MODULE=on go mod vendor  && GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models

serve-swagger: 
	swagger serve -F=swagger swagger.yaml

dev-up:
	docker-compose up --build -d
	
dev-down:
	docker-compose down
    