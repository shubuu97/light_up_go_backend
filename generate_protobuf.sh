#!/usr/bin/env bash

protoc -I light-up-backend/proto  --go_out=. --micro_out=. light-up-backend/proto/common.proto
protoc -I light-up-backend/proto  --go_out=. --micro_out=. light-up-backend/proto/authentication.proto
protoc -I light-up-backend/proto  --go_out=. --micro_out=. light-up-backend/proto/lighter.proto
protoc -I light-up-backend/proto  --go_out=. --micro_out=. light-up-backend/proto/lightSeeker.proto
protoc -I light-up-backend/proto  --go_out=. --micro_out=. light-up-backend/proto/admin.proto
protoc -I light-up-backend/proto  --go_out=. --micro_out=. light-up-backend/proto/entity.proto