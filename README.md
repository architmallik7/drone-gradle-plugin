# Gradle-plugin

- [Synopsis](#Synopsis)
- [Plugin Image](#Plugin-Image)
- [Parameters](#Parameters)
- [Building](#building)
- [Examples](#Examples)


## Synopsis

This plugin is used to build Java application using [Gradle](https://gradle.org). 


## Plugin Image

The plugin `Rakshitharness/drone-gradle-plugin` is available for the following architectures:

| OS            | Tag                                |
| ------------- | ---------------------------------- |
| linux/amd64   | `linux-amd64`                      |
| linux/arm64   | `linux-arm64`                      |
| windows/amd64 | `windows-amd64`                    |

## Parameters

| Parameter                                                                        | Comments                                                                                                                                  |
|:---------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------|
| Goals <span style="font-size: 10px"></span>                       | An array of gradle goals to run.     


## Building

Build the plugin image:

```text
./scripts/build.sh
```

## Examples

```
# Plugin YAML
- step:
    type: Plugin
    name: gradle-plugin-arm64
    identifier: gradle-plugin-arm64
    spec:
        connectorRef: harness-docker-connector
        image: rakshitagar/drone-gradle-plugin:linux-arm64
       

- step:
    type: Plugin
    name: gradle-plugin-amd64
    identifier: gradle-plugin-amd64
    spec:
        connectorRef: harness-docker-connector
        image: rakshitagar/drone-gradle-plugin:linux-amd64
        