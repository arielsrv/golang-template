#!/bin/bash
appName=myapp
echo $appName

#heroku container:push web -a $appName
#heroku container:release web -a $appName
#heroku logs --tail -a $appName

docker buildx build --platform linux/amd64 -t $appName .
docker tag $appName registry.heroku.com/$appName/web
docker push registry.heroku.com/$appName/web
heroku container:release web -a $appName
