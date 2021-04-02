# SQL migration files

Place here sql files to be applied to the dataabse. They will be applied on numeric order. 

Files should follow the naming conventions specified below (unless changed via configuration):

### One time migrations:

These are changes that need to be executed one time, and will be applied in the natural order of the version.
Due to the size of the team, and potential for collisions, `outOfOrder` mode is enabled. To help with this,
the following naming convention *needs* to be followed for the sql scripts:

`V[year][month][day].[military time]__[nice_title_snake_case]`
`V190903.1527__add_audit_timestamps.sql` would be an acceptable filename. HUGE NOTE: you need to add two underscores after
the version, or everything breaks. 


### Repeatable migrations:

These are scripts that should be executed repateably (i.e. creation of views or materialized views). They will be executed if there has been a change on the file (verified via checksum)

R__[description].sql

Examples

R__refresh views.sql
R__rebuild_indexes.sql

# Secret management

No sensitive information should be written on the SQL migration files (i.e. user passwords, API keys, etc). Instead place a placeholder of the form ${property_name} with the property name on lowercase.

These values will be fed to the container running the migration via environment variables. For local development an env file can be created and referenced on docker-compose (sensitive values shouldn't be the same as on live environments)
