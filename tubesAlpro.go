package main
import "fmt"

type historiTransaksi struct {
	pelanggan dataPelanggan
	sukuCadang sparePart
	tarif float64
	jumlahSparepart int
	waktuTransaksi waktu
}

type sparePart struct {
	namaSparePart string
	merekSparePart string
	hargaBeli float64
	tarifServis float64
	jumlahPadaServis int
}


type dataPelanggan struct {
	nama  string
	umur int
}
type waktu struct {
	tanggal int
	bulan int
	tahun int
}
type spareParts struct {
	sParts [1000]sparePart
	nSparts int
}
type transanctionRecords struct{
	hTransaksi [3500]historiTransaksi
	nTransaksi int
}
type customerData struct {
	customers [3000]dataPelanggan
	nCustomers int
}

func main () {
menu()
//func main hanya akan berisi prosedur menu saja	
}

func menu () {
	var pilihan int
	var parts spareParts
	var transaksi transanctionRecords
	var pelanggan customerData
	Header()
	fmt.Println("Pilihan Anda:")
	fmt.Scan(&pilihan)
	for pilihan!=6 {
		if pilihan==1 {
			menuInputData(&parts,&transaksi,&pelanggan)
		}else if pilihan==2 {
			menuUbahData(&transaksi,&pelanggan,&parts)
		}else if pilihan==3{
			sortingData(&transaksi)
			mencariData(transaksi)
		}else if pilihan==4 {
			menampilkanSparepart(&parts)
		}else if pilihan==5 {
			sortingData(&transaksi)
			menampilkanData(transaksi,parts,pelanggan)
			//opsi: menunjugan kembali header beserta isi pilihan yang ada
		}else if pilihan!=6{
			fmt.Println("Pilihan Tidak Sesuai")
		}
		
		/*menampilkanTransaksi(transaksi)
		for i:=0;i<transaksi.nTransaksi;i++{
			fmt.Println(transaksi.hTransaksi[i])
		}
		for i:=0;i<pelanggan.nCustomers;i++{
			fmt.Println(pelanggan.customers[i])
		}
		for i:=0;i<parts.nSparts;i++{
			fmt.Println(parts.sParts[i])
		}*/
		Header()
		fmt.Println("Pilihan Anda:")
		fmt.Scan(&pilihan)
	}
	
}
func menampilkanData(T transanctionRecords,S spareParts,C customerData) {
	var pilihan string
	Header5()
	fmt.Println("Pilihan Anda :")
	fmt.Scan(&pilihan)
	for pilihan!="4"{
		if pilihan=="1"{
			Header51()
			menampilkanTransaksi(T)
			fmt.Println()
		}else if pilihan =="2"{
			Header52()
			menampilkanCustomer(C)
			fmt.Println()
		}else if pilihan=="3"{
			Header53()
			menampilkanSparepart1(S)
			fmt.Println()
		}
		if pilihan!="4"{
			Header5()
			fmt.Println("Pilihan Anda : ")
			fmt.Scan(&pilihan)
		}
	}

}
func menampilkanSparepart(S *spareParts) {
var pilihan string
	Header4()
	fmt.Println("Pilihan Anda : ")
	fmt.Scan(&pilihan)
	for pilihan!="3"{
	if pilihan =="1"{
		Header41()
		sortingSparePartDSC(S)
		menampilkanSparepart1(*S)
	}else if pilihan=="2"{
		Header42()
	sortingSparePartAsc(S)
		menampilkanSparepart1(*S)
	}
	if pilihan!="3"{
		Header4()
		fmt.Println("Pilihan Anda : ")
		fmt.Scan(&pilihan)
	}
	
	}

}

func sortingSparePartAsc (S *spareParts) {
	var j int
	var temp sparePart
	for i:=1;i<S.nSparts;i++{
		j=i
		temp =S.sParts[i]
		for j>=1 && temp.jumlahPadaServis<S.sParts[j-1].jumlahPadaServis{
			S.sParts[j]=S.sParts[j-1]
			j=j-1
		}
		S.sParts[j]=temp
		
	
	
	}


}
func sortingSparePartDSC (S *spareParts) {
	var j int
	var temp sparePart
	for i:=1;i<S.nSparts;i++{
		j=i
		temp =S.sParts[i]
		for j>=1 && temp.jumlahPadaServis>S.sParts[j-1].jumlahPadaServis{
			S.sParts[j]=S.sParts[j-1]
			j=j-1
		}
		S.sParts[j]=temp
		
	
	
	}


}

//Memasukan Pilihan user
//1 Input data 
//2 mengubah data (mengakah ke menu selanjutnya yaitu meng-edit,menghapus atau menambah data)
//3 Mencari data (retang transaksi)
//4 Menampilkan sparepart yang sering digunakan
//5 selesai 
//Note : masing =masing pilihan akan memiliki fungsi/prosedur tersendiri

