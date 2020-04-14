#!/usr/bin/env bash

run_path=$(pwd)
shell_run_name=$(basename $0)
shell_run_path=$(
  cd $(dirname $0)
  pwd
)

build_version=v1.11.1
build_root_path=$(dirname ${shell_run_path})
build_root_name=example-temp-gin
build_depend_path=
build_need_proxy=0

go_proxy_url=https://goproxy.cn/
alpinelinux_proxy=mirrors.aliyun.com
docker_none_mark=none

# mysql
db_mysql_maintain_path=${shell_run_path}/db/mysql
db_mysql_config_maintain_path=${shell_run_path}/db/mysql/conf
db_mysql_target_path=${build_depend_path}/data/tmp/db
db_mysql_config_target_path=${build_depend_path}/data/mysql-${build_root_name}-conf

db_mysql_db_basic_file_name=admin
db_mysql_db_basic_file_base_name=${db_mysql_db_basic_file_name}.sql

# maintain_biz.sh
db_mysql_db_basic_maintain_biz_file_list=(
  demo.sql
)

db_mysql_root_pwd=3a5f549be630a477
db_mysql_db_basic_name=GoAdmin
db_mysql_user=golang
db_mysql_pwd=golang123456

# redis
cache_redis_maintain_path=${shell_run_path}/db/redis
cache_redis_maintain_config=${cache_redis_maintain_path}/redis.conf
cache_redis_target_path=${build_depend_path}/data/redis-${build_root_name}-etc/
cache_redis_target_config=${cache_redis_target_path}/redis.conf

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
  if [[ -z $1 ]]; then
    pW "Want find contain is empty"
    echo "-1"
  else
    c_status=$(docker inspect $1)
    if [[ ! $? -eq 0 ]]; then
      echo "1"
    else
      echo "0"
    fi
  fi
}

dockerStopContainWhenRunning() {
  if [[ -z $1 ]]; then
    pW "Want stop contain is empty"
  else
    c_status=$(docker inspect --format='{{ .State.Status}}' $1)
    if [[ "running" == ${c_status} ]]; then
      pD "-> docker stop contain [ $1 ]"
      docker stop $1
      checkFuncBack "docker stop $1"
    fi
  fi
}

dockerRemoveContainSafe() {
  if [[ -z $1 ]]; then
    pW "Want remove contain is empty"
  else
    has_contain=$(dockerIsHasContainByName $1)
    if [[ ${has_contain} -eq 0 ]]; then
      dockerStopContainWhenRunning $1
      c_status=$(docker inspect --format='{{ .State.Status}}' $1)
      if [[ "exited" == ${c_status} ]]; then
        pD "-> docker rm contain [ $1 ]"
        docker rm $1
        checkFuncBack "docker rm $1"
      fi
      if [[ "created" == ${c_status} ]]; then
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
checkBinary docker-compose

while getopts "hpb:n:" arg; do #after param has ":" need option
  case $arg in
  p) # -p open proxy of build
    build_need_proxy=1
    ;;
  b) # -b [v1.0.0] build version of contains
    build_version=${OPTARG}
    ;;
  n) # -n [example-temp-gin] name of build
    build_root_name=${OPTARG}
    build_depend_path=${build_root_path}/${build_root_name}-dev
    db_mysql_target_path=${build_depend_path}/data/tmp/db
    db_mysql_config_target_path=${build_depend_path}/data/mysql-${build_root_name}-conf
    cache_redis_target_path=${build_depend_path}/data/redis-${build_root_name}-etc/
    cache_redis_target_config=${cache_redis_target_path}/redis.conf
    ;;
  h)
    echo -e "this script to mark docker build file
    use as ${shell_run_name} -p
ars:
  -p open proxy of build
  -b [v1.0.0] build version of contains
  -n [example-temp-gin] name of build
"
    ;;
  ?) # other param?
    echo "unkonw argument, plase use -h to show help"
    exit 1
    ;;
  esac
done

if [[ -z "${build_depend_path}" ]]; then
  pE "build_depend_path is empty"
  exit 128
fi

