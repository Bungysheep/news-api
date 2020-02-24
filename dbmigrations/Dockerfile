FROM migrate/migrate

RUN apk --update add postgresql-client

COPY /migrations /migrations

COPY postgres_custom_migration.sh .

ENTRYPOINT [ "" ]