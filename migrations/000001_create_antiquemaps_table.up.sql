CREATE TABLE IF NOT EXISTS antiquemaps (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    year integer NOT NULL,
    country text NOT NULL,
    condition text NOT NULL,
    type text NOT NULL,
    version integer NOT NULL DEFAULT 1
);