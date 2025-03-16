# data-provider
queries a QuestDB database and returns results via gRPC

as recommended in [QuestDB documentation](https://questdb.com/docs/reference/api/postgres/) we will use [pq](https://github.com/lib/pq) for querying QuestDB.

We will keep all .proto files under the same folder (proto) following the standads, but as I am following a vertical slice architecture, the generated classes should be generated in its own folder (usecase/feature)

from the root of the project (data-provider) we will execute for each usecase (feature):

```
protoc --proto_path=proto   --go_out=app/usecases/cpu/get_cpu_user_usage/grpc   --go-grpc_out=app/usecases/cpu/get_cpu_user_usage/grpc   proto/get_cpu_user_usage.proto

protoc --proto_path=proto   --go_out=app/usecases/cpu/get_cpu_system_usage/grpc   --go-grpc_out=app/usecases/cpu/get_cpu_system_usage/grpc   proto/get_cpu_system_usage.proto
```

so:

### Generating gRPC Code
Run the following command to regenerate gRPC files:

```
make generate

```

