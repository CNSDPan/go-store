#!/bin/bash
docker stop api
docker stop rpc-api
docker stop rpc-socket
docker stop websocket-1
docker stop etcd1
docker stop etcd2