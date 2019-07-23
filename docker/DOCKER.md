`cd expenditure-trello` # you should be inside project folder

### How to build image
```bash
docker build -f docker/Dockerfile --build-arg PROJECT="expenditure-trello" --build-arg USER_ID=`id -u` -t expenditure-trello .
```

### How to run container to get expenditure-trello executable
```bash
docker run -e USER_ID=`id -u` -v $PWD:/builder/src/gitlab.picus.io/picussecurity/expenditure-trello --rm expenditure-trello
```

### How to run container with custom build.sh
```bash
docker run -e USER_ID=`id -u` -v $PWD:/builder/src/gitlab.picus.io/picussecurity/expenditure-trello -v "$PWD"/build.sh:/build.sh --entrypoint=/build.sh --rm expenditure-trello
```

The build result(binary and frontend) will be inside project folder.
