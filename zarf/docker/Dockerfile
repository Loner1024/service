FROM golang:1.17 as build_service
ENV CGO_ENABLE 0
ARG BUILD_REF

COPY . /service

WORKDIR /service
RUN go build -ldflags "-X main.build=${BUILD_REF}"

FROM alpine:latest
ARG BUILD_DATE
ARG BUILD_REF
COPY --from=build_service /service/service /
CMD ["./service"]

# LABEL org.opencontainers.image.created="${BUILD_DATE}" \
#       org.opencontainers.image.title="service" \
#       org.opencontainers.image.authors="Loner Xie <xth12138@gmail.com>" \
#       org.opencontainers.image.source="https://github.com/Loner1024/service" \
#       org.opencontainers.image.revision="${BUILD_REF}" \
#       org.opencontainers.image.vendor="Loner Xie"