#!/bin/bash
appName=go-fiber-app
echo $appName
heroku container:push web -a $appName
heroku container:release web -a $appName
heroku logs --tail -a $appName