if [[ ! -d ${build_depend_path} ]]; then
  mkdir -p ${build_depend_path}
fi

# set git ignore
echo -e "# ignore at ${build_root_name}-dev
data/
" >${build_depend_path}/.gitignore

# db_mysql_config
if [[ -z "${db_mysql_config_target_path}" ]]; then
  pE "db_mysql_config_target_path is empty"
  exit 128
fi

if [[ ! -d ${db_mysql_config_target_path} ]]; then
  mkdir -p ${db_mysql_config_target_path}
fi
cp ${db_mysql_config_maintain_path}/my_custom.cnf ${db_mysql_config_target_path}

# mysql file
if [[ ! -d ${db_mysql_target_path} ]]; then
  mkdir -p ${db_mysql_target_path}
fi

echo -e "#!/bin/bash

ENV_SQL_BASE_HOST=127.0.0.1
ENV_SQL_BASE_PORT=3306
ENV_SQL_BASE_USER=${db_mysql_user}
ENV_SQL_BASE_PWD=${db_mysql_pwd}

mysql \\
	--host=\${ENV_SQL_BASE_HOST} \\
	--port=\${ENV_SQL_BASE_PORT} \\
	--user=\${ENV_SQL_BASE_USER} \\
	--password=\${ENV_SQL_BASE_PWD} \\
	--database=${db_mysql_db_basic_name} \\
	< ${db_mysql_db_basic_file_base_name}
" >${db_mysql_target_path}/init-${db_mysql_db_basic_file_name}.sh
cp ${db_mysql_maintain_path}/${db_mysql_db_basic_file_base_name} ${db_mysql_target_path}

if [[ -z "${db_mysql_db_basic_maintain_biz_file_list}" ]]; then
  echo "empty of db_mysql_db_basic_maintain_biz_file_list"
else
  echo -e "#!/bin/bash
ENV_SQL_BASE_HOST=127.0.0.1
ENV_SQL_BASE_PORT=3306
ENV_SQL_BASE_USER=${db_mysql_user}
ENV_SQL_BASE_PWD=${db_mysql_pwd}
" >${db_mysql_target_path}/maintain-biz.sh

  for file_name in ${db_mysql_db_basic_maintain_biz_file_list[@]}; do
    if [[ -f ${db_mysql_maintain_path}/${file_name} ]]; then
      echo -e "mysql \\
	--host=\${ENV_SQL_BASE_HOST} \\
	--port=\${ENV_SQL_BASE_PORT} \\
	--user=\${ENV_SQL_BASE_USER} \\
	--password=\${ENV_SQL_BASE_PWD} \\
	--database=${db_mysql_db_basic_name} \\
	< ${file_name}
" >> ${db_mysql_target_path}/maintain-biz.sh
      cp ${db_mysql_maintain_path}/${file_name} ${db_mysql_target_path}
    fi
  done
fi

echo -e "#!/bin/bash
ENV_SQL_BASE_HOST=127.0.0.1
ENV_SQL_BASE_PORT=3306
ENV_SQL_BASE_USER=${db_mysql_user}
ENV_SQL_BASE_PWD=${db_mysql_pwd}

# --skip-disable-keys --add-drop-table --skip-extended-insert --skip-add-locks --skip-lock-tables

mysqldump \\
  ${db_mysql_db_basic_name} \\
	--host=\${ENV_SQL_BASE_HOST} \\
	--port=\${ENV_SQL_BASE_PORT} \\
	--user=\${ENV_SQL_BASE_USER} \\
	--password=\${ENV_SQL_BASE_PWD} \\
	--result-file=\"db-${db_mysql_db_basic_name}-dump.sql\" \\
	\"goadmin_menu\" \\
	\"goadmin_permissions\" \\
	\"goadmin_role_menu\" \\
	\"goadmin_role_permissions\" \\
	\"goadmin_role_users\" \\
	\"goadmin_roles\" \\
	\"goadmin_session\" \\
	\"goadmin_user_permissions\" \\
	\"goadmin_users\" \\
	\"goadmin_site\" \\
  --skip-disable-keys --add-drop-table --skip-extended-insert --add-locks --lock-tables

