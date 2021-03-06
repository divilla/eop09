#!/bin/bash -e

exec > >(tee -a /var/log/app/entry.log|logger -t server -s 2>/dev/console) 2>&1

APP_ENV=${APP_ENV:-local}

echo "[`date`] Running entrypoint script in the '${APP_ENV}' environment..."

CONFIG_FILE=./config/${APP_ENV}.yml

if [[ -z ${APP_PORT} ]]; then
  export APP_PORT=`sed -n 's/^server_port: : *\(.*\)/\1/p' ${CONFIG_FILE}`
fi

if [[ -z ${APP_DSN} ]]; then
  export APP_DSN=`sed -n 's/^ports_dsn: *\(.*\)/\1/p' ${CONFIG_FILE}`
fi

echo "[`date`] Starting server..."
./server -mode ${APP_ENV} >> /var/log/app/server.log 2>&1
