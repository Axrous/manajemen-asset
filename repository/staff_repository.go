package repository

import (
	"database/sql"
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"
	"math"
)

type StaffRepository interface {
	Save(payload model.Staff) error
	FindById(nik_staff string) (model.Staff, error)
	FindByName(name string) ([]model.Staff, error)
	FindByAll() ([]model.Staff, error)
	Update(payload model.Staff) error
	Delete(nik_staff string) error
	Paging(payload dto.PageRequest) ([]model.Staff, dto.Paging, error)
}

type staffRepository struct {
	db *sql.DB
}

// Delete implements StaffRepository.
func (s *staffRepository) Delete(nik_staff string) error {
	_, err := s.db.Exec("DELETE FROM staff WHERE nik_staff=$1", nik_staff)
	if err != nil {
		return err
	}
	return nil
}

// FindByAll implements StaffRepository.
func (s *staffRepository) FindByAll() ([]model.Staff, error) {
	//nik_staff, name, phone_number, address, birth_date, img_url, divisi
	rows, err := s.db.Query("SELECT nik_staff, name, phone_number, address, birth_date, img_url, divisi FROM staff")
	if err != nil {
		return nil, err
	}
	var staffs []model.Staff
	for rows.Next() {
		var staff model.Staff
		rows.Scan(&staff.Nik_Staff, &staff.Name, &staff.Phone_number, &staff.Address, &staff.Birth_date, &staff.Img_url, &staff.Divisi)
		staffs = append(staffs, staff)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return staffs, nil
}

// FindById implements StaffRepository.
func (s *staffRepository) FindById(nik_staff string) (model.Staff, error) {
	row := s.db.QueryRow("SELECT nik_staff, name, phone_number, address, birth_date, img_url, divisi FROM staff WHERE nik_staff=$1", nik_staff)
	var staff model.Staff
	err := row.Scan(&staff.Nik_Staff, &staff.Name, &staff.Phone_number, &staff.Address, &staff.Birth_date, &staff.Img_url, &staff.Divisi)
	if err != nil {
		return model.Staff{}, err
	}
	return staff, nil
}

// FindByName implements StaffRepository.
func (s *staffRepository) FindByName(name string) ([]model.Staff, error) {
	rows, err := s.db.Query(`SELECT nik_staff, name, phone_number, address, birth_date, img_url, divisi FROM staff WHERE name ILIKE $1`, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	var staffs []model.Staff
	for rows.Next() {
		var staff model.Staff
		rows.Scan(&staff.Nik_Staff, &staff.Name, &staff.Phone_number, &staff.Address, &staff.Birth_date, &staff.Img_url, &staff.Divisi)
		staffs = append(staffs, staff)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return staffs, nil
}

// Paging implements StaffRepository.
func (s *staffRepository) Paging(payload dto.PageRequest) ([]model.Staff, dto.Paging, error) {
	if payload.Page <= 0 {
		payload.Page = 1
	}
	q := `SELECT nik_staff, name, phone_number, address, birth_date, img_url, divisi FROM staff LIMIT $2 OFFSET $1`
	rows, err := s.db.Query(q, (payload.Page-1)*payload.Size, payload.Size)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	var staffs []model.Staff
	for rows.Next() {
		var staff model.Staff
		err := rows.Scan(&staff.Nik_Staff, &staff.Name, &staff.Phone_number, &staff.Address, &staff.Birth_date, &staff.Img_url, &staff.Divisi)
		if err != nil {
			return nil, dto.Paging{}, err
		}
		staffs = append(staffs, staff)
	}
	var count int
	row := s.db.QueryRow("SELECT COUNT(nik_staff) FROM staff")
	if err := row.Scan(&count); err != nil {
		return nil, dto.Paging{}, err
	}

	paging := dto.Paging{
		Page:       payload.Page,
		Size:       payload.Size,
		TotalRows:  count,
		TotalPages: int(math.Ceil(float64(count) / float64(payload.Size))), // (totalrow / size)
	}

	return staffs, paging, nil

}

// Save implements StaffRepository.
func (s *staffRepository) Save(payload model.Staff) error {
	_, err := s.db.Exec("INSERT INTO staff (nik_staff, name, phone_number, address, birth_date, img_url, divisi) VALUES ($1, $2, $3, $4, $5, $6, $7)", payload.Nik_Staff, payload.Name, payload.Phone_number, payload.Address, payload.Birth_date, payload.Img_url, payload.Divisi)
	if err != nil {
		return err
	}
	return nil
}

// Update implements StaffRepository.
func (s *staffRepository) Update(payload model.Staff) error {
	_, err := s.db.Exec("UPDATE staff SET nik_staff=$1, name=$2, phone_number=$3, address=$4, birth_date=$5, img_url=$6, divisi=$7 WHERE nik_staff=$1", payload.Nik_Staff, payload.Name, payload.Phone_number, payload.Address, payload.Birth_date, payload.Img_url, payload.Divisi)
	if err != nil {
		return err
	}
	return nil
}

func NewStaffRepository(db *sql.DB) StaffRepository {
	return &staffRepository{
		db: db,
	}
}
