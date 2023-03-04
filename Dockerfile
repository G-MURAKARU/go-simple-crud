# # In Dockerfiles, the order of commands is important because each command translates to a 'layer' in the image
# # Meaning if one layer changes, all subsequent layers inclusive will be rebuilt/rerun even if it is unnecessary
# # Therefore, to optimise caching image layers, best practice is to order each command from the least frequently-changing command
# # to the most frequently-changing command

# # TO REDUCE THE FINAL IMAGE SIZE, USE MULTI-STAGE BUILDS
# # Since Go is a compiled language, the only thing needed to run the app is the final executable
# # everything else is redundant during runtime
# # Therefore, it is best practice to build an image from just the executable

# # We specify the base image we need for our Go application
# # 3 best practices - use a (1.) specific version of an (2.) official, (3.) small-sized image as the base image
# # golang - official image;; based on alpine - small size;; included image tag - specific version

# # FIRST STAGE - BUILD # #

# original:
# FROM golang:1.20-alpine3.17

# multistage - can take advantage of larger OS images, some features may be needed:
FROM golang:1.20-alpine3.17 AS builder

# # Create /app directory within the application to hold our application's source code
WORKDIR /app

# # Pre-copy/Cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
# # Best practice - optimising caching image layers
# # ./ below signifies WORKDIR above
COPY go.mod go.sum ./

# # Install application's dependencies
RUN go mod download && go mod verify

# # Copy the rest of our source files to the image's working directory
COPY . .

# # Build the Go application (with optional dependencies/flags)
RUN CGO_ENABLED=0 GOOS=linux go build -o main

# # SECOND STAGE - RUN # #

# # Specify a more lightweight base image
FROM alpine:3.17

# # Define a new workdir for the new image
WORKDIR /app

# # 'Import' output from previous stage
COPY --from=builder /app/main /app/.env ./

# # Configure the port that the container listens to at runtime
# # NOTE: this does not actually open up the port
# # it is only for documentation purposes i.e. tells anyone using the image that this is the designed/intended port to use
EXPOSE 3000

# # Creating a new low-privilege user
RUN addgroup -S goUsers && adduser -S goUser -G goUsers

# # Assigning permissions to the new user
RUN chown -R goUser:goUsers /app

# # Switch to the newly-created user
USER goUser

# original:
# # Set the command to execute when the image is used to start a container
# CMD [ "/app/go-docker" ]

# multistage:
# # Set the entrypoint for the image
CMD [ "/app/main" ]