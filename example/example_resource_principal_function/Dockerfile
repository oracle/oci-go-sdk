FROM fnproject/go:dev as build-stage
ARG GOFLAGS=-mod=vendor
ADD . /func/
RUN cd /func/ && go mod vendor -v && go build -o func
FROM fnproject/go
WORKDIR /function
COPY --from=build-stage /func/func /function/
ENTRYPOINT ["./func"]
