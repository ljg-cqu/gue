CREATE TABLE IF NOT EXISTS gue_jobs
 (
     job_id      bigserial   NOT NULL PRIMARY KEY,
     job_type    text        NOT NULL,
     queue       text        NOT NULL,
     args        json        NOT NULL,

     priority    smallint    NOT NULL,
     run_at      timestamptz NOT NULL,

     error_count integer     NOT NULL DEFAULT 0,
     last_error  text,

     created_at  timestamptz NOT NULL,
     updated_at  timestamptz NOT NULL,
     finished_at  timestamptz
 );

CREATE INDEX IF NOT EXISTS "idx_gue_jobs_selector" ON "gue_jobs" ("queue", "run_at", "priority");

COMMENT ON TABLE gue_jobs IS '1';

CREATE TABLE IF NOT EXISTS gue_jobs_finished
(
    job_id      bigserial   NOT NULL PRIMARY KEY,
    job_type    text        NOT NULL,
    queue       text        NOT NULL,
    args        json        NOT NULL,

    priority    smallint    NOT NULL,
    run_at      timestamptz NOT NULL,

    error_count integer     NOT NULL DEFAULT 0,
    last_error  text,

    created_at  timestamptz NOT NULL,
    updated_at  timestamptz NOT NULL,
    finished_at  timestamptz NOT NULL
);

CREATE INDEX IF NOT EXISTS "idx_gue_jobs_finished_selector" ON "gue_jobs_finished" ("job_type", "priority", "queue", "finished_at");

COMMENT ON TABLE gue_jobs_finished IS 'Finished gue jobs log';