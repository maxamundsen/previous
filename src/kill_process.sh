#!/bin/sh

kill $(lsof -nP -iTCP -sTCP:LISTEN | grep 8080 | awk '{print $2}')