## mockgen-example
This repo demos mockgen and why we might wanna use it.



### Setup
- Install mockgen
    - `go install go.uber.org/mock/mockgen@latest`
> Note: Uber-go mock succeeds golang-mock. golang-mock is deprecated and should not be used.
- Mocks can be regenerated with `go generate ./...`

### Walkthough
- Give `service.go` a look to understand the code that will be undertest
- Look at `dependency.go` to see what'll get mocked and how it gets mocked automatically with mockgen & go:generate
- Read through `service_test.go` to see the basic nuts & bolts and the problems with it , then `service_suite_test.go` for the full enchilada.