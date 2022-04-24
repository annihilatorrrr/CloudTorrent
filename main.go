package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const (
	root = "/home/roseloverx/ct/downloads/"
        // os.pwd()
	// dir to serve
)

func main() {
	fmt.Println("Server started on port " + PORT())
	http.Handle("/", http.FileServer(http.Dir("./static/")))
	http.HandleFunc("/api/v1/status", SystemStats)
	http.HandleFunc("/api/v1/torrents", TorrentsStats)
	http.HandleFunc("/downloads/", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("./static/downloads.html"))
		template.Execute(w, nil)
	})
	http.HandleFunc("/stream/", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("./static/player.html"))
		template.Execute(w, nil)
	})
	http.Handle("/torrents/update", SSEFeed)
	go streamTorrentUpdate()
	http.HandleFunc("/dir/", GetDirContents)
	http.HandleFunc("/home/", MainPage)
	http.HandleFunc("/add", AddTorrent)
	http.HandleFunc("/torrents/add", AddTorrent)
	http.HandleFunc("/torrents/delete", DeleteTorrent)
	http.HandleFunc("/torrents", TorrentsServe)
	http.HandleFunc("/torrents/details", GetTorrDir)
	http.HandleFunc("/torrents/search/", TorrentSearchPage)
	fmt.Println(http.ListenAndServe(":"+PORT(), nil))
}
