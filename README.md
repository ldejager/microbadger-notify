# Microbadger Notify

Simple go application which can be used in build pipelines to notify microbadger that a new docker image has been pushed.

### Usage

*Drone*

```
build:
  image: ldejager/microbadger-notify
  environment:
    - MB_REPOSITORY=your/repository
    - MB_TOKEN=sometoken
  commands:
    - microbadger-notify
  when:
    status: success
```
