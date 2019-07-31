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

## Task
- [ ] Create Dockerfile for building Golang application
- [ ] Create Ansible role for building container
- [ ] Create Ansible role for running migration
- [ ] Create Ansible role for deploying container
- [ ] Create Ansible playbooks for: deploy, migrate