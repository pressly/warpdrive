# Warpdrive

### Postgres Setup

before running warpdrive, make sure you have a right role and database. you can run the follwoing sql in Postgres terminal

```bash
CREATE USER warpdrive WITH PASSWORD 'warpdrive';
CREATE DATABASE warpdrivedb;
```

and make sure to set the correct username, password and database in warpdrive.conf.

# Warpfile

server:
  addr: 192.168.0.1:3000
cycles:
  production:
    build: react-native build $PLATFORM

# APIs

- [x] GET       /session
- [x] POST      /session/start
- [x] GET       /session/end

- [x] GET       /users?name=&email=
- [x] GET       /users/:userId
- [X] POST      /users
- [X] PUT       /users

- [x] GET       /apps?name=
- [x] GET       /apps/:appId
- [x] POST      /apps
- [X] PUT       /apps/:appId

- [x] GET       /apps/:appId/users?name=&email=
- [x] POST      /apps/:appId/users/:userId
- [x] DELETE    /apps/:appId/users/:userId

- [ ] GET       /apps/:appId/cycles?name=
- [ ] GET       /apps/:appId/cycles/:cycleId
- [ ] POST      /apps/:appId/cycles
- [ ] GET       /apps/:appId/cycles/:cycleId/key
- [ ] PUT       /apps/:appId/cycles/:cycleId
- [ ] DELETE    /apps/:appId/cycles/:cycleId

- [ ] GET       /apps/:appId/cycles/:cycleId/releases?platforn=&version=&note=
- [ ] GET       /apps/:appId/cycles/:cycleId/releases/:releaseId
- [ ] POST      /apps/:appId/cycles/:cycleId/releases
- [ ] PUT       /apps/:appId/cycles/:cycleId/releases/:releaseId
- [ ] DELETE    /apps/:appId/cycles/:cycleId/releases/:releaseId
- [ ] POST      /apps/:appId/cycles/:cycleId/releases/:releaseId/bundles
- [ ] GET       /apps/:appId/cycles/:cycleId/releases/:releaseId/bundles?name=
- [ ] POST      /apps/:appId/cycles/:cycleId/releases/:releaseId/lock
- [ ] DELETE    /apps/:appId/cycles/:cycleId/releases/:releaseId/lock

## the following apis call in clinet native, also they are public

- [ ] GET       /apps/:appId/cycles/:cycleId/releases/latest/:version
- [ ] POST      /apps/:appId/cycles/:cycleId/releases/:releaseId/download

## this following api calls for any auditing

- [ ] GET       /apps/:appId/logs
