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
	mockery --all --recursive --output ./mocks

db-up:
	docker-compose up -d db

db-down:
	docker-compose down