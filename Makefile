check-swagger:
	./install.sh

swagger: check-swagger
	swagger generate spec -o ./swagger.yaml --scan-models

serve-swagger: swagger
	swagger serve -F=swagger swagger.yaml

dev-up:
	docker-compose up --build -d
	
dev-down:
	docker-compose down

create-mock:
	mockery --all --dir api --dir repo --dir service --recursive --output ./mocks

test: create-mock
	go test ./...

db-up:
	docker-compose up -d db

db-down:
	docker-compose down