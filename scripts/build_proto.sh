#!/usr/bin/env bash

# protoc --proto_path=src/zk_leader_election/idl --go_out=src/zk_leader_election/gen_src


protoc \
--go_out=:src/zk_leader_election/idl \
--go_opt=module=github.com/cxp863/zk_leader_election/src/zk_leader_election/idl \
--go-grpc_out=require_unimplemented_servers=false,:src/zk_leader_election/idl \
--go-grpc_opt=module=github.com/cxp863/zk_leader_election/src/zk_leader_election/idl \
src/zk_leader_election/idl/service/*.proto \
#src/zk_leader_election/idl/admin_service/*.proto

