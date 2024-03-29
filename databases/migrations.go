package databases

type Reservasi struct {
	IdReservasi  		int    `gorm:"primaryKey;autoIncrement" json:"id_reservasi"`
	TglKunjungan 		string `gorm:"column:tgl_kunjungan" json:"tgl_kunjungan"`
	Nama         		string `gorm:"column:nama" json:"nama"`
	NoRS         		int    `gorm:"column:no_rs" json:"no_rs"`
	Poli         		string `gorm:"column:poli" json:"poli"`
	Dokter       		string `gorm:"column:dokter" json:"dokter"`
	Pembayaran   		string `gorm:"column:pembayaran" json:"pembayaran"`
	NoAsuransi   		int    `gorm:"column:no_asuransi" json:"no_asuransi"`
	NIK          		int    `gorm:"column:nik" json:"nik"`
	TglLahir     		string `gorm:"column:tgl_lahir" json:"tgl_lahir"`
	Email        		string `gorm:"column:email" json:"email"`
	JenisKelamin 		string `gorm:"column:jenis_kelamin" json:"jenis_kelamin"`
	NoTelp       		string `gorm:"column:no_telp" json:"no_telp"`
	RecaptchaResponse 	string `json:"recaptchaResponse"`
}

type ProfilDokter struct {
	IdDokter      int    		`gorm:"primaryKey" json:"id_dokter"`
	Nama          string 		`gorm:"column:nama_dokter" json:"nama_dokter"`
	Poli		  string 		`gorm:"column:poli" json:"poli"`
	Gambar        []byte 		`gorm:"column:foto_dokter" json:"gambar"`
	EncodedGambar string 		`json:"-"`
	JadwalDokter  []JadwalDokter`gorm:"foreignKey:IdDokter" json:"JadwalDokter"`
}

type JadwalDokter struct {
	IdJadwal     string `gorm:"primaryKey" json:"id_jadwal"`
	IdDokter     int    `gorm:"column:id_dokter" json:"id_dokter"`
	Hari_Praktek string `gorm:"column:hari_praktek" json:"hari"`
	JamMulai     string `gorm:"column:jam_mulai" json:"jam_mulai"`
	JamSelesai   string `gorm:"column:jam_selesai" json:"jam_selesai"`
}

type User struct {
	NoRS     	 int 	`gorm:"primaryKey;autoIncrement" json:"no_rs"`
	NIK      	 string `gorm:"column:nik" json:"nik"`
	Password 	 string `gorm:"column:password" json:"-"`
	Nama         string `gorm:"column:nama" json:"nama"`
	TglLahir     string `gorm:"column:tgl_lahir" json:"tgl_lahir"`
	JenisKelamin string `gorm:"column:jenis_kelamin" json:"jenis_kelamin"`
}

func (Reservasi) TableName() string {
	return "reservasi"
}

func (ProfilDokter) TableName() string {
	return "profil_dokter"
}

func (JadwalDokter) TableName() string {
	return "jadwal_dokter"
}

func (User) TableName() string {
	return "akun_user"
}