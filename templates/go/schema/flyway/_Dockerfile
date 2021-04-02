# Use flyway runtime as base
FROM boxfuse/flyway:5.2.4-alpine

# Map configuration
COPY conf /flyway/conf

# Map migration scripts to expected locations
COPY sql /flyway/sql/migrations

CMD ["migrate"]