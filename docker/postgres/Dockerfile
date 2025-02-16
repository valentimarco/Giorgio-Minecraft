FROM postgres:16

RUN apt-get update 
RUN apt-get install -y curl
RUN curl -L "https://github.com/pksunkara/pgx_ulid/releases/download/v0.2.0/pgx_ulid-v0.2.0-pg16-amd64-linux-gnu.deb" -o /tmp/pgx_ulid.deb
RUN apt-get install -y /tmp/*.deb

COPY ./plugins.sql /docker-entrypoint-initdb.d/