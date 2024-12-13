#!/bin/bash

kill $(lsof -nP -iTCP -sTCP:LISTEN | grep 9090 | awk '{print $2}')
