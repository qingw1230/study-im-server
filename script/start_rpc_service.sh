#!/usr/bin/env bash

source ./style_info.cfg
source ./path_info.cfg
source ./function.sh

service_filename=(
    ${msg_name}
    api
    account
    friend
    group
)

service_port_name=(
    offlineMessagePort
    apiPort
    accountPort
    friendPort
    groupPort
)

for ((i = 0; i < ${#service_filename[*]}; i++)); do
    # 检查服务是否存在
    service_name="ps -aux | grep -w ${service_filename[$i]} | grep -v grep"
    count="${service_name} | wc -l"

    if [ $(eval ${count}) -gt 0 ]; then
        # 该命令用于提取进程 PID
        pid="${service_name} | awk '{print \$2}'"
        echo -e "${SKY_BLUE_PREFIX}${service_filename[$i]} service has been started, pid:$(eval $pid)$COLOR_SUFFIX"
        echo -e "${SKY_BLUE_PREFIX}Killing the service ${service_filename[$i]} pid:$(eval $pid)${COLOR_SUFFIX}"
        kill -9 $(eval $pid)
        sleep 0.5
    fi

    cd ../bin && echo -e "${SKY_BLUE_PREFIX}${service_filename[$i]} service is starting${COLOR_SUFFIX}"
    # 从配置文件中读取 rpc port
    portList=$(cat $config_path | grep ${service_port_name[$i]} | awk -F '[:]' '{print $NF}')
    list_to_string ${portList}

    for j in ${ports_array}; do
        echo -e "${SKY_BLUE_PREFIX}${service_filename[$i]} Service is starting, port number:$j $COLOR_SUFFIX"
        ./${service_filename[$i]} -port $j &
        sleep 5
        pid="netstat -ntlp | grep $j | awk '{printf \$7}'| cut -d/ -f1"
        echo -e "${RED_PREFIX}${service_filename[$i]} Service is started, port number:$j pid:$(eval $pid)$COLOR_SUFFIX"
    done
done
