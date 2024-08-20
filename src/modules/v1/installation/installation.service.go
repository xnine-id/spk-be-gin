package installation

import (
	"encoding/csv"
	"fmt"
	"strconv"
	"time"

	"github.com/amuhajirs/gin-gorm/src/helpers/customerror"
	"github.com/amuhajirs/gin-gorm/src/helpers/pagination"
	"github.com/amuhajirs/gin-gorm/src/helpers/upload"
	"github.com/amuhajirs/gin-gorm/src/models"
	"github.com/xuri/excelize/v2"
)

type Service interface {
	find(qs *findInstallationQs) (*pagination.Pagination[models.Installation], error)
	findById(id string) (*models.Installation, error)
	create(body *installationBody) (*models.Installation, error)
	update(body *installationBody, id string) error
	delete(id string) error

	importData(body *importInstallationBody) error
	importExcel(body *importInstallationBody) error
	importCsv(body *importInstallationBody) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) find(qs *findInstallationQs) (*pagination.Pagination[models.Installation], error) {
	result := pagination.New(models.Installation{})

	if err := s.repo.find(result, qs); err != nil {
		return nil, customerror.GormError(err, "Pemasangan")
	}

	return result, nil
}

func (s *service) findById(id string) (*models.Installation, error) {
	var installation models.Installation

	if err := s.repo.findById(&installation, id); err != nil {
		return nil, customerror.GormError(err, "Pemasangan")
	}

	return &installation, nil
}

func (s *service) create(body *installationBody) (*models.Installation, error) {
	var installation models.Installation

	spkDate, err := time.Parse(time.DateOnly, body.SpkDate)
	if err != nil {
		return nil, customerror.New("Format waktu tidak valid", 400)
	}

	installationDate, err := time.Parse(time.DateOnly, body.InstallationDate)
	if err != nil {
		return nil, customerror.New("Format waktu tidak valid", 400)
	}

	installation.SpkNumber = body.SpkNumber
	installation.SpkDate = spkDate
	installation.StoreId = body.StoreId
	installation.InstallationDate = installationDate
	installation.SalesId = body.SalesId
	installation.Status = *body.Status

	if err := s.repo.save(&installation); err != nil {
		return nil, customerror.GormError(err, "Pemasangan")
	}

	return &installation, nil
}

func (s *service) update(body *installationBody, id string) error {
	var installation models.Installation

	if err := s.repo.findById(&installation, id); err != nil {
		return customerror.GormError(err, "Pemasangan")
	}

	spkDate, err := time.Parse(time.DateOnly, body.SpkDate)
	if err != nil {
		return customerror.New("Format waktu tidak valid", 400)
	}

	installationDate, err := time.Parse(time.DateOnly, body.InstallationDate)
	if err != nil {
		return customerror.New("Format waktu tidak valid", 400)
	}

	installation.SpkNumber = body.SpkNumber
	installation.SpkDate = spkDate
	installation.StoreId = body.StoreId
	installation.InstallationDate = installationDate
	installation.SalesId = body.SalesId
	installation.Status = *body.Status

	if err := s.repo.save(&installation); err != nil {
		return customerror.GormError(err, "Pemasangan")
	}

	return nil
}

func (s *service) delete(id string) error {
	if err := s.repo.delete(id); err != nil {
		return customerror.GormError(err, "Pemasangan")
	}
	return nil
}

func (s *service) importData(body *importInstallationBody) error {
	file := upload.ExtractFileName(body.File.Filename)

	if file.Ext == "csv" {
		return s.importCsv(body)
	}

	return s.importExcel(body)
}

func (s *service) importExcel(body *importInstallationBody) error {
	fileContent, err := body.File.Open()
	if err != nil {
		return err
	}

	defer fileContent.Close()

	xlsx, err := excelize.OpenReader(fileContent)
	if err != nil {
		return err
	}

	rows, err := xlsx.GetRows(xlsx.GetSheetName(0))

	if err != nil {
		return err
	}

	var columnMap map[string]int
	var installations []models.Installation

	for i, row := range rows {
		if i == 0 {
			fmt.Printf("%v", row)
			columnMap = mapHeadersToFields(row, body.ColumnMap)
			// Skip header row
			continue
		}

		var inst models.Installation

		storeId, err := strconv.ParseUint(row[columnMap["StoreId"]], 10, 64)
		if err != nil {
			return customerror.New(fmt.Sprintf("Pastikan nilai pada kolom %s berupa angka", rows[0][columnMap["StoreId"]]), 400)
		}

		// MM-DD-YY
		spkDate, err := time.Parse("01-02-06", row[columnMap["SpkDate"]])
		if err != nil {
			// YYYY-MM-DD
			spkDate, err = time.Parse(time.DateOnly, row[columnMap["SpkDate"]])

			if err != nil {
				// DD/MM/YYYY
				spkDate, err = time.Parse("02/01/2006", row[columnMap["SpkDate"]])

				if err != nil {
					return customerror.New(fmt.Sprintf("Pastikan nilai pada kolom %s berupa format tanggal yang valid (MM-DD-YY) atau (YYYY-MM-DD) atau (DD/MM/YYYY)", rows[0][columnMap["SpkDate"]]), 400)
				}
			}
		}

		inst.SpkNumber = row[columnMap["SpkNumber"]]
		inst.SpkDate = spkDate
        inst.StoreId = uint(storeId)

        installations = append(installations, inst)
	}

	if err := s.repo.importData(&installations); err != nil {
		return customerror.GormError(err, "Pemasangan")
	}

	return nil
}

func (s *service) importCsv(body *importInstallationBody) error {
	fileContent, err := body.File.Open()
	if err != nil {
		return err
	}

	defer fileContent.Close()

	reader := csv.NewReader(fileContent)

	// Read the header row
	headers, err := reader.Read()
	if err != nil {
		return err
	}

	// Map headers to struct fields based on ColumnMap
    columnMap := mapHeadersToFields(headers, body.ColumnMap)

	// Read the rest of the rows
	records, err := reader.ReadAll()

	if err != nil {
		return err
	}

	var installations []models.Installation

	for _, record := range records {
		var inst models.Installation

		storeId, err := strconv.ParseUint(record[columnMap["StoreId"]], 10, 64)
		if err != nil {
			return customerror.New(fmt.Sprintf("Pastikan nilai pada kolom %s berupa angka", headers[columnMap["StoreId"]]), 400)
		}

		spkDate, err := time.Parse(time.DateOnly, record[columnMap["SpkDate"]])
		if err != nil {
			return customerror.New(fmt.Sprintf("Pastikan nilai pada kolom %s berupa format tanggal yang valid (yyyy-mm-dd)", headers[columnMap["SpkDate"]]), 400)
		}

		inst.SpkNumber = record[columnMap["SpkNumber"]]
		inst.SpkDate = spkDate
        inst.StoreId = uint(storeId)

        installations = append(installations, inst)
	}

	if err := s.repo.importData(&installations); err != nil {
		return customerror.GormError(err, "Pemasangan")
	}

	return nil
}

// mapHeadersToFields creates a mapping between CSV headers and struct fields
func mapHeadersToFields(headers []string, columnMap ColumnMap) map[string]int {
    mapping := make(map[string]int)
    for i, header := range headers {
        switch header {
        case columnMap.SpkNumber:
            mapping["SpkNumber"] = i
        case columnMap.SpkDate:
            mapping["SpkDate"] = i
        case columnMap.StoreId:
            mapping["StoreId"] = i
        }
    }
    return mapping
}
