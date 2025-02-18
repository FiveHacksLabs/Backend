package models

import "time"

// KYCData mewakili data KYC nasabah.
type KYCData struct {
	// Menggunakan CIF sebagai primary key (disesuaikan dengan tabel master nasabah)
	CIF                     string  `json:"cif"`
	JenisIdentitas          string  `json:"jenis_identitas"` // misalnya: KTP, Passport
	NoIdentitas             string  `json:"no_identitas"`
	NPWP                    string  `json:"npwp"`
	Pekerjaan               string  `json:"pekerjaan"`
	SumberPenghasilan       string  `json:"sumber_penghasilan"`
	PendapatanBulanan       float64 `json:"pendapatan_bulanan"`
	TujuanPembukaanRekening string  `json:"tujuan_pembukaan_rekening"`
	FotoKTP                 string  `json:"foto_ktp"`
	FotoSelfie              string  `json:"foto_selfie"`
	StatusVerifikasi        string  `json:"status_verifikasi"`
	CreatedAt               time.Time `json:"created_at"`
	CreatedBy               string    `json:"created_by"`
	UpdatedAt               time.Time `json:"updated_at"`
	UpdatedBy               string    `json:"updated_by"`
}
