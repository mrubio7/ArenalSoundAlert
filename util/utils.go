package util

import (
	"ArenalSoundAlert/models"
	"ArenalSoundAlert/util/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"time"
)

func GetMediasFromApi() []models.Media {
	var medias []models.Media

	mediasData := getMediaDataApi()

	for _, mD := range mediasData {
		var temp models.Media
		temp.Date = mD.Date
		temp.ID = mD.ID
		temp.Image = mD.GUID.Rendered

		medias = append(medias, temp)
	}
	return medias
}

func LoadAllInDatabase() {
	medias := GetMediasFromApi()

	for _, media := range medias {
		db.NuevoElemento(media)
	}
}

func LoadInDatabase(media models.Media) {
	db.NuevoElemento(media)
}

func GetMediaDb() models.Media {
	return db.GetUltimoElemento()
}

func getMediaDataApi() []models.MediaData {
	var media []models.MediaData

	response, err := http.Get("https://www.arenalsound.com/wp-json/wp/v2/media")
	if err != nil {
		log.Panic("Error al obtener respuesta")
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panic("Error al obtener body")
	}

	json.Unmarshal(body, &media)
	for i, m := range media {
		newDate, err := time.Parse("2006-01-02T15:04:05", m.DateString)
		if err != nil {
			log.Panic(err)
		}
		media[i].Date = newDate
	}

	return media
}

func SendMail(media models.Media) {
	smtpHost := "smtp.ionos.es"
	smtpPort := "587"
	email := "arenalsound@s908820863.mialojamiento.es"

	auth := smtp.PlainAuth("", email, "****", smtpHost)

	to := []string{"arenalsoundalert@psonder.com"}

	msg := []byte("Subject: ArenalSoundAlert found new media\nFrom: " + email + "\n" + fmt.Sprintf("Fecha: %s\n Imagen: %s", media.Date.Format("02-01-2006 15:04:05"), media.Image))

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, email, to, msg)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("email sent with id: ", media.ID)
}
