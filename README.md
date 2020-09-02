 #pg-log-extractor
 ## watch postgreSQL table
 to watch postgres database table in the project path execute this command
## Definitions
* Logical Decoding
```
Logical decoding is the process of extracting all persistent changes to a database's tables into a coherent, easy to understand format which can be interpreted without detailed knowledge of the database's internal state.

In PostgreSQL, logical decoding is implemented by decoding the contents of the write-ahead log, which describe changes on a storage level, into an application-specific form such as a stream of tuples or SQL statements.
```

* Replication Slots

``` 
In the context of logical replication, a slot represents a stream of changes that can be replayed to a client in the order they were made on the origin server. Each slot streams a sequence of changes from a single database.

A replication slot has an identifier that is unique across all databases in a PostgreSQL cluster. Slots persist independently of the connection using them and are crash-safe.

A logical slot will emit each change just once in normal operation. The current position of each slot is persisted only at checkpoint, so in the case of a crash the slot may return to an earlier LSN, which will then cause recent changes to be resent when the server restarts. Logical decoding clients are responsible for avoiding ill effects from handling the same message more than once. Clients may wish to record the last LSN they saw when decoding and skip over any repeated data or (when using the replication protocol) request that decoding start from that LSN rather than letting the server determine the start point. The Replication Progress Tracking feature is designed for this purpose, refer to replication origins.

Multiple independent slots may exist for a single database. Each slot has its own state, allowing different consumers to receive changes from different points in the database change stream. For most applications, a separate slot will be required for each consumer.

A logical replication slot knows nothing about the state of the receiver(s). It's even possible to have multiple different receivers using the same slot at different times; they'll just get the changes following on from when the last receiver stopped consuming them. Only one receiver may consume changes from a slot at any given time.
```

* Output Plugins
```
Output plugins transform the data from the write-ahead log's internal representation into the format the consumer of a replication slot desires.

In our case wi are using  wal2json  https://github.com/eulerto/wal2json.
```

* Exported Snapshots
```
When a new replication slot is created using the streaming replication interface (see CREATE_REPLICATION_SLOT), a snapshot is exported (see Section 9.26.5), which will show exactly the state of the database after which all changes will be included in the change stream. This can be used to create a new replica by using SET TRANSACTION SNAPSHOT to read the state of the database at the moment the slot was created. This transaction can then be used to dump the database's state at that point in time, which afterwards can be updated using the slot's contents without losing any changes.

Creation of a snapshot is not always possible. In particular, it will fail when connected to a hot standby. Applications that do not require snapshot export may suppress it with the NOEXPORT_SNAPSHOT option.

```

## PostgreSQL Server Configuration
Once the wal2json plug-in has been installed, the database server should be configured.

Setting up libraries, WAL and replication parameters
Add the following lines at the end of the postgresql.conf PostgreSQL configuration file in order to include the plug-in at the shared libraries and to adjust some WAL and streaming replication settings. The configuration is extracted from postgresql.conf.sample. You may need to modify it, if for example you have additionally installed shared_preload_libraries.

postgresql.conf , configuration file parameters settings
############ REPLICATION ##############
# MODULES
shared_preload_libraries = 'wal2json'   

# REPLICATION
wal_level = logical                     
max_wal_senders = 10                     
max_replication_slots = 10    


1. tells the server that it should load at startup the wal2json (use decoderbufs for protobuf) logical decoding plug-in(s) (the names of the plug-ins are set in protobuf and wal2json Makefiles)
2. tells the server that it should use logical decoding with the write-ahead log
3. tells the server that it should use a maximum of 4 separate processes for processing WAL changes
4. tells the server that it should allow a maximum of 4 replication slots to be created for streaming WAL changes

> sudo systemctl restart postgresql 



to try logical replication  and  stream the changes of specifique  database follows those steps :

 ### step 1
 ```
 go run cmd/main.go rest --db="profile" --user="postgres" --password="postgres" --topic="tagTopic" --project-id="target-datalake-ng" --service-account="./service-account-file.json" 
 ```
 > Change the bd flag by the name of your database. That you need to stream change it data from it.

 ### step 2
    ```
    GET  `/v1/api/init`

    * response:


    {
        "slotName": "delta_ablaze5"
    }
    ```
 ### step 3
 ```
   POST /v1/api/snapshot/data
   body:
   
        {
            "slotName": "delta_ablaze5",
            "table":"user_profile",
            "offset":0,
            "limit":200,
            "order_by":{
                "column":"full_name",
                "order":"ASC"
            }
        }
  ```  
    

### step 4
 to test postgres databse table extractor go to `http://localhost:8080/pg-stream`
> entre the `slotName` and start watch for the new data on that was inserted in that table

* that use :

