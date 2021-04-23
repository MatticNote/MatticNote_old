create table if not exists ap_url
(
    uuid         uuid default gen_random_uuid() not null
        constraint ap_url_pk
            primary key,
    inbox        varchar(256),
    outbox       varchar(256),
    shared_inbox varchar(256),
    featured     varchar(256),
    uri          varchar(256),
    url          varchar(256)
);
