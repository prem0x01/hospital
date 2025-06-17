-- name: GetPatients :many
SELECT id, first_name, last_name, email, phone, date_of_birth, gender,
       address, medical_history, allergies, emergency_contact_name,
       emergency_contact_phone, created_by, created_at, updated_at
FROM patients
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: GetPatientByID :one
SELECT id, first_name, last_name, email, phone, date_of_birth, gender,
       address, medical_history, allergies, emergency_contact_name,
       emergency_contact_phone, created_by, created_at, updated_at
FROM patients
WHERE id = $1;

-- name: CreatePatient :one
INSERT INTO patients (
    first_name, last_name, email, phone, date_of_birth,
    gender, address, medical_history, allergies,
    emergency_contact_name, emergency_contact_phone, created_by
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING id, first_name, last_name, email, phone, date_of_birth, gender,
          address, medical_history, allergies, emergency_contact_name,
          emergency_contact_phone, created_by, created_at, updated_at;

-- name: UpdatePatient :one
UPDATE patients
SET
    first_name = COALESCE(sqlc.narg(first_name), first_name),
    last_name = COALESCE(sqlc.narg(last_name), last_name),
    email = COALESCE(sqlc.narg(email), email),
    phone = COALESCE(sqlc.narg(phone), phone),
    date_of_birth = COALESCE(sqlc.narg(date_of_birth), date_of_birth),
    gender = COALESCE(sqlc.narg(gender), gender),
    address = COALESCE(sqlc.narg(address), address),
    medical_history = COALESCE(sqlc.narg(medical_history), medical_history),
    allergies = COALESCE(sqlc.narg(allergies), allergies),
    emergency_contact_name = COALESCE(sqlc.narg(emergency_contact_name), emergency_contact_name),
    emergency_contact_phone = COALESCE(sqlc.narg(emergency_contact_phone), emergency_contact_phone),
    updated_at = NOW()
WHERE id = $1
RETURNING id, first_name, last_name, email, phone, date_of_birth, gender,
          address, medical_history, allergies, emergency_contact_name,
          emergency_contact_phone, created_by, created_at, updated_at;

-- name: DeletePatient :exec
DELETE FROM patients WHERE id = $1;

-- name: CountPatients :one
SELECT COUNT(*) FROM patients;

-- name: SearchPatients :many
SELECT id, first_name, last_name, email, phone, date_of_birth, gender,
       address, medical_history, allergies, emergency_contact_name,
       emergency_contact_phone, created_by, created_at, updated_at
FROM patients
WHERE
    LOWER(first_name) LIKE LOWER('%' || $1 || '%') OR
    LOWER(last_name) LIKE LOWER('%' || $1 || '%') OR
    LOWER(email) LIKE LOWER('%' || $1 || '%') OR
    phone LIKE '%' || $1 || '%'
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;
