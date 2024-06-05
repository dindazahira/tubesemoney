package main

import (
	"fmt"
	"time"
)

type Akun struct {
	ID              int
	Nama            string
	Saldo           float64
	StatusDisetujui bool
}

type Transaksi struct {
	Jenis        string
	Jumlah       float64
	Tanggal      time.Time
	AkunPengirim int
	AkunPenerima int
}

const MAX_AKUN = 100
const MAX_TRANSAKSI = 1000

var DaftarAkun [MAX_AKUN]Akun
var DaftarTransaksi [MAX_TRANSAKSI]Transaksi
var jumlahAkun int
var jumlahTransaksi int

func RegistrasiAkun(nama string, saldo float64) {
	if jumlahAkun < MAX_AKUN {
		DaftarAkun[jumlahAkun] = Akun{
			ID:              jumlahAkun,
			Nama:            nama,
			Saldo:           saldo,
			StatusDisetujui: false,
		}
		jumlahAkun++
		SelectionSort(true)
		fmt.Println("Akun berhasil didaftarkan, menunggu persetujuan admin.")
		fmt.Printf("ID : %d\nNama : %s\nSaldo : %.2f\n", DaftarAkun[jumlahAkun-1].ID, DaftarAkun[jumlahAkun-1].Nama, DaftarAkun[jumlahAkun-1].Saldo)
	} else {
		fmt.Println("Batas maksimal akun tercapai.")
	}
}

func SetujuiAkun(nama string) {
	var idx int
	idx = SequentialSearch(nama)
	if idx < jumlahAkun && idx != -1 {
		DaftarAkun[idx].StatusDisetujui = true
		fmt.Println("Akun disetujui.")
	} else {
		fmt.Println("Akun tidak ditemukan.")
	}
}

func TolakAkun(nama string) {
	var idx int
	idx = BinarySearch(nama)
	if idx < jumlahAkun && idx != -1 {
		DaftarAkun[idx].StatusDisetujui = false
		fmt.Println("Akun ditolak.")
	} else {
		fmt.Println("Akun tidak ditemukan.")
	}
}

func CetakDaftarAkun() {
	SelectionSort(false)
	fmt.Println("Daftar Akun:")
	for i := 0; i < jumlahAkun; i++ {
		fmt.Printf("ID: %d, Nama: %s, Saldo: %.2f, Disetujui: %v\n", DaftarAkun[i].ID, DaftarAkun[i].Nama, DaftarAkun[i].Saldo, DaftarAkun[i].StatusDisetujui)
	}
}

func TransferUang(namaPengirim, namaPenerima string, jumlah float64) {
	var idxPengirim, idxPenerima int
	idxPengirim = SequentialSearch(namaPengirim)
	idxPenerima = BinarySearch(namaPenerima)
	if idxPengirim < jumlahAkun && idxPenerima < jumlahAkun {
		if DaftarAkun[idxPengirim].StatusDisetujui && DaftarAkun[idxPenerima].StatusDisetujui {
			if DaftarAkun[idxPengirim].Saldo >= jumlah {
				DaftarAkun[idxPengirim].Saldo -= jumlah
				DaftarAkun[idxPenerima].Saldo += jumlah
				DaftarTransaksi[jumlahTransaksi] = Transaksi{
					Jenis:        "Transfer",
					Jumlah:       jumlah,
					Tanggal:      time.Now(),
					AkunPengirim: idxPengirim,
					AkunPenerima: idxPenerima,
				}
				jumlahTransaksi++
				fmt.Println("Transfer berhasil.")
			} else {
				fmt.Println("Saldo tidak mencukupi.")
			}
		} else {
			fmt.Println("Salah satu akun belum disetujui.")
		}
	} else {
		fmt.Println("Akun tidak ditemukan.")
	}
}

func Pembayaran(namaPengirim string, jumlah float64, jenis string) {
	var idxPengirim int
	idxPengirim = SequentialSearch(namaPengirim)
	if idxPengirim < jumlahAkun {
		if DaftarAkun[idxPengirim].StatusDisetujui {
			if DaftarAkun[idxPengirim].Saldo >= jumlah {
				DaftarAkun[idxPengirim].Saldo -= jumlah
				DaftarTransaksi[jumlahTransaksi] = Transaksi{
					Jenis:        jenis,
					Jumlah:       jumlah,
					Tanggal:      time.Now(),
					AkunPengirim: idxPengirim,
					AkunPenerima: -1,
				}
				jumlahTransaksi++
				fmt.Println("Pembayaran berhasil.")
			} else {
				fmt.Println("Saldo tidak mencukupi.")
			}
		} else {
			fmt.Println("Akun belum disetujui.")
		}
	} else {
		fmt.Println("Akun tidak ditemukan.")
	}
}

