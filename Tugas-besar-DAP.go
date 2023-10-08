package main

import "fmt"

type point struct {
	id, nama, password string
	teman              [N]string
	jml                int
}

const N int = 200

type tabel struct {
	array [N]point
	n     int
}

type chat struct {
	chat [N]isi
	n    int
}

type isi struct {
	isi [N]point2
	n   int
}

type point2 struct {
	user, isichat string
}

func calonTeman(tab *tabel) {
	var i int

	i = 0
	for i < tab.n {
		i++
	}
	fmt.Print("Nama : ")
	fmt.Scan(&tab.array[i].nama)
	fmt.Print("id : ")
	fmt.Scan(&tab.array[i].id)
	fmt.Print("password : ")
	fmt.Scan(&tab.array[i].password)
}

func menu1(tab *tabel, t *chat) {
	var klik int

	fmt.Println(" ")
	fmt.Println("===== Menu Utama =====")
	fmt.Println("1. Insert")
	fmt.Println("2. View") 
	fmt.Scan(&klik)
	if klik == 1 {
		menu()
	} else if klik == 2 {
		view(&*tab, &*t)
	}
}

func menu() {
	fmt.Println(" ")
	fmt.Println("===== Insert =====")
	fmt.Println("Tekan 1 untuk daftar")
	fmt.Println("Tekan 2 untuk login")
	fmt.Println("Tekan 3 untuk kembali ke menu awal")
	fmt.Println("Tekan 4 untuk Exit")

}

func view(tab *tabel, t *chat) {
	var klik int

	fmt.Println(" ")
	fmt.Println("===== View =====")
	fmt.Println("Tekan 1 untuk Edit")
	fmt.Println("Tekan 2 untuk Delete")
	fmt.Println("Tekan 3 untuk cari ")
	fmt.Println("Tekan 4 untuk sorting ")
	fmt.Println("Tekan 5 untuk Tampilkan Teman")
	fmt.Println("Tekan 6 untuk tampilkan chat")

	fmt.Scan(&klik)
	if klik == 1 {
		edit(&*tab)
	} else if klik == 2 {
		delete(&*tab)
	} else if klik == 3 {
		cari1(&*tab)
	} else if klik == 4 {
		sorting(*tab)

	} else if klik == 5 {
		tampilinTeman(&*tab)
	} else if klik == 6 {
		tampilChatting(*t)
	}
}
func edit(tab *tabel) {
	var idbaru, id string
	var i int
	var found bool

	fmt.Println(" ")
	fmt.Print("Masukkan ID lama : ")
	fmt.Scan(&id)
	i = 0
	found = false
	for (i <= tab.n) && (!found) {
		if tab.array[i].id == id {
			found = true
		} else {
			i = i + 1
		}
	}
	if (found) && (i <= tab.n) {
		fmt.Print("Masukkan ID baru : ")
		fmt.Scan(&idbaru)
		tab.array[i].id = idbaru
		fmt.Println(tab.array[i].id)
	} else {
		fmt.Println("ID salah ")
		fmt.Println("Masukkan ID lama : ")
		fmt.Scan(&id)
	}
}

func delete(tab *tabel) {
	var idx, i  int
	var found bool
	var nama string

	fmt.Println("\n=====Delete Kontak=====")
	fmt.Println("Input Kontak : ")
	fmt.Scan(&nama)
	for (!found) && (i <= tab.n) {
		if tab.array[i].nama == nama {
			found = true
			 idx = i
		} else {
			i = i + 1
		}
	}
	for i = idx; i < 6-1; i++ {
		tab.array[i].nama = tab.array[i+1].nama
		tab.array[i].teman[tab.array[i].jml] = tab.array[i+1].teman[tab.array[i+1].jml] 
	}
	
	
}

func cari1(tab *tabel) {
	var klik int
	var idx int

	fmt.Println(" ")
	fmt.Println("===== Cari =====")
	fmt.Println("1. Nama")
	fmt.Println("2. ID")

	fmt.Scan(&klik)
	if klik == 1 {
		cari(&*tab, &idx)
	} else if klik == 2 {
		cari2(&*tab, &idx)
	}
}

