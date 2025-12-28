create table "endpoint_export"
(
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires TIMESTAMPTZ NOT NULL DEFAULT (NOW() + INTERVAL '24 hours'),
    id      SERIAL PRIMARY KEY,
    code    VARCHAR(16) NOT NULL UNIQUE,
    data    JSONB NOT NULL
);

create trigger set_timestamp
    before update
    on "endpoint_export"
    for each row
execute procedure trigger_set_timestamp();

-- Index for fast lookups by the short code.
CREATE INDEX idx_endpoint_export_code
    ON endpoint_export (code);

-- Index to efficiently find and clean up expired records.
CREATE INDEX idx_endpoint_export_expires_at
    ON endpoint_export (expires);
