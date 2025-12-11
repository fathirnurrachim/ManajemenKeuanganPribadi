package main

import "fmt"

func main() {
	// Deklarasi variabel
	var income int
	var expense int
	var totalExpense int
	var remainingIncome int
	var choice string

	// Inisialisasi variabel
	totalExpense = 0

	// Pengguna memasukkan income
	fmt.Println("Masukkan Income :")
	fmt.Scan(&income)

	// Validasi saat income error
	if income < 0 {
		for income < 0 {
			fmt.Println("Error : Income tidak boleh negatif!")
			fmt.Println("Masukkan Income :")
			fmt.Scan(&income)
		}
	}

	// Pengguna memasukkan expense
	for {
		fmt.Println("Masukkan Expense :")
		fmt.Scan(&expense)

		// Validasi saat expense error
		for expense < 0 {
			fmt.Println("Error: Expense tidak boleh negatif!")
			fmt.Println("Masukkan Expense :")
			fmt.Scan(&expense)
		}

		// Menambahkan expense ke total expense
		totalExpense = totalExpense + expense

		// Menanyakan pengguna, apakah ingin menambah expense lagi?
		fmt.Println("Tambah expense lagi? (y/n)")
		fmt.Scan(&choice)

		// Validasi jawaban pengguna
		if choice == "n" {
			break
		}
	}

	// Menghitung sisa pemasukan
	remainingIncome = income - totalExpense

	// Output berupa list keuangan
	fmt.Println("------------------------------")
	fmt.Println("Total Income:", income)
	fmt.Println("Total Expense:", totalExpense)
	fmt.Println("Remaining Income:", remainingIncome)
	fmt.Println("------------------------------")
}
