#!/bin/bash
docker exec go1.19 /bin/bash -c "sh /var/www/store/build/dev-build.sh"
#重启容器
docker restart go1.19
docker restart api
docker restart rpc-api
docker restart rpc-socket
docker restart websocket-1