func cari(tab *tabel, idx *int) {
	var i int
	var found bool
	var nama string

	i = 0
	found = false
	fmt.Print("Input Nama : ")
	fmt.Scan(&nama)
	for (!found) && (i <= tab.n) {
		if tab.array[i].nama == nama {
			found = true
			*idx = i
		} else {
			i = i + 1
		}
	}
	if found {
		fmt.Println("ketemu")
	} else {
		fmt.Println("Nama Tidak Ditemukan")
	}
}
func cari2(tab *tabel, idx *int) {
	var i int
	var found bool
	var id string

	fmt.Println("Cari ID")
	i = 0
	found = false
	fmt.Scan(&id)
	for (!found) && (i <= tab.n) {
		if tab.array[i].id == id {
			found = true
		} else {
			i = i + 1
		}
	}
	if found {
		fmt.Println("ID Ditemukan")
	} else {
		fmt.Println("ID Tidak Ditemukan")
	}
}
func sorting(tab tabel) {
	var klik, tekan, pass, i, j, idx_min, idx_max int
	var temp string

	fmt.Println(" ")
	fmt.Println("===== Sorting =====")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")

	fmt.Scan(&klik)
	if klik == 1 {
		fmt.Println("===== Selection Sort =====")
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")
		fmt.Scan(&tekan)
		if tekan == 1 {
			for j < tab.n {
				pass = 0
				for pass < tab.array[j].jml-1 {
					idx_min = pass
					i = pass + 1
					for i < tab.array[j].jml {
						if tab.array[j].teman[idx_min] > tab.array[j].teman[i] {
							idx_min = i
						}
						i++
					}
					temp = tab.array[j].teman[idx_min]
					tab.array[j].teman[idx_min] = tab.array[j].teman[pass]
					tab.array[j].teman[pass] = temp
					pass++

				}
				j++
			}
			i = 0
			for i < tab.n {
				fmt.Println("User:", tab.array[i].nama)
				fmt.Println("Teman:")
				for j := 0; j < tab.array[i].jml; j++ {
					fmt.Println(j+1, ". Nama:", tab.array[i].teman[j])
				}
				i++
			}

		} else if tekan == 2 {
			for j < tab.n {
				pass = 0
				i = pass + 1
				for pass < tab.array[j].jml-1 {
					idx_max = pass
					i = pass + 1
					for i < tab.array[j].jml {
						if tab.array[j].teman[idx_max] < tab.array[j].teman[i] {
							idx_max = i
						}
						i++
					}
					temp = tab.array[j].teman[idx_max]
					tab.array[j].teman[idx_max] = tab.array[j].teman[pass]
					tab.array[j].teman[pass] = temp
					pass++
				}
				j++
			}
			i = 0
			for i < tab.n {
				fmt.Println("User:", tab.array[i].nama)
				fmt.Println("Teman:")
				for j := 0; j < tab.array[i].jml; j++ {
					fmt.Println(j+1, ". Nama:", tab.array[i].teman[j])
				}
				i++
			}
		}
	} else if klik == 2 {
		fmt.Println(" ")
		fmt.Println("===== Insertion Sort =====")
		fmt.Println("1. Ascending :")
		fmt.Println("2. Descending : ")
		fmt.Scan(&tekan)
		if tekan == 1 {
			for j < tab.n {
				pass = 0
				for pass < tab.array[j].jml-1 {
					i = pass + 1
					temp = tab.array[j].teman[i]
					for i > 0 && temp < tab.array[j].teman[i-1] {
						tab.array[j].teman[i] = tab.array[j].teman[i-1]
						i--
					}
					fmt.Println(temp)
					tab.array[j].teman[i] = temp
					pass++
				}
				j++
			}
			i = 0
			for i < tab.n {
				fmt.Println("User:", tab.array[i].nama)
				fmt.Println("Teman:")
				for j := 0; j < tab.array[i].jml; j++ {
					fmt.Println(j+1, ". Nama:", tab.array[i].teman[j])
				}
				i++
			}
		} else if tekan == 2 {
			for j < tab.n {
				pass = 0
				for pass < tab.array[j].jml-1 {
					i = pass + 1
					temp = tab.array[j].teman[i]
					for i > 0 && temp > tab.array[j].teman[i-1] {
						tab.array[j].teman[i] = tab.array[j].teman[i-1]
						i--
					}
					fmt.Println(temp)
					tab.array[j].teman[i] = temp
					pass++
				}
				j++
			}
			i = 0
			for i < tab.n {
				fmt.Println("User:", tab.array[i].nama)
				fmt.Println("Teman:")
				for j := 0; j < tab.array[i].jml; j++ {
					fmt.Println(j+1, ". Nama:", tab.array[i].teman[j])
				}
				i++
			}
		}
	}
}

