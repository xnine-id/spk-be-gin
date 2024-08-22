package installation

import (
	"encoding/csv"
	"fmt"
	"strconv"
	"time"

	"github.com/amuhajirs/gin-gorm/src/database"
	"github.com/amuhajirs/gin-gorm/src/helpers"
	"github.com/amuhajirs/gin-gorm/src/helpers/customerror"
	"github.com/amuhajirs/gin-gorm/src/helpers/pagination"
	"github.com/amuhajirs/gin-gorm/src/helpers/upload"
	"github.com/amuhajirs/gin-gorm/src/helpers/validation"
	"github.com/amuhajirs/gin-gorm/src/models"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type Service interface {
	find(qs *findInstallationQs) (*pagination.Pagination[models.Installation], error)
	findById(id string) (*models.Installation, error)
	create(body *installationBody) (*models.Installation, error)
	update(body *updateInstallationBody, id string) error
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
	// Validate date
	spkDate, err := time.Parse(time.DateOnly, body.SpkDate)
	if err != nil {
		return nil, customerror.New("Pastikan nilai pada kolom Tanggal SPK merupakan format tanggal yang valid (yyyy-mm-dd)", 400)
	}

	installationDate, err := time.Parse(time.DateOnly, body.InstallationDate)
	if err != nil {
		return nil, customerror.New("Pastikan nilai pada kolom Tanggal Pemasangan merupakan format tanggal yang valid (yyyy-mm-dd)", 400)
	}

	// Validate images
	for _, image := range body.Images {
		fileName := upload.ExtractFileName(image.Filename)

		if found := helpers.Find(&validation.ImageExtensions, func(t *string) bool {
			return fileName.Ext == *t
		}); found == nil {
			return nil, customerror.New("File harus berupa gambar", 400)
		}
	}

	var installation models.Installation

	if err := database.DB.Transaction(func(tx *gorm.DB) error {
		installation.SpkNumber = body.SpkNumber
		installation.SpkDate = &spkDate
		installation.StoreId = body.StoreId
		installation.InstallationDate = &installationDate
		installation.SalesId = body.SalesId
		installation.Status = *body.Status

		if err := s.repo.save(&installation); err != nil {
			return customerror.GormError(err, "Pemasangan")
		}

		var images []models.InstallationImage
		for i, image := range body.Images {
			file, err := upload.New(&upload.Option{
				Folder:      "installation",
				File:        image,
				NewFilename: strconv.FormatUint(uint64(*installation.Id), 10) + "-" + strconv.Itoa(i+1),
			})

			if err != nil {
				return customerror.New("Gagal saat mengupload file", 500)
			}

			images = append(images, models.InstallationImage{
				InstallationId: *installation.Id,
				Image:          file.Url,
			})
		}

		if err := s.repo.saveImages(&images); err != nil {
			return customerror.GormError(err, "Gambar Pemasangan")
		}

		installation.Images = &images
		return nil
	}); err != nil {
		return nil, err
	}

	return &installation, nil
}

func (s *service) update(body *updateInstallationBody, id string) error {
	var installation models.Installation

	if err := s.repo.findById(&installation, id); err != nil {
		return customerror.GormError(err, "Pemasangan")
	}

	// Validation date
	spkDate, err := time.Parse(time.DateOnly, body.SpkDate)
	if err != nil {
		return customerror.New("Pastikan nilai pada kolom Tanggal SPK merupakan format tanggal yang valid (yyyy-mm-dd)", 400)
	}

	installationDate, err := time.Parse(time.DateOnly, body.InstallationDate)
	if err != nil {
		return customerror.New("Pastikan nilai pada kolom Tanggal Pemasangan merupakan format tanggal yang valid (yyyy-mm-dd)", 400)
	}

	// Validate images
	for _, image := range body.Images {
		fileName := upload.ExtractFileName(image.Filename)

		if found := helpers.Find(&validation.ImageExtensions, func(t *string) bool {
			return fileName.Ext == *t
		}); found == nil {
			return customerror.New("File harus berupa gambar", 400)
		}
	}

	if err := database.DB.Transaction(func(tx *gorm.DB) error {
		installation.SpkNumber = body.SpkNumber
		installation.SpkDate = &spkDate
		installation.StoreId = body.StoreId
		installation.InstallationDate = &installationDate
		installation.SalesId = body.SalesId
		installation.Status = *body.Status

		if err := s.repo.save(&installation); err != nil {
			return customerror.GormError(err, "Pemasangan")
		}

		if len(body.Images) != 4 && len(*installation.Images) != 4 {
			return customerror.New("Foto-foto harus diisi dan berjumlah 4", 400)
		}

		for i, image := range body.Images {
			file, err := upload.New(&upload.Option{
				Folder:      "installation",
				File:        image,
				NewFilename: strconv.FormatUint(uint64(*installation.Id), 10) + "-" + strconv.Itoa(i+1),
			})

			if err != nil {
				return customerror.New("Gagal saat mengupload file", 500)
			}

			if installation.Images != nil && len(*installation.Images) == 4 {
				for idx, currImage := range *installation.Images {
					if *currImage.Id == body.ImageIds[i] {
						(*installation.Images)[idx].Image = file.Url
					}
				}
			} else {
				if len(body.Images) != 4 {
					return customerror.New("Jumlah file yang diunggah harus 4", 400)
				}

				*installation.Images = append(*installation.Images, models.InstallationImage{
					InstallationId: *installation.Id,
					Image: file.Url,
				})
			}
		}

		if len(body.Images) > 0 {
			if err := s.repo.saveImages(installation.Images); err != nil {
				return customerror.GormError(err, "Gambar Pemasangan")
			}
		}

		return nil
	}); err != nil {
		return err
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
		inst.SpkDate = &spkDate
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
		inst.SpkDate = &spkDate
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
