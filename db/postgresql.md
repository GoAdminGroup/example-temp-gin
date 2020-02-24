## use postgresql

- in main.go code must set

```go
import _ "github.com/lib/pq" // postgresql driver
```

## need cli 

install psql as exec

```bash
# ubuntu
$ sudo apt show postgresql-client
$ sudo apt-get install -y postgresql-client

$ sudo curl -s -L --fail https://raw.githubusercontent.com/bridgewwater/docker-exec-tools/master/pgcli/psql/9.6.16-alpine/run.sh -o /usr/local/bin/psql
$ sudo chmod +x /usr/local/bin/psql

# check
$ psql --help
```

install pgcli

```bash
# ubuntu
$ sudo apt-get install pgcli

$ sudo curl -s -L --fail https://raw.githubusercontent.com/bridgewwater/docker-exec-tools/master/pgcli/pygmy/run.sh -o /usr/local/bin/pgcli
$ sudo chmod +x /usr/local/bin/pgcli
```

### macOS-install

```bash
# psql management
brew install -v postgresql
```

### db-docker-compose

```yaml
# more info see https://docker.github.io/compose/compose-file/
version: '3.7'
networks:
  default:
    # Use a custom driver
    #driver: custom-driver-1
  # app-tier: # use as networks: - app-tier
    #driver: bridge
services:
  # https://hub.docker.com/r/bitnami/postgresql
  fix-DevGoAdminPostgresql-permissions:
    container_name: 'Dev-GoAdmin-Postgresql-fix'
    image: 'bitnami/postgresql:9.6.16'
    user: root
    command: chown -R 1001:1001 /bitnami
    volumes:
      - ./data/Dev/GoAdmin-postgresql:/bitnami
  DevGoAdminPostgresql:
    container_name: 'Dev-GoAdmin-Postgresql'
    image: 'bitnami/postgresql:9.6.16'
    depends_on:
      - fix-DevGoAdminPostgresql-permissions
    ports:
      - '27019:5432'
    volumes:
      - './data/Dev/GoAdmin-postgresql:/bitnami/postgresql'
    environment:
      - POSTGRESQL_USERNAME=golang
      - POSTGRESQL_PASSWORD=golang
      - POSTGRESQL_DATABASE=GoAdmin
    # always, on-failure:3 or unless-stopped default "no" https://docs.docker.com/compose/compose-file/#restart
    restart: on-failure:3
```

then `conf/config.yaml` must use as

```yaml
goAdmin:
  color_scheme: skin-black
  language: CN
  indexUrl: /
  urlPrefix: admin
  store:
    path: ./uploads
    prefix: uploads
  dataBases:
    default:
      driver: postgresql # mysql sqlite postgresql mssql
      host: 127.0.0.1
      port: 27019
      user: golang
      pwd: golang
      name: GoAdmin
      maxIdleCon: 50
      MaxOpenCon: 150
  dashBoard:
    title: GoAdmin-temple
```