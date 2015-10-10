#!/bin/bash
cd $(dirname $0)

gox -os linux &
gox -os windows &
gox -os darwin &
wait
echo "gox done"

for f in `ls mysql-schema-sync_*`
do
    tar -czvf $f.tar.gz $f
done

echo "all done"