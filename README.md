# Microbadger Notify

[![Build Status](https://droneci.rwxlabs.io/api/badges/ldejager/microbadger-notify/status.svg)](https://droneci.rwxlabs.io/ldejager/microbadger-notify) [![](https://images.microbadger.com/badges/image/ldejager/microbadger-notify.svg)](https://microbadger.com/images/ldejager/microbadger-notify "Get your own image badge on microbadger.com")

Simple utility container which can be used in drone build pipelines to notify microbadger that a new docker image has been pushed.

### Usage

```
build:
  image: ldejager/microbadger-notify
  environment:
    - MB_REPOSITORY=your/repository
    - MB_TOKEN=sometoken
  when:
    status: success
```
