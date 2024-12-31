package server

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/antibaloo/imageGenerator/configs"
	"github.com/antibaloo/imageGenerator/pkg/img"
)

func rend(w http.ResponseWriter, msg string) {
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Println(err)
	}
}

func rendImage(w http.ResponseWriter, buffer *bytes.Buffer) {
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println(err)
	}
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	buffer, err := img.Generate(strings.Split(r.URL.Path, "/"))
	if err != nil {
		log.Println(err)
	}
	rendImage(w, buffer)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	buffer, err := img.GenerateFavicon()
	if err != nil {
		log.Println(err)
	}
	rendImage(w, buffer)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "pong")
}

func robotsHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "robots")
}

func Run(conf configs.ConfI) {
	http.HandleFunc("/", imgHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/robots.txt", robotsHandler)
	log.Println("Server starting ...")

	// Создаем канал для получения сигналов остановки приложения
	sigs := make(chan os.Signal, 1)
	// Указываем какие сигнал хотим получать через этот канал
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// Запускаем сервер в горутине
	go func() {
		if err := http.ListenAndServe(":"+conf.GetPort(), nil); err != nil {
			log.Fatal(err)
		}
	}()
	// Читаем канал с сигналами
	signalValue := <-sigs
	// При получении сигнала останавливаем уведомления
	signal.Stop(sigs)
	// ЛОгируем полученный сигнал
	log.Println("stop signal: ", signalValue)
}