func daftar(tab *tabel) {
	var i int
	var id, password, nama string
	var found bool

	fmt.Println("=== Daftar ===")
	fmt.Print("Nama : ")
	fmt.Scan(&nama)
	fmt.Print("ID : ")
	fmt.Scan(&id)
	fmt.Print("Password : ")
	fmt.Scan(&password)

	i = 0
	found = false
	for (!found) && (i <= tab.n) {
		if tab.array[i].id == id {
			found = true
		} else {
			i = i + 1
		}
	}
	if found {
		fmt.Println("Data sudah ada")
	} else {
		tab.array[tab.n].nama = nama
		tab.array[tab.n].id = id
		tab.array[tab.n].password = password
		tab.n++
	}
}

func login(tab *tabel, t *chat) {
	var id, password string
	var i int
	var found bool

	fmt.Println("===== Login =====")
	fmt.Print("ID : ")
	fmt.Scan(&id)
	fmt.Print("Password : ")
	fmt.Scan(&password)

	i = 0
	found = false
	for (!found) && (i <= tab.n) {
		if tab.array[i].id == id && tab.array[i].password == password {
			found = true
		} else {
			i = i + 1
		}
	}
	if !found {
		fmt.Println("Data tidak ada")
		menu()

	} else {
		fmt.Println("Login Berhasil")
		fmt.Println("1. Tambah Teman ")
		fmt.Println("2. Chatting")
		fmt.Println("3. Logout")
		tambahTeman(tab, i, t)
	}
}

func tambahTeman(tab *tabel, indeks int, t *chat) {
	var tekan int
	var nama string

	fmt.Scan(&tekan)
	if tekan == 1 {
		fmt.Print("Nama :")
		fmt.Scan(&nama)
		tab.array[indeks].teman[tab.array[indeks].jml] = nama
		tab.array[indeks].jml++
		fmt.Println("Tambah teman berhasil")
		fmt.Println("1. Tambah Teman ")
		fmt.Println("2. Chatting")
		fmt.Println("3. Logout")
		tambahTeman(tab, indeks, t)
	} else if tekan == 2 {
		chatting(*tab, &*t)
	}
}

func tampilinTeman(tab *tabel) {
	var i, j int

	//loop pertamaa buat akunnya
	i = 0
	for i <= tab.n {
		fmt.Println("User : ", tab.array[i].nama)

		j = 0
		for j <= tab.array[i].jml {
			fmt.Println("Teman :", tab.array[i].teman[j])
			j++
		}
		i = i + 1
	}
}
func chatting(tab tabel, t *chat) {
	var isichat, nama string
	var i, j int
	var found bool

	for nama != "stop" {
		fmt.Print("User :")
		fmt.Scan(&nama)
		i = 0
		found = false
		for (!found) && (i <= tab.n) {
			if tab.array[i].nama == nama || tab.array[i].teman[j] == nama {
				found = true
			} else {
				i++
			}
		}
		if found {
			fmt.Println("Isi chat : ")
			fmt.Scan(&isichat)
			t.chat[t.n].isi[t.chat[t.n].n].user = nama
			t.chat[t.n].isi[t.chat[t.n].n].isichat = isichat
			t.chat[t.n].n++
			
		}
	}
	t.n++
}

func tampilChatting(t chat) {
	var i, j int

	for i < t.n {
		fmt.Println("Chat ke :", i)
		for j < t.chat[i].n {
			fmt.Println("User :", t.chat[i].isi[j].user)
			fmt.Println("Isi chat :", t.chat[i].isi[j].isichat)
			j++
		}
		i++
	}
}

func main() {
	var tab tabel
	var klik int
	var t chat

	tab.array[0].nama = "rahma"
	tab.array[0].id = "rahma"
	tab.array[0].password = "rahma"
	tab.array[1].nama = "egi"
	tab.array[1].id = "egi"
	tab.array[1].password = "egi"
	tab.array[2].nama = "dika"
	tab.array[2].id = "dika"
	tab.array[2].password = "dika"
	tab.array[3].nama = "dita"
	tab.array[3].id = "dita"
	tab.array[3].password = "dita"
	tab.array[4].nama = "azriel"
	tab.array[4].id = "azriel"
	tab.array[4].password = "azriel"
	tab.n = 5
	for klik != 4 {
		menu1(&tab, &t)
		fmt.Println("Main")
		fmt.Scan(&klik)
		if klik == 1 {
			daftar(&tab)
		} else if klik == 2 {
			login(&tab, &t)
		} else if klik == 3 {
			menu1(&tab, &t)
		} else {
			//i = 4
		}
	}
}
