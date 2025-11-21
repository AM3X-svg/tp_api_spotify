package controller

import (
	"apispotify/pages"
	Struct "apispotify/struct"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func renderPage(w http.ResponseWriter, filename string, data any) {
	err := pages.Temp.ExecuteTemplate(w, filename, data)
	if err != nil {
		http.Error(w, "Erreur rendu template : "+err.Error(), http.StatusInternalServerError)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Damso":  GetDamsoAlbums(),
		"Laylow": GetLaylowTrack(),
	}
	renderPage(w, "index.html", data)
}

func Damso(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Damso": GetDamsoAlbums(),
	}
	renderPage(w, "damso.html", data)
}

func Laylow(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Laylow": GetLaylowTrack(),
	}
	renderPage(w, "laylow.html", data)
}

func GetDamsoAlbums() []Struct.Album {
	urlApi := "https://api.spotify.com/v1/artists/2UwqpfQtNuhBwviIC0f2ie/albums"

	token := ReloadToken()
	if token == "" {
		fmt.Println("Erreur : token vide")
		return nil
	}

	req, err := http.NewRequest(http.MethodGet, urlApi, nil)
	if err != nil {
		fmt.Println("Erreur création requête :", err)
		return nil
	}

	req.Header.Add("Authorization", "Bearer "+token)

	client := http.Client{Timeout: 5 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Erreur requête HTTP :", err)
		return nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Erreur lecture corps :", err)
		return nil
	}

	var data Struct.ApiData
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Erreur décodage JSON :", err)
		return nil
	}

	return data.Items
}

func GetLaylowTrack() []Struct.TrackData {
	urlApi := "https://api.spotify.com/v1/tracks/67Pf31pl0PfjBfUmvYNDCL"

	token := ReloadToken()
	if token == "" {
		fmt.Println("Erreur : token vide")
		return nil
	}

	req, err := http.NewRequest(http.MethodGet, urlApi, nil)
	if err != nil {
		fmt.Println("Erreur création requête :", err)
		return nil
	}

	req.Header.Add("Authorization", "Bearer "+token)

	client := http.Client{Timeout: 5 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Erreur requête HTTP :", err)
		return nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Erreur lecture corps :", err)
		return nil
	}

	var track Struct.TrackData
	if err := json.Unmarshal(body, &track); err != nil {
		fmt.Println("Erreur décodage JSON track :", err)
		return nil
	}

	return []Struct.TrackData{track}
}

func ReloadToken() string {
	urlApi := "https://accounts.spotify.com/api/token"

	form := "grant_type=client_credentials&client_id=188e8a12982e43ffa8e5e0875a9b030a&client_secret=800b0a1e2eeb4293875d8301425b4a43"
	reqBody := bytes.NewBufferString(form)

	req, err := http.NewRequest(http.MethodPost, urlApi, reqBody)
	if err != nil {
		fmt.Println("Erreur création requête token :", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{Timeout: 5 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Erreur requête token :", err)
		return ""
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Erreur lecture token :", err)
		return ""
	}

	fmt.Println("Body token:", string(body))

	var data Struct.Token
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Erreur décodage JSON token :", err)
		return ""
	}

	if data.AccessToken == "" {
		fmt.Println("Erreur : token vide après décodage")
	} else {
		fmt.Println("Token récupéré :", data.AccessToken)
	}

	return data.AccessToken
}
