#!/bin/sh

content=$(echo -ne "hello world" | base64)
token="<TOKEN>"
username="<USERNAME>"
repository="<REPOSITORY>"
timestamp=$(date +%s)

curl -vi -H "Accept: application/json" -H "Content-Type: application/json" -X PUT -d '{"content":"'$content'", "message":"random message"}' https://api.github.com/repos/$username/$repository/contents/output/LOG-$timestamp.log?access_token=$token