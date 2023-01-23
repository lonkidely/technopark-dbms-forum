FROM golang:latest AS build

ADD . /app
WORKDIR /app
RUN go build ./cmd/main.go


FROM ubuntu:20.04

RUN apt-get -y update && apt-get install -y tzdata

ENV PGVER 12

RUN apt-get -y update && apt-get install -y postgresql-$PGVER

USER postgres

RUN /etc/init.d/postgresql start &&\
    psql --command "CREATE USER lonkidely WITH SUPERUSER PASSWORD 'lonkidely';" &&\
    createdb -O lonkidely tech_db_forum &&\
    /etc/init.d/postgresql stop

RUN echo "listen_addresses='*'" >> /etc/postgresql/$PGVER/main/postgresql.conf
RUN echo "host all all 0.0.0.0/0 md5" >> /etc/postgresql/$PGVER/main/pg_hba.conf

EXPOSE 5432

VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

USER root

WORKDIR /usr/src/app

COPY . .
COPY --from=build /app/main .

EXPOSE 5000
EXPOSE 5432
ENV PGPASSWORD lonkidely
CMD service postgresql start && psql -h localhost -d tech_db_forum -U lonkidely -p 5432 -a -q -f ./db/db.sql && ./main