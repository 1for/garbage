#!/bin/sh

cfile="/home/liuyang1-iri/etc/server.yaml"
visitor="visitor"

if [ -z $1 ];then
$visitor -c $cfile
exit 1
fi

if [ -z $2 ];then
$visitor -c $cfile -p $1
exit 1
fi

ip=$($visitor -c $cfile -p $1 -i $2)
if [ -z $ip ];then
echo "ip empty"
fi
echo "ssh "$ip
ssh $ip
