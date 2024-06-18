package repository

import (
	"context"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProfileRepo struct {
	db *pgxpool.Pool
}

func NewProfileRepo(db *pgxpool.Pool) *ProfileRepo {
	return &ProfileRepo{db: db}
}

func (r *ProfileRepo) Create(ctx context.Context, profile domain.Profile) error {
	query := `INSERT INTO profiles (specialty_code, education_type_id, profile_name)
              VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, profile.SpecialtyCode, profile.EducationTypeID, profile.ProfileName)

	return err
}

func (r *ProfileRepo) Put(ctx context.Context, profile domain.Profile) error {
	query := `UPDATE profiles SET specialty_code=$1, education_type_id=$2, profile_name=$3 WHERE profile_id=$4`
	_, err := r.db.Exec(ctx, query, profile.SpecialtyCode, profile.EducationTypeID, profile.ProfileName, profile.ProfileID)

	return err
}

func (r *ProfileRepo) Patch(ctx context.Context, profileID int64, updates map[string]interface{}) error {
	query := `UPDATE profiles SET`
	args := make([]interface{}, 0, len(updates)+1)
	argsCounter := 1

	for col, val := range updates {
		query += " " + col + " = $" + strconv.Itoa(argsCounter) + ","
		args = append(args, val)
		argsCounter++
	}

	query = query[:len(query)-1]
	query += " WHERE profile_id = $" + strconv.Itoa(argsCounter)
	args = append(args, profileID)
	_, err := r.db.Exec(ctx, query, args...)

	return err
}

func (r *ProfileRepo) Delete(ctx context.Context, profileID int64) error {
	query := `DELETE FROM profiles WHERE profile_id = $1`
	_, err := r.db.Exec(ctx, query, profileID)

	return err
}

func (r *ProfileRepo) GetByID(ctx context.Context, profileID int64) (domain.ProfileInfo, error) {
	query := `SELECT 
			p.profile_id, p.specialty_code, p.education_type_id, p.profile_name,
			et.education_type_name
		FROM 
			profiles p
		LEFT JOIN 
			educationTypes et ON p.education_type_id = et.education_type_id
		WHERE p.profile_id = $1`

	profileInfo := domain.ProfileInfo{}
	err := r.db.QueryRow(ctx, query, profileID).Scan(
		&profileInfo.Profile.ProfileID,
		&profileInfo.Profile.SpecialtyCode,
		&profileInfo.Profile.EducationTypeID,
		&profileInfo.Profile.ProfileName,
		&profileInfo.ProfileSub.EducationTypeName,
	)

	return profileInfo, err
}

func (r *ProfileRepo) GetByName(ctx context.Context, profileName string) (domain.ProfileInfo, error) {
	query := `SELECT 
			p.profile_id, p.specialty_code, p.education_type_id, p.profile_name,
			et.education_type_name
		FROM 
			profiles p
		LEFT JOIN 
			educationTypes et ON p.education_type_id = et.education_type_id
		WHERE p.profile_name = $1`

	profileInfo := domain.ProfileInfo{}
	err := r.db.QueryRow(ctx, query, profileName).Scan(
		&profileInfo.Profile.ProfileID,
		&profileInfo.Profile.SpecialtyCode,
		&profileInfo.Profile.EducationTypeID,
		&profileInfo.Profile.ProfileName,
		&profileInfo.ProfileSub.EducationTypeName,
	)

	return profileInfo, err
}

func (r *ProfileRepo) GetAll(ctx context.Context) ([]domain.ProfileInfo, error) {
	query := `SELECT 
			p.profile_id, p.specialty_code, p.education_type_id, p.profile_name,
			et.education_type_name
		FROM 
			profiles p
		LEFT JOIN 
			educationTypes et ON p.education_type_id = et.education_type_id`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	/*profileQuantity, err := r.getCountProfiles(ctx)
	if err != nil {
		return nil, err
	}*/
	const defaultCapacity = 100
	profiles := make([]domain.ProfileInfo, 0, defaultCapacity)
	for rows.Next() {
		var profileInfo domain.ProfileInfo
		err := rows.Scan(
			&profileInfo.Profile.ProfileID,
			&profileInfo.Profile.SpecialtyCode,
			&profileInfo.Profile.EducationTypeID,
			&profileInfo.Profile.ProfileName,
			&profileInfo.ProfileSub.EducationTypeName,
		)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, profileInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return profiles, nil
}

func (r *ProfileRepo) GetAllBySpecialtyCode(ctx context.Context, specialtyCode string) ([]domain.ProfileInfo, error) {
	query := `SELECT 
			p.profile_id, p.specialty_code, p.education_type_id, p.profile_name,
			et.education_type_name
		FROM 
			profiles p
		LEFT JOIN 
			educationTypes et ON p.education_type_id = et.education_type_id
		WHERE p.specialty_code = $1`

	rows, err := r.db.Query(ctx, query, specialtyCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	/*profileQuantity, err := r.getCountProfilesBySpecialtyCode(ctx, specialtyCode)
	if err != nil {
		return nil, err
	}*/
	const defaultCapacity = 100
	profiles := make([]domain.ProfileInfo, 0, defaultCapacity)
	for rows.Next() {
		var profileInfo domain.ProfileInfo
		err := rows.Scan(
			&profileInfo.Profile.ProfileID,
			&profileInfo.Profile.SpecialtyCode,
			&profileInfo.Profile.EducationTypeID,
			&profileInfo.Profile.ProfileName,
			&profileInfo.ProfileSub.EducationTypeName,
		)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, profileInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return profiles, nil
}

func (r *ProfileRepo) GetByEducationTypeID(ctx context.Context, educationTypeID int64) ([]domain.ProfileInfo, error) {
	query := `SELECT 
			p.profile_id, p.specialty_code, p.education_type_id, p.profile_name,
			et.education_type_name
		FROM 
			profiles p
		LEFT JOIN 
			educationTypes et ON p.education_type_id = et.education_type_id
		WHERE p.education_type_id = $1`

	rows, err := r.db.Query(ctx, query, educationTypeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	/*profileQuantity, err := r.getCountProfilesByEducationTypeID(ctx, educationTypeID)
	if err != nil {
		return nil, err
	}*/
	const defaultCapacity = 100
	profiles := make([]domain.ProfileInfo, 0, defaultCapacity)
	for rows.Next() {
		var profileInfo domain.ProfileInfo
		err := rows.Scan(
			&profileInfo.Profile.ProfileID,
			&profileInfo.Profile.SpecialtyCode,
			&profileInfo.Profile.EducationTypeID,
			&profileInfo.Profile.ProfileName,
			&profileInfo.ProfileSub.EducationTypeName,
		)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, profileInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return profiles, nil
}

func (r *ProfileRepo) getCountProfiles(ctx context.Context) (int64, error) {
	query := `SELECT COUNT(*) FROM profiles;`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var count int64
	if rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	return count, nil
}

func (r *ProfileRepo) getCountProfilesBySpecialtyCode(ctx context.Context, specialtyCode string) (int64, error) {
	query := `SELECT COUNT(*) FROM profiles WHERE specialty_code = $1;`
	rows, err := r.db.Query(ctx, query, specialtyCode)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var count int64
	if rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	return count, nil
}

func (r *ProfileRepo) getCountProfilesByEducationTypeID(ctx context.Context, educationTypeID int64) (int64, error) {
	query := `SELECT COUNT(*) FROM profiles WHERE education_type_id = $1;`
	rows, err := r.db.Query(ctx, query, educationTypeID)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var count int64
	if rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	return count, nil
}
