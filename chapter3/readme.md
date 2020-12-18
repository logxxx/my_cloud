echo "hehe=haha" | md5sum

curl 127.0.0.1:8008/objects/hehe -XPUT -d'hehe=haha' -H'digest:SHA-256=41b6c429d4f83f62bea8127cef1b62a4'

curl 127.0.0.1:8008/objects/hehe -H'digest:SHA-256=41b6c429d4f83f62bea8127cef1b62a4'

curl http://49.232.219.233:9200/metadata/objects/hehe_1?type=create -XPUT -H 'content-Type:application/json' -d'{"name":"hehe", "version":1, "size":9, "hash":"41b6c429d4f83f62bea8127cef1b62a4"}'