func menuInputData(parts *spareParts,transaksi *transanctionRecords,pelanggan *customerData) {
var cek bool
var umur,pilihan int
var nama, namaSparePart,merekSparePart string
var idx int
//Header menu input data
Header1()
fmt.Println("Pilihan Anda:")
fmt.Scan(&pilihan)
for pilihan!=4{
if pilihan ==1 {
	Header11()
	fmt.Print("	Nama Pelanggan : ")
	fmt.Scan(&nama)
	fmt.Println()
	fmt.Print("	Umur Pelanggan : ")
	fmt.Scan(&umur)
	fmt.Println()
	for nama!="Selesai" && umur!=0{
		
		if ada(*pelanggan,nama) == -1 || ada2(*pelanggan,umur)==-1{
			pelanggan.customers[pelanggan.nCustomers].nama=nama
			pelanggan.customers[pelanggan.nCustomers].umur=umur
			pelanggan.nCustomers++
			
		}
		fmt.Print("	Nama Spare-part : ")
		fmt.Scan(&namaSparePart)
		fmt.Println()
		fmt.Print("	Merek Spare-part :")
		fmt.Scan(&merekSparePart)
		fmt.Println()
		idx=ada3(*parts,namaSparePart, merekSparePart)
		if idx ==-1{
			fmt.Println("	Nama Dan Merek Spare-part Belum Ada ")
			fmt.Println("	Mohon Masukan Data Terkait Spare-part")
			fmt.Print("	Harga Beli Spare-part : ")
			
			fmt.Scan(&transaksi.hTransaksi[transaksi.nTransaksi].sukuCadang.hargaBeli)
			fmt.Println()
			fmt.Print("	Tarif Servis Spare-part : ")
			fmt.Scan(&transaksi.hTransaksi[transaksi.nTransaksi].sukuCadang.tarifServis)
			fmt.Println()
			fmt.Println("	Data Sudah Ditambahkan ")
			fmt.Println()
			parts.sParts[parts.nSparts].hargaBeli=transaksi.hTransaksi[transaksi.nTransaksi].sukuCadang.hargaBeli
			parts.sParts[parts.nSparts].tarifServis=transaksi.hTransaksi[transaksi.nTransaksi].sukuCadang.tarifServis
			parts.sParts[parts.nSparts].namaSparePart=namaSparePart
			parts.sParts[parts.nSparts].merekSparePart=merekSparePart
			parts.nSparts++
		}else {
			fmt.Println()
			fmt.Println("	Data Spare-part Sudah terdaftar	")
			fmt.Println()
			transaksi.hTransaksi[transaksi.nTransaksi].sukuCadang.hargaBeli=parts.sParts[idx].hargaBeli
			transaksi.hTransaksi[transaksi.nTransaksi].sukuCadang.tarifServis=parts.sParts[idx].tarifServis
		
		}
		idx=ada3(*parts,namaSparePart, merekSparePart)
		
		transaksi.hTransaksi[transaksi.nTransaksi].pelanggan.nama=nama
		transaksi.hTransaksi[transaksi.nTransaksi].pelanggan.umur=umur
		
		transaksi.hTransaksi[transaksi.nTransaksi].sukuCadang.namaSparePart=namaSparePart
		transaksi.hTransaksi[transaksi.nTransaksi].sukuCadang.merekSparePart=merekSparePart
		
		
		fmt.Println("	Mohon Masukan Data Berikut :	")
		fmt.Print("	Banyak Spare-part yang digunakan : ")
		fmt.Scan(&transaksi.hTransaksi[transaksi.nTransaksi].jumlahSparepart)
		fmt.Println()
		fmt.Print("	Tanggal Transaksi :	")
		fmt.Scan(&transaksi.hTransaksi[transaksi.nTransaksi].waktuTransaksi.tanggal)
		fmt.Println()
		fmt.Print("	Bulan Transaksi	  :	")
		fmt.Scan(&transaksi.hTransaksi[transaksi.nTransaksi].waktuTransaksi.bulan)
		fmt.Println()
		fmt.Print("	Tahun Transaksi	  :	")
		fmt.Scan(&transaksi.hTransaksi[transaksi.nTransaksi].waktuTransaksi.tahun)
		fmt.Println()
		cek=cekValidWaktu(transaksi.hTransaksi[transaksi.nTransaksi].waktuTransaksi.tanggal,transaksi.hTransaksi[transaksi.nTransaksi].waktuTransaksi.bulan,transaksi.hTransaksi[transaksi.nTransaksi].waktuTransaksi.tahun)
		for !cek {
		fmt.Printf("\n Data Waktu Tidak Valid Mohon Input Ulang\n\n")
		fmt.Print("	Tanggal Transaksi :	")
		fmt.Scan(&transaksi.hTransaksi[transaksi.nTransaksi].waktuTransaksi.tanggal)
		fmt.Println()
		fmt.Print("	Bulan Transaksi	  :	")
		fmt.Scan(&transaksi.hTransaksi[transaksi.nTransaksi].waktuTransaksi.bulan)
		fmt.Println()
		fmt.Print("	Tahun Transaksi	  :	")
		fmt.Scan(&transaksi.hTransaksi[transaksi.nTransaksi].waktuTransaksi.tahun)
		fmt.Println()
		cek=cekValidWaktu(transaksi.hTransaksi[transaksi.nTransaksi].waktuTransaksi.tanggal,transaksi.hTransaksi[transaksi.nTransaksi].waktuTransaksi.bulan,transaksi.hTransaksi[transaksi.nTransaksi].waktuTransaksi.tahun)
		
		}
		transaksi.hTransaksi[transaksi.nTransaksi].tarif=hitungTarif (float64(transaksi.hTransaksi[transaksi.nTransaksi].jumlahSparepart),transaksi.hTransaksi[transaksi.nTransaksi].sukuCadang.tarifServis)
		fmt.Println("	Tarif Pada Transaksi ini Sebesar : ",transaksi.hTransaksi[transaksi.nTransaksi].tarif)
		fmt.Println()
		//parts.sParts[parts.nSparts-1].jumlahPadaServis=parts.sParts[parts.nSparts-1].jumlahPadaServis+transaksi.hTransaksi[transaksi.nTransaksi].jumlahSparepart
		parts.sParts[idx].jumlahPadaServis=parts.sParts[idx].jumlahPadaServis+transaksi.hTransaksi[transaksi.nTransaksi].jumlahSparepart
		transaksi.hTransaksi[transaksi.nTransaksi].sukuCadang.jumlahPadaServis=parts.sParts[idx].jumlahPadaServis
		transaksi.nTransaksi=transaksi.nTransaksi+1
		Header11()
			fmt.Print("	Nama Pelanggan : ")
		fmt.Scan(&nama)
		fmt.Println()
		fmt.Print("	Umur Pelanggan : ")
		fmt.Scan(&umur)
		fmt.Println()
	
	}
	
}else if pilihan == 2{
	Header12()
	fmt.Print("	Nama Pelanggan : ")
	fmt.Scan(&nama)
	fmt.Println()
	fmt.Print("	Umur Pelanggan : ")
	fmt.Scan(&umur)
	fmt.Println()
	for nama!="Selesai" && umur!=0 {
		for ada(*pelanggan,nama) != -1 && ada2(*pelanggan,umur)!=-1{
		fmt.Println()
		fmt.Println("	Nama dan Umur sudah ada tolong input kembali	")	
		fmt.Println()
		Header12()
		fmt.Print("	Nama Pelanggan : ")
		fmt.Scan(&nama)
		fmt.Println()
		fmt.Println("	Umur Pelanggan : ")
		fmt.Scan(&umur)
		fmt.Println()	
		}
		pelanggan.customers[pelanggan.nCustomers].nama=nama
		pelanggan.customers[pelanggan.nCustomers].umur=umur
		pelanggan.nCustomers=pelanggan.nCustomers+1
		Header12()
		fmt.Print("	Nama Pelanggan : ")
		fmt.Scan(&nama)
		fmt.Println()
		fmt.Print("	Umur Pelanggan : ")
		fmt.Scan(&umur)
		fmt.Println()
		
		
		
		
	}

	
	
}else if pilihan == 3 {
	Header13()
		fmt.Print("	Nama Spare-part : ")
		fmt.Scan(&namaSparePart)
		fmt.Println()
		fmt.Print("	Merek Spare-part :")
		fmt.Scan(&merekSparePart)
		fmt.Println()
	for namaSparePart!="Selesai" && merekSparePart!="Selesai"{
		for ada3(*parts,namaSparePart, merekSparePart)!=-1 {
		
		fmt.Println("	Nama dan merek Spare-part sudah ada tolong input kembali	")
			fmt.Println()
		Header13()
		fmt.Print("	Nama Spare-part : ")
		fmt.Scan(&namaSparePart)
		fmt.Println()
		fmt.Print("	Merek Spare-part :")
		fmt.Scan(&merekSparePart)
		fmt.Println()	
		}
		parts.sParts[parts.nSparts].namaSparePart=namaSparePart
		parts.sParts[parts.nSparts].merekSparePart=merekSparePart
		fmt.Print("	Harga Beli Spare-part : ")
		fmt.Scan(&parts.sParts[parts.nSparts].hargaBeli)
		fmt.Println()
		fmt.Print("	Tarif Servis Spare-part : ")
		fmt.Scan(&parts.sParts[parts.nSparts].tarifServis)
		fmt.Println()
		
		parts.nSparts=parts.nSparts+1
		
		Header13()
		fmt.Print("	Nama Spare-part : ")
		fmt.Scan(&namaSparePart)
		fmt.Println()
		fmt.Print("	Merek Spare-part :")
		fmt.Scan(&merekSparePart)
		fmt.Println()
	}

}
if pilihan !=4 {
	Header1()
	fmt.Println("Pilihan Anda:")
	fmt.Scan(&pilihan)
}

}
}
	
func hitungTarif(a,b float64) float64 {
	return a*b
	
}
	
	
	/*type transanctionRecords {
	hTransaksi [3500]historiTransaksi
	nTransaksi int
	
	type historiTransaksi struct {
	pelanggan dataPelanggan
	jenisServis string
	sukuCadang sparePart
	biayaServis float64
	waktuTransaksi waktu
	statusTransaksi string
}
}*/

func ada (pelanggan customerData,nama string) int {
	
	var i,index int
	index=-1
	for i<pelanggan.nCustomers && index==-1 {
		if pelanggan.customers[i].nama==nama{
			index=i
		}
		i=i+1
	}
	return index

}
func ada2 (pelanggan customerData, umur int) int {
	
	var i,index int
	index=-1
	for i<pelanggan.nCustomers && index==-1 {
		if pelanggan.customers[i].umur==umur{
			index=i
		}
		i=i+1
	}
	return index

}
func ada3 (parts spareParts ,nama, merek string) int {
	
	var i,index int
	index=-1
	for i<parts.nSparts&& index==-1 {
		if parts.sParts[i].namaSparePart==nama&&parts.sParts[i].merekSparePart==merek{
			index=i
		}
		i=i+1
	}
	return index

}
func ada4(pelanggan customerData,umur int,nama string) int {
	var i,index int
	index=-1
	for i<pelanggan.nCustomers && index==-1 {
		if pelanggan.customers[i].nama==nama &&pelanggan.customers[i].umur==umur{
			index=i
		}
		i=i+1
	}
	return index

}
/*
}
//1 menginput history transaksi
//2 menginput dataPelanggan
//3 menginput data sparePart
//4 kembali/selesai


func header input data*/








