## Install Etcd:

Run the etcd container:

    docker run --name etcd -p 2379:2379 -d etcd
    docker exec -it etcd /bin/bash

Restart Etcd:

    docker restart etcd

## Use etcdctl

``` bash
export ETCDCTL_API=3
# write
etcdctl put /micro/config "{ \"Network\": \"172.30.0.0/16\", \"Backend\": \"vxlan\"}"

# list
etcdctl ls /micro/config

# read value
etcdctl get /micro/config
or
etcdctll get --prefix /micro/config

# delete key
etcdctl del /micro/config

# deep del
etcdctl rm --recursive registry

# modify
etcdctl mk /micro/config "{ \"Network\": \"127.0.0.1/16\", \"Backend\": \"xlan\"}"

# backup 
etcdctl backup

# health
etcdctl endpoint health
```

## User etcdctl backup shapshot

``` bash
etcdctl snapshot save test.db
etcdctl  snapshot restore test.db --data-dir=default.etcd
```

## Run:

``` go
go run main
or
go test
```
