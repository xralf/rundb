#!/usr/bin/env bash

curl http://localhost:8080/suppliers/
curl --header "Content-Type: application/json" --request POST --data '{"name":"s1","address":"address1"}' http://localhost:8080/suppliers/
curl --header "Content-Type: application/json" --request POST --data '{"name":"s2","address":"address2"}' http://localhost:8080/suppliers/
curl --header "Content-Type: application/json" --request POST --data '{"name":"s3","address":"address3"}' http://localhost:8080/suppliers/
curl http://localhost:8080/suppliers/
curl --request DELETE http://localhost:8080/suppliers/s2
curl http://localhost:8080/suppliers/
curl --request DELETE http://localhost:8080/suppliers/s1
curl --request DELETE http://localhost:8080/suppliers/s3
curl http://localhost:8080/suppliers/

curl "http://localhost:8080/products"
curl "http://localhost:8080/products?category=paint"
curl "http://localhost:8080/products?name=pink"
curl "http://localhost:8080/products?name=pink&category=paint"