func menuUbahData (T *transanctionRecords,C *customerData,S *spareParts) {
var pilihan string
var nama,namaSparePart,merekSparePart string 
var umur,jServis,tgl,bln,thn,jmlSparepart ,index,temp,index2 int
var hargaSparepart,tarifSparepart,tarif float64
fmt.Println("Pilihan Anda: (Ketik 'Selesai' jika sudah selesai)")
//Q menambah (akan masuk ke menu memilih data apa yang ditambah (historiTransaksi,dataPelanggan,atau data sparePart))
//1 mengubah (akan masuk ke menu memilih data apa yang diubah (historiTransaksi,dataPelanggan,atau data sparePart))
//2 menghapus (akan masuk ke menu memilih data apa yang dihapus (historiTransaksi,dataPelanggan,atau data sparePart))
//3 selesai
Header2()
fmt.Println("Pilihan Anda :	")
fmt.Scan(&pilihan)
//fmt.Println(C.nCustomers,pilihan=="2")
for pilihan!="3"{
    if pilihan=="1"{
	
	Header21()
	//fmt.Println("Pilihan Data yang diubah : (Ketik 'Selesai' jika sudah selesai)")
	fmt.Println("Pilihan Anda :")
	fmt.Scan(&pilihan)
	for pilihan!="4"{
	//bikin kondisi kalau data blm diisi
		if pilihan=="1"{
			if T.nTransaksi==0{
				fmt.Println()
				fmt.Println("Data Histori Transaksi Masih Kosong Mohon Input Data Terlebih Dahulu")
				fmt.Println()
				pilihan="4"
			}else {
			Header211()
			menampilkanTransaksi(*T)
			fmt.Print("	Nama Pelanggan : ")
		fmt.Scan(&nama)
		fmt.Println()
		fmt.Print("	Umur Pelanggan : ")
		fmt.Scan(&umur)
		fmt.Println()
		fmt.Print("	Nama Spare-part : ")
		fmt.Scan(&namaSparePart)
		fmt.Println()
		fmt.Print("	Merek Spare-part :")
		fmt.Scan(&merekSparePart)
		fmt.Println()
		fmt.Print("	Harga Beli Spare-part : ")
		fmt.Scan(&hargaSparepart)
		fmt.Println()
		fmt.Print("	Tarif Servis Spare-part : ")
		fmt.Scan(&tarifSparepart)
		fmt.Println()
		fmt.Print("	Total Spare-part yang telah digunakan : ")
		fmt.Scan(&jServis)
		fmt.Println()
		fmt.Print("   	Total Tarif Transaksi : ")
		fmt.Scan(&tarif)
		fmt.Println()
		fmt.Print("	Banyak Spare-part yang digunakan saat servis : ")
		fmt.Scan(&jmlSparepart)
		fmt.Println()
		fmt.Print("	Tanggal Transaksi :	")
		fmt.Scan(&tgl)
		fmt.Println()
		fmt.Print("	Bulan Transaksi	  :	")
		fmt.Scan(&bln)
		fmt.Println()
		fmt.Print("	Tahun Transaksi	  :	")
		fmt.Scan(&thn)
		fmt.Println()

			//fmt.Println("Silahkan masukan data historiTransaksi yang ingin diubah (Nama Pelanggan,umur pelanggan,nama spare part,merek spare part, jumlah spare part, waktu transaksi (tanggal,bulan tahun))")
			//menunjukan array histori transaksi 
			//fmt.Scan(&nama,&umur,&namaSparePart,&merekSparePart,&hargaSparepart,&tarifSparepart,&jServis,&tarif,&jmlSparepart,&tgl,&bln,&thn)
			index =cariData(*T,nama,namaSparePart,merekSparePart,umur,jServis,tgl,bln,thn,jmlSparepart,hargaSparepart,tarifSparepart,tarif)
			for index==-1 {
				fmt.Println("data tidak ditemukan silahkan isi ulang data yang dicari")
				menampilkanTransaksi(*T)
			fmt.Print("	Nama Pelanggan : ")
		fmt.Scan(&nama)
		fmt.Println()
		fmt.Print("	Umur Pelanggan : ")
		fmt.Scan(&umur)
		fmt.Println()
		fmt.Print("	Nama Spare-part : ")
		fmt.Scan(&namaSparePart)
		fmt.Println()
		fmt.Print("	Merek Spare-part :")
		fmt.Scan(&merekSparePart)
		fmt.Println()
		fmt.Print("	Harga Beli Spare-part : ")
		fmt.Scan(&hargaSparepart)
		fmt.Println()
		fmt.Print("	Tarif Servis Spare-part : ")
		fmt.Scan(&tarifSparepart)
		fmt.Println()
		fmt.Print("	Total Spare-part yang telah digunakan : ")
		fmt.Scan(&jServis)
		fmt.Println()
		fmt.Print("   	Total Tarif Transaksi : ")
		fmt.Scan(&tarif)
		fmt.Println()
		fmt.Print("	Banyak Spare-part yang digunakan saat servis : ")
		fmt.Scan(&jmlSparepart)
		fmt.Println()
		fmt.Print("	Tanggal Transaksi :	")
		fmt.Scan(&tgl)
		fmt.Println()
		fmt.Print("	Bulan Transaksi	  :	")
		fmt.Scan(&bln)
		fmt.Println()
		fmt.Print("	Tahun Transaksi	  :	")
		fmt.Scan(&thn)
		fmt.Println()
				index =cariData(*T,nama,namaSparePart,merekSparePart,umur,jServis,tgl,bln,thn,jmlSparepart,hargaSparepart,tarifSparepart,tarif)
			}
			Header211()
			fmt.Println("Data telah ditemukan, data apa yang ingin anda ubah ?")
			fmt.Println("1 Nama Pelanggan")
			fmt.Println("2 Umur Pelanggan")
			fmt.Println("3 Nama Sparepart")
			fmt.Println("4 Merek Sparepart")
			fmt.Println("5 Harga Beli Sparepart")
			fmt.Println("6 Tarif Servis Sparepart")
			fmt.Println("7 Jumlah Sparepart yang telah digunakan")
			fmt.Println("8 tarif transaksi")
			fmt.Println("9 Banyak spare part yang digunakan saat transaksi")
			fmt.Println("10 tanggal transaksi")
			fmt.Println("11 bulan transaksi")
			fmt.Println("12 tahun transaksi")
			fmt.Println("13 Selesai")
			
			fmt.Println("Pilihan Anda:")
			fmt.Scan(&pilihan)
			for pilihan!="13"{
		
			if pilihan =="1" {
				fmt.Println("Masukan Data Nama Pelanggan Baru :")
				fmt.Scan(&T.hTransaksi[index].pelanggan.nama)
			}else if pilihan=="2"{
				fmt.Println("Masukan Data Umur Pelanggan Baru :")
				fmt.Scan(&T.hTransaksi[index].pelanggan.umur)
			}else if pilihan=="3"{
				fmt.Println("Masukan Data Nama Spare-part Baru :")
				fmt.Scan(&T.hTransaksi[index].sukuCadang.namaSparePart)
				
				
				
			}else if pilihan=="4"{
				fmt.Println("Masukan Data Merek Spare-part Baru :")
				fmt.Scan(&T.hTransaksi[index].sukuCadang.merekSparePart)
				
				
				
			}else if pilihan=="5"{
				fmt.Println("Masukan Data Harga Beli Spare-part Baru :")
				fmt.Scan(&T.hTransaksi[index].sukuCadang.hargaBeli)
			}else if pilihan=="6"{
				fmt.Println("Masukan Data Tarif Servis Spare-part Baru :")
				fmt.Scan(&T.hTransaksi[index].sukuCadang.tarifServis)
				T.hTransaksi[index].tarif=hitungTarif(T.hTransaksi[index].sukuCadang.tarifServis,float64(T.hTransaksi[index].jumlahSparepart))
			}else if pilihan=="7"{
				fmt.Println("Masukan Data Total Spare-part yang telah digunakan dalam Servis yang Baru :")
				fmt.Scan(&T.hTransaksi[index].sukuCadang.jumlahPadaServis)
				/*index2=ada3(*S,T.hTransaksi[index].sukuCadang.namaSparePart, T.hTransaksi[index].sukuCadang.merekSparePart)
				if index2!=-1{
					S.sParts[index2].jumlahPadaServis=T.hTransaksi[index].sukuCadang.jumlahPadaServis
				}*/
			}else if pilihan=="9"{
				fmt.Println("Masukan Banyak Spare-part yang digunakan dalam Servis yang Baru :")
				fmt.Scan(&temp)
				
				
				
				
				index2=ada3(*S,T.hTransaksi[index].sukuCadang.namaSparePart, T.hTransaksi[index].sukuCadang.merekSparePart)
				//fmt.Println(index2,T.hTransaksi[index].sukuCadang.namaSparePart, T.hTransaksi[index].sukuCadang.merekSparePart)
				if index2!=-1{
				//fmt.Println("Halo")
				S.sParts[index2].jumlahPadaServis=S.sParts[index2].jumlahPadaServis+(temp-T.hTransaksi[index].jumlahSparepart)
				
				//T.hTransaksi[index].sukuCadang.jumlahPadaServis=S.sParts[index2].jumlahPadaServis
				}
				T.hTransaksi[index].jumlahSparepart=temp
				T.hTransaksi[index].tarif=hitungTarif(T.hTransaksi[index].sukuCadang.tarifServis,float64(T.hTransaksi[index].jumlahSparepart))
			
			
			
			
			
			}else if pilihan=="10"{
				fmt.Println("Masukan Tanggal Baru : ")
				fmt.Scan(&T.hTransaksi[index].waktuTransaksi.tanggal)
			}else if pilihan=="11"{
				fmt.Println("Masukan Bulan Baru : ")
				fmt.Scan(&T.hTransaksi[index].waktuTransaksi.bulan)
			}else if pilihan=="12"{
				fmt.Println("Masukan Tahun Baru	:	")
				fmt.Scan(&T.hTransaksi[index].waktuTransaksi.tahun)
			}else if pilihan=="8"{
				fmt.Println("Masukan Tarif Transaksi Baru : ")
				fmt.Scan(&T.hTransaksi[index].tarif)
			}
			
			if pilihan!="13"{
			//	if pilihan=="3"||pilihan=="4"{
					//menampilkanSparepart1(*S)
				
				
				//}
				menampilkanTransaksi1(*T,index)
			Header211()
			fmt.Println("Data telah ditemukan, data apa yang ingin anda ubah ?")
			fmt.Println("1 Nama Pelanggan")
			fmt.Println("2 Umur Pelanggan")
			fmt.Println("3 Nama Sparepart")
			fmt.Println("4 Merek Sparepart")
			fmt.Println("5 Harga Beli Sparepart")
			fmt.Println("6 Tarif Servis Sparepart")
			fmt.Println("7 Jumlah Sparepart yang telah digunakan")
			fmt.Println("8 tarif transaksi")
			fmt.Println("9 Banyak spare part yang digunakan saat transaksi")
			fmt.Println("10 tanggal transaksi")
			fmt.Println("11 bulan transaksi")
			fmt.Println("12 tahun transaksi")
			fmt.Println("13 Selesai")
				fmt.Println("Pilihan Anda:")
				fmt.Scan(&pilihan)
				if pilihan=="13"{
				index2=ada3(*S,T.hTransaksi[index].sukuCadang.namaSparePart, T.hTransaksi[index].sukuCadang.merekSparePart)
				if index2==-1{
					menampilkanSparepart1(*S)
					fmt.Printf("\nData Spare-part Tidak Ditemukan\nApakah Ingin Ditambahkan Ke Dalam Data Spare-part?\nKetik 1 jika Ya\n")
					fmt.Scan(&temp)
					if temp==1{
					S.sParts[S.nSparts]=T.hTransaksi[index].sukuCadang
					S.nSparts=S.nSparts+1
					}
				
				}
			
			
			}
			}
			}
			//sisissisi
			}
	
			
		}else if pilihan=="2" {
		//TULIS KONDISI KALAU N DATA 0 ATAU KOSONG
		if C.nCustomers==0{
			fmt.Println()
			fmt.Println("Data Pelanggan Masih Kosong Mohon Input Data Terlebih Dahulu")
			fmt.Println()
			pilihan="4"
		
		}else {
		Header212()
		menampilkanCustomer(*C)
		//fmt.Println("Silahkan Masukan nama dan umur pelanggan yang ingin diubah")
		fmt.Print("	Nama Pelanggan : ")
		fmt.Scan(&nama)
		fmt.Println()
		fmt.Print("	Umur Pelanggan : ")
		fmt.Scan(&umur)
		fmt.Println()
		index=ada4(*C,umur,nama)
		for index==-1 {
		fmt.Println("Data tidak ditemukan mohon input kembali:")
		fmt.Print("	Nama Pelanggan : ")
		fmt.Scan(&nama)
		fmt.Println()
		fmt.Print("	Umur Pelanggan : ")
		fmt.Scan(&umur)
		fmt.Println()
		index=ada4(*C,umur,nama)
		}
	
		fmt.Println("Data telah ditemukan, silahkan tentukan data yang ingin diubah	:")
		fmt.Println("1 Nama Pelanggan")
		fmt.Println("2 Umur Pelanggan")
		fmt.Println("3 Selesai")
		fmt.Scan(&pilihan)
		if pilihan =="3"{
			pilihan="10"
		
		}
		for pilihan!="10"{
		if pilihan=="1"{
			fmt.Println("Masukan Nama Pelanggan Baru ")
			fmt.Scan(&C.customers[index].nama)
		}else if pilihan=="2"{
			fmt.Println("Masukan Umur Pelanggan Baru ")
			fmt.Scan(&C.customers[index].umur)
		
		}
		fmt.Println("Data ke ",index+1)
		fmt.Printf("Nama: %s Umur: %d\n",C.customers[index].nama,C.customers[index].umur)
		if pilihan!="3"{
			Header212()
			fmt.Println("1 Nama Pelanggan")
			fmt.Println("2 Umur Pelanggan")
			fmt.Println("3 Selesai")
			fmt.Println("Pilihan Anda:")
			fmt.Scan(&pilihan)
			
		}
		if pilihan =="3"{
			pilihan="10"
		
		}
		}
		//skdijfisjdfi
		}

}else if pilihan =="3"{
	if S.nSparts==0{
		fmt.Println()
		fmt.Println("Data Spare-part Masih Kosong Mohon Input Data Terlebih Dahulu")
		fmt.Println()
		pilihan="4"
	
	}else {
	Header213()
	menampilkanSparepart1(*S)
	//fmt.Println("Tentukan data mana yang mau diubah namaSparePart merekSparePart hargaBeli tarifServis jumlahPadaServis ")
	fmt.Print("	Nama Spare-part : ")
		fmt.Scan(&namaSparePart)
		fmt.Println()
		fmt.Print("	Merek Spare-part :")
		fmt.Scan(&merekSparePart)
		fmt.Println()
		fmt.Print("	Harga Beli Spare-part : ")
		fmt.Scan(&hargaSparepart)
		fmt.Println()
		fmt.Print("	Tarif Servis Spare-part : ")
		fmt.Scan(&tarifSparepart)
		fmt.Println()
		fmt.Print("	Total Spare-part yang telah digunakan : ")
		fmt.Scan(&jmlSparepart)
		fmt.Println()
	//fmt.Scan(&namaSparePart,&merekSparePart,&hargaSparepart,&tarifSparepart,&jmlSparepart)
	index=cariDataParts(*S,namaSparePart,merekSparePart,hargaSparepart,tarifSparepart,jmlSparepart)
	for index==-1 {
		fmt.Println("Data tidak ditemukan mohon input kembali:")
			Header213()
		menampilkanSparepart1(*S)
	
		fmt.Print("	Nama Spare-part : ")
		fmt.Scan(&namaSparePart)
		fmt.Println()
		fmt.Print("	Merek Spare-part :")
		fmt.Scan(&merekSparePart)
		fmt.Println()
		fmt.Print("	Harga Beli Spare-part : ")
		fmt.Scan(&hargaSparepart)
		fmt.Println()
		fmt.Print("	Tarif Servis Spare-part : ")
		fmt.Scan(&tarifSparepart)
		fmt.Println()
		fmt.Print("	Total Spare-part yang telah digunakan : ")
		fmt.Scan(&jmlSparepart)
		fmt.Println()
		//fmt.Scan(&namaSparePart,&merekSparePart,&hargaSparepart,&tarifSparepart,&jmlSparepart)
		index=cariDataParts(*S,namaSparePart,merekSparePart,hargaSparepart,tarifSparepart,jmlSparepart)
	}
	fmt.Println("Data telah ditemukan, tentukan data apa yang ingin diubah")
	fmt.Println("1 Nama Spare Part")
	fmt.Println("2 Merek Spare Part")
	fmt.Println("3 Harga Beli Spare Part")
	fmt.Println("4 Tarif Servis Spare Part")
	fmt.Println("5 Jumlah Spare Part Yang Telah Digunakan")
	fmt.Println("6 Selesai")
	
	fmt.Println("Pilihan Anda :")
	fmt.Scan(&pilihan)
	for pilihan!="6"{
	if pilihan=="1"{
		fmt.Println("Masukan Nama Spare part Baru :")
		fmt.Scan(&S.sParts[index].namaSparePart)
	}else if pilihan=="2"{
		fmt.Println("Masukan Merek Spare part Baru :")
		fmt.Scan(&S.sParts[index].merekSparePart)
	}else if pilihan=="3"{
		fmt.Println("Masukan Harga Beli Spare part Baru :")
		fmt.Scan(&S.sParts[index].hargaBeli)
	
	}else if pilihan =="4"{
		fmt.Println("Masukan Tarif Servis Spare part Baru :")
		fmt.Scan(&S.sParts[index].tarifServis)
	
	}else if pilihan =="5"{
		fmt.Println("Masukan Data Total Spare-part yang telah digunakan dalam Servis yang Baru :")
		fmt.Scan(&S.sParts[index].jumlahPadaServis)
	
	}
	
	if pilihan !="6"{
		//fmt.Println(S.sParts[index])
		fmt.Println("Data ke ",index+1)
		fmt.Printf("Nama Spare-part: %s Merek Spare-part: %s Harga Beli Spare-part: %.3f Tarif Servis Spare-part: %.3f Total Spare-part Yang Telah DIgunakan Dalam Servis: %d\n",S.sParts[index].namaSparePart,S.sParts[index].merekSparePart,S.sParts[index].hargaBeli,S.sParts[index].tarifServis,S.sParts[index].jumlahPadaServis)
			
		fmt.Println("1 Nama Spare Part")
		fmt.Println("2 Merek Spare Part")
		fmt.Println("3 Harga Beli Spare Part")
		fmt.Println("4 Tarif Servis Spare Part")
		fmt.Println("5 Jumlah Spare Part Yang Telah Digunakan")
		fmt.Println("6 Selesai")
		fmt.Println("Pilihan Anda : ")
		fmt.Scan(&pilihan)
	
	}
	
	}
	}
	//sjidjsijdi

}else if pilihan!="4"{
	Header21()
	fmt.Println("Pilihan Anda :")
	fmt.Scan(&pilihan)


}

}
}else if pilihan =="2"{
	Header22()
	fmt.Println("Pilihan Anda :")
	fmt.Scan(&pilihan)
	for pilihan!="4"{
		if pilihan=="1"{
		Header221()
		menampilkanTransaksi(*T)
		fmt.Print("	Nama Pelanggan : ")
		fmt.Scan(&nama)
		fmt.Println()
		fmt.Print("	Umur Pelanggan : ")
		fmt.Scan(&umur)
		fmt.Println()
		fmt.Print("	Nama Spare-part : ")
		fmt.Scan(&namaSparePart)
		fmt.Println()
		fmt.Print("	Merek Spare-part :")
		fmt.Scan(&merekSparePart)
		fmt.Println()
		fmt.Print("	Harga Beli Spare-part : ")
		fmt.Scan(&hargaSparepart)
		fmt.Println()
		fmt.Print("	Tarif Servis Spare-part : ")
		fmt.Scan(&tarifSparepart)
		fmt.Println()
		fmt.Print("	Total Spare-part yang telah digunakan : ")
		fmt.Scan(&jServis)
		fmt.Println()
		fmt.Print("   	Total Tarif Transaksi : ")
		fmt.Scan(&tarif)
		fmt.Println()
		fmt.Print("	Banyak Spare-part yang digunakan saat servis : ")
		fmt.Scan(&jmlSparepart)
		fmt.Println()
		fmt.Print("	Tanggal Transaksi :	")
		fmt.Scan(&tgl)
		fmt.Println()
		fmt.Print("	Bulan Transaksi	  :	")
		fmt.Scan(&bln)
		fmt.Println()
		fmt.Print("	Tahun Transaksi	  :	")
		fmt.Scan(&thn)
		fmt.Println()
		//fmt.Println("Silahkan masukan data historiTransaksi yang ingin dihapus (Nama Pelanggan,umur pelanggan,nama spare part,merek spare part, jumlah spare part, waktu transaksi (tanggal,bulan tahun))")
		
		//fmt.Scan(&nama,&umur,&namaSparePart,&merekSparePart,&hargaSparepart,&tarifSparepart,&jServis,&tarif,&jmlSparepart,&tgl,&bln,&thn)
		index =cariData(*T,nama,namaSparePart,merekSparePart,umur,jServis,tgl,bln,thn,jmlSparepart,hargaSparepart,tarifSparepart,tarif)
		if index ==-1 {
			fmt.Println("Data tidak ditemukan")
		}else {
			for i:=index;i<T.nTransaksi-1;i++{
				T.hTransaksi[i]=T.hTransaksi[i+1]
			
			}
			T.nTransaksi=T.nTransaksi-1
		
		}
		fmt.Println("Data Baru :")
		menampilkanTransaksi(*T)
		
	}else if pilihan =="2"{
		Header222()
		menampilkanCustomer(*C)
		//fmt.Println("Silahkan Masukan nama dan umur pelanggan yang ingin diubah")
		//fmt.Scan(&nama,&umur)
		fmt.Print("	Nama Pelanggan : ")
		fmt.Scan(&nama)
		fmt.Println()
		fmt.Print("	Umur Pelanggan : ")
		fmt.Scan(&umur)
		fmt.Println()
		index=ada4(*C,umur,nama)
		if index==-1{
			fmt.Println("Data tidak ditemukan")
		}else {
			for i:=index;i<C.nCustomers-1;i++{
				C.customers[i]=C.customers[i+1]
			}
			C.nCustomers=C.nCustomers-1
		}
		fmt.Println("Data Baru:")
		menampilkanCustomer(*C)
		
		
	
	}else if pilihan=="3"{
		Header223()
		menampilkanSparepart1(*S)
		//fmt.Println("Tentukan data mana yang mau diubah namaSparePart merekSparePart hargaBeli tarifServis jumlahPadaServis ")
		//fmt.Scan(&namaSparePart,&merekSparePart,&hargaSparepart,&tarifSparepart,&jmlSparepart)
		fmt.Print("	Nama Spare-part : ")
		fmt.Scan(&namaSparePart)
		fmt.Println()
		fmt.Print("	Merek Spare-part :")
		fmt.Scan(&merekSparePart)
		fmt.Println()
		fmt.Print("	Harga Beli Spare-part : ")
		fmt.Scan(&hargaSparepart)
		fmt.Println()
		fmt.Print("	Tarif Servis Spare-part : ")
		fmt.Scan(&tarifSparepart)
		fmt.Println()
		fmt.Print("	Total Spare-part yang telah digunakan : ")
		fmt.Scan(&jmlSparepart)
		fmt.Println()
		index=cariDataParts(*S,namaSparePart,merekSparePart,hargaSparepart,tarifSparepart,jmlSparepart)
		if index==-1{
			fmt.Println("Data tidak ditemukan")
		
		}else {
			for i:=index;i<S.nSparts;i++{
				S.sParts[i]=S.sParts[i+1]
			}
			S.nSparts=S.nSparts-1
			fmt.Println("Data Baru :")
			menampilkanSparepart1(*S)
		
		}
	
	}
	if pilihan !="4"{
		Header22()
		fmt.Println("Pilihan Anda :")
		fmt.Scan(&pilihan)
	
	}
	
	}



}else if pilihan !="3"{
	Header2()
fmt.Println("Pilihan Anda :")
		fmt.Scan(&pilihan)
}
	//1 ubah historiTransaksi
	//2 ubah dataPelanggan
	//3 ubah sparePart
}
}
func cariDataParts(S spareParts,nama,merek string,harga,tarif float64,jservis int) int {
	var index int=-1
	var i int =0
	for i<S.nSparts && index==-1{
		if S.sParts[i].namaSparePart==nama&&S.sParts[i].merekSparePart==merek&&S.sParts[i].hargaBeli==harga&&S.sParts[i].tarifServis==tarif&&S.sParts[i].jumlahPadaServis==jservis {
			index=i
			
		
		}
		//fmt.Println(S.sParts[i].namaSparePart==nama,S.sParts[i].merekSparePart==merek,S.sParts[i].hargaBeli==harga,S.sParts[i].tarifServis==tarif,S.sParts[i].jumlahPadaServis==jservis,S.sParts[i].tarifServis,tarif )
		i=i+1
	}
	return index
}
func cariData(T transanctionRecords,nama,namaSparePart,merekSparePart string,umur,jServis,tgl,bln,thn,jmlSparepart int,hargaSparepart,tarifSparepart,tarif float64) int {
	var index =-1
	var i int
	for i<T.nTransaksi && index==-1{
		if T.hTransaksi[i].pelanggan.nama==nama&&T.hTransaksi[i].pelanggan.umur==umur&&T.hTransaksi[i].sukuCadang.namaSparePart==namaSparePart&&T.hTransaksi[i].sukuCadang.merekSparePart==merekSparePart&&T.hTransaksi[i].sukuCadang.hargaBeli==hargaSparepart&&T.hTransaksi[i].sukuCadang.tarifServis==tarifSparepart&&T.hTransaksi[i].sukuCadang.jumlahPadaServis==jServis&&T.hTransaksi[i].tarif==tarif&&T.hTransaksi[i].jumlahSparepart==jmlSparepart&&T.hTransaksi[i].waktuTransaksi.tanggal==tgl&&T.hTransaksi[i].waktuTransaksi.bulan==bln&&T.hTransaksi[i].waktuTransaksi.tahun==thn{
			index=i
		}
		i=i+1
	}
	return index

}

