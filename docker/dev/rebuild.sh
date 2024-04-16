#!/bin/bash
docker exec k-1.19 /bin/bash -c "sh /var/www/k/build/dev-build.sh"
#重启容器
docker restart k-api
docker restart k-r-api