This app is a micro service based app made with Go lang using Gin gonic.

To run it pls execute the below commands, 

1) Build the dependency tree. Install [wire](https://github.com/google/wire/tree/main) using the command 

    ```go install github.com/google/wire/cmd/wire@latest```

2) Create the mocks folder for running the tests. Install [Mockery](https://github.com/vektra/mockery)
    ```go install github.com/vektra/mockery/v2@latest```

3) Run the tests:

    ```go test ./...```

4) run it 

    ```go run main.go```

Using docker:

1) Simply type ```docker compose up``` to bring up the services.
2) ```docker compose down``` to stop the services.