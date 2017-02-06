FROM golang:1.7

RUN mkdir /src
WORKDIR /src

# this will ideally be built by the ONBUILD below ;)
CMD ["go-wrapper", "run"]

ONBUILD COPY . /src
ONBUILD RUN go-wrapper download
ONBUILD RUN go-wrapper install
