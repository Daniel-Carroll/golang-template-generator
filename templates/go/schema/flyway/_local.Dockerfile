# Use flyway runtime as base
FROM boxfuse/flyway:5.2.4-alpine

# Map configuration
COPY conf /flyway/conf

# Map migration scripts to expected locations
COPY sql /flyway/sql

# Remove partner db access scripts
RUN rm -rf /flyway/sql/migrations/*_db_access.sql

CMD ["migrate"]