# Playneta DevOps test task

## Prerequisites
 - [ ] Golang 1.12
 - [ ] PostgreSQL
 
## Building
```bash
mkdir /tmp/devops
cd /tmp/devops
git clone git@github.com:playneta/devops-test.git .
go build .
```

## Migration
```bash
psql -U postgres -d db -a -f migrations/init.sql
```

## Help
```bash
Usage of main:
  -db string
        database db (default "devops")
  -db-addr string
        database address (default "127.0.0.1:5432")
  -db-password string
        database password (default "password")
  -db-user string
        database username (default "postgres")
  -host string
        listening host (default "127.0.0.1")
  -port string
        listening port (default "12345")
```

## Checking
```bash
curl http://localhost:12345/users
> [{"id":1,"name":"Pavel Makarenko","email":"p.makarenko@playneta.gg"},{"id":2,"name":"Maksim Moskalik","email":"m.moskalik@playneta.gg"},{"id":3,"name":"Sergey Fugin","email":"s.fugin@playneta.gg"}]
```

## Task
- [ ] Create Dockerfile for building Golang application
- [ ] Create Ansible role for building container
- [ ] Create Ansible role for running migration
- [ ] Create Ansible role for deploying container
- [ ] Create Ansible playbooks for: deploy, migrate