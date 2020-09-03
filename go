#!/usr/bin/env bash

set -e
script_dir=$(cd "$(dirname "$0")" ; pwd -P)
app_name=example-app

function echob {
    echo -e "\033[1m$1\033[0m"
}

goal_install() {
    pushd "${script_dir}" > /dev/null
       npm install
    popd > /dev/null
}


goal_run() {
    pushd "${script_dir}" > /dev/null
       node app.js
    popd > /dev/null
}

goal_build-container() {
  pushd "${script_dir}" > /dev/null
      docker build -t "${app_name}" .
  popd > /dev/null
}

goal_run-container() {
  pushd "${script_dir}" > /dev/null
      docker run -p 8080:8080 "${app_name}:latest"
  popd > /dev/null
}

if type -t "goal_$1" &>/dev/null; then
  goal_$1 ${@:2}
else
  echo "usage: $0 <goal>

goal:
    install                      - Installs dependencies
    run                          - Runs application
    build-container              - Builds image
    run-container                - Runs image
"
  exit 1
fi
