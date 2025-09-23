#!/bin/bash

Script to collect the status of lshw output from home servers
Dependencies:
*LSHW: http://ezix.org/project/wiki/HardwareLiSter
*JQ: http://stedolan.github.io

On each machine you can run something like this from cron
0 0 * * * /usr/sbin/lshw -json -quiet > /var/log/lshw-dump.json
Author: Jose Vicente Nunez

set -o errtrace # enable the err trap, code will get called when an error is detected
trap "echo ERROR: there was an error in ${FUNCNAME-main context}, details to follow" ERR

#check if dependencies present
declare -a dependencies=(/usr/bin/timeout /usr/bin/ssh /usr/bin/jq)
for dependency in "${dependencies[@]}"; do
    if [ ! -x $dependency ]; then
        echo "ERROR: missing dependency"
        exit 100
    fi
done


# declare command in Bash is used to modify the attributes of variables or to define variables with specific properties
declare -a servers=(
    dmaf5
)

declare REMOTE_FILE="var/log/lshw-dump.json"
declare MAX_RETRIES=3

# Uses timeout to make sure the scp finishes no later than 45s
function remote_copy {
    local server=$1
    local retries=$2
    local now=1
    status=0

    while [ $now -le $retries ]; do 
        echo "INFO: Trying to copy file from: $server, attempt $now"
        /usr/bin/timeout --kill-after 25.0s 20.0s \
            /usr/bin/scp \
                -o BatchMode=yes \
                -o logLevel=Error \
                -o ConnectTimeout=5 \
                -o ConnectionAttempts=3 \
                ${server}:/var/log/lshw-dump.json ${DATADIR}/lshw-$server-dump.json
        status=$?
        if [ $status -ne 0 ]; then
            sleep_time=$(((RANFOM % 60 ) + 1 ))
            echo "WARNING: Copy failed for $server:$REMOTE_FULE, retrying in $sleep_time seconds"
            /usr/bin/sleep ${sleep_time}s
        else 
            break
        fi
    done 
    return $status
}

DATADIR="$HOME/Documents/lshw-dump"
if [ ! -d "$DATADIR" ]; then
    /usr/bin/mkdir -p -v "$DATADIR" || "FATAL: Failed to create $DATADIR" && exit 100
fi

declare -A server_pid
for server in ${servers[*]}; do
    remote_copy $server $MAX_RETRIES &
    server_pid[$server]=$!
done


# Iterate through all the servers and:
wait for the return code of each and check the exist code from each scp

for server in ${!server_pid[*]}; do
    wait "${server_pid[$server]}"
    test $? -ne 0 && echo "ERROR: Copy from $server had problems, will not continue" && exit 100
done

for lshw in $(/usr/bin/find $DATADIR -type f -name 'lshw-*-dump.json'); do
    /usr/bin/jq '.["product","vendor","configuration"]' $lshw
done

