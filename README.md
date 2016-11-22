# Microbadger Notify

[![Build Status](https://droneci.rwxlabs.io/api/badges/ldejager/microbadger-notify/status.svg)](https://droneci.rwxlabs.io/ldejager/microbadger-notify) [![](https://images.microbadger.com/badges/image/ldejager/microbadger-notify.svg)](https://microbadger.com/images/ldejager/microbadger-notify "Get your own image badge on microbadger.com")

Simple utility container which can be used in drone build pipelines to notify Microbadger via an HTTP POST that a new docker image has been pushed which triggers a metadata refresh on Microbadger for the docker image.

### Configuration

You can get your repository token from signing into Microbadger and clicking the "Get the webhook" link. The token is the last part of the URL.

### Usage

```
notify:
  image: ldejager/microbadger-notify
  environment:
    - MB_REPOSITORY=your/repository
    - MB_TOKEN=sometoken
  when:
    status: success
```
