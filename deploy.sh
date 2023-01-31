#!/bin/bash
appName=go-fiber-app
echo "appName: " $appName

docker buildx build --platform linux/amd64 -t $appName .
docker tag $appName registry.heroku.com/$appName/web
docker push registry.heroku.com/$appName/web
heroku container:release web -a $appName
