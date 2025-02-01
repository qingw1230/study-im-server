#!/usr/bin/env bash

need_to_start_server_shell=(
  start_rpc_service.sh
  msg_gateway_start.sh
  push_start.sh
  msg_transfer_start.sh
)

for i in ${need_to_start_server_shell[*]}; do
  chmod +x $i
  ./$i
done