func menampilkanTransaksi(T transanctionRecords) {
	
	for i:=0;i<T.nTransaksi;i++{
		fmt.Println("Data ke: ",i+1)
		fmt.Printf("Nama: %s Umur: %d\n",T.hTransaksi[i].pelanggan.nama,T.hTransaksi[i].pelanggan.umur)
		fmt.Printf("Nama Spare-part: %s Merek Sparepart: %s\n",T.hTransaksi[i].sukuCadang.namaSparePart,T.hTransaksi[i].sukuCadang.merekSparePart)
		fmt.Printf("Harga Beli: %.3f Tarif Servis Sparepart %.3f Total Servis Spare-part: %d\n" ,T.hTransaksi[i].sukuCadang.hargaBeli,T.hTransaksi[i].sukuCadang.tarifServis,T.hTransaksi[i].sukuCadang.jumlahPadaServis)
		fmt.Printf("Tarif Total: %.3f Jumlah Spare-part yang digunakan: %d\n",T.hTransaksi[i].tarif,T.hTransaksi[i].jumlahSparepart)
		fmt.Printf("Waktu(tanggal,bulan,tahun): %d-%d-%d\n",T.hTransaksi[i].waktuTransaksi.tanggal,T.hTransaksi[i].waktuTransaksi.bulan,T.hTransaksi[i].waktuTransaksi.tahun)
		fmt.Println()
	}
}
func menampilkanTransaksi1(T transanctionRecords,n int) {
	
		var i int =n
		fmt.Println("Data ke: ",i+1)
		fmt.Printf("Nama: %s Umur: %d\n",T.hTransaksi[i].pelanggan.nama,T.hTransaksi[i].pelanggan.umur)
		fmt.Printf("Nama Spare-part: %s Merek Sparepart: %s\n",T.hTransaksi[i].sukuCadang.namaSparePart,T.hTransaksi[i].sukuCadang.merekSparePart)
		fmt.Printf("Harga Beli: %.3f Tarif Servis Sparepart %.3f Total Servis Spare-part: %d\n" ,T.hTransaksi[i].sukuCadang.hargaBeli,T.hTransaksi[i].sukuCadang.tarifServis,T.hTransaksi[i].sukuCadang.jumlahPadaServis)
		fmt.Printf("Tarif Total: %.3f Jumlah Spare-part yang digunakan: %d\n",T.hTransaksi[i].tarif,T.hTransaksi[i].jumlahSparepart)
		fmt.Printf("Waktu(tanggal,bulan,tahun): %d-%d-%d\n",T.hTransaksi[i].waktuTransaksi.tanggal,T.hTransaksi[i].waktuTransaksi.bulan,T.hTransaksi[i].waktuTransaksi.tahun)
		fmt.Println()
	
}
//
//
//
//

