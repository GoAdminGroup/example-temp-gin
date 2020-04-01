#!/usr/bin/env bash

run_path=$(pwd)
shell_run_name=$(basename $0)
shell_run_path=$(
  cd $(dirname $0)
  pwd
)

build_need_proxy=0
build_version=v1.0.0
build_root_name=example-temp-gin
build_root_path=$(dirname $(dirname ${shell_run_path}))
build_clean_root_path=data

build_clean_path_list=(
  biz/
  mysql-${build_root_name}-conf/
  redis-${build_root_name}/etc
  tmp/db/
)

docker_none_mark=none

pV() {
  echo -e "\033[;36m$1\033[0m"
}
pI() {
  echo -e "\033[;32m$1\033[0m"
}
pD() {
  echo -e "\033[;34m$1\033[0m"
}
pW() {
  echo -e "\033[;33m$1\033[0m"
}
pE() {
  echo -e "\033[;31m$1\033[0m"
}

checkFuncBack() {
  if [[ $? -ne 0 ]]; then
    echo -e "\033[;31mRun [ $1 ] error exit code 1\033[0m"
    exit 1
  fi
}

checkBinary() {
  binary_checker=$(which $1)
  checkFuncBack "which $1"
  if [[ ! -n "${binary_checker}" ]]; then
    echo -e "\033[;31mCheck binary [ $1 ] error exit\033[0m"
    exit 1
    #  else
    #    echo -e "\033[;32mCli [ $1 ] event check success\033[0m\n-> \033[;34m$1 at Path: ${evn_checker}\033[0m"
  fi
}

check_root() {
  if [[ ${EUID} != 0 ]]; then
    echo "no not root user"
  fi
}

dockerIsHasContainByName() {
  if [[ ! -n $1 ]]; then
    pW "Want find contain is empty"
    echo "-1"
  else
    c_status=$(docker inspect $1)
    if [ ! $? -eq 0 ]; then
      echo "1"
    else
      echo "0"
    fi
  fi
}

dockerStopContainWhenRunning() {
  if [[ ! -n $1 ]]; then
    pW "Want stop contain is empty"
  else
    c_status=$(docker inspect --format='{{ .State.Status}}' $1)
    if [ "running" == ${c_status} ]; then
      pD "-> docker stop contain [ $1 ]"
      docker stop $1
      checkFuncBack "docker stop $1"
    fi
  fi
}

dockerRemoveContainSafe() {
  if [[ ! -n $1 ]]; then
    pW "Want remove contain is empty"
  else
    has_contain=$(dockerIsHasContainByName $1)
    if [[ ${has_contain} -eq 0 ]]; then
      dockerStopContainWhenRunning $1
      c_status=$(docker inspect --format='{{ .State.Status}}' $1)
      if [ "exited" == ${c_status} ]; then
        pD "-> docker rm contain [ $1 ]"
        docker rm $1
        checkFuncBack "docker rm $1"
      fi
      if [ "created" == ${c_status} ]; then
        pD "-> docker rm contain [ $1 ]"
        docker rm $1
        checkFuncBack "docker rm $1"
      fi
    else
      pE "dockerRemoveContainSafe Not found contain [ $1 ]"
    fi
  fi
}

# checkenv
checkBinary docker

while getopts "hpc:b:n:" arg; do #after param has ":" need option
  case $arg in
  p) # -p open proxy of build
    build_need_proxy=1
    ;;
  c) # -c [clean path at build root] name of clean path root
    build_clean_root_path=${OPTARG}
    ;;
  b) # -b [v1.0.0] build version of contains
    build_version=${OPTARG}
    ;;
  n) # -n [example-temp-gin] name of build
    build_root_name=${OPTARG}
    build_clean_path_list=(
      biz/
      mysql-${build_root_name}-conf/
      redis-${build_root_name}/etc
      tmp/db/
    )
    ;;
  h)
    echo -e "this script to mark docker build file
    use as ${shell_run_name} -p
ars:
  -p open proxy of build
  -b [v1.0.0] build version of contains
  -n [example-temp-gin] name of build
  -c [clean path at build root] name of clean path root
"
    ;;
  ?) # other param?
    echo "unkonw argument, plase use -h to show help"
    exit 1
    ;;
  esac
done

if [[ -z "${build_clean_path_list}" ]]; then
  pI "=> empty of build_clean_path_list"
else
  for clean_path in ${build_clean_path_list[@]}; do
    if [[ -d ${build_root_path}/${build_clean_root_path}/${clean_path} ]]; then
#      pD "-> try to clean ${build_root_path}/${build_clean_root_path}/${clean_path}"
      rm -rf ${build_root_path}/${build_clean_root_path}/${clean_path}
      checkFuncBack "rm -rf ${build_root_path}/${build_clean_root_path}/${clean_path}"
    fi
  done
fi

# for remove docker images which no tag mark by <none>
docker images | grep ${docker_none_mark} | awk '{print $3}' | xargs docker rmi

exit 0
