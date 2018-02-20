# kubicorn-example-server


## Starting the server
To start a server please run

```sbtshell
go run ./cmd/clusterserver/main.go
```


## Running a client
You may write a client as you wish, here we have two sample clients, each can be run as follows

### bring up a cluster, (creating and apply)
```sbtshell
CLUSTER_CREATOR_STEP=up go run ./cmd/clientkubicorn/client.go
```
or

```sbtshell
CLUSTER_CREATOR_STEP=up go run ./cmd/clientkubicorncli/client.go

```

### bring down a cluster, (creating and apply)
```sbtshell
CLUSTER_CREATOR_STEP=down go run ./cmd/clientkubicorn/client.go
```
or

```sbtshell
CLUSTER_CREATOR_STEP=down go run ./cmd/clientkubicorncli/client.go

```


## gRPC-gateway
Not implemented yet, to be done soon. 


## Development details


### Adding a new provider
The code allows you to add more providers and not just use kubicorn. To do so you have to inherit the methods of the 
provider interface, namely:

```sbtshell
type Provider interface {
	Apply() (*api.ClusterStatusMsg, error) // should allow you to execute on cluster state, actual state should be reconciled.
	Create() (*api.ClusterStatusMsg, error) // create a default cluster config file, usually local.
	Delete() (*api.ClusterStatusMsg, error) // delete a cluster, currently also destroys config file. 
}
```

