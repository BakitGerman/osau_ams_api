package repository

import (
	"context"
	"strconv"

	"github.com/BeRebornBng/OsauAmsApi/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GroupRepo struct {
	db *pgxpool.Pool
}

func NewGroupRepo(db *pgxpool.Pool) *GroupRepo {
	return &GroupRepo{db: db}
}

func (r *GroupRepo) Create(ctx context.Context, group domain.Group) error {
	query := `INSERT INTO groups (group_id, profile_id)
              VALUES ($1, $2)`
	_, err := r.db.Exec(ctx, query, group.GroupID, group.ProfileID)

	return err
}

func (r *GroupRepo) Put(ctx context.Context, group domain.Group) error {
	query := `UPDATE groups SET profile_id=$1 WHERE group_id=$2`
	_, err := r.db.Exec(ctx, query, group.ProfileID, group.GroupID)

	return err
}

func (r *GroupRepo) Patch(ctx context.Context, groupID string, updates map[string]interface{}) error {
	query := `UPDATE groups SET`
	args := make([]interface{}, 0, len(updates)+1)
	argsCounter := 1

	for col, val := range updates {
		query += " " + col + " = $" + strconv.Itoa(argsCounter) + ","
		args = append(args, val)
		argsCounter++
	}

	query = query[:len(query)-1]
	query += " WHERE group_id = $" + strconv.Itoa(argsCounter)
	args = append(args, groupID)
	_, err := r.db.Exec(ctx, query, args...)

	return err
}

func (r *GroupRepo) Delete(ctx context.Context, groupID string) error {
	query := `DELETE FROM groups WHERE group_id = $1`
	_, err := r.db.Exec(ctx, query, groupID)

	return err
}

func (r *GroupRepo) GetByID(ctx context.Context, groupID string) (domain.GroupInfo, error) {
	query := `SELECT 
			g.group_id, g.profile_id,
			p.profile_name
		FROM 
			groups g
		LEFT JOIN 
			profiles p ON g.profile_id = p.profile_id
		WHERE g.group_id = $1`

	groupInfo := domain.GroupInfo{}
	err := r.db.QueryRow(ctx, query, groupID).Scan(
		&groupInfo.Group.GroupID,
		&groupInfo.Group.ProfileID,
		&groupInfo.GroupSub.ProfileName,
	)

	return groupInfo, err
}

func (r *GroupRepo) GetByName(ctx context.Context, profileName string) (domain.GroupInfo, error) {
	query := `SELECT 
			g.group_id, g.profile_id,
			p.profile_name
		FROM 
			groups g
		LEFT JOIN 
			profiles p ON g.profile_id = p.profile_id
		WHERE p.profile_name = $1`

	groupInfo := domain.GroupInfo{}
	err := r.db.QueryRow(ctx, query, profileName).Scan(
		&groupInfo.Group.GroupID,
		&groupInfo.Group.ProfileID,
		&groupInfo.GroupSub.ProfileName,
	)

	return groupInfo, err
}

func (r *GroupRepo) GetAll(ctx context.Context) ([]domain.GroupInfo, error) {
	query := `SELECT 
			g.group_id, g.profile_id,
			p.profile_name
		FROM 
			groups g
		LEFT JOIN 
			profiles p ON g.profile_id = p.profile_id`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	/*groupQuantity, err := r.getCountGroups(ctx)
	if err != nil {
		return nil, err
	}*/
	const defaultCapacity = 100
	groups := make([]domain.GroupInfo, 0, defaultCapacity)
	for rows.Next() {
		var groupInfo domain.GroupInfo
		err := rows.Scan(
			&groupInfo.Group.GroupID,
			&groupInfo.Group.ProfileID,
			&groupInfo.GroupSub.ProfileName,
		)
		if err != nil {
			return nil, err
		}
		groups = append(groups, groupInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}

func (r *GroupRepo) GetAllByProfileID(ctx context.Context, profileID int64) ([]domain.GroupInfo, error) {
	query := `SELECT 
			g.group_id, g.profile_id,
			p.profile_name
		FROM 
			groups g
		LEFT JOIN 
			profiles p ON g.profile_id = p.profile_id
		WHERE g.profile_id = $1`

	rows, err := r.db.Query(ctx, query, profileID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	/*groupQuantity, err := r.getCountGroupsByProfileID(ctx, profileID)
	if err != nil {
		return nil, err
	}*/
	const defaultCapacity = 100
	groups := make([]domain.GroupInfo, 0, defaultCapacity)
	for rows.Next() {
		var groupInfo domain.GroupInfo
		err := rows.Scan(
			&groupInfo.Group.GroupID,
			&groupInfo.Group.ProfileID,
			&groupInfo.GroupSub.ProfileName,
		)
		if err != nil {
			return nil, err
		}
		groups = append(groups, groupInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}

func (r *GroupRepo) getCountGroups(ctx context.Context) (int64, error) {
	query := `SELECT COUNT(*) FROM groups;`
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

func (r *GroupRepo) getCountGroupsByProfileID(ctx context.Context, profileID int64) (int64, error) {
	query := `SELECT COUNT(*) FROM groups WHERE profile_id = $1;`
	rows, err := r.db.Query(ctx, query, profileID)
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
