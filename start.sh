#!/bin/bash

rm -rf /tmp/light-up-*

listVar="authentication-service entity-service light-seeker-service lighter-service admin-service"
for i in $listVar; do
    go run "$i"/main.go &
done