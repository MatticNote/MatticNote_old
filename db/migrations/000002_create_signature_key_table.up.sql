create table if not exists signature_key
(
    uuid        uuid default gen_random_uuid() not null
        constraint signature_key_pk
            primary key,
    public_key  text,
    private_key text
);
