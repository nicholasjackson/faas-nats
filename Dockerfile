FROM alpine

RUN adduser -h /home/faasnats -D faasnats faasnats

COPY ./faas-natsd /home/faasnats/
RUN chmod +x /home/faasnats/faas-natsd

USER faasnats

ENTRYPOINT ["/home/faasnats/faas-natsd"]
