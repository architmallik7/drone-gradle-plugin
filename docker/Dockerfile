# Use Alpine Linux as the base image for the final stage
FROM alpine:latest

# # Install Gradle and Java using apk
RUN apk update && \
    apk add --no-cache openjdk11 && \
    apk add --no-cache gradle

# Copy the built Go executable from the builder stage to the final stage
COPY release/linux/amd64/drone-gradle-plugin /bin/

# Set the entrypoint to the Go executable
ENTRYPOINT ["/bin/drone-gradle-plugin"]
