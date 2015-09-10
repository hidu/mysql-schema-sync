#!/bin/bash
set -x
cd $(dirname $0)
if [ ! -d "log" ];then
   mkdir -p log
fi

exec 1>>log/sync.log.`date +"%Y%m%d"` 2>&1


for f in `ls *.json`
do
        ./mysql-schema-sync -conf $f -sync
done


cd log/
DAY_MAX=15
find ./ -type f -name "*.log*" -mtime +$DAY_MAX 