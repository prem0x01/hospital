-- name: GetAppointments :many
SELECT
    a.id, a.patient_id, a.doctor_id, a.appointment_date, a.status,
    a.notes, a.diagnosis, a.treatment_plan, a.created_by,
    a.created_at, a.updated_at,
    p.first_name || ' ' || p.last_name as patient_name,
    COALESCE(u.first_name || ' ' || u.last_name, '') as doctor_name
FROM appointments a
JOIN patients p ON a.patient_id = p.id
LEFT JOIN users u ON a.doctor_id = u.id
ORDER BY a.appointment_date DESC
LIMIT $1 OFFSET $2;

-- name: GetAppointmentsByDoctor :many
SELECT
    a.id, a.patient_id, a.doctor_id, a.appointment_date, a.status,
    a.notes, a.diagnosis, a.treatment_plan, a.created_by,
    a.created_at, a.updated_at,
    p.first_name || ' ' || p.last_name as patient_name,
    COALESCE(u.first_name || ' ' || u.last_name, '') as doctor_name
FROM appointments a
JOIN patients p ON a.patient_id = p.id
LEFT JOIN users u ON a.doctor_id = u.id
WHERE a.doctor_id = $1
ORDER BY a.appointment_date DESC
LIMIT $2 OFFSET $3;

-- name: GetAppointmentByID :one
SELECT
    a.id, a.patient_id, a.doctor_id, a.appointment_date, a.status,
    a.notes, a.diagnosis, a.treatment_plan, a.created_by,
    a.created_at, a.updated_at,
    p.first_name || ' ' || p.last_name as patient_name,
    COALESCE(u.first_name || ' ' || u.last_name, '') as doctor_name
FROM appointments a
JOIN patients p ON a.patient_id = p.id
LEFT JOIN users u ON a.doctor_id = u.id
WHERE a.id = $1;

-- name: CreateAppointment :one
INSERT INTO appointments (patient_id, doctor_id, appointment_date, notes, created_by)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, patient_id, doctor_id, appointment_date, status, notes,
          diagnosis, treatment_plan, created_by, created_at, updated_at;

-- name: UpdateAppointment :one
UPDATE appointments
SET
    doctor_id = COALESCE(sqlc.narg(doctor_id), doctor_id),
    appointment_date = COALESCE(sqlc.narg(appointment_date), appointment_date),
    status = COALESCE(sqlc.narg(status), status),
    notes = COALESCE(sqlc.narg(notes), notes),
    diagnosis = COALESCE(sqlc.narg(diagnosis), diagnosis),
    treatment_plan = COALESCE(sqlc.narg(treatment_plan), treatment_plan),
    updated_at = NOW()
WHERE id = $1
RETURNING id, patient_id, doctor_id, appointment_date, status, notes,
          diagnosis, treatment_plan, created_by, created_at, updated_at;

-- name: DeleteAppointment :exec
DELETE FROM appointments WHERE id = $1;

-- name: CountAppointments :one
SELECT COUNT(*) FROM appointments;

-- name: CountAppointmentsByStatus :one
SELECT COUNT(*) FROM appointments WHERE status = $1;

-- name: GetTodaysAppointments :many
SELECT
    a.id, a.patient_id, a.doctor_id, a.appointment_date, a.status,
    a.notes, a.diagnosis, a.treatment_plan, a.created_by,
    a.created_at, a.updated_at,
    p.first_name || ' ' || p.last_name as patient_name,
    COALESCE(u.first_name || ' ' || u.last_name, '') as doctor_name
FROM appointments a
JOIN patients p ON a.patient_id = p.id
LEFT JOIN users u ON a.doctor_id = u.id
WHERE DATE(a.appointment_date) = CURRENT_DATE
ORDER BY a.appointment_date;

-- name: GetAppointmentsByDateRange :many
SELECT
    a.id, a.patient_id, a.doctor_id, a.appointment_date, a.status,
    a.notes, a.diagnosis, a.treatment_plan, a.created_by,
    a.created_at, a.updated_at,
    p.first_name || ' ' || p.last_name as patient_name,
    COALESCE(u.first_name || ' ' || u.last_name, '') as doctor_name
FROM appointments a
JOIN patients p ON a.patient_id = p.id
LEFT JOIN users u ON a.doctor_id = u.id
WHERE a.appointment_date BETWEEN $1 AND $2
ORDER BY a.appointment_date;

-- name: GetPatientAppointments :many
SELECT
    a.id, a.patient_id, a.doctor_id, a.appointment_date, a.status,
    a.notes, a.diagnosis, a.treatment_plan, a.created_by,
    a.created_at, a.updated_at,
    p.first_name || ' ' || p.last_name as patient_name,
    COALESCE(u.first_name || ' ' || u.last_name, '') as doctor_name
FROM appointments a
JOIN patients p ON a.patient_id = p.id
LEFT JOIN users u ON a.doctor_id = u.id
WHERE a.patient_id = $1
ORDER BY a.appointment_date DESC;
