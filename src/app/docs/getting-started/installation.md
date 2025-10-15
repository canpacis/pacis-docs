# Installation

To install, and get started, your best bet is creating a new repository from Pacis' [Github Template](https://github.com/canpacis/pacis-template). Go ahead create yours and clone it to your local environment.

```shell
git clone <your-repository>
```
You can follow the simple readme guide in the template repository or follow this one.

## Install Dependencies

Install your go and bun dependencies. 

> Pacis uses bun and vite for bundling your frontend assets. You can replace it with another package manager but the template heavily relies on it.

```shell
bun install
```

```shell
go mod download
```

Pacis also uses [air](https://github.com/air-verse/air) and [taskfile](https://taskfile.dev/) to streamline development and build processes. If you don't have them already, follow their installation guides:

- [Air](https://github.com/air-verse/air?tab=readme-ov-file#installation)
- [Taskfile](https://taskfile.dev/docs/installation)

## Development

You are ready to run your development server. There is a task called `dev` already defined for you in the `Taskfile.yml` file. 

```shell
task dev
```

Running this command will spin up a server on localhost port 8080. It will also run a vite dev server on port 5173.

> The dev server actually runs on port 8081 but it is proxied to port 8080 with hot reloading using air. Check the `.air.toml` file for its configuration.