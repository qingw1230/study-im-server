#!/usr/bin/env bash

source ./style_info.cfg
source ./path_info.cfg
source ./function.sh

list1=$(cat $config_path | grep pushPort | awk -F '[:]' '{print $NF}')
list_to_string $list1
rpc_ports=($ports_array)


check=$(ps aux | grep -w ./${push_name} | grep -v grep | wc -l)
if [ $check -ge 1 ]; then
    oldPid=$(ps aux | grep -w ./${push_name} | grep -v grep | awk '{print $2}')
    kill -9 $oldPid
fi
sleep 1

cd ${push_binary_root}
for ((i = 0; i < ${#rpc_ports[@]}; i++)); do
    ./${push_name} -port ${rpc_ports[$i]} &
done

sleep 3
check=$(ps aux | grep -w ./${push_name} | grep -v grep | wc -l)
if [ $check -ge 1 ]; then
    newPid=$(ps aux | grep -w ./${push_name} | grep -v grep | awk '{print $2}')
    ports=$(netstat -netulp | grep -w ${newPid} | awk '{print $4}' | awk -F '[:]' '{print $NF}')
    allPorts=""

    for i in $ports; do
        allPorts=${allPorts}"$i "
    done
    echo -e ${SKY_BLUE_PREFIX}"SERVICE START SUCCESS "${COLOR_SUFFIX}
    echo -e ${SKY_BLUE_PREFIX}"SERVICE_NAME: "${COLOR_SUFFIX}${YELLOW_PREFIX}${push_name}${COLOR_SUFFIX}
    echo -e ${SKY_BLUE_PREFIX}"PID: "${COLOR_SUFFIX}${YELLOW_PREFIX}${newPid}${COLOR_SUFFIX}
    echo -e ${SKY_BLUE_PREFIX}"LISTENING_PORT: "${COLOR_SUFFIX}${YELLOW_PREFIX}${allPorts}${COLOR_SUFFIX}
else
    echo -e ${YELLOW_PREFIX}${push_name}${COLOR_SUFFIX}${RED_PREFIX}"SERVICE START ERROR, PLEASE CHECK openIM.log"${COLOR_SUFFIX}
fi
