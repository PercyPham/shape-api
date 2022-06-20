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

## Getting Started

### Setup Env Variables

The default environment variables are defined in `.default.env` file at the root of the project.

To overwrite those env variables, create a new file `.env` and overwrite the desired env variables.

### Development

To start in Dev mode:

```
make dev
```

### Production

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
- `.default.env` and `.env` must be located at the current command execution path (e.g. located at root of repo when running above command).
