FROM mariadb

WORKDIR /docker-entrypoint-initdb.d

COPY ./api/sql/sql.sql .

EXPOSE 3306


# docker container run --name devbookdb --rm --env MARIADB_ROOT_PASSWORD=secret --env MARIADB_DATABASE=mariadbbase -it -p 3306:3306 devbookdb