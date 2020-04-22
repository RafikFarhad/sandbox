FROM alpine:latest

LABEL maintainer=RafikFarhad

RUN addgroup -S sandbox

RUN adduser -D -u 1000 -G sandbox sandbox

WORKDIR /home/sandbox/raw

RUN apk update && \
    apk add --no-cache --virtual \
    build-dependencies \
    build-base \
    libseccomp-dev \
    gcc \
    go

# GoLang
RUN export PATH="/usr/local/go/bin:$PATH" \
    export GOPATH=/opt/go/ \
    export PATH=$PATH:$GOPATH/bin

COPY ./container ./raw

USER sandbox

CMD ["sh", "./run.sh"]
