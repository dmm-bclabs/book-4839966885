#!/bin/sh

set -e

host="$1"
shift
cmd="$@"

until mysqladmin ping -h $host --silent; do
  echo 'Waiting for mysql'
  sleep 1
done

echo "MySQL is up!"
exec $cmd