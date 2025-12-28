CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SELECT uuid_generate_v4();

create function trigger_set_timestamp() returns trigger
    language plpgsql
as
$$
BEGIN
    NEW.updated = NOW();
    RETURN NEW;
END;
$$;