CREATE TABLE IF NOT EXISTS observation_services
(
    id                  SERIAL PRIMARY KEY,
    project_id          BIGINT,
    name                varchar(64) NOT NULL,
    source              jsonb,
    created_at          timestamp NOT NULL default current_timestamp,
    updated_at          timestamp NOT NULL default current_timestamp,
    status              varchar(32),
    error               varchar(2048),
    UNIQUE              (project_id, name)
);
