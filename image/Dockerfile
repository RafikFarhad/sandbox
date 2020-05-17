FROM alpine:latest

LABEL maintainer=RafikFarhad

RUN addgroup -S sandbox

RUN adduser -D -u 1000 -G sandbox sandbox

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

COPY run.sh /home/sandbox/

RUN chmod +x /home/sandbox/run.sh

WORKDIR /home/sandbox

USER sandbox

CMD ["sh", "run.sh"]
