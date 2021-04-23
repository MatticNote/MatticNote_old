create table if not exists host
(
    host        varchar(64) not null
                constraint host_pk primary key,
    found_at    timestamp with time zone default now(),
    is_suspend  boolean default false not null
);

create unique index host_host_uindex
    on host (host);
