#!/bin/bash
docker-compose -f docker-compose-build.yml up -d --build
docker exec k-build /bin/bash -c "sh /var/www/k/build/dev-build.sh"
docker rm -f k-build
docker-compose -f docker-compose.yml up -d --build