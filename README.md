# rundb

This little REST API server project assumes a local deployment and uses PostgreSQL.

## Create a local PostgreSQL database `rundb`

~~~
createdb -U postgres rundb
~~~

This project assumes that the `hba.conf` file has this `md5` autehtication mode:
~~~
# Database administrative login by Unix domain socket
local   all             postgres                                md5
~~~
... instead the default `peer` mode:
~~~
local   all             postgres                                peer
~~~

### Populate the data base

~~~
make setup-db
~~~

### Build and run the API server in Terminal 1

~~~
make run
~~~

### Submit sample requests from Terminal 2

~~~
make requests
~~~

## Example output

### Terminal 1

~~~
~/git/xralf/rundb (main ✗) make run
go build api.go
./api
&{<nil> false Server listening at port :8080 ... <nil> map[Server listening at port :8080 ...:0xc000119360] {false false false {<nil> <nil> []} []  <nil>}}
Getting suppliers ...
Creating supplier ...
        {Name:s1 Address:address1}
Creating supplier ...
        {Name:s2 Address:address2}
Creating supplier ...
        {Name:s3 Address:address3}
Getting suppliers ...
        {Name:s1 Address:address1}
        {Name:s2 Address:address2}
        {Name:s3 Address:address3}
Deleting supplier ...
        name: s2
Getting suppliers ...
        {Name:s1 Address:address1}
        {Name:s3 Address:address3}
Deleting supplier ...
        name: s1
Deleting supplier ...
        name: s3
Getting suppliers ...
Querying products ...
        select name, category, sku from products
Querying products ...
        select name, category, sku from products where category like '%paint%'
Querying products ...
        select name, category, sku from products where name like '%pink%'
Querying products ...
        select name, category, sku from products where name like '%pink%' and category like '%paint%'
~~~

### Terminal 2

~~~
~/git/xralf/rundb (main ✗) make requests
./requests.sh
{"type":"success","message":"all suppliers","data":null}
{"type":"success","message":"the supplier was added successfully","data":null}
{"type":"success","message":"the supplier was added successfully","data":null}
{"type":"success","message":"the supplier was added successfully","data":null}
{"type":"success","message":"all suppliers","data":[{"name":"s1","address":"address1"},{"name":"s2","address":"address2"},{"name":"s3","address":"address3"}]}
{"type":"success","message":"the supplier was deleted successfully","data":null}
{"type":"success","message":"all suppliers","data":[{"name":"s1","address":"address1"},{"name":"s3","address":"address3"}]}
{"type":"success","message":"the supplier was deleted successfully","data":null}
{"type":"success","message":"the supplier was deleted successfully","data":null}
{"type":"success","message":"all suppliers","data":null}
{"type":"success","message":"all qualifying products","data":[{"name":"pink super-gloss","category":"paint","sku":"sku1"},{"name":"glossy yellow","category":"paint","sku":"sku2"},{"name":"outdoors yellow","category":"paint","sku":"sku3"},{"name":"feeling so indoorsy","category":"paint","sku":"sku4"},{"name":"blue concrete","category":"tiles","sku":"sku5"},{"name":"pink porcelain","category":"tiles","sku":"sku6"}]}
{"type":"success","message":"all qualifying products","data":[{"name":"pink super-gloss","category":"paint","sku":"sku1"},{"name":"glossy yellow","category":"paint","sku":"sku2"},{"name":"outdoors yellow","category":"paint","sku":"sku3"},{"name":"feeling so indoorsy","category":"paint","sku":"sku4"}]}
{"type":"success","message":"all qualifying products","data":[{"name":"pink super-gloss","category":"paint","sku":"sku1"},{"name":"pink porcelain","category":"tiles","sku":"sku6"}]}
{"type":"success","message":"all qualifying products","data":[{"name":"pink super-gloss","category":"paint","sku":"sku1"}]}
~~~