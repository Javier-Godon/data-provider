expose questdb:
sql: k port-forward svc/questdb-single-local -n questdb 31881:8812
web console: k port-forward svc/questdb-single-local -n questdb 31333:9000

over app:
go mod init data-provider

It will create go.mod

to build the executable. over app
go build .

then we can run the application executing
./data-provider

we can run it withot building it:
go run main.go

to add dependencies to go.mod from imports run:
go get
