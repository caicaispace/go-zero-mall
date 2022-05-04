#!/bin/bash

# log
log() {
  local msg=$1
  local type=$2
  local time=$(date '+%F %T')
  if [ ! -n "$2" ]; then
    type="info"
  fi
  if [ "$type" == "info" ]; then
    echo "time: ${time} msg: ${msg}"
  elif [ "$type" == "waring" ]; then
    echo "⚠️  time: ${time} msg: ${msg}"
  elif [ "$type" == "error" ]; then
    echo "⚡ time: ${time} msg: ${msg}"
  fi
}

# command exists
commandExists() {
	local command=$1
  local ret='0'
  command -v $command >/dev/null 2>&1 || { local ret='1'; }
  # fail on non-zero return value
  if [ "$ret" -ne 0 ]; then
    log "command not found: $command" error
    exit
  fi
}

if [ "$1" == "log" ]; then
  log "------------ git clear commit ------------"
  log helloworld
  log helloworld waring
  log helloworld error
fi