package collector

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/tenntenn/natureremo"
)

type Message struct {
	RoomNames []string `json:"RoomNames"`
}

func CollectorHandler(w http.ResponseWriter, r *http.Request) {
	accessToken := os.Getenv("NATURE_REMO_ACCESS_TOKEN")
	projectID := os.Getenv("GCP_PROJECT")
	rootPath := os.Getenv("FIRESTORE_ROOT_PATH")

	var m Message
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Printf("ioutil.ReadAll: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(b, &m); err != nil {
		log.Printf("json.Unmarshal: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	errorChannel := make(chan error, len(m.RoomNames))
	for _, roomName := range m.RoomNames {
		go collect(accessToken, roomName, projectID, rootPath, errorChannel)
	}
	for range m.RoomNames {
		err := <-errorChannel
		if err != nil {
			log.Printf("collect: %v", err)
			if IsNoDevice(err) {
				http.Error(w,
					"Bad Request",
					http.StatusBadRequest)
			} else {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}
	}
}

func collect(accessToken, deviceID, projectID, rootPath string, c chan error) {
	natureremoClient := natureremo.NewClient(accessToken)
	fetcher := NewFetcher(natureremoClient, deviceID)

	ctx := context.Background()
	firestoreClient, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		c <- err
		return
	}
	defer firestoreClient.Close()
	repository, err := NewRepository(firestoreClient, rootPath, deviceID)
	if err != nil {
		c <- err
		return
	}

	service := NewCollectorService(fetcher, repository)
	err = service.Collect()
	if err != nil {
		c <- err
		return
	}
	c <- nil
}