mysqldump \\
  ${db_mysql_db_basic_name} \\
	--host=\${ENV_SQL_BASE_HOST} \\
	--port=\${ENV_SQL_BASE_PORT} \\
	--user=\${ENV_SQL_BASE_USER} \\
	--password=\${ENV_SQL_BASE_PWD} \\
	\"goadmin_operation_log\" \\
	--no-data --skip-disable-keys --skip-extended-insert --skip-add-locks --skip-lock-tables --add-drop-table \\
	>> db-${db_mysql_db_basic_name}-dump.sql
" > ${db_mysql_target_path}/dump-${db_mysql_db_basic_name}.sh

# redis
if [[ ! -d ${cache_redis_target_path} ]]; then
  mkdir -p ${cache_redis_target_path}
fi
cp ${cache_redis_maintain_config} ${cache_redis_target_config}

echo -e "# copy right
# Licenses http://www.apache.org/licenses/LICENSE-2.0
# more info see https://docs.docker.com/compose/compose-file/ or https://docker.github.io/compose/compose-file/
version: '3.7'

networks:
  default:
#volumes:
#  web-data:
services:
  # https://hub.docker.com/r/bitnami/mysql
  ${build_root_name}-mysql-fix:
    container_name: '${build_root_name}-mysql-fix'
    image: bitnami/mysql:5.7.27
    user: root
    command: chown -R 1001:1001 /bitnami/
    volumes:
      - ./data/mysql-${build_root_name}:/bitnami/
  # https://hub.docker.com/r/bitnami/mysql
  ${build_root_name}-mysql:
    container_name: '${build_root_name}-mysql'
    image: bitnami/mysql:5.7.27
    depends_on:
      - ${build_root_name}-mysql-fix
    ports:
      - '39005:3306'
    volumes:
      - './data/mysql-${build_root_name}:/bitnami/mysql/data'
      - './data/mysql-${build_root_name}-conf/my_custom.cnf:/opt/bitnami/mysql/conf/bitnami/my_custom.cnf:ro'
      - './data/tmp/:/tmp'
    environment:
      MYSQL_ROOT_PASSWORD: '${db_mysql_root_pwd}'
      MYSQL_DATABASE: '${db_mysql_db_basic_name}'
      MYSQL_USER: '${db_mysql_user}'
      MYSQL_PASSWORD: '${db_mysql_pwd}'
    restart: always # on-failure:3 or unless-stopped always default no
  # https://hub.docker.com/r/bitnami/redis/
  ${build_root_name}-redis-fix: # use https://github.com/bitnami/bitnami-docker-mongodb/issues/103#issuecomment-424833086 fix ubunut error
    container_name: '${build_root_name}-redis-fix'
    image: 'bitnami/redis:5.0.0'
    user: root
    command: chown -R 1001:1001 /bitnami
    volumes:
      - ./data/redis-${build_root_name}:/bitnami
  ${build_root_name}-redis:
    container_name: '${build_root_name}-redis'
    # for fix debian folder error
    image: 'bitnami/redis:5.0.0'
    depends_on:
      - ${build_root_name}-redis-fix
    environment:
      # ALLOW_EMPTY_PASSWORD is recommended only for development.
      - ALLOW_EMPTY_PASSWORD=yes
      # - REDIS_REPLICATION_MODE=slave
      # - REDIS_PASSWORD=[pwd]
      # - REDIS_MASTER_HOST=[pwd]
      # - REDIS_MASTER_PORT_NUMBER=[pwd]
      # - REDIS_MASTER_PASSWORD=[pwd]
      - REDIS_EXTRA_FLAGS=--maxmemory 300mb
    ports:
      - '39006:6379'
    volumes:
      - './data/redis-${build_root_name}-etc/redis.conf:/opt/bitnami/redis/etc/redis.conf'
      - './data/redis-${build_root_name}/data:/bitnami/redis/data'
    restart: always # on-failure:3 or unless-stopped always default no
" >${build_depend_path}/docker-compose.yml

docker images | grep ${docker_none_mark} | awk '{print $3}' | xargs docker rmi

exit 0
