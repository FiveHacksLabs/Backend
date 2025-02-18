package models

import "time"

// TransaksiDetail mewakili detail transaksi.
type TransaksiDetail struct {
	NoUrut           int       `json:"no_urut"`
	TanggalTransaksi time.Time `json:"tanggal_transaksi"`
	NomorJurnal      int       `json:"nomor_jurnal"`
	NomorRekening    string    `json:"nomor_rekening"`
	// 'D' untuk Debit, 'K' untuk Kredit
	JenisTransaksi string  `json:"jenis_transaksi"`
	Nominal        float64 `json:"nominal"`
	Keterangan1    string  `json:"keterangan1"`
	Keterangan2    string  `json:"keterangan2"`
	CreatedAt      time.Time `json:"created_at"`
	CreatedBy      string    `json:"created_by"`
	UpdatedAt      time.Time `json:"updated_at"`
	UpdatedBy      string    `json:"updated_by"`
}
