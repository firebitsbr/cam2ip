package handlers

import (
	"log"
	"net/http"

	"github.com/gen2brain/cam2ip/camera"
)

// JPEG handler.
type JPEG struct {
	camera *camera.Camera
}

// NewJPEG returns new JPEG handler.
func NewJPEG(camera *camera.Camera) *JPEG {
	return &JPEG{camera}
}

// ServeHTTP handles requests on incoming connections.
func (j *JPEG) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Add("Connection", "close")
	w.Header().Add("Cache-Control", "no-store, no-cache")
	w.Header().Add("Content-Type", "image/jpeg")

	img, err := j.camera.Read()
	if err != nil {
		log.Printf("jpeg: read: %v", err)
		return
	}

	enc := camera.NewEncoder(w)

	err = enc.Encode(img)
	if err != nil {
		log.Printf("jpeg: encode: %v", err)
		return
	}
}
