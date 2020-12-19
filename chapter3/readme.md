echo "hehe=haha" | md5sum

curl 127.0.0.1:8008/objects/hehe -XPUT -d'hehe=haha' -H'digest:SHA-256=41b6c429d4f83f62bea8127cef1b62a4'

curl 127.0.0.1:8008/objects/hehe -H'digest:SHA-256=41b6c429d4f83f62bea8127cef1b62a4'

curl http://49.232.219.233:9200/metadata/objects/hehe_1?type=create -XPUT -H 'content-Type:application/json' -d'{"name":"hehe", "version":1, "size":9, "hash":"41b6c429d4f83f62bea8127cef1b62a4"}'

curl -XPUT http://49.232.219.233:9200/version -H 'content-Type:application/json'-d '{
  "mappings": {
    "type": {
      "properties": {
        "name": {
          "type": "text",
          "fielddata": true }
      }
      }
    }
  }
}'

curl -XPUT 'http://localhost:9200/metadata/objects/_mapping' -d '
 {       
   "properties": {
         "version": {  
             "type": "text",
             "fielddata": true
         }       
     }         
 }'
 
 curl http://49.232.219.233:9200/metadata/objects/_mapping
 返回:
 {
   "metadata": {
     "mappings": {
       "objects": {
         "properties": {
           "hash": {
             "type": "text",
             "fields": {
               "keyword": {
                 "type": "keyword",
                 "ignore_above": 256
               }
             }
           },
           "name": {
             "type": "text",
             "fields": {
               "keyword": {
                 "type": "keyword",
                 "ignore_above": 256
               }
             }
           },
           "size": {
             "type": "long"
           },
           "version": {
             "type": "long"
           }
         }
       }
     }
   }
 }
 
curl -XPUT http://49.232.219.233:9200/metadata/objects/_mapping -H'content-Type:application/json' -d'
 {
   "metadata": {
     "mappings": {
       "objects": {
         "properties": {
           "hash": {
             "type": "keyword",
             "fields": {
               "keyword": {
                 "type": "keyword",
                 "ignore_above": 256
               }
             }
           },
           "name": {
             "type": "keyword",
             "fields": {
               "keyword": {
                 "type": "keyword",
                 "ignore_above": 256
               }
             }
           },
           "size": {
             "type": "long"
           },
           "version": {
             "type": "long"
           }
         }
       }
     }
   }
 }
'