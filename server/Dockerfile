FROM golang

ADD ./../bin /opt

RUN chmod +x /opt/compilebox_linux

ENV COMPILEBOX_PORT=80

ENTRYPOINT /opt/compilebox_linux

EXPOSE 80