package main

import (
	"fmt"
	"strings"
)

const NMAX = 20

type reference struct {
	kode, penulis, buku, serial, penerbit string
	tahun                                 int
}

type tabReference struct {
	references [NMAX]reference
	n          int
}

func main() {
	var dataPaper tabReference
	var menuInput, n int
	var penerbit, judul string

	menuInput = 0
	for menuInput != 8 {
		fmt.Println("========================================================")
		fmt.Println("                      BIBLIOGRAPHY                      ")
		fmt.Println("========================================================")
		fmt.Println("1. Input data reference")
		fmt.Println("2. Tampilkan penerbit paling produktif")
		fmt.Println("3. Tampilkan semua buku oleh penerbit tertentu")
		fmt.Println("4. Tampilkan informasi buku dari judul")
		fmt.Println("5. Menampilkan referensi yang paling tua")
		fmt.Println("6. Menampilkan data mengurut membesar berdasarkan tahun")
		fmt.Println("7. Menampilkan data mengurut mengecil berdasarkan tahun")
		fmt.Println("8. Exit")
		fmt.Println("========================================================")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&menuInput)

		switch menuInput {
		case 1:
			// Transfer data ke array
			fmt.Print("Masukkan jumlah data: ")
			fmt.Scan(&n)
			inputData(&dataPaper, n)
			fmt.Print("\nData Reference: \n")
			printData(dataPaper)
		case 2:
			// Menampilkan penerbit paling produktif
			penerbitProduktif(dataPaper)
		case 3:
			// Menampilkan semua buku yang diterbitkan penerbit tertentu
			fmt.Print("Masukkan penerbit: ")
			fmt.Scan(&penerbit)
			stringSeperator(&penerbit)
			semuaBukuPenerbit(dataPaper, penerbit)
		case 4:
			// Menampilkan informasi lengkap buku dengan judul tertentu
			fmt.Print("Masukkan judul buku: ")
			fmt.Scan(&judul)
			stringSeperator(&judul)
			dataJudul(dataPaper, judul)
		case 5:
			// Menampilkan referensi yang paling tua (tahun terkecil)
			tahunTertua(dataPaper)
		case 6:
			// Selection Sort
			// Mengurutkan data berdasarkan tahun ascending
			fmt.Println("Data reference yang diurutkan berdasarkan tahun (ascending):")
			dataBerdasarkanTahunAscending(dataPaper)
		case 7:
			// Insertion Sort
			// Mengurutkan data berdasarkan tahun descending
			fmt.Println("Data reference yang diurutkan berdasarkan tahun (descending):")
			dataBerdasarkanTahunDescending(dataPaper)
		}
	}
}

func stringSeperator(j *string) {
	*j = strings.Join(strings.Split(*j, "_"), " ")
}

func inputData(A *tabReference, n int) {
	var i, index int
	for i = 0; i < n && A.n < NMAX; i++ {
		index = A.n
		fmt.Printf("Paper ke-%d: \n", index+1)
		fmt.Print("Masukkan Kode Paper: ")
		fmt.Scan(&A.references[index].kode)
		fmt.Print("Masukkan Penulis Paper: ")
		fmt.Scan(&A.references[index].penulis)
		fmt.Print("Masukkan Tahun Paper: ")
		fmt.Scan(&A.references[index].tahun)
		fmt.Print("Masukkan Judul Buku Paper: ")
		fmt.Scan(&A.references[index].buku)
		fmt.Print("Masukkan Judul Serial Paper: ")
		fmt.Scan(&A.references[index].serial)
		fmt.Print("Masukkan Penerbit Paper: ")
		fmt.Scan(&A.references[index].penerbit)

		stringSeperator(&A.references[index].penulis)
		stringSeperator(&A.references[index].buku)
		stringSeperator(&A.references[index].serial)
		stringSeperator(&A.references[index].penerbit)
		A.n++
	}
}

func printData(A tabReference) {
	var i int
	for i = 0; i < A.n; i++ {
		fmt.Printf("[%s] %s. %d. %s. %s. %s.\n", A.references[i].kode, A.references[i].penulis, A.references[i].tahun, A.references[i].buku, A.references[i].serial, A.references[i].penerbit)
	}
}

func penerbitProduktif(A tabReference) {
	var i, j, jumlah, maxJumlah int
	var penerbitProduktif string

	maxJumlah = 0
	for i = 0; i < A.n; i++ {
		jumlah = 0
		for j = 0; j < A.n; j++ {
			if A.references[j].penerbit == A.references[i].penerbit {
				jumlah++
			}
		}
		if jumlah > maxJumlah {
			maxJumlah = jumlah
			penerbitProduktif = A.references[i].penerbit
		}
	}

	fmt.Printf("Penerbit paling produktif adalah %s dengan %d buku.\n", penerbitProduktif, maxJumlah)
}

func semuaBukuPenerbit(A tabReference, penerbit string) {
	var i int
	var found bool
	found = false
	i = 0
	for !found && i < A.n {
		if A.references[i].penerbit == penerbit {
			found = true
		}
		i++
	}

	if found {
		fmt.Printf("Buku yang diterbitkan oleh %s:\n", penerbit)
		for i = 0; i < A.n; i++ {
			if A.references[i].penerbit == penerbit {
				fmt.Printf("[%s] %s. %d. %s. %s. %s.\n", A.references[i].kode, A.references[i].penulis, A.references[i].tahun, A.references[i].buku, A.references[i].serial, A.references[i].penerbit)
			}
		}
	} else {
		fmt.Println("Penerbit tidak ada dalam array.")
	}
}

func dataJudul(A tabReference, judul string) {
	var i int
	i = 0
	for i < A.n && A.references[i].buku != judul {
		i++
	}

	if A.references[i].buku == judul {
		fmt.Printf("[%s] %s. %d. %s. %s. %s.\n", A.references[i].kode, A.references[i].penulis, A.references[i].tahun, A.references[i].buku, A.references[i].serial, A.references[i].penerbit)
	} else {
		fmt.Println("Judul tidak ada dalam array.")
	}
}

func tahunTertua(A tabReference) {
	var i, min int
	min = 0
	for i = 1; i < A.n; i++ {
		if A.references[i].tahun < A.references[min].tahun {
			min = i
		}
	}
	fmt.Print("Referensi yang paling tua adalah: \n")
	fmt.Printf("[%s] %s. %d. %s. %s. %s.\n", A.references[min].kode, A.references[min].penulis, A.references[min].tahun, A.references[min].buku, A.references[min].serial, A.references[min].penerbit)
}

func dataBerdasarkanTahunAscending(A tabReference) {
	var pass, i, idx_min int
	var t reference
	pass = 1
	for pass <= A.n-1 {
		idx_min = pass - 1
		i = pass
		for i < A.n {
			if A.references[idx_min].tahun > A.references[i].tahun {
				idx_min = i
			}
			i = i + 1
		}
		t = A.references[idx_min]
		A.references[idx_min] = A.references[pass-1]
		A.references[pass-1] = t
		pass = pass + 1
	}
	printData(A)
}

func dataBerdasarkanTahunDescending(A tabReference) {
	var i, j int
	var temp reference
	i = 1
	for i <= A.n-1 {
		j = i
		temp = A.references[j]
		for j > 0 && temp.tahun > A.references[j-1].tahun {
			A.references[j] = A.references[j-1]
			j = j - 1
		}
		A.references[j] = temp
		i = i + 1
	}
	printData(A)
}
