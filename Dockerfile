FROM centurylink/ca-certs
ADD microbadger-notify /
CMD ["/microbadger-notify"]
