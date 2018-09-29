FROM drone/ca-certs

ADD src/campaign /

CMD ["/campaign"]
