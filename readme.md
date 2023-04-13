# Go io.writer Exercise
## Introduction
Based on https://golang.cafe/blog/golang-writer-example.html

Demonstrates the mega basic concept of using io.Writer interface.

* os.File.Write() implements the io.Writer interface when writing to a file 
* bytes.Buffer implements the io.Writer interface which is passed into NewEncoder

## Docker
Used `docker init` to get Docker to create:

* .dockerignore
* docker-compose.yaml
* Dockerfile

In order to allow the container to create and write a text file, added `COPY . .` to `Dockerfile` to ensure the source files were added to the image during the `build` stage.

At end of `Dockerfile`, copied the text file to the `final` stage with `COPY --chown=appuser:appuser hi.txt /bin/`.

This copies the file to the container's `bin` directory, applying the user created here:

```
ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser
USER appuser
```
... so Docker has permission to open and write to the file.

When opening the file in Go code, the path is:

`/bin/hi.txt`