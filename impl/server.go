package impl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SimpleDiagram struct {
	ProcessName string `json:"process_name"`
	VersionTag  int    `json:"version_tag"`
	ProcessId   int    `json:"process_id"`
}

func handleNewDiagram(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Println(r.Header)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	_, err = w.Write(body)
	if err != nil {
		log.Println("handleNewDiagram w.Write error: ", err)
	}
}

func handleGetDiagram(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Println(r.Header)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	_, err = w.Write(body)
	if err != nil {
		log.Println("handleNewDiagram w.Write error: ", err)

	}
	w.WriteHeader(http.StatusOK)
}

func handleTestGetAllDiagram(w http.ResponseWriter, r *http.Request) {
	var simpleDiagrams []*SimpleDiagram
	simpleDiagrams = append(simpleDiagrams, &SimpleDiagram{
		ProcessId:   1,
		ProcessName: "First_process",
		VersionTag:  1,
	})
	simpleDiagrams = append(simpleDiagrams, &SimpleDiagram{
		ProcessId:   2,
		ProcessName: "First_process",
		VersionTag:  2,
	})
	simpleDiagrams = append(simpleDiagrams, &SimpleDiagram{
		ProcessId:   3,
		ProcessName: "Second_process",
		VersionTag:  1,
	})
	jsonData, err := json.Marshal(simpleDiagrams)
	if err != nil {
		log.Println("handleTestGetAllDiagram json.Marshal error: ", err)
	}
	w.Header().Set("Content-Type", "application/json charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonData)
	if err != nil {
		log.Println("handleTestGetAllDiagram w.Write error: ", err)
	}
}

func Serve(addr string) {
	log.Println("test")
	http.HandleFunc("/api/new", handleNewDiagram)
	http.HandleFunc("/api/get", handleGetDiagram)
	http.HandleFunc("/api/test/getAll", handleTestGetAllDiagram)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
