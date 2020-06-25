package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var dbMahasiswa = []mahasiswa{}
var paramPanjangNamaMin int = 3
var paramPanjangNamaMax int = 20
var paramUmur int = 17
var paramJurusan int = 20

type mahasiswa struct {
	nama, jurusan string
	umur          int
}

func main() {
	mainMenu()
	for {
		var selectMenu string

		fmt.Scanln(&selectMenu)
		switch selectMenu {
		case "1":
			addMahasiswa()
			break
		case "2":
			deleteMahasiswa()
			break
		case "3":
			viewMahasiswa()
			break
		case "4":
			os.Exit(0)
		default:
			clearScreen()
		}
	}

}

func mainMenu() {
	fmt.Println("--------------------------------------")
	fmt.Println("MAIN MENU")
	fmt.Println("--------------------------------------")
	fmt.Println("1. Add Mahasiswa")
	fmt.Println("2. Delete Mahasiswa")
	fmt.Println("3. View Mahasiswa")
	fmt.Println("4. Exit")
	fmt.Println("Pilihan menu")
}

func clearScreen() {
	osRunning := runtime.GOOS
	if osRunning == "linux" || osRunning == "darwin" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if osRunning == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	mainMenu()
}

func addMahasiswa() {
	var inputNama, inputJurusan, saveConfirm string
	var inputUmur int
	fmt.Println("--------------------------------------")
	fmt.Println("ADD MAHASISWA")
	fmt.Println("--------------------------------------")
	fmt.Print("Nama : ")
	fmt.Scanln(&inputNama)
	fmt.Print("Umur : ")
	fmt.Scanln(&inputUmur)
	fmt.Print("Jurusan : ")
	fmt.Scanln(&inputJurusan)

	isNameValid := validasiPanjangNama(inputNama, paramPanjangNamaMin, paramPanjangNamaMax)
	isUmurValid := validasiUmurMin(inputUmur, paramUmur)
	isJurusanValid := validasiPanjangJurusanMax(inputJurusan, paramJurusan)

	if isNameValid && isUmurValid && isJurusanValid {
		fmt.Print("Data Mahasiswa baru akan disimpan? (y/n) ")
		fmt.Scanln(&saveConfirm)
		if strings.ToLower(saveConfirm) == "y" {
			mahasiswaBaru := mahasiswa{
				nama:    inputNama,
				umur:    inputUmur,
				jurusan: inputJurusan,
			}
			fmt.Println(mahasiswaBaru)
			dbMahasiswa = append(dbMahasiswa, mahasiswaBaru)
			fmt.Println(dbMahasiswa)
			f, err := os.OpenFile("text.log",
				os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			if _, err := f.WriteString("\nNama :" + mahasiswaBaru.nama + "Umur :" + mahasiswaBaru.umur + "Jurusan : " + mahasiswaBaru.jurusan); err != nil {
				log.Println(err)
			}
			clearScreen()
		} else {
			clearScreen()
		}

	} else {
		fmt.Println("\n Data yang anda masukkan tidak sesuai")
		// clearScreen()
		addMahasiswa()
	}

}
func validasiPanjangNama(nama string, paramPanjangNamaMin, paramPanjangNamaMax int) (validasi bool) {
	if len(nama) >= paramPanjangNamaMin && len(nama) <= paramPanjangNamaMax {
		validasi = true
	} else {
		validasi = false
	}
	return
}
func validasiUmurMin(umur, paramUmur int) (validasi bool) {
	if umur >= 17 {
		validasi = true
	} else {
		validasi = false
	}
	return
}
func validasiPanjangJurusanMax(jurusan string, paramJurusan int) (validasi bool) {
	if len(jurusan) <= 10 {
		validasi = true
	} else {
		validasi = false
	}
	return
}

func viewMahasiswa() {
	var choose int
	fmt.Println("--------------------------------------")
	fmt.Println("VIEW MAHASISWA")
	fmt.Println("--------------------------------------")
	fmt.Println("1. Index Mahasiswa")
	fmt.Println("2. Tampilkan Semua")
	fmt.Println("Pilih menu : ")
	fmt.Scanln(&choose)

	if choose == 1 {
		var inputIndex int
		fmt.Print("Masukkan index mahasiswa : ")
		fmt.Scanln(&inputIndex)
		for index, isi := range dbMahasiswa {
			if index == inputIndex-1 {
				fmt.Println("--------------------------------------")
				fmt.Println("VIEW BY INDEX")
				fmt.Println("--------------------------------------")
				fmt.Println("Nama : ", isi.nama)
				fmt.Println("Umur : ", isi.umur)
				fmt.Println("Jurusan : ", isi.jurusan)
				fmt.Println("--------------------------------------")
			}
		}
		fmt.Println("Index kosong")

	} else if choose == 2 {
		fmt.Println("--------------------------------------")
		fmt.Println("VIEW ALL MAHASISWA")
		fmt.Println("--------------------------------------")
		for index, isi := range dbMahasiswa {
			fmt.Println(index + 1)
			fmt.Println("Nama : ", isi.nama)
			fmt.Println("Umur : ", isi.umur)
			fmt.Println("Jurusan : ", isi.jurusan)
			fmt.Println("--------------------------------------")

		}
	}

}

func deleteMahasiswa() {

	dbMahasiswa[len(dbMahasiswa)-1] = mahasiswa{"", "", 0}
	dbMahasiswa = dbMahasiswa[0 : len(dbMahasiswa)-1]
	clearScreen()

}
