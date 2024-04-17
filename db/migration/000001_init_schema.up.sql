CREATE TABLE "candidate_profile"
(
    "user_id"            bigserial PRIMARY KEY,
    "google_id"          bigserial,
    "email"              VARCHAR(255) NOT NULL,
    "first_name"         VARCHAR(255),
    "last_name"          VARCHAR(255),
    "profile_image"      TEXT,
    "first_name_profile" VARCHAR(255),
    "last_name_profile"  VARCHAR(255),
    "phone"              VARCHAR(20),
    "address"            TEXT,
    "location_lat"       DOUBLE PRECISION,
    "location_lon"       DOUBLE PRECISION,
    "visa"               BOOLEAN,
    "description"        TEXT,
    "position"           VARCHAR(255),
    "start_date"         DATE,
    "work_whenever"      BOOLEAN,
    "work_shift"         VARCHAR(50),
    "share_profile"      BOOLEAN,
    "resume_link"        TEXT,
    "updated_at"         timestamptz  NOT NULL DEFAULT (now()),
    "created_at"         timestamptz  NOT NULL DEFAULT (now()),
    "auth_method"        VARCHAR(255)
);

CREATE INDEX ON "candidate_profile" ("google_id");
CREATE UNIQUE INDEX ON "candidate_profile" ("email");

-- ALTER TABLE "applications"
--     ADD CONSTRAINT "candidate_job_key" UNIQUE ("candidate_id", "job_id");