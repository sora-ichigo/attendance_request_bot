# build stage
# ------------------------------------------
FROM golang:1.19.3-bullseye AS build

WORKDIR /app
COPY . .

RUN apt-get update && apt-get install -y \
  ca-certificates \
  cmake \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/* \
  && make
# ------------------------------------------

# application stage
# ------------------------------------------
FROM debian:bullseye-slim

WORKDIR /app

ENV TZ Asia/Tokyo
ENV PORT 80

RUN apt-get update && apt-get install -y \
  ca-certificates \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

COPY --from=build /app/bin/ /app/bin/
ENV PATH /app/bin:$PATH
RUN chmod +x /app/bin/*

EXPOSE 80

CMD ["request_attendance_server"]
# ------------------------------------------
