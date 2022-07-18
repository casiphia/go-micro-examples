# Install Etcd

Run the etcd container:

    docker run --name etcd -d -p 2379:2379 -p 2380:2380 -e ALLOW_NONE_AUTHENTICATION=yes bitnami/etcd:3.3.11 etcd
    docker exec -it etcd /bin/bash

Use V3:

    export ETCDCTL_API=3

Get All key:

    etcdctl get  / --prefix

Restart Etcd:

    docker restart etcd

## Use etcdctl

    # write
    etcdctl put /micro/config/demo "{ \"network\": \"172.30.0.0/16\", \"backend\": \"vxlan\"}"

    # list
    etcdctl ls /micro/config/demo

    # read value
    etcdctl get /micro/config/demo
    or
    etcdctll get --prefix /micro/config/demo

    # delete key
    etcdctl del /micro/config/demo

    # deep del
    etcdctl rm --recursive registry

    # modify
    etcdctl mk /micro/config/demo "{ \"network\": \"127.0.0.1/16\", \"backend\": \"xlan\"}"

    # backup
    etcdctl backup

    # health
    etcdctl endpoint health

## User etcdctl backup shapshot

    etcdctl snapshot save test.db
    etcdctl  snapshot restore test.db --data-dir=default.etcd

## Run

    export MICRO_REGISTRY=etcd
    export MICRO_REGISTRY_ADDRESS="127.0.0.1:2379" 
    go run main
    or
    go test
