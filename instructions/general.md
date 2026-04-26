## Overview

### Phase 1

The goal of this project is to learn Golang and timescale Databse.
I've installed everything inlcuding Go and postgres SQL which is installed both with docker and locally.

The databse can be activated thru this command:

```
docker run -d --name timescaledb -p 5432:5432  -v /Users/Eleanor/Projects/playground/timescale-db:/pgdata -e PGDATA=/pgdata -e POSTGRES_PASSWORD=password timescale/timescaledb-ha:pg18

```

I want to build a simple server and interact with the timescale databse. Help me come up with a implementation plan, and create a new implementation plan file under this folder.

The implementation plan must include:

1. formatted steps outlined what to do
2. key learning points and key takeaways

### Phase 2

PoC backend design in the knowledge foldeer.
