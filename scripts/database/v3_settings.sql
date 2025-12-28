create table setting
(
    created      timestamp with time zone default now()              not null,
    updated      timestamp with time zone default now()              not null,
    deleted      timestamp with time zone,
    id           serial unique primary key,
    uuid         uuid                     default uuid_generate_v4() not null,
    key          varchar(64)                                         not null,
    value        varchar(512)                                        not null
);

create trigger set_timestamp
    before update
    on setting
    for each row
execute procedure trigger_set_timestamp();

INSERT INTO setting (key, value) VALUES ('health_status', 'healthy');