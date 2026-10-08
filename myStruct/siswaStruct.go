package mystruct

import "fmt"


type Domisili struct {
	Kota string
	Provinsi string
}

type Siswa struct{
	Nama string
	Nilai int
	Alamat Domisili
}

func (s *Siswa) Lulus() string{
	if s.Nilai > 75 {return "Lulus"}
	return "Tidak lulus"
}


func MySiswa(){
	/*
	siswa := Siswa{
		Nama: "Alfred",
		Nilai: 999,
		Alamat: Domisili{Kota: "Sleman", Provinsi: "DIY"},
	}
*/
	daftar := []Siswa{
		{Nama: "John Doe", Nilai: 77},
		{Nama: "Jane Doe", Nilai: 88},
		{Nama: "John Smith", Nilai: 67},
		{Nama: "Jane Smith", Nilai: 76},
	}

	for i, v := range daftar {
		fmt.Printf("Nama siswa %d: %-15sNilai: %d ",i+1, v.Nama, v.Nilai)
		fmt.Printf("%s\n", v.Lulus())
	}
/*
	fmt.Printf("Nama: %s\nNilai: %d\n", siswa.Nama, siswa.Nilai)
	fmt.Printf("Domisili, Kota %s, Provinsi %s\n\n", siswa.Alamat.Kota, siswa.Alamat.Provinsi)

	fmt.Println(daftar)
*/
}
