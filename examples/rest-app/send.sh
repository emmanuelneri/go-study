#!/usr/bin/env bash

URL=http://localhost:8080/orders
TOTAL=1

START=$(date +%s)

echo "send starting... ${TOTAL} registers"
for i in $(seq 1 $TOTAL);
do
  customer="Customer ${i}"

  json="{\"customer\": \"${customer}\",\"total\": 100,\"items\": [{\"description\": \"Product 1\", \"value\": 30},{\"description\": \"Product 2\",\"value\": 10},{\"description\": \"Product 3\", \"value\": 60}]}"

  curl -d "${json}" -H "Content-Type: application/json" -X POST ${URL} &
done