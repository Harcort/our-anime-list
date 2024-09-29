#!/usr/bin/env sh

for e in $(cat /env);
    do export $e;
done

if [ -d /src ]; then
    cd /src
    while true; do
        go build
        ./backend
    done
fi

./app