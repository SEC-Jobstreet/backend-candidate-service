-- -- name: CreateApplication :one
-- INSERT INTO applications (
--     candidate_id,
--     job_id,
--     status
-- ) VALUES (
--     $1, $2, $3
-- ) RETURNING *;
--
-- -- name: GetApplication :one
-- SELECT * FROM applications
-- WHERE id = $1 LIMIT 1;
--
-- -- name: ListApplications :many
-- SELECT * FROM applications
-- WHERE
--     (candidate_id = @candidate_id OR @candidate_id = 0)
--     AND (job_id = @job_id OR @job_id = 0)
--     AND (status = @status OR @status = '')
-- ORDER BY id
-- LIMIT @l
-- OFFSET @o;
--
-- -- name: UpdateStatusApplication :one
-- UPDATE applications
-- SET
--     status = sqlc.arg(status),
--     message = COALESCE(sqlc.narg(message), message),
--     updated_at = sqlc.arg(updated_at)
-- WHERE
--     id = sqlc.arg(id)
-- RETURNING *;

-- name: GetCandidateProfiles :many
SELECT user_id,
       google_id,
       COALESCE(email, '')                AS email,
       COALESCE(first_name, '')           AS first_name,
       COALESCE(last_name, '')            AS last_name,
       COALESCE(profile_image, '')        AS profile_image,
       COALESCE(first_name_profile, '')   AS first_name_profile,
       COALESCE(last_name_profile, '')    AS last_name_profile,
       COALESCE(phone, '')                AS phone,
       COALESCE(phone_number_country, '') AS phone_number_country,
       COALESCE(address, '')              AS address,
       COALESCE(current_location, '')     AS current_location,
       COALESCE(privacy_setting, '')      AS privacy_setting,
       work_eligibility,
       COALESCE(resume_link, '')          AS resume_link,
       COALESCE("current_role", '')       AS current_role,
       work_whenever,
       work_shift                         AS work_shift,
       COALESCE(location_lat, 0.0)        AS location_lat,
       COALESCE(location_lon, 0.0)        AS location_lon,
       visa,
       COALESCE(description, '')          AS description,
       COALESCE(position, '') AS position,
    COALESCE(start_date, '1970-01-01') AS start_date, -- Assuming DATE type, default to UNIX epoch start
    share_profile,
    updated_at,
    created_at
FROM public.candidate_profile
WHERE user_id = @user_id;

-- name: CreateCandidateProfile :one
INSERT INTO candidate_profile (user_id,
                               created_at,
                               updated_at)
VALUES ($1, now(), now()) RETURNING *;