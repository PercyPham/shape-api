# shape-api

## Intro

This repo is the answer to this Mini Test Assignment:

```
Create a RESTful API that will allow:
● User Registration
● User Login
● User can CRUD shapes of type:
  - Triangle
  - Rectangle
  - Square
  - Diamond
● User can request for shapes’ (note: these should be computed on the fly):
  - Area
  - Perimeter
```

APIs' structures developed in this repository are described in [`api/postman`](./api/postman/).

Note about project:

- This project is structured as Clean Architecture.
- Entities and usecases are tested with automation tests.
- Since the time constraints, API tests are tested manually, not automatically.
- CRUD operations for shapes contains:
  - Create single shape
  - Get single shape by id
  - Update single shape
  - Delete single shape

## Getting Started

### Requirements

- Docker & Docker Compose
- Golang 1.17

### Setup Env Variables

The environment variables are illustrated in `.env.example` file at the root of the project.

To overwrite those env variables, create a new file `.env` and overwrite the desired env variables.

### Development

To setup dev environment:

```
make dev_up
```
- Note: This will run docker-compose and spin up mysql db with an `phpMyAdmin` as DB GUI at `localhost:8081`.
- Please be patient, it will take few minutes to get everything prepared.

To stop db and db gui:
```
make dev_down
```
- Note: This will stop the db and db gui containers. However, it does not wipe out all data, and can be turned on again with "`make dev_up`" command.

To wipe out all database of dev mode, run:
```
make dev_db_clean
```
- Note: must run "`make dev_down`" first.

To run code in Dev mode:

```
make dev
```

### Test

To run test:

```
make test
```

### Production

<details>
<summary>Setup DB</summary>
Must have a DB created with the name `shape`, the sql script is located at `./init/db/schemas.sql`.
</details>

To build the executable file:

```
make build
```
- The built file will be located at `build/main`.

To run the built file:

```
./build/main
```

Note:
- The `.env` file must be created with an `APP_ENV` equals to `prod`.
- `APP_DOMAINS` must be defined, since this will run with SSL certificate provided by `Let's Encrypt`.
- `.env` must be located at the current command execution path (e.g. located at root of repo when running above command).
