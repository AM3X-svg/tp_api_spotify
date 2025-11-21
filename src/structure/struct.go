package structure

type ApiData struct {
    Href     string  `json:"href"`
    Limit    int     `json:"limit"`
    Next     *string `json:"next"`
    Offset   int     `json:"offset"`
    Previous *string `json:"previous"`
    Total    int     `json:"total"`
    Items    []Album `json:"items"`
}

type Album struct {
    AlbumType       string   `json:"album_type"`
    TotalTracks     int      `json:"total_tracks"`
    AvailableMarkets []string `json:"available_markets"`
    ExternalURLs    struct {
        Spotify string `json:"spotify"`
    } `json:"external_urls"`
    Href    string  `json:"href"`
    ID      string  `json:"id"`
    Images  []Image `json:"images"`
    Name    string  `json:"name"`
    ReleaseDate         string `json:"release_date"`
    ReleaseDatePrecision string `json:"release_date_precision"`
    Type    string  `json:"type"`
    URI     string  `json:"uri"`
    Artists []Artist `json:"artists"`
}

type Image struct {
    URL    string `json:"url"`
    Height int    `json:"height"`
    Width  int    `json:"width"`
}

type Artist struct {
    ExternalURLs struct {
        Spotify string `json:"spotify"`
    } `json:"external_urls"`
    Href string `json:"href"`
    ID   string `json:"id"`
    Name string `json:"name"`
    Type string `json:"type"`
    URI  string `json:"uri"`
}

type Token struct {
    AccessToken string `json:"access_token"`
    TokenType   string `json:"token_type"`
    ExpiresIn   int    `json:"expires_in"`
}
