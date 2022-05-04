#!/bin/bash
source utils.sh
echo -------- "$1" "$2" "$3" "$4" --------

### git
# ./cmd.sh git push
# ./cmd.sh git clear
if [ "$1" == "git" ]; then
  commandExists git
  if [ "$2" == "push" ]; then
    log "------------ git push ${branch}  ------------"
    branch=$(git symbolic-ref --short -q HEAD)
    time=$(date '+%F %T')
    msg='auto push at '${time}
    git add .
    git commit -m "${msg}"
    git push origin "${branch}"
  elif [ "$2" == "clear" ]; then
    log "------------ git clear commit ------------"
    remoteUrl=$(git config --get remote.origin.url)
    log "$remoteUrl"
    rm -rf .git
    git init
    git add .
    git commit -am "init"
    git remote add origin "${remoteUrl}"
    git push origin master --force
  fi
fi

### docker
# ./cmd.sh docker build order
# ./cmd.sh docker run order 9800
# ./cmd.sh docker push order zero
if [ "$1" == "docker" ]; then
  commandExists docker
  runType=$2
  project=$3
  if [ "$runType" == "build" ]; then
    log "------------ docker build $project ------------"
    if [[ "$(docker images -q $project 2> /dev/null)" != "" ]]; then
      docker image rm $project || (echo "Image $project didn't exist so not removed."; exit 0)
    fi
    cd ./code
    apiPath="./service/$project/api"
    rpcPath="./service/$project/rpc"
    if [ -d "$apiPath" ]; then
      docker build -t $project:latest -f $apiPath/Dockerfile .
    fi
    if [ -d "$rpcPath" ]; then
      docker build -t $project:latest -f $rpcPath/Dockerfile .
    fi
  elif [ "$runType" == "run" ]; then
    log "------------ docker run $project ------------"
    imageName = $2
    imagePort = $3
    docker run -p $imagePort:$imagePort $imageName
  elif [ "$runType" == "push" ]; then
    namespace=$3
    project=$4
    log "------------ docker push $namespace/$project ------------"
    imageId=$(docker images -q $project:latest)
    docker tag $imageId registry.cn-shanghai.aliyuncs.com/$namespace/$project:latest
    docker push registry.cn-shanghai.aliyuncs.com/$namespace/$project:latest
  fi
fi

### go-zero code generate
# ./cmd.sh gen api order
# ./cmd.sh gen rpc order
# ./cmd.sh gen model order
# ./cmd.sh gen dockerfile order
if [ "$1" == "gen" ]; then
  commandExists goctl
  if [ ! -n "$2" ]; then
    log the second argument must be required
    exit
  fi
  if [ ! -n "$3" ]; then
    log the third argument must be required
    exit
  fi
  genType=$2
  genPath=$3
  if [ "$genType" == "api" ]; then
    cd ./code/service/"$genPath"
    goctl api go -api ./api/"$genPath".api -dir ./api -style goZero
  elif [ "$genType" == "rpc" ]; then
    cd ./code/service/"$genPath"
    goctl rpc proto -src ./rpc/"$genPath".proto -dir ./rpc -style goZero
  elif [ "$genType" == "model" ]; then
    serviceName="$4"
    cd ./code/service/"$genPath"
    if [ "$4" ]; then
      serviceName="$4"
    fi
    goctl model mysql ddl -src ./model/ddl/"$serviceName".sql -dir ./model -c -style goZero
  elif [ "$genType" == "dockerfile" ]; then
    apiPath="./code/service/$genPath/api"
    rpcPath="./code/service/$genPath/rpc"
    if [ -d "$apiPath" ]; then
      cd $apiPath
      goctl docker -go $genPath.go
      cd -
    fi
    if [ -d "$rpcPath" ]; then
      cd ./code/service/"$genPath"/rpc
      goctl docker -go $genPath.go
    fi
  fi
fi