FROM scratch
ADD release/linux/amd64/microbadger-notify /
CMD ["/microbadger-notify"]