GET    `/v1/api/lr/stream?slotName=slot_name`
> WebSocket Endpoint that Stream the all changed data to all subscribed clients.   


## USER_PROFILE APIs

GET    `/v1/api/user_profile/bigquery/init`

> Create profile DB schema in BigQuery.

    ```
    * response:
    {
        "data": null,
        "success": true
    }
    ```

GET    `/v1/api/user_profile/bigquery/Dump-user-profile`

> Dump all user_profile rows into BigQuery.

    ```
    * response:
    {
        "data": true,
        "success": true
    }
    ```

GET    `/v1/api/user_profile/bigquery/Dump-career-aspiration`

> Dump all user_profile rows into BigQuery.

    ```
    * response:
    {
        "data": true,
        "success": true
    }
    ```

GET    `/v1/api/user_profile/bigquery/Dump-experience`


> Dump all user_profile rows into BigQuery.

    ```
    * response:
    {
        "data": true,
        "success": true
    }
    ```

GET    `/v1/api/user_profile/bigquery/Dump-user-competency`

> Dump all user_competency rows into BigQuery.

    ```
    * response:
    {
        "data": true,
        "success": true
    }
    ```

GET    `/v1/api/user_profile/bigquery/Dump-education-specialization` 


> Dump all user_education_specialization  rows into BigQuery.

    ```
    * response:
    {
        "data": true,
        "success": true
    }
    ```

GET    `/v1/api/user_profile/bigquery/Dump-user-language`


> Dump all user_language  rows into BigQuery.

    ```
    * response:
    {
        "data": true,
        "success": true
    }
    ```


GET    `/v1/api/user_profile/bigquery/Dump-user-education`


> Dump all user_education  rows into BigQuery.

    ```
    * response:
    {
        "data": true,
        "success": true
    }
    ```


GET    `/v1/api/user_profile/bigquery/Dump-degree-level`


> Dump all degree_level  rows into BigQuery.

    ```
    * response:
    {
        "data": true,
        "success": true
    }
    ```


GET    `/v1/api/user_profile/bigquery/Stream`

> insert all stream it data into bigquery. 
 



## hcm database extractor cmd
>  go run cmd/main.go hcm-extractor --db="hcm" --user="postgres" --password="postgres" --project-id="target-datalake-ng" --service-account="./service-account-file.json"

### HCM APIs

GET `/v1/api/hcm/bigquery/init`

> init hcm schema in bigquery.

GET `/v1/api/hcm/bigquery/Dump-job`

> Dump all job  rows into BigQuery.

    ```
    * response:
    {
        "data": true,
        "success": true
    }
    ```

GET `/v1/api/hcm/bigquery/Dump-job-competency`

> Dump all job_competency  rows into BigQuery.

    ```
    * response:
    {
        "data": true,
        "success": true
    }
    ```

GET `/v1/api/hcm/bigquery/Dump-job-education-specialization`

> Dump all job_education_specialieation  rows into BigQuery.

    ```
    * response:
    {
        "data": true,
        "success": true
    }
    ```


GET `/v1/api/hcm/bigquery/Dump-job-language`

> Dump all job_language  rows into BigQuery.

    ```
    * response:
    {
        "data": true,
        "success": true
    }
    ```

GET `/v1/api/hcm/bigquery/Dump-job-nationality`

> Dump all job_nationality  rows into BigQuery.

    ```
    * response:
    {
        "data": true,
        "success": true
    }
    ```

GET `/v1/api/hcm/bigquery/Dump-degree-level`

> Dump all degree_level  rows into BigQuery.

    ```
    * response:
    {
        "data": true,
        "success": true
    }
    ```

GET `/v1/api/hcm/bigquery/Stream`

> insert all stream it data into bigquery. 



## watch mongodb collection:

 * cd mongodb-replicaset
 * docker-compose up

docker ps 

### to verifier if mongodb is on replicaset mode 
> docker exec mongodb-replicaset_mongo-rs0-1_1_3d80121ac8ef bash -c 'mongo --eval "rs.status();"'

### to activate mongo shell.
> docker exec -it  mongodb-replicaset_mongo-rs0-3_1_3151c936b77f bash -c 'mongo' 

```
rs0:PRIMARY> use db1
switched to db db1
rs0:PRIMARY> db.c1.insertOne( { x: 1 } );
{
	"acknowledged" : true,
	"insertedId" : ObjectId("5d6e31747b1cd2d703b34983")
}
rs0:PRIMARY> db.c1.insertOne( { x: 2 } );
{
	"acknowledged" : true,
	"insertedId" : ObjectId("5d6e31947b1cd2d703b34984")
}

```
 to test mongobd database collection change stream go to `http://localhost:8080/mongo-stream`

 > Entre `database-name` and `collection-name` and start watch stream on that collection. 
