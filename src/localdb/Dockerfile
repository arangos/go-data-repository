FROM postgres:16-bookworm

ENV POSTGRES_PASSWORD=postgres

RUN mkdir /docker-entrypoint-initdb.d || true
COPY mccusa-intialization.sql /docker-entrypoint-initdb.d/mccusa-intialization.sql

RUN chown -R postgres:postgres /docker-entrypoint-initdb.d/