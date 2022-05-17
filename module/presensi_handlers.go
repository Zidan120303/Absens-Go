package module

import (
	"APLIKASI_1/model"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetPresen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var presen []model.Presen
	DB.Where("id_pengguna = ? ", params["id"]).Find(&presen)
	var response = model.PresenJResponses{Type: true, Message: "Success Get data presen by id pengguna", Data: presen}
	json.NewEncoder(w).Encode(response)
}

func Presen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var presen model.Presen
	var response = model.PresenJResponses{}
	json.NewDecoder(r.Body).Decode(&presen)

	Sekarang := time.Now()
	Id_pengguna := presen.Id_pengguna
	In := presen.Lokasi_in
	Out := presen.Lokasi_out
	Status := "Working"

	rows, _ := DB.Raw("select (cast( ? as date))-max((cast(time_in as date))) as time_in from presens where id_pengguna=?", Sekarang, Id_pengguna).Rows()
	for rows.Next() {
		var Time int16
		rows.Scan(&Time)
		if Time <= 1 {
			//cek query id record tanggal terakhir ada atau tidak
			result2 := DB.Where("id_pengguna = ? AND TO_CHAR(time_in, 'YYYY-MM-DD') = LEFT(TO_CHAR(CURRENT_TIMESTAMP, 'YYYY-MM-DD') , 10)", Id_pengguna).First(&presen)
			if result2.RowsAffected > 0 {
				//ambil data id dan tgl hari ini
				DB.Where("id_pengguna = ? AND TO_CHAR(time_in, 'YYYY-MM-DD') = LEFT(TO_CHAR(CURRENT_TIMESTAMP, 'YYYY-MM-DD') , 10)", Id_pengguna).First(&presen)
				presen.Time_out = Sekarang
				presen.Lokasi_out = Out
				DB.Save(&presen)
				response = model.PresenJResponses{Type: true, Message: "success, Check out"}
			} else {
				presen.Time_in = Sekarang
				presen.Lokasi_in = In
				presen.Status = Status
				DB.Select("id_pengguna", "time_in", "lokasi_in", "status").Create(&presen)
				response = model.PresenJResponses{Type: true, Message: "Success, Check in"}
			}
		} else {
			i := (-24 * (Time - 1))
			for ; i < 0; i += 24 {
				Kemaren := Sekarang.Add(time.Duration(i) * time.Hour).UTC()
				presen.Time_in = Kemaren
				presen.Lokasi_in = "Null"
				presen.Time_out = Kemaren
				presen.Lokasi_out = "Null"
				presen.Status = "Alpa"
				DB.Select("id_pengguna", "time_in", "lokasi_in", "time_out", "lokasi_out", "status").Create(&presen)
				response = model.PresenJResponses{Type: false, Message: "Sorry please click button absen in again"}
			}
			presen.Time_in = Sekarang
			presen.Lokasi_in = In
			presen.Status = Status
			DB.Select("id_pengguna", "time_in", "lokasi_in", "status").Create(&presen)
			response = model.PresenJResponses{Type: true, Message: "Success, Check in"}
		}
	}
	json.NewEncoder(w).Encode(response)
}
