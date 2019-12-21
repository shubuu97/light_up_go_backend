#!/bin/bash

rm -rf /tmp/light-up-*

listVar="authentication-service lighter-service light-seeker-service "
for i in $listVar; do
    go run "$i"/main.go &
done