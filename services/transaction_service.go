package services

import (
	"database/sql"
	"time"
	"fmt"
	"users/models"
)

// PaymentRequest digunakan untuk memproses transaksi pembayaran.
type PaymentRequest struct {
	NoRekening string  `json:"nomor_rekening"`
	Amount     float64 `json:"amount"`
	Keterangan string  `json:"keterangan"`
}

// TransactionService interface untuk operasi transaksi.
type TransactionService interface {
	// InquiryTransaction mengembalikan header dan detail transaksi berdasarkan nomor jurnal.
	InquiryTransaction(nomorJurnal int) (models.TransaksiHeader, []models.TransaksiDetail, error)
	// PaymentTransaction membuat transaksi pembayaran baru.
	PaymentTransaction(req PaymentRequest) (models.TransaksiHeader, error)
	// ReportTransaction mengembalikan daftar transaksi antara rentang tanggal.
	ReportTransaction(startDate, endDate time.Time) ([]models.TransaksiHeader, error)
}

type transactionService struct {
	db *sql.DB
}

// NewTransactionService mengembalikan instance TransactionService.
func NewTransactionService(db *sql.DB) TransactionService {
	return &transactionService{db: db}
}

// InquiryTransaction mengambil header dan detail transaksi berdasarkan nomor jurnal.
func (s *transactionService) InquiryTransaction(nomorJurnal int) (models.TransaksiHeader, []models.TransaksiDetail, error) {
	var header models.TransaksiHeader
	err := s.db.QueryRow("SELECT * FROM transaksi_header WHERE nomor_jurnal = ?", nomorJurnal).Scan(&header.NomorJurnal, &header.TanggalTransaksi, &header.SaldoAwal, &header.TotalDebet, &header.TotalKredit, &header.SaldoAkhir, &header.KeteranganTransaksi, &header.CreatedAt, &header.CreatedBy, &header.UpdatedAt, &header.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			// Mengembalikan error custom dengan kode 12
			return header, nil, fmt.Errorf("ERR_12: Data transaksi tidak ditemukan")
		}
		return header, nil, err
	}
	var details []models.TransaksiDetail
	rows, err := s.db.Query("SELECT * FROM transaksi_detail WHERE nomor_jurnal = ?", nomorJurnal)
	if err != nil {
		return header, nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var detail models.TransaksiDetail
		err = rows.Scan(&detail.NomorJurnal, &detail.TanggalTransaksi, &detail.NomorRekening, &detail.JenisTransaksi, &detail.Nominal, &detail.Keterangan1, &detail.CreatedAt, &detail.CreatedBy, &detail.UpdatedAt, &detail.UpdatedBy)
		if err != nil {
			return header, nil, err
		}
		details = append(details, detail)
	}
	return header, details, nil
}

// PaymentTransaction membuat transaksi pembayaran dengan header dan detail secara atomik.
func (s *transactionService) PaymentTransaction(req PaymentRequest) (models.TransaksiHeader, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return models.TransaksiHeader{}, err
	}
	defer tx.Rollback()
	// Misal, ambil saldo awal dari rekening (di sini disederhanakan sebagai 0)
	header := models.TransaksiHeader{
		TanggalTransaksi:    time.Now(),
		SaldoAwal:           0,
		TotalDebet:          req.Amount,
		TotalKredit:         0,
		SaldoAkhir:          req.Amount, // Simplified, real logic akan melibatkan update saldo rekening
		KeteranganTransaksi: req.Keterangan,
		CreatedAt:           time.Now(),
		CreatedBy:           "system",
	}
	_, err = tx.Exec("INSERT INTO transaksi_header (tanggal_transaksi, saldo_awal, total_debet, total_kredit, saldo_akhir, keterangan_transaksi, created_at, created_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING nomor_jurnal", header.TanggalTransaksi, header.SaldoAwal, header.TotalDebet, header.TotalKredit, header.SaldoAkhir, header.KeteranganTransaksi, header.CreatedAt, header.CreatedBy)
	if err != nil {
		return header, err
	}
	// Buat satu detail transaksi sebagai contoh (Debit)
	detail := models.TransaksiDetail{
		TanggalTransaksi: time.Now(),
		NomorJurnal:      header.NomorJurnal,
		NomorRekening:    req.NoRekening,
		JenisTransaksi:   "D",
		Nominal:          req.Amount,
		Keterangan1:      req.Keterangan,
		CreatedAt:        time.Now(),
		CreatedBy:        "system",
	}
	_, err = tx.Exec("INSERT INTO transaksi_detail (nomor_jurnal, tanggal_transaksi, nomor_rekening, jenis_transaksi, nominal, keterangan1, created_at, created_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", detail.NomorJurnal, detail.TanggalTransaksi, detail.NomorRekening, detail.JenisTransaksi, detail.Nominal, detail.Keterangan1, detail.CreatedAt, detail.CreatedBy)
	if err != nil {
		return header, err
	}
	// Jika diperlukan, update saldo pada master_rekening (tidak ditampilkan di sini)
	if err := tx.Commit(); err != nil {
		return header, err
	}
	return header, nil
}

// ReportTransaction mengembalikan daftar header transaksi berdasarkan rentang tanggal.
func (s *transactionService) ReportTransaction(startDate, endDate time.Time) ([]models.TransaksiHeader, error) {
	var headers []models.TransaksiHeader
	rows, err := s.db.Query("SELECT * FROM transaksi_header WHERE tanggal_transaksi BETWEEN $1 AND $2 ORDER BY tanggal_transaksi ASC", startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var header models.TransaksiHeader
		err = rows.Scan(&header.NomorJurnal, &header.TanggalTransaksi, &header.SaldoAwal, &header.TotalDebet, &header.TotalKredit, &header.SaldoAkhir, &header.KeteranganTransaksi, &header.CreatedAt, &header.CreatedBy, &header.UpdatedAt, &header.UpdatedBy)
		if err != nil {
			return nil, err
		}
		headers = append(headers, header)
	}
	return headers, nil
}

