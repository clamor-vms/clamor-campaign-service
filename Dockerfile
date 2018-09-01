FROM drone/ca-certs

ADD src/operation /

CMD ["/operation"]
