package models

import "time"

// TransaksiHeader mewakili header transaksi.
type TransaksiHeader struct {
	NomorJurnal         int       `json:"nomor_jurnal"`
	TanggalTransaksi    time.Time `json:"tanggal_transaksi"`
	SaldoAwal           float64   `json:"saldo_awal"`
	TotalDebet          float64   `json:"total_debet"`
	TotalKredit         float64   `json:"total_kredit"`
	SaldoAkhir          float64   `json:"saldo_akhir"`
	KeteranganTransaksi string    `json:"keterangan_transaksi"`
	CreatedAt           time.Time `json:"created_at"`
	CreatedBy           string    `json:"created_by"`
	UpdatedAt           time.Time `json:"updated_at"`
	UpdatedBy           string    `json:"updated_by"`
}
