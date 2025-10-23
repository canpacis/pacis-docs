# Deploying

## Building

### Building a Binary

You can build a binary using the `task build` command, or manually run the following:

```sh
bun run build
```

This command builds the web assets and places them in the `build` directory, ready to be embedded in the final binary.

```sh
go build -ldflags "-s -w" -tags=prod -o ./build/server .
```

This builds the main package with the `prod` tag, embedding the static assets.
The `-ldflags "-s -w"` options strip debug symbols to reduce the binary size.

> The final binary will be located at `build/server`.
> The static files inside `build/` have already been embedded at this stage and are now just build artefacts.

To cross-compile for another platform, use Go’s standard environment variables.
For example, to compile for ARM64 Linux:

```sh
GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -tags=prod -o ./build/server .
```

### Building an Image

You can build a container image using the included `Dockerfile`.
The build process is multi-staged to optimize layer caching:

1. Download web asset dependencies
2. Build static assets
3. Download Go module dependencies
4. Build the final application

The resulting image is based on Docker’s `scratch` base image, making it extremely lightweight.

## Going to Production

You can run the final artefact as you would any other binary or container image.

Run the binary directly:

```sh
./server
```

Or run the Docker image:

```sh
docker run -p 8080:8080 <image-name>
```
