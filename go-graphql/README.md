## Project with Go && GraphQL

Steps:
1. Create a directory for project and initialize go modules file:
    `go mod init github.com/[username]/hackernews`

2. after that use ‍‍gqlgen init command to setup a gqlgen project.
    `go run github.com/99designs/gqlgen init`

3. Define schema

4. Now run the following command to regenerate files;
    `go run github.com/99designs/gqlgen generate`


Test server:

1. Clone repository
    `git clone <SSH>`

2. Install relative package  
    `git mod tidy`

3. Start server
    `go run server.go`