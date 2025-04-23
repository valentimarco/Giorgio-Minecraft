CREATE TABLE IF NOT EXISTS servers (
    id ulid NOT NULL DEFAULT gen_ulid() PRIMARY KEY,
    name text not null,
    version text not null,
    type text not null,
    memory INT not null,
    max_player INT not null,
    aika_flag BOOLEAN not null
);