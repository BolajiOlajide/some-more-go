################## Create GoLang dev pipeline ##################
FROM golang:1.16 as dev

WORKDIR /work

################## Create GoLang build pipeline ##################
FROM golang:1.16 as build

WORKDIR /app
COPY ./app/* /app/
RUN go build -o app

################## Create GoLang runtime pipeline ##################
FROM alpine as runtime
COPY --from=build /app/app /
CMD ./app
