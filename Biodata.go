package main

import (
	"fmt"
	"os"
)

type Friend struct {
	Name         string
	Address      string
	Job          string
	GolangReason string
}

func main() {
	// Membuat data teman
	friends := []Friend{
		Friend{Name: "Thomas", Address: "Jakarta", Job: "Developer", GolangReason: "Ingin mempelajari bahasa pemrograman baru yang powerful dan scalable."},
		Friend{Name: "Dito", Address: "Jakarta", Job: "Software Engineer", GolangReason: "Switch Carrerr."},
		Friend{Name: "Sarah", Address: "Surabaya", Job: "Data Scientist", GolangReason: "Ingin belajar pemrograman untuk analisis data dan machine learning."},
	}

	// Ambil argumen nama dari CLI
	name := os.Args[1]

	// Cari data teman berdasarkan nama
	var friend Friend
	for _, f := range friends {
		if f.Name == name {
			friend = f
			break
		}
	}

	// Tampilkan data teman yang sesuai dengan nama
	if friend.Name != "" {
		fmt.Printf("Data teman dengan nama '%s':\n", friend.Name)
		fmt.Printf("- Alamat: %s\n", friend.Address)
		fmt.Printf("- Pekerjaan: %s\n", friend.Job)
		fmt.Printf("- Alasan memilih kelas Golang: %s\n", friend.GolangReason)
	} else {
		fmt.Printf("Tidak ada data teman dengan nama '%s'\n", name)
	}
}
