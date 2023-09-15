package helper

import (
	"bytes"
	"encoding/csv"
	"final-project-enigma-clean/model"
)

func ConvertToCSV(staffs []model.Staff) ([]byte, error) {
	var csvData [][]string

	// Header nya
	header := []string{"Nik_Staff", "Name", "Phone Number", "Address", "Birth Date", "Img_URL", "Divisi"}

	//SET HEADERRRR !!!!!
	csvData = append(csvData, header)

	// convert to csv
	for _, staff := range staffs {
		birthDatestr := staff.Birth_date.Format("2006-01-02")
		row := []string{
			staff.Nik_Staff,
			staff.Name,
			staff.Phone_number,
			staff.Address,
			birthDatestr,
			staff.Img_url,
			staff.Divisi,
		}
		csvData = append(csvData, row)
	}

	var csvBuffer bytes.Buffer
	writer := csv.NewWriter(&csvBuffer)

	// write data to csv
	err := writer.WriteAll(csvData)
	if err != nil {
		return nil, err
	}
	return csvBuffer.Bytes(), nil
}
