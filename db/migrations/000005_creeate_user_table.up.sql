create table "user"
(
    uuid               uuid                     default gen_random_uuid() not null
        constraint user_pk
            primary key,
    username           varchar(32)                                        not null,
    host               varchar(64)
        constraint user_host_host_fk
            references host
            on update restrict on delete restrict,
    email              varchar(128),
    display_name       varchar(64),
    summary            text,
    password           varchar(64),
    created_at         timestamp with time zone default now(),
    updated_at         timestamp with time zone default now(),
    is_active          boolean                  default true              not null,
    is_silence         boolean                  default false             not null,
    is_suspend         boolean                  default false             not null,
    accept_manually    boolean                  default false             not null,
    is_superuser       boolean                  default false             not null,
    ap_url_uuid        uuid
        constraint user_ap_url_uuid_fk
            references ap_url
            on update restrict on delete set null,
    signature_key_uuid uuid
        constraint user_signature_key_uuid_fk
            references signature_key
            on update restrict on delete set null,
    avatar_uuid        uuid,
    header_uuid        uuid,
    is_mail_verified   boolean                  default false,
    is_bot             boolean                  default false             not null,
    two_fa             uuid
        constraint user_user_2fa_uuid_fk
            references user_2fa
            on update restrict on delete set null,
    constraint acct_pk
        unique (username, host)
);

create unique index user_email_uindex
    on "user" (email);

create unique index user_ap_url_uuid_uindex
    on "user" (ap_url_uuid);

create unique index user_signature_key_uuid_uindex
    on "user" (signature_key_uuid);

create unique index user_two_fa_uindex
    on "user" (two_fa);
