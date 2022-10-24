# Thirumathikart-Messaging-Service
___

### Requirements
* [Go](https://go.dev/)
* [golangci](https://golangci-lint.run/usage/install/)
* [Docker](https://www.docker.com/)
* [protobuff](https://grpc.io/docs/protoc-installation/)

### Setup
* Configure .vscode/settings.json
    ```
    {
        "go.lintTool":"golangci-lint",
        "go.lintFlags": [
        "--fast"
        ],
        "go.lintOnSave": "package",
        "go.formatTool": "goimports",
        "go.useLanguageServer": true,
        "[go]": {
            "editor.formatOnSave": true,
            "editor.codeActionsOnSave": {
                "source.organizeImports": true
            }
        },
        "go.docsTool": "gogetdoc"
    }
    ```
* Create .env file
    ``` sh
    cp .env.example .env
    ```
* Enable githooks
    ``` sh
    git config core.hooksPath .githooks
    ```
* Add the following in ~/.bash_profile
    ```sh
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOPATH/bin
    ```
* Install Protobuff Dependencies
   ```sh
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   ```
   ```sh
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```       


* Configure database settings in config.json

### Run
- #### On Docker
    ``` sh
    docker-compose up
    ```    
   