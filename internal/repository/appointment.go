package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/prem0x01/hospital/internal/database/queries"
	"github.com/prem0x01/hospital/internal/domain"
)

type AppointmentRepository struct {
	q      *queries.Queries
	dbConn queries.DBTX
}

func NewAppointmentRepository(q *queries.Queries, dbConn queries.DBTX) *AppointmentRepository {
	return &AppointmentRepository{q: q, dbConn: dbConn}
}

func (r *AppointmentRepository) GetAll(ctx context.Context, limit, offset int32, userRole string, userID int32) ([]domain.Appointment, error) {
	if userRole == "doctor" {
		doctorAppointments, err := r.q.GetAppointmentsByDoctor(ctx, queries.GetAppointmentsByDoctorParams{
			DoctorID: &userID,
			Limit:    limit,
			Offset:   offset,
		})
		if err != nil {
			return nil, err
		}

		result := make([]domain.Appointment, 0, len(doctorAppointments))
		for _, a := range doctorAppointments {
			appointment := toDomainAppointmentFromRow(queries.GetAppointmentsRow(*a))
			result = append(result, *appointment)
		}
		return result, nil
	}

	appointments, err := r.q.GetAppointments(ctx, queries.GetAppointmentsParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	result := make([]domain.Appointment, 0, len(appointments))
	for _, a := range appointments {
		appointment := toDomainAppointmentFromRow(*a)
		result = append(result, *appointment)
	}
	return result, nil
}

func (r *AppointmentRepository) GetByID(ctx context.Context, id int32) (*domain.Appointment, error) {
	a, err := r.q.GetAppointmentByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toDomainAppointmentFromRow(*convertGetAppointmentByIDRowToGetAppointmentsRow(a)), nil
}

func convertGetAppointmentByIDRowToGetAppointmentsRow(a *queries.GetAppointmentByIDRow) *queries.GetAppointmentsRow {
	if a == nil {
		return nil
	}
	return &queries.GetAppointmentsRow{
		ID:              a.ID,
		PatientID:       a.PatientID,
		DoctorID:        a.DoctorID,
		AppointmentDate: a.AppointmentDate,
		Status:          a.Status,
		Notes:           a.Notes,
		Diagnosis:       a.Diagnosis,
		TreatmentPlan:   a.TreatmentPlan,
		CreatedBy:       a.CreatedBy,
		CreatedAt:       a.CreatedAt,
		UpdatedAt:       a.UpdatedAt,
		PatientName:     a.PatientName,
		DoctorName:      a.DoctorName,
	}
}

func (r *AppointmentRepository) Create(ctx context.Context, a *domain.Appointment) error {
	result, err := r.q.CreateAppointment(ctx, queries.CreateAppointmentParams{
		PatientID:       a.PatientID,
		DoctorID:        a.DoctorID,
		AppointmentDate: a.AppointmentDate,
		Notes:           a.Notes,
		CreatedBy:       a.CreatedBy,
	})
	if err != nil {
		return err
	}

	a.ID = result.ID
	a.Status = result.Status
	a.CreatedAt = result.CreatedAt
	a.UpdatedAt = result.UpdatedAt
	return nil
}

func (r *AppointmentRepository) Delete(ctx context.Context, id int32) error {
	return r.q.DeleteAppointment(ctx, id)
}

func (r *AppointmentRepository) Count(ctx context.Context) (int64, error) {
	return r.q.CountAppointments(ctx)
}

func toDomainAppointmentFromRow(a queries.GetAppointmentsRow) *domain.Appointment {
	return &domain.Appointment{
		ID:              a.ID,
		PatientID:       a.PatientID,
		DoctorID:        a.DoctorID,
		AppointmentDate: a.AppointmentDate,
		Status:          a.Status,
		Notes:           a.Notes,
		Diagnosis:       a.Diagnosis,
		TreatmentPlan:   a.TreatmentPlan,
		CreatedBy:       a.CreatedBy,
		CreatedAt:       a.CreatedAt,
		UpdatedAt:       a.UpdatedAt,
		PatientName:     a.PatientName.(string),
		DoctorName:      a.DoctorName.(string),
	}
}

func (r *AppointmentRepository) Update(ctx context.Context, id int32, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return fmt.Errorf("no updates provided")
	}

	setParts := make([]string, 0, len(updates))
	args := make([]interface{}, 0, len(updates)+1)
	argIndex := 1

	for field, value := range updates {
		setParts = append(setParts, fmt.Sprintf("%s = $%d", field, argIndex))
		args = append(args, value)
		argIndex++
	}

	setParts = append(setParts, fmt.Sprintf("updated_at = $%d", argIndex))
	args = append(args, time.Now())
	argIndex++

	args = append(args, id)

	query := fmt.Sprintf(`
		UPDATE appointments
		SET %s
		WHERE id = $%d`, strings.Join(setParts, ", "), argIndex)
	_, err := r.dbConn.Exec(ctx, query, args...)
	_, err = r.dbConn.Exec(ctx, query, args...)
	return err
}