func CetakRiwayatTransaksi(namaAkun string) {
	var idxAkun int
	idxAkun = BinarySearch(namaAkun)
	fmt.Printf("Riwayat Transaksi untuk Akun %s:\n", namaAkun)
	for i := 0; i < jumlahTransaksi; i++ {
		if DaftarTransaksi[i].AkunPengirim == idxAkun || DaftarTransaksi[i].AkunPenerima == idxAkun {
			fmt.Printf("Jenis: %s, Jumlah: %.2f, Tanggal: %s\n", DaftarTransaksi[i].Jenis, DaftarTransaksi[i].Jumlah, DaftarTransaksi[i].Tanggal.Format("2006-01-02 15:04:05"))
		}
	}
}

func SequentialSearch(nama string) int {
	for i := 0; i < jumlahAkun; i++ {
		if DaftarAkun[i].Nama == nama {
			return i
		}
	}
	return -1
}

func BinarySearch(nama string) int {
	InsertionSort(true)
	left, right := 0, jumlahAkun-1
	for left <= right {
		mid := left + (right-left)/2
		if DaftarAkun[mid].Nama == nama {
			return mid
		}
		if DaftarAkun[mid].Nama < nama {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func SelectionSort(ascending bool) {
	for i := 0; i < jumlahAkun-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlahAkun; j++ {
			if (ascending && DaftarAkun[j].Saldo < DaftarAkun[minIdx].Saldo) || (!ascending && DaftarAkun[j].Saldo > DaftarAkun[minIdx].Saldo) {
				minIdx = j
			}
		}
		DaftarAkun[i], DaftarAkun[minIdx] = DaftarAkun[minIdx], DaftarAkun[i]
	}
}

func InsertionSort(ascending bool) {
	for i := 1; i < jumlahAkun; i++ {
		key := DaftarAkun[i]
		j := i - 1
		for j >= 0 && ((ascending && DaftarAkun[j].Nama > key.Nama) || (!ascending && DaftarAkun[j].Nama < key.Nama)) {
			DaftarAkun[j+1] = DaftarAkun[j]
			j--
		}
		DaftarAkun[j+1] = key
	}
}

func main() {
	var pilihan int
	var user int
	for {
		fmt.Println("=== APLIKASI E-MONEY ===")
		fmt.Println()
		fmt.Println("Pilih User:")
		fmt.Println("1. Admin")
		fmt.Println("2. Pengguna")
		fmt.Println("3. Keluar")
		fmt.Scanln(&user)
		if user == 1 {
			fmt.Println("Menu:")
			fmt.Println("1. Persetujuan Akun")
			fmt.Println("2. Cetak Daftar Akun")
			fmt.Scanln(&pilihan)
			if pilihan == 1 {
				var nama string
				var setuju bool
				fmt.Print("Masukkan Nama akun: ")
				fmt.Scanln(&nama)
				fmt.Print("Setujui akun? (1: ya, 0: tidak): ")
				fmt.Scanln(&setuju)
				if setuju {
					SetujuiAkun(nama)
				} else {
					TolakAkun(nama)
				}
			} else if pilihan == 2 {
				CetakDaftarAkun()
			} else {
				fmt.Println("Pilihan tidak valid.")
			}
		} else if user == 2 {
			fmt.Println("Menu:")
			fmt.Println("1. Registrasi Akun")
			fmt.Println("2. Transfer Uang")
			fmt.Println("3. Pembayaran")
			fmt.Println("4. Cetak Riwayat Transaksi")
			fmt.Println("5. Cetak Daftar Akun")
			fmt.Print("Pilih menu: ")
			fmt.Scanln(&pilihan)

			if pilihan == 1 {
				var nama string
				var saldo float64
				fmt.Print("Masukkan Nama: ")
				fmt.Scanln(&nama)
				fmt.Print("Masukkan saldo awal: ")
				fmt.Scanln(&saldo)
				RegistrasiAkun(nama, saldo)
				InsertionSort(true)
			} else if pilihan == 2 {
				var namaPengirim, namaPenerima string
				var jumlah float64
				fmt.Print("Masukkan Nama pengirim: ")
				fmt.Scanln(&namaPengirim)
				fmt.Print("Masukkan Nama penerima: ")
				fmt.Scanln(&namaPenerima)
				fmt.Print("Masukkan jumlah: ")
				fmt.Scanln(&jumlah)
				TransferUang(namaPengirim, namaPenerima, jumlah)
			} else if pilihan == 3 {
				var namaPengirim string
				var jumlah float64
				var jenis string
				fmt.Print("Masukkan Nama pengirim: ")
				fmt.Scanln(&namaPengirim)
				fmt.Print("Masukkan jumlah: ")
				fmt.Scanln(&jumlah)
				fmt.Print("Masukkan jenis pembayaran: ")
				fmt.Scanln(&jenis)
				Pembayaran(namaPengirim, jumlah, jenis)
			} else if pilihan == 4 {
				var namaAkun string
				fmt.Print("Masukkan Nama akun: ")
				fmt.Scanln(&namaAkun)
				CetakRiwayatTransaksi(namaAkun)
			} else if pilihan == 5 {
				CetakDaftarAkun()
			} else {
				fmt.Println("Pilihan tidak valid.")
			}
		} else if user == 3 {
			fmt.Println("Keluar dari aplikasi.")
			return
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
