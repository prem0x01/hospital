package repository

import (
	"context"

	"github.com/prem0x01/hospital/internal/database/queries"
	"github.com/prem0x01/hospital/internal/domain"
)

type PatientRepository struct {
	q *queries.Queries
}

func NewPatientRepository(q *queries.Queries) *PatientRepository {
	return &PatientRepository{q: q}
}

func (r *PatientRepository) Create(ctx context.Context, p models.Patient) (*models.Patient, error) {
	arg := queries.CreatePatientParams{
		FirstName:             p.FirstName,
		LastName:              p.LastName,
		Email:                 p.Email,
		Phone:                 p.Phone,
		DateOfBirth:           p.DateOfBirth,
		Gender:                p.Gender,
		Address:               p.Address,
		MedicalHistory:        p.MedicalHistory,
		Allergies:             p.Allergies,
		EmergencyContactName:  p.EmergencyContactName,
		EmergencyContactPhone: p.EmergencyContactPhone,
		CreatedBy:             p.CreatedBy,
	}

	result, err := r.q.CreatePatient(ctx, arg)
	if err != nil {
		return nil, err
	}

	return toDomainPatient(result), nil
}

func (r *PatientRepository) GetByID(ctx context.Context, id int32) (*models.Patient, error) {
	p, err := r.q.GetPatientByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toDomainPatient(p), nil
}

func (r *PatientRepository) GetAll(ctx context.Context, limit, offset int32) ([]models.Patient, error) {
	patients, err := r.q.GetPatients(ctx, queries.GetPatientsParams{Limit: limit, Offset: offset})
	if err != nil {
		return nil, err
	}

	var result []domain.Patient
	for _, p := range patients {
		result = append(result, *toDomainPatient(p))
	}
	return result, nil
}

func (r *PatientRepository) Update(ctx context.Context, p domain.Patient) (*models.Patient, error) {
	arg := queries.UpdatePatientParams{
		ID:                    int(p.ID),
		FirstName:             p.FirstName,
		LastName:              p.LastName,
		Email:                 p.Email,
		Phone:                 p.Phone,
		DateOfBirth:           convertPgTypeDateToNullDate(p.DateOfBirth),
		Gender:                p.Gender,
		Address:               p.Address,
		MedicalHistory:        p.MedicalHistory,
		Allergies:             p.Allergies,
		EmergencyContactName:  p.EmergencyContactName,
		EmergencyContactPhone: p.EmergencyContactPhone,
	}

	updated, err := r.q.UpdatePatient(ctx, arg)
	if err != nil {
		return nil, err
	}
	return toDomainPatient(updated), nil
}

func (r *PatientRepository) Delete(ctx context.Context, id int32) error {
	return r.q.DeletePatient(ctx, id)
}

func (r *PatientRepository) Count(ctx context.Context) (int64, error) {
	return r.q.CountPatients(ctx)
}

func (r *PatientRepository) Search(ctx context.Context, keyword string, limit, offset int32) ([]models.Patient, error) {
	arg := queries.SearchPatientsParams{
		Column1: &keyword,
		Limit:   limit,
		Offset:  offset,
	}

	patients, err := r.q.SearchPatients(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result []domain.Patient
	for _, p := range patients {
		result = append(result, *toDomainPatient(p))
	}
	return result, nil
}

func toDomainPatient(p *queries.Patient) *domain.Patient {
	return &domain.Patient{
		ID:                    int(p.ID),
		FirstName:             p.FirstName,
		LastName:              p.LastName,
		Email:                 p.Email,
		Phone:                 p.Phone,
		DateOfBirth:           convertPgTypeDateToNullDate(p.DateOfBirth),
		Gender:                p.Gender,
		Address:               p.Address,
		MedicalHistory:        p.MedicalHistory,
		Allergies:             p.Allergies,
		EmergencyContactName:  p.EmergencyContactName,
		EmergencyContactPhone: p.EmergencyContactPhone,
		CreatedBy:             convertInt32PtrToIntPtr(p.CreatedBy),
		CreatedAt:             convertPgTypeTimestampToTime(p.CreatedAt),
		UpdatedAt:             convertPgTypeTimestampToTime(p.UpdatedAt),
	}
}

func convertInt32PtrToIntPtr(i *int32) *int {
	if i == nil {
		return nil
	}
	result := int(*i)
	return &result
}

func convertPgTypeTimestampToTime(ts pgtype.Timestamp) time.Time {
	if ts.Valid {
		return ts.Time
	}
	return time.Time{}
}

func convertPgTypeDateToNullDate(d pgtype.Date) *domain.NullDate {
	if d.Valid {
		return &domain.NullDate{
			Date:  d.Time,
			Valid: true,
		}
	}
	return &domain.NullDate{Valid: false}
}
