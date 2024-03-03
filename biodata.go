package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	biodata := []Biodata{
		{
			Nama:      "Adittiya Asril",
			Alamat:    "West Sumatra",
			Pekerjaan: "Software Developer",
			Alasan:    "Want to learn Golang to improve backend skills ",
		},
		{
			Nama:      "Sophie Dupont",
			Alamat:    "Paris",
			Pekerjaan: "Graphic Designer",
			Alasan:    "passionate about art and design",
		},
		{
			Nama:      "Marco Rossi",
			Alamat:    "Rome",
			Pekerjaan: "Architect",
			Alasan:    "interested in creating beautiful structures",
		},
		{
			Nama:      "Elena Petrova",
			Alamat:    "Moscow",
			Pekerjaan: "Marketing Manager",
			Alasan:    "interested in expanding marketing skills",
		},
		{
			Nama:      "Anna Kowalski",
			Alamat:    "Warsaw",
			Pekerjaan: "Project Manager",
			Alasan:    "interested in leading and organizing projects",
		},
	}

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run biodata.go [nomor_absen]")
		return
	}

	arg, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Nomor absen harus berupa angka.")
		return
	}

	if arg < 1 || arg > len(biodata) {
		fmt.Println("Nomor absen tidak valid. Harus dalam rentang 1 -", len(biodata))
		return
	}

	arg--

	fmt.Printf("Nama: %s\n", biodata[arg].Nama)
	fmt.Printf("Alamat: %s\n", biodata[arg].Alamat)
	fmt.Printf("Pekerjaan: %s\n", biodata[arg].Pekerjaan)
	fmt.Printf("Alasan: %s\n", biodata[arg].Alasan)
}
