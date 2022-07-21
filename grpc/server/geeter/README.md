# Geeter Service

This is the Geeter service

Generated with

```
micro new --namespace=go.micro --type=service geeter
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.service.geeter
- Type: service
- Alias: geeter

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./geeter-service
```

Build a docker image
```
make docker
```