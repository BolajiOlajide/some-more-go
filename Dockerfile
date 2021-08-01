################## Create GoLang dev pipeline ##################
FROM golang:1.16 as dev

WORKDIR /work

################## Create GoLang build pipeline ##################
FROM golang:1.16 as build

COPY ./app/* /app/
RUN go build -o app

FROM alpine as runtime
COPY --from=build /app/app /
CMD ./app