func mencariData(T transanctionRecords) {
	var pilihan,namaSparePart,merekSparePart string
	var tgl,bln,thn,tgl2,bln2,thn2 int
	Header3()
	//fmt.Println("1. Mencari Data berdasarkan Waktu")
	//fmt.Println("2. Mencari data pelnggan berdasarkan nama sparepart")
	fmt.Println("Pilihan Anda :")
	fmt.Scan(&pilihan)
	for pilihan!="3"{
		if pilihan=="1"{
		Header31()
		//fmt.Println("1.Pencarian Data Dengan Waktu Tertentu")
		//fmt.Println("2.Pencarian Data Denngan Rentang Waktu Tertentu")
		fmt.Println("Pilihan Anda :")
		fmt.Scan(&pilihan)
	
		if pilihan=="1"{
			Header311()
			menampilkanTransaksi(T)
			fmt.Print("	Tanggal Transaksi :	")
			fmt.Scan(&tgl)
			fmt.Println()
			fmt.Print("	Bulan Transaksi	  :	")
			fmt.Scan(&bln)
			fmt.Println()
			fmt.Print("	Tahun Transaksi	  :	")
			fmt.Scan(&thn)
			fmt.Println()
			//fmt.Println("Tolong inputkan waktu(tanggal,bulan,tahun) yang ingin dicari")
			//fmt.Scan(&tgl,&bln,&thn)
			cariDataTanggal2(T,tgl,bln,thn)
		
		}else if pilihan=="2"{
			Header312()
			menampilkanTransaksi(T)
			fmt.Print("	Tanggal Transaksi Sesudah :	")
			fmt.Scan(&tgl)
			fmt.Println()
			fmt.Print("	Bulan Transaksi	Sesudah  :	")
			fmt.Scan(&bln)
			fmt.Println()
			fmt.Print("	Tahun Transaksi	Sesudah :	")
			fmt.Scan(&thn)
			fmt.Println()
			fmt.Print("	Tanggal Transaksi Sebelum :	")
			fmt.Scan(&tgl2)
			fmt.Println()
			fmt.Print("	Bulan Transaksi	 Sebelum :	")
			fmt.Scan(&bln2)
			fmt.Println()
			fmt.Print("	Tahun Transaksi	 Sebelum :	")
			fmt.Scan(&thn2)
			fmt.Println()
			//fmt.Println("Tolong inputkan Rentang waktu(tanggal,bulan,tahun) yang ingin dicari")
			//fmt.Println("Rentang Waktu Sesudah :")
			//fmt.Scan(&tgl,&bln,&thn)
			//fmt.Println("Rentang Waktu Sebelum :")
			//fmt.Scan(&tgl2,&bln2,&thn2)
			cariRentangWaktu(T,tgl,bln,thn,tgl2,bln2,thn2) 
			//data pelanggan ya 
		}else {
			fmt.Println("Pilihan Tidak Sesuai")
		}
		}else if pilihan=="2"{
		Header32()
		//fmt.Println("1.Mencari Pelanggan Yang menggunakan Nama Sparepart tertentu")
		//fmt.Println("2.Mencari Pelanggan Yang menggunakan Nama Sparepart dan Merek Sparepart tertentu")
		fmt.Println("Pilihan Anda :")
		fmt.Scan(&pilihan)
		if pilihan =="1"{
			Header321()
			menampilkanTransaksi(T)
			//fmt.Println("Silahkan Masukan Nama Spare Part:")
			fmt.Print("	Nama Spare-part : ")
			fmt.Scan(&namaSparePart)
			cariNamaSparepart(T,namaSparePart)
		
		}else if pilihan=="2"{
			Header322()
			menampilkanTransaksi(T)
			//fmt.Println("Silahkan Masukan Nama Sparepart dan merek Sparepart:")
			//fmt.Scan(&namaSparePart,&merekSparePart)
			fmt.Print("	Nama Spare-part : ")
			fmt.Scan(&namaSparePart)
			fmt.Println()
			fmt.Print("	Merek Spare-part : ")
			fmt.Scan(&merekSparePart)
			fmt.Println()
			cariNamaMerekSparepart(T,namaSparePart,merekSparePart)
		}
	}
	if pilihan!="3"{
		Header3()
		fmt.Println("Pilihan Anda : ")
		fmt.Scan(&pilihan)
	}
	}
	//1mencari data transaksi dengan waktu tertentu
	//2mencari data pelanggan yang melakukan membeli sparepart tertentu
}
func cariNamaMerekSparepart(T transanctionRecords,nama,merek string){
	var x bool
		fmt.Println()
		fmt.Println("       ","Data Pelanggan :")
		for i:=0;i<T.nTransaksi;i++{
			if nama==T.hTransaksi[i].sukuCadang.namaSparePart&&merek==T.hTransaksi[i].sukuCadang.merekSparePart{
				fmt.Println("       ","Nama : ",T.hTransaksi[i].pelanggan.nama," ","Umur : ",T.hTransaksi[i].pelanggan.umur)
				x=true
			}
		}
	if !x{
		fmt.Println("Data tidak ditemukan")
	}
}
func cariNamaSparepart(T transanctionRecords,nama string) {
	var x bool
		fmt.Println()
		fmt.Println("       ","Data Pelanggan :")
		for i:=0;i<T.nTransaksi;i++{
			if nama==T.hTransaksi[i].sukuCadang.namaSparePart{
			
				fmt.Println("       ","Nama : ",T.hTransaksi[i].pelanggan.nama," ","Umur : ",T.hTransaksi[i].pelanggan.umur)
				x=true
			}
		}
	if !x{
		fmt.Println("Data tidak ditemukan")
	}
}
func cariRentangWaktu (T transanctionRecords,tgl,bln,thn,tgl2,bln2,thn2 int) {
	var total int
	var x bool
	fmt.Println()
	fmt.Println("       ","Data Pelanggan :")
	for i:=0;i<T.nTransaksi;i++{
		total=(T.hTransaksi[i].waktuTransaksi.tanggal+((T.hTransaksi[i].waktuTransaksi.bulan)*31)+((T.hTransaksi[i].waktuTransaksi.tahun)*366))
		if	total>(tgl+(bln*31)+(thn*366))&&total<(tgl2+(bln2*31)+(thn2*366)){
			fmt.Println("       ","Nama : ",T.hTransaksi[i].pelanggan.nama," ","Umur : ",T.hTransaksi[i].pelanggan.umur)
			x=true
		
		}
		
	}
	if !x{
		fmt.Println("Data tidak ditemukan")
	}
	fmt.Println()

}
func cariDataTanggal2(T transanctionRecords,tgl,bln,thn int){
	var total,total2 int 
	var x bool
	total2=(tgl+(bln*31)+(thn*366))
	fmt.Println()
	fmt.Println("       ","Data Pelanggan :")
	for i:=0;i<T.nTransaksi;i++{
		total=(T.hTransaksi[i].waktuTransaksi.tanggal+((T.hTransaksi[i].waktuTransaksi.bulan)*31)+((T.hTransaksi[i].waktuTransaksi.tahun)*366))
		if total==total2{
			fmt.Println("       ","Nama : ",T.hTransaksi[i].pelanggan.nama," ","Umur : ",T.hTransaksi[i].pelanggan.umur)
			x=true
		}
	
	}
	
	if !x{
		fmt.Println("Data tidak ditemukan")
	}
	fmt.Println()
}
/*
func cariDataTanggal(T transanctionRecords,tgl,bln,thn) int {
	var mid ,l,r,index int
	index=-1
	l=0
	r=T.nTransaksi-1
	mid=(l+r)/2
	for l<=r&&index==-1{
		if (tgl+(bln*31)+thn*366)<(T.hTransaksi[mid].waktuTransaksi.tanggal+((T.hTransaksi[mid].waktuTransaksi.bulan)*31)+((T.hTransaksi[mid].waktuTransaksi.tahun)*366)){
			r=mid-1
		}else if (tgl+(bln*31)+thn*366)>(T.hTransaksi[mid].waktuTransaksi.tanggal+((T.hTransaksi[mid].waktuTransaksi.bulan)*31)+((T.hTransaksi[mid].waktuTransaksi.tahun)*366)){
			l=mid+1
		
		}else {
			index=mid
		}
	
	}
	return index

}
*/
func sortingData(T *transanctionRecords){
	var temp transanctionRecords
	var i int
	for j:=1;j<T.nTransaksi;j++{
		temp.hTransaksi[0]=T.hTransaksi[j]
		i=j
		//	fmt.Println(T.hTransaksi[i])
		//fmt.Println()
		//fmt.Println(((T.hTransaksi[i].waktuTransaksi.tanggal)+((T.hTransaksi[i].waktuTransaksi.bulan)*31)+((T.hTransaksi[i].waktuTransaksi.tahun)*366)),((T.hTransaksi[i-1].waktuTransaksi.tanggal)+((T.hTransaksi[i-1].waktuTransaksi.bulan)*31)+((T.hTransaksi[i-1].waktuTransaksi.tahun)*366)))
		//fmt.Println(((T.hTransaksi[i].waktuTransaksi.tanggal)+((T.hTransaksi[i].waktuTransaksi.bulan)*31)+((T.hTransaksi[i].waktuTransaksi.tahun)*366))<((T.hTransaksi[i-1].waktuTransaksi.tanggal)+((T.hTransaksi[i-1].waktuTransaksi.bulan)*31)+((T.hTransaksi[i-1].waktuTransaksi.tahun)*366)))
		for i>=1 && ((temp.hTransaksi[0].waktuTransaksi.tanggal)+((temp.hTransaksi[0].waktuTransaksi.bulan)*31)+((temp.hTransaksi[0].waktuTransaksi.tahun)*366))<((T.hTransaksi[i-1].waktuTransaksi.tanggal)+((T.hTransaksi[i-1].waktuTransaksi.bulan)*31)+((T.hTransaksi[i-1].waktuTransaksi.tahun)*366)){
			
			
			T.hTransaksi[i]=T.hTransaksi[i-1]
			i=i-1
		}
		T.hTransaksi[i]=temp.hTransaksi[0]
		
	}

}

