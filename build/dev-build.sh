#!/bin/bash
cd /var/www/store/build/dev && ./api.sh && ./rpc_api.sh && ./rpc_socket.sh && ./websocket.sh
