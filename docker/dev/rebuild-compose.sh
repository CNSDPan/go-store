#!/bin/bash
docker exec go1.19 /bin/bash -c "sh /var/www/store/build/dev-build.sh"
docker-compose -f docker-compose.yml up -d --build