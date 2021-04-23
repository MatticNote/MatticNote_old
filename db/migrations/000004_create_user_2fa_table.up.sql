create table if not exists user_2fa
(
    uuid        uuid    default gen_random_uuid() not null
        constraint user_2fa_pk
            primary key,
    is_enable   boolean default false             not null,
    secret_code char(32)                          not null,
    backup_code jsonb   default '[]'::jsonb
);
