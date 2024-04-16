#!/bin/bash
docker exec k-1.19 /bin/bash -c "sh /var/www/k/build/dev-build.sh"
docker-compose -f docker-compose.yml up -d --build