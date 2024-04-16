#!/bin/bash
docker-compose -f docker-compose-build.yml up -d --build
docker exec go1.19-build /bin/bash -c "sh /var/www/store/build/dev-build.sh"
docker rm -f go1.19-build
docker-compose -f docker-compose.yml up -d --build