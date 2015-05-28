FROM flynn/busybox

ADD ./kempctl /

ENTRYPOINT ["/kempctl"]