func menampilkanSparepart1 (S spareParts) {
//1menampilkan sparepart yang sering digunakan
//2menampilkan sparepart dengan jumlah tertentu 
	for i:=0;i<S.nSparts;i++{
		fmt.Println("Data ke ",i+1)
		fmt.Printf("Nama Spare-part: %s Merek Spare-part: %s\n",S.sParts[i].namaSparePart,S.sParts[i].merekSparePart)
		fmt.Printf("Harga Beli Spare-part: %.3f Tarif Servis Spare-part: %.3f\n",S.sParts[i].hargaBeli,S.sParts[i].tarifServis)
		fmt.Printf("Total Spare-part Yang Telah Digunakan Dalam Servis: %d\n",S.sParts[i].jumlahPadaServis)
		fmt.Println()
	}
	
}

func menampilkanCustomer (C customerData) {
	for i:=0;i<C.nCustomers;i++{
		fmt.Println("Data ke ",i+1)
		fmt.Printf("Nama: %s Umur: %d\n",C.customers[i].nama,C.customers[i].umur)
	}
	


}

func cekValidWaktu(tgl,bln,thn int) bool {
	var jmlHari int
	var x bool =false
	if bln==1||bln==3||bln==5||bln==7||bln==8||bln==10||bln==12 {
		jmlHari=31
		x=true
	}else if bln==4||bln==6||bln==9||bln==11{
		jmlHari=30
		x=true
	}else if bln==2{
		if thn%400==0 || (thn%100!=0 && thn%4==0){
			jmlHari=29
		}else{
			jmlHari=28
		}
		x=true
	}
	//fmt.Println(jmlHari<=tgl && tgl>0 && bln>0 && thn>=0," Q ",jmlHari,tgl,jmlHari<=tgl , tgl>0 , bln>0 , thn>=0)
	return tgl>0 && tgl<=jmlHari && bln>0 && thn>=0&&x
}



