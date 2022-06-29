package repositories

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// gorm.Model definition
type SystemsMember struct {
	ID             uuid.UUID `gorm:"type:char(36);primary_key"`
	FirebaseAuthID string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (member *SystemsMember) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	member.ID = id
	return err
}

type MemberRepo struct {
	db *gorm.DB
}

func NewMemberRepo(db *gorm.DB) *MemberRepo {
	return &MemberRepo{
		db: db,
	}
}

func (r *MemberRepo) GetAll(page, pagesize int, order string) (results []*SystemsMember, totalRows int64, err error) {

	resultOrm := r.db.Model(&SystemsMember{})

	resultOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		resultOrm = resultOrm.Offset(offset).Limit(pagesize)
	} else {
		resultOrm = resultOrm.Limit(pagesize)
	}

	if order != "" {
		resultOrm = resultOrm.Order(order)
	}

	if err = resultOrm.Find(&results).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return results, totalRows, nil
}

func (r *MemberRepo) Get(argID uuid.UUID) (record *SystemsMember, err error) {
	record = &SystemsMember{}
	if err = r.db.Debug().Preload("BTags").Where(&SystemsMember{ID: argID}).First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

func (r *MemberRepo) Create(record *SystemsMember) (result *SystemsMember, RowsAffected int64, err error) {
	db := r.db.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

func (r *MemberRepo) Update(argID uuid.UUID, updated *SystemsMember) (result *SystemsMember, RowsAffected int64, err error) {
	record := &SystemsMember{}
	db := r.db.Where(&SystemsMember{ID: argID}).First(record)
	db.Delete(record)

	result = &SystemsMember{}
	db = r.db.First(result, argID)
	if err = db.Error; err != nil {
		return nil, -1, ErrNotFound
	}

	db.Where(&SystemsMember{ID: argID}).Updates(updated)
	if err = db.Error; err != nil {
		return nil, -1, ErrUpdateFailed
	}

	return result, db.RowsAffected, nil
}

func (r *MemberRepo) Delete(argID uint32) (rowsAffected int64, err error) {
	record := &SystemsMember{}
	db := r.db.First(record, argID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
