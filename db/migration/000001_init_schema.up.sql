CREATE TABLE "candidate_profile"
(
    "user_id"              VARCHAR(255) PRIMARY KEY,
    "google_id"            bigserial,
    "email"                VARCHAR(255),
    "first_name"           VARCHAR(255),
    "last_name"            VARCHAR(255),
    "profile_image"        TEXT,
    "first_name_profile"   VARCHAR(255),
    "last_name_profile"    VARCHAR(255),
    "phone"                VARCHAR(20),
    "phone_number_country" VARCHAR(10),
    "address"              TEXT,
    "current_location"     TEXT,
    "privacy_setting"      VARCHAR(50),
    "work_eligibility"     JSONB,
    "resume_link"          TEXT,
    "resume"               TEXT,
    "current_role"         VARCHAR(100),
    "work_whenever"        BOOLEAN,
    "work_shift"           JSONB,
    "location_lat"         DOUBLE PRECISION,
    "location_lon"         DOUBLE PRECISION,
    "visa"                 BOOLEAN,
    "description"          TEXT,
    "position"             VARCHAR(255),
    "start_date"           DATE,
    "share_profile"        BOOLEAN,
    "updated_at"           timestamptz NOT NULL DEFAULT (now()),
    "created_at"           timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "candidate_profile" ("google_id");
-- CREATE UNIQUE INDEX ON "candidate_profile" ("email");

-- ALTER TABLE "applications"
--     ADD CONSTRAINT "candidate_job_key" UNIQUE ("candidate_id", "job_id");