# Configuration file

flyway.conf stores the default configuration for the migration library. Environment variables should be used to overwrite any values that should be different by environment or that are sensitive, as this file will be commited unencrypted

If you want flyway to own the schema (create and destroy it as necessary), add the schema name to the key:
flyway.schemas=