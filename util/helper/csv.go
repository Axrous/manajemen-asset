package helper

import (
	"bytes"
	"encoding/csv"
	"final-project-enigma-clean/model"
	"fmt"
)

func ConvertToCSVForStaff(staffs []model.Staff) ([]byte, error) {
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

// for assets endpoint
func ConvertToCSVForAssets(assets []model.ManageAsset) ([]byte, error) {
	var csvData [][]string

	// Header nya
	header := []string{"Asset ID", "User ID", "User Name", "Staff Nik", "Staff Name", "Staff BirthDate", "Submission Date", "Return Date", "Detail Asset"}

	//SET HEADERRRR !!!!!
	csvData = append(csvData, header)

	// convert to csv
	for _, asset := range assets {
		birthDatestr := asset.Staff.Birth_date.Format("2006-01-02")
		submissionDatestr := asset.SubmissionDate.Format("2006-01-02")
		returnDatestr := asset.ReturnDate.Format("2006-01-02")

		var detailStr string
		for _, detail := range asset.Detail {
			detailStr += fmt.Sprintf("ID: %s, Asset ID: %s, Total Item: %d, Status: %s \n",
				detail.Id, detail.Asset.Id, detail.TotalItem, detail.Status)
		}

		// Menambahkan header dan isinya ke dalam row
		row := []string{
			asset.Id,
			asset.User.ID,
			asset.User.Name,
			asset.Staff.Nik_Staff,
			asset.Staff.Name,
			birthDatestr,
			submissionDatestr,
			returnDatestr,
			detailStr,
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

// convert to csv
//for _, asset := range assets {
//	birthDatestr := asset.Staff.Birth_date.Format("2006-01-02")
//	submissionDatestr := asset.SubmissionDate.Format("2006-01-02")
//	returnDatestr := asset.ReturnDate.Format("2006-01-02")
//
//	// cv assett detail to json
//	assetDetailJson, err := json.Marshal(asset.Detail)
//	if err != nil {
//		return nil, err
//	}
//	row := []string{
//		asset.Id,
//		asset.User.ID,
//		asset.User.Name,
//		asset.Staff.Nik_Staff,
//		asset.Staff.Name,
//		birthDatestr,
//		submissionDatestr,
//		returnDatestr,
//		string(assetDetailJson),
//	}
//	csvData = append(csvData, row)
//}