func Header() {
	fmt.Println("=========================================================")
	fmt.Println("| 	       Aplikasi Servis Motor			|")
	fmt.Println("=========================================================")
	fmt.Println("|		     Kelompok 3			        |")
	fmt.Println("|	  Sutan Rifky Tedjasukmana 1301223465		|")
	fmt.Println("|	  Eka Setyanto Harsono 1301223485		|")
	fmt.Println("|	 	      IF-46-07	 		 	|")
	fmt.Println("|=======================================================|")
	fmt.Println("| 			MENU			   	|")
	fmt.Println("|=======================================================|")
	fmt.Println("|Silahkan Pilih Pilihan Diantara Berikut  	  	|")
	fmt.Println("|Pilihan :						|")
	fmt.Println("|1.Menambah atau Menginput Data				|")
	fmt.Println("|2.Mengubah atau Menghapus Data				|")
	fmt.Println("|3.Mencari Data Pelanggan				|")
	fmt.Println("|4.Mengurutkan Data Spare-part			        |")
	fmt.Println("|5.Menampilkan	Data					|")
	fmt.Println("|6.Selesai						|")
	fmt.Println("|							|")
	fmt.Println("=========================================================")
	
}
func Header1(){
	fmt.Println("========================================================")
	fmt.Println("| 	        Aplikasi Servis Motor		       |")
	fmt.Println("| 		     Menambah Data		       |")
	fmt.Println("========================================================")
	fmt.Println("|Silakan Tentukan Jenis Data yang Ingin Anda Tambahkan |")
	fmt.Println("|Pilihan : 					       |")
	fmt.Println("|1.Histori Transaksi				       |")
	fmt.Println("|2.Data Pelanggan				       |")
	fmt.Println("|3.Spare-part					       |")
	fmt.Println("|4.Selesai					       |")
	fmt.Println("========================================================")
}
func Header11(){
	
	fmt.Println(" 	        Aplikasi Servis Motor		       ")
	fmt.Println(" 		     Menambah Data		       ")
	fmt.Println(" 		   Histori Transaksi		       ")
	fmt.Println("       ========================================         ")
	fmt.Println("	Ketik 'Selesai' pada Nama dan 0 Pada Umur Untuk Selesai	")
	
	
}
func Header12(){
	fmt.Println(" 	        Aplikasi Servis Motor		       ")
	fmt.Println(" 		     Menambah Data		       ")
	fmt.Println(" 		   Data Pelanggan		       ")
	fmt.Println("       ========================================         ")
	fmt.Println("	Ketik 'Selesai' pada Nama dan 0 Pada Umur Untuk Selesai	")




}
func Header13() {
	fmt.Println(" 	        Aplikasi Servis Motor		       ")
	fmt.Println(" 		     Menambah Data		       ")
	fmt.Println(" 		   Data Spare-part		       ")
	fmt.Println("       ========================================         ")
	fmt.Println("	Ketik 'Selesai' pada Nama dan 'Selesai' Pada Merek dari Spare-part Untuk Selesai	")

	
}
func Header2(){
	fmt.Println("========================================================")
	fmt.Println("| 	        Aplikasi Servis Motor		       |")
	fmt.Println("| 	     Mengubah atau Menghapus Data	       |")
	fmt.Println("========================================================")
	fmt.Println("|Silakan Tentukan Diantara Opsi Berikut 	       |")
	fmt.Println("|Pilihan : 					       |")
	fmt.Println("|1.Mengubah Data				       |")
	fmt.Println("|2.Menghapus Data				       |")
	fmt.Println("|3.Selesai					       |")
	fmt.Println("========================================================")
}
func Header21() {
	fmt.Println(" 	        Aplikasi Servis Motor		       ")
	fmt.Println(" 	    Mengubah atau Menghapus Data		       ")
	fmt.Println(" 		   Mengubah Data		       ")
	fmt.Println("       ======================================           ")
	fmt.Println("	1.Histori Transaksi	")
	fmt.Println("	2.Pelanggan	")
	fmt.Println("	3.Spare-part ")
	fmt.Println("	4.Selesai")
	
}
func Header211() {
	fmt.Println(" 	        Aplikasi Servis Motor		       ")
	fmt.Println(" 	    Mengubah atau Menghapus Data		       ")
	fmt.Println(" 		   Mengubah Data		       ")
	fmt.Println("		 Histori Transaksi")
	fmt.Println("       ======================================           ")

}
func Header212() {
	fmt.Println(" 	        Aplikasi Servis Motor		       ")
	fmt.Println(" 	    Mengubah atau Menghapus Data		       ")
	fmt.Println(" 		   Mengubah Data		       ")
	fmt.Println("		   Data Pelanggan")
	fmt.Println("       ======================================           ")
	


}
func Header213() {
	fmt.Println(" 	        Aplikasi Servis Motor		       ")
	fmt.Println(" 	    Mengubah atau Menghapus Data		       ")
	fmt.Println(" 		   Mengubah Data		       ")
	fmt.Println("		  Data Spare-part")
	fmt.Println("       ======================================           ")

}
func Header22() {
	
	fmt.Println(" 	        Aplikasi Servis Motor		       ")
	fmt.Println(" 	    Mengubah atau Menghapus Data		       ")
	fmt.Println(" 		   Menghapus Data		       ")
	fmt.Println("       ======================================           ")
	fmt.Println("	1.Histori Transaksi	")
	fmt.Println("	2.Pelanggan	")
	fmt.Println("	3.Spare-part ")
	fmt.Println("	4.Selesai")
	


}
func Header221() {
	fmt.Println(" 	        Aplikasi Servis Motor		       ")
	fmt.Println(" 	    Mengubah atau Menghapus Data		       ")
	fmt.Println(" 		   Menghapus Data		       ")
	fmt.Println("		 Histori Transaksi")
	fmt.Println("       ======================================           ")

}
func Header222() {
	
	fmt.Println(" 	        Aplikasi Servis Motor		       ")
	fmt.Println(" 	    Mengubah atau Menghapus Data		       ")
	fmt.Println(" 		   Menghapus Data		       ")
	fmt.Println("		   Data Pelanggan")
	fmt.Println("       ======================================           ")

}
func Header223() {
	fmt.Println(" 	        Aplikasi Servis Motor		       ")
	fmt.Println(" 	    Mengubah atau Menghapus Data		       ")
	fmt.Println(" 		   Menghapus Data		       ")
	fmt.Println("		   Data Spare-part")
	fmt.Println("       ======================================           ")

}
func Header3(){
	
	fmt.Println("========================================================")
	fmt.Println("| 	        Aplikasi Servis Motor		       |")
	fmt.Println("| 	        Mencari Data Pelanggan	               |")
	fmt.Println("========================================================")
	fmt.Println("|Silakan Tentukan Diantara Opsi Berikut 	       |")
	fmt.Println("|Pilihan : 					       |")
	fmt.Println("|1.Mencari Data Berdasarkan Waktu 		       |")
	fmt.Println("|2.Mencari Data Berdasarkan Spare-part	       	       |")
	fmt.Println("|3.Selesai					       |")
	fmt.Println("========================================================")



}
func Header31(){
	
	
	fmt.Println(" 	        Aplikasi Servis Motor		       ")
	fmt.Println(" 	        Mencari Data Pelanggan		       ")
	fmt.Println(" 	      Mencari Berdasarkan Waktu 		       ")
	fmt.Println("       ======================================           ")
	fmt.Println("	1.Berdasarkan Waktu Tertentu	")
	fmt.Println("	2.Berdasarkan Rentang Waktu Tertentu	")
	fmt.Println("	3.Selesai ")
	




}
func Header311(){
	
	
	fmt.Println(" 	        Aplikasi Servis Motor		       ")
	fmt.Println(" 	        Mencari Data Pelanggan		       ")
	fmt.Println(" 	      Mencari Berdasarkan Waktu 		       ")
	fmt.Println(" 	    	    Waktu Tertentu		       ")
	fmt.Println("       ======================================           ")
	
}
func Header312(){
	
	
	fmt.Println(" 	        Aplikasi Servis Motor		       ")
	fmt.Println(" 	        Mencari Data Pelanggan		       ")
	fmt.Println(" 	      Mencari Berdasarkan Waktu 		       ")
	fmt.Println(" 	        Rentang Waktu Tertentu		       ")
	fmt.Println("       ======================================           ")
	
}
func Header32(){
	fmt.Println(" 	        Aplikasi Servis Motor		       ")
	fmt.Println(" 	        Mencari Data Pelanggan		       ")
	fmt.Println(" 	    Mencari Berdasarkan Spare-part 		       ")
	fmt.Println("       ======================================           ")
	fmt.Println("	1.Berdasarkan Nama Spare-part 	")
	fmt.Println("	2.Berdasarkan Nama dan Merek Spare-part	")
	fmt.Println("	3.Selesai ")


}
func Header321() {
	fmt.Println(" 	        Aplikasi Servis Motor		       ")
	fmt.Println(" 	        Mencari Data Pelanggan		       ")
	fmt.Println(" 	    Mencari Berdasarkan Spare-part 		       ")
	fmt.Println("	            Nama Spare-part	")
	fmt.Println("       ======================================           ")

}
func Header322() {
	fmt.Println(" 	        Aplikasi Servis Motor		       ")
	fmt.Println(" 	        Mencari Data Pelanggan		       ")
	fmt.Println(" 	    Mencari Berdasarkan Spare-part 		       ")
	fmt.Println("	       Nama dan Merek Spare-part	")
	fmt.Println("       ======================================           ")

}
func Header4(){
	
	fmt.Println("========================================================")
	fmt.Println("| 	        Aplikasi Servis Motor		       |")
	fmt.Println("| 	     Mengurutkan Data Spare-part  	       |")
	fmt.Println("========================================================")
	fmt.Println("|Data Diurutkan Berdasarkan Jumlah Servis     	       |")
	fmt.Println("|Silakan Tentukan Diantara Opsi Berikut 	       |")
	fmt.Println("|Pilihan : 					       |")
	fmt.Println("|1.Diurutkan Dari Tertinggi Ke Terkecil                |")
	fmt.Println("|2.Diurutkan Dari Terkecil ke Tertinggi	       	       |")
	fmt.Println("|3.Selesai					       |")
	fmt.Println("========================================================")



}
func Header41() {
	fmt.Println(" 	        Aplikasi Servis Motor		       ")
	fmt.Println(" 	        Mencari Data Pelanggan		       ")
	fmt.Println(" 	     Mengurutkan Data Spare-part  	       ")
	fmt.Println("	       	      Descending	")
	fmt.Println("       ======================================           ")
	fmt.Println("Data Terurut :")
	fmt.Println()
}
func Header42() {
	fmt.Println(" 	        Aplikasi Servis Motor		       ")
	fmt.Println(" 	        Mencari Data Pelanggan		       ")
	fmt.Println(" 	     Mengurutkan Data Spare-part  	       ")
	fmt.Println("	       	      Ascending	")
	fmt.Println("       ======================================           ")
	fmt.Println("Data Terurut :")
	fmt.Println()
}
func Header5() {
		fmt.Println("========================================================")
	fmt.Println("| 	        Aplikasi Servis Motor		       |")
	fmt.Println("| 	        Menampilkan Data 	               |")
	fmt.Println("========================================================")
	fmt.Println("|Silakan Tentukan Diantara Opsi Berikut 	       |")
	fmt.Println("|Pilihan : 					       |")
	fmt.Println("|1.Histori Transaksi    		               |")
	fmt.Println("|2.Data Pelanggan          	       		       |")
	fmt.Println("|3.Spare-part           	       		       |")
	fmt.Println("|4.Selesai					       |")
	fmt.Println("========================================================")

}
func Header51() {
	fmt.Println(" 	        Aplikasi Servis Motor		       ")
	fmt.Println(" 	          Menampilkan Data 		       ")
	fmt.Println(" 	          Histori Transaksi 		       ")
	fmt.Println("       ======================================           ")
	fmt.Println("Data Histori Transaksi :")
}
func Header52() {
	fmt.Println(" 	        Aplikasi Servis Motor		       ")
	fmt.Println(" 	          Menampilkan Data 		       ")
	fmt.Println(" 	           Data Pelanggan 		       ")
	fmt.Println("       ======================================           ")
	fmt.Println("Data Pelanggan  :")
}
func Header53() {
	fmt.Println(" 	        Aplikasi Servis Motor		       ")
	fmt.Println(" 	          Menampilkan Data 		       ")
	fmt.Println(" 	             Spare-part  		       ")
	fmt.Println("       ======================================           ")
	fmt.Println("Data Spare-part  :")
}
/*
1
1
Ahmad 29
radiator honda
3500 10000
5 20 5 2005 
juber 29 
Sekring yamaha
23000 2000000
10 2 5 2004
Siti 32
Stang Harley 
2300 2430000
10 10 10 2010
Boby 22
Oli Honda
234 3463453
5 21 5 2005
juyun 23
Oli Ferrari 
232 3232333
10 20 5 2005
Jabrig 33
Oli Honda
25 20 10 2003
Ahsan 22
Sekring yamaha
10 12 12 2012
Selesai 0


Ahmad 32
radiator honda
3500 10000
5 20 5 2005 
Selesai

Ahmad 29 radiator honda 3500 10000 5 50000 5 20 5 2005
Ahmad 29 Oli Honda 3500 10000 5 50000 5 20 5 2005
//
//fmt.Println(((T.hTransaksi[i].waktuTransaksi.tanggal)+((T.hTransaksi[i].waktuTransaksi.bulan)*31)+((T.hTransaksi[i].waktuTransaksi.tahun)*366))<((T.hTransaksi[i-1].waktuTransaksi.tanggal)+((T.hTransaksi[i-1].waktuTransaksi.bulan)*31)+((T.hTransaksi[i-1].waktuTransaksi.tahun)*366)))
*/
/*
INGATTTTTTTTTTT
!!!!!!!!!!!!!!!!!
!!!!!!!!!!!!!!!!!


BUAT KONDISI DI MENGUBAH DATA DAN YANG LAIN JIKA PERLU 
KONDISI DIMANA DATA BELUM ADA 
BERARTI JIKA MASIH KOSONG
PENGGUNA PERLU MENGISI DATA YAITU PILIHAN 1
INGATTTTTTTTTTT 
!!!!!!!!!!
!!!!!!!!!
!!!!!!!!
!!!!!!!
!!!!!!
!!!!!
!!!!
!!!
!!
!

Baru !!!!!!!!
tambahin kondisi 
kalau data yang diubah dalam data histori transak si adalah tarif servis atau banyak servis yang dipakai maka 
lakukan kembali perhitungan tari f ama banyak nya jumlah pada servis
*/
/*
		for i:=0;i<T.nTransaksi;i++{
			fmt.Println(T.hTransaksi[i])
		}
		fmt.Println(i)
		fmt.Println()
		if i==1{
				
				if(temp.hTransaksi[i].waktuTransaksi.tanggal+((temp.hTransaksi[i].waktuTransaksi.bulan)*31)+((temp.hTransaksi[i].waktuTransaksi.tahun)*366))<(T.hTransaksi[i-1].waktuTransaksi.tanggal+((T.hTransaksi[i-1].waktuTransaksi.bulan)*31)+((T.hTransaksi[i-1].waktuTransaksi.tahun)*366)){
				T.hTransaksi[i]=T.hTransaksi[i-1]
				
			
			}
			
			}
			fmt.Println("Baru")
			/*for i:=0;i<T.nTransaksi;i++{
			fmt.Println(T.hTransaksi[i])
			}
			fmt.Println()
			*/
