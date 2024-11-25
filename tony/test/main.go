package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// Some()
	router := http.NewServeMux()

	router.Handle("/static", http.StripPrefix("./static", http.FileServer(http.Dir("static"))))
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/picture.html")
	})
	// router.HandleFunc("GET /", mainHandler)

	http.ListenAndServe(":3001", router)
}

// func imageHandler(w http.ResponseWriter, r *http.Request) {
// 	// Serve static files from the "images" directory
// 	http.ServeFile(w, r, r.URL.Path[len("./static/default"):])
// }

// func mainHandler(w http.ResponseWriter, r *http.Request) {
// 	// Generate a simple HTML page with an image tag

// 	imgPath := GetAvatarPath(os.Getenv("AVATARDEF")) // 假设图片文件名为 example.jpg
// 	fmt.Println(imgPath)
// 	html := fmt.Sprintf(`
// 		<!DOCTYPE html>
// 		<html>
// 		<head>
// 			<title>Image Display</title>
// 		</head>
// 		<body>
// 			<h1>Displaying Image from Local File System</h1>
// 			<img src="%s" alt="Example Image">
// 		</body>
// 		</html>
// 	`, imgPath)
// 	w.Header().Set("Content-Type", "text/html")
// 	fmt.Fprint(w, html)
// }

// 	// Ensure the images directory exists
// 	err := os.MkdirAll("images", os.ModePerm)
// 	if err != nil && !os.IsExist(err) {
// 		log.Fatalf("Failed to create images directory: %v", err)
// 	}

// 	// Register handlers
// 	http.HandleFunc("/", mainHandler)
// 	http.HandleFunc("/image/", imageHandler)

// 	// Start the server
// 	fmt.Println("Starting server on :8080...")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatalf("Server failed: %v", err)
// 	}
// }

func GetAvatarPath(avatarPath string) string {
	avatar := os.Getenv("AVATARDEF")
	aslice := strings.Split(avatar, " ")
	randi := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(4)
	return aslice[randi]
}

func Some() {
	image := GetAvatarPath(os.Getenv("AVATARDEF"))
	fmt.Println(image)
}
