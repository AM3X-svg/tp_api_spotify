package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"tpspotify/pages"
	Struct "tpspotify/structure"
)

func renderPage(w http.ResponseWriter, filename string, data any) {
	err := pages.Temp.ExecuteTemplate(w, filename, data)
	if err != nil {
		http.Error(w, "Erreur rendu template : "+err.Error(), http.StatusInternalServerError)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	damso()
	renderPage(w, "index.html", nil)
}

func damso() {
	urlApi := "https://api.spotify.com/v1/artists/2UwqpfQtNuhBwviIC0f2ie/albums"
	token := ReloadToken()

	req, err := http.NewRequest(http.MethodGet, urlApi, nil)
	if err != nil {
		fmt.Println("Erreur requête:", err)
		return
	}

	req.Header.Add("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Erreur response:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Erreur lecture body:", err)
		return
	}

	var data Struct.ApiData
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Erreur JSON:", err)
		return
	}

	fmt.Println("Total albums:", data.Total)
	for _, album := range data.Items {
		fmt.Println("Nom album:", album.Name)
		fmt.Println("Type:", album.AlbumType)
		fmt.Println("Nombre de tracks:", album.TotalTracks)
		fmt.Println("URL Spotify:", album.ExternalURLs.Spotify)
		fmt.Println("Première image:", album.Images[0].URL)
		fmt.Println("-----")
	}
}

func ReloadToken() string {

	urlApi := "https://accounts.spotify.com/api/token"

	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	reqBody := bytes.NewBufferString(
		"grant_type=client_credentials&client_id=188e8a12982e43ffa8e5e0875a9b030a&client_secret=800b0a1e2eeb4293875d8301425b4a43",
	)

	req, errReq := http.NewRequest(http.MethodPost, urlApi, reqBody)
	if errReq != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errReq.Error())
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, errResp := httpClient.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	} else {
		fmt.Println("Error creating response:", errResp.Error())
		os.Exit(2)
	}

	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errBody.Error())
		return ""
	}

	var data Struct.Token
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		return ""
	}

	token := data.AccessToken

	fmt.Println(token)

	return token
}
