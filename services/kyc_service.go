package services

import (
	"database/sql"
	"time"

	"users/models"
)

// KYCService interface untuk operasi data KYC.
type KYCService interface {
	// GetKYCByCIF mengambil data KYC berdasarkan CIF.
	GetKYCByCIF(cif string) (models.KYCData, error)
	// UpdateKYC memperbarui data KYC untuk CIF tertentu.
	UpdateKYC(cif string, data models.KYCData) error
	// CreateKYC membuat data KYC baru.
	CreateKYC(data *models.KYCData) error
}

type kycService struct {
	db *sql.DB
}

// NewKYCService mengembalikan instance KYCService.
func NewKYCService(db *sql.DB) KYCService {
	return &kycService{db: db}
}

func (s *kycService) GetKYCByCIF(cif string) (models.KYCData, error) {
	var kyc models.KYCData
	err := s.db.QueryRow("SELECT * FROM kyc_data WHERE cif = ?", cif).Scan(
		&kyc.CIF, &kyc.JenisIdentitas, &kyc.NoIdentitas, &kyc.NPWP, &kyc.Pekerjaan,
		&kyc.SumberPenghasilan, &kyc.PendapatanBulanan, &kyc.TujuanPembukaanRekening,
		&kyc.FotoKTP, &kyc.FotoSelfie, &kyc.StatusVerifikasi, &kyc.CreatedAt,
		&kyc.CreatedBy, &kyc.UpdatedAt, &kyc.UpdatedBy,
	)
	return kyc, err
}

func (s *kycService) UpdateKYC(cif string, data models.KYCData) error {
	_, err := s.db.Exec(`
		UPDATE kyc_data
		SET jenis_identitas = $2, no_identitas = $3, npwp = $4, pekerjaan = $5,
			sumber_penghasilan = $6, pendapatan_bulanan = $7, tujuan_pembukaan_rekening = $8,
			foto_ktp = $9, foto_selfie = $10, status_verifikasi = $11,
			updated_at = $12, updated_by = $13
		WHERE cif = $1
	`, cif, data.JenisIdentitas, data.NoIdentitas, data.NPWP, data.Pekerjaan,
		data.SumberPenghasilan, data.PendapatanBulanan, data.TujuanPembukaanRekening,
		data.FotoKTP, data.FotoSelfie, data.StatusVerifikasi, time.Now(), data.UpdatedBy,
	)
	return err
}

func (s *kycService) CreateKYC(data *models.KYCData) error {
	_, err := s.db.Exec(`
		INSERT INTO kyc_data (
			cif, jenis_identitas, no_identitas, npwp, pekerjaan, sumber_penghasilan,
			pendapatan_bulanan, tujuan_pembukaan_rekening, foto_ktp, foto_selfie,
			status_verifikasi, created_at, created_by
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
		)
	`, data.CIF, data.JenisIdentitas, data.NoIdentitas, data.NPWP, data.Pekerjaan,
		data.SumberPenghasilan, data.PendapatanBulanan, data.TujuanPembukaanRekening,
		data.FotoKTP, data.FotoSelfie, data.StatusVerifikasi, time.Now(), data.CreatedBy,
	)
	return err
}

