package main

import (
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"
)

type Assist_points struct {
	bonus1 string
	bonus2 string
}

type Score struct {
	Power          int `json:"power"`
	Agile          int `json:"agile"`
	Techni         int `json:"techni"`
	Breaking_point int `json:"breaking_point"`
	Mental         int `json:"mental"`
}

type Outbound_Data struct {
	Helper                   string `json:"helper"`
	Hit                      Score  `json:"hit"`
	Twobase                  Score  `json:"twobase"`
	Threebase                Score  `json:"threebase"`
	Homerun                  Score  `json:"homerun"`
	SacrificeBunt            Score  `json:"SacrificeBunt"`
	Sacrificefly             Score  `json:"Sacrificefly"`
	Steal                    Score  `json:"Steal"`
	Innings                  Score  `json:"Innings"`
	Outwithfastball          Score  `json:"Outwithfastball"`
	Grounderwithfastball     Score  `json:"Grounderwithfastball"`
	Popflywithfastball       Score  `json:"Popflywithfastball"`
	Outwithbreakingball      Score  `json:"Outwithbreakingball"`
	Grounderwithbreakingball Score  `json:"Grounderwithbreakingball"`
	Popflywithbreaking       Score  `json:"Popflywithbreaking"`
	Doubleplay               Score  `json:"Doubleplay"`
}

func main() {
	//本番環境
	//prdServer()
	//開発環境
	stgServer()
}

func prdServer() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set.")
	}

	go http.ListenAndServe(":80", http.HandlerFunc(redirect))
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	mux.HandleFunc("/", NotfoundHandler)
	mux.HandleFunc("/index", IndexHandler)
	mux.HandleFunc("/execute", ExecuteHandler)
	mux.HandleFunc("/announce", AnnounceHandler)
	mux.HandleFunc("/feedback", FeedbackHandler)
	mux.HandleFunc("/policy", PolicyHandler)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
		TLSConfig: &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			},
			PreferServerCipherSuites: true,
			InsecureSkipVerify:       true,
			MinVersion:               tls.VersionTLS12,
			MaxVersion:               tls.VersionTLS12,
		},
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServeTLS("/etc/letsencrypt/live/ilovegameandcomputer.tk/fullchain.pem", "/etc/letsencrypt/live/ilovegameandcomputer.tk/privkey.pem"); err != nil {
		log.Fatal(err)
	}
}

func stgServer() {
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	log.Fatal("$PORT must be set.")
	// }

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("dist/")))
	//mux.HandleFunc("/", NotfoundHandler)
	//mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/execute", ExecuteHandler)
	//mux.HandleFunc("/announce", AnnounceHandler)
	//mux.HandleFunc("/feedback", FeedbackHandler)
	//mux.HandleFunc("/policy", PolicyHandler)

	if err := http.ListenAndServe(":80", mux); err != nil {
		log.Fatal(err)
	}
}

func redirect(w http.ResponseWriter, req *http.Request) {
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target, 301)
}

func NotfoundHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/page_root.html"))
	tmpl.Execute(w, nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("dist/index.html"))
	tmpl.Execute(w, nil)
}

func ExecuteHandler(w http.ResponseWriter, r *http.Request) {
	var result Outbound_Data
	var data map[string]string

	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err = json.Unmarshal(body[:length], &data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper := data["helper"]
	result.Helper = helper
	isBehave(&result, helper)

	res, err := json.MarshalIndent(result, "", "	")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-Type", "application/json")
	w.Write(res)
}

func AnnounceHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/announce.html"))
	tmpl.Execute(w, nil)
}

func FeedbackHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/feedback.html"))
	tmpl.Execute(w, nil)
}

func PolicyHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/policy.html"))
	tmpl.Execute(w, nil)
}

func isHelper(helper string) Assist_points {
	var AP Assist_points

	switch helper {
	case "明星雪華":
		AP.bonus1 = "Power"
		AP.bonus2 = "Breaking-point"
	case "木場静香":
		AP.bonus1 = "Power"
		AP.bonus2 = "Technical"
	case "七瀬はるか":
		AP.bonus1 = "none"
		AP.bonus2 = "Breaking-point"
	case "倉家凪":
		AP.bonus1 = "Technical"
		AP.bonus2 = "Breaking-point"
	case "須神絵久":
		AP.bonus1 = "Power"
		AP.bonus2 = "Agile"
	case "エミリ":
		AP.bonus1 = "Power"
		AP.bonus2 = "mental"
	case "神良美砂":
		AP.bonus1 = "Agile"
		AP.bonus2 = "mental"
	case "嵐山美鈴":
		AP.bonus1 = "Agile"
		AP.bonus2 = "none"
	case "鴨川しぐれ":
		AP.bonus1 = "Agile"
		AP.bonus2 = "Breaking-point"
	case "虹谷彩理":
		AP.bonus1 = "Technical"
		AP.bonus2 = "Breaking-point"
	case "我間摩夕":
		AP.bonus1 = "none"
		AP.bonus2 = "none"
	case "姫野カレン":
		AP.bonus1 = "mental"
		AP.bonus2 = "Breaking-point"
	case "紺野美崎":
		AP.bonus1 = "none"
		AP.bonus2 = "none"
	case "黒沢愛":
		AP.bonus1 = "Power"
		AP.bonus2 = "none"
	case "四条澄香":
		AP.bonus1 = "Power"
		AP.bonus2 = "none"
	default:
		AP.bonus1 = "none"
		AP.bonus2 = "none"
	}

	return AP
}

func isBehave(data *Outbound_Data, helper string) {
	fmt.Printf("%p\n", &data)
	var mag1 float32
	var mag2 float32
	var mag3 float32
	var mag4 float32
	var mag5 float32

	reflectValue := reflect.ValueOf(*data)
	reflectType := reflectValue.Type()
	ap := isHelper(helper)
	mag1 = 1.0
	mag2 = 1.0
	mag3 = 1.0
	mag4 = 1.0
	mag5 = 1.0

	if ap.bonus1 == "Power" || ap.bonus2 == "Power" {
		mag1 = 1.6
	} else if ap.bonus1 == "Technical" || ap.bonus2 == "Technical" {
		mag2 = 1.6
	} else if ap.bonus1 == "Agile" || ap.bonus2 == "Agile" {
		mag3 = 1.6
	} else if ap.bonus1 == "Breaking-point" || ap.bonus2 == "Breaking-point" {
		mag4 = 1.6
	} else if ap.bonus1 == "Mental" || ap.bonus2 == "Mental" {
		mag5 = 1.6
	}

	for i := 0; i < reflectValue.NumField(); i++ {
		field := reflectType.Field(i)
		switch field.Name {
		case "Hit":
			data.Hit.Power = int(7 * mag1)
			data.Hit.Techni = int(5 * mag2)
			data.Hit.Agile = int(8 * mag3)
			data.Hit.Breaking_point = int(0 * mag4)
			data.Hit.Mental = int(0 * mag5)
		case "Twobase":
			data.Twobase.Power = int(7 * mag1)
			data.Twobase.Techni = int(5 * mag2)
			data.Twobase.Agile = int(8 * mag3)
			data.Twobase.Breaking_point = int(0 * mag4)
			data.Twobase.Mental = int(0 * mag5)
		case "Threebase":
			data.Threebase.Power = int(8 * mag1)
			data.Threebase.Techni = int(5 * mag2)
			data.Threebase.Agile = int(10 * mag3)
			data.Threebase.Breaking_point = int(0 * mag4)
			data.Threebase.Mental = int(0 * mag5)
		case "Homerun":
			data.Homerun.Power = int(10 * mag1)
			data.Homerun.Techni = int(5 * mag2)
			data.Homerun.Agile = int(0 * mag3)
			data.Homerun.Breaking_point = int(0 * mag4)
			data.Homerun.Mental = int(0 * mag5)
		case "SacrificeBunt":
			data.SacrificeBunt.Power = int(0 * mag1)
			data.SacrificeBunt.Techni = int(4 * mag2)
			data.SacrificeBunt.Agile = int(4 * mag3)
			data.SacrificeBunt.Breaking_point = int(0 * mag4)
			data.SacrificeBunt.Mental = int(0 * mag5)
		case "Sacrificefly":
			data.Sacrificefly.Power = int(3 * mag1)
			data.Sacrificefly.Techni = int(3 * mag2)
			data.Sacrificefly.Agile = int(5 * mag3)
			data.Sacrificefly.Breaking_point = int(0 * mag4)
			data.Sacrificefly.Mental = int(0 * mag5)
		case "Steal":
			data.Steal.Power = int(2 * mag1)
			data.Steal.Techni = int(2 * mag2)
			data.Steal.Agile = int(8 * mag3)
			data.Steal.Breaking_point = int(0 * mag4)
			data.Steal.Mental = int(0 * mag5)
		case "Innings":
			data.Innings.Power = int(3 * mag1)
			data.Innings.Techni = int(1 * mag2)
			data.Innings.Agile = int(1 * mag3)
			data.Innings.Breaking_point = int(0 * mag4)
			data.Innings.Mental = int(0 * mag5)
		case "Outwithfastball":
			data.Outwithfastball.Power = int(12 * mag1)
			data.Outwithfastball.Techni = int(8 * mag2)
			data.Outwithfastball.Agile = int(3 * mag3)
			data.Outwithfastball.Breaking_point = int(5 * mag4)
			data.Outwithfastball.Mental = int(0 * mag5)
		case "Grounderwithfastball":
			data.Grounderwithfastball.Power = int(9 * mag1)
			data.Grounderwithfastball.Techni = int(9 * mag2)
			data.Grounderwithfastball.Agile = int(4 * mag3)
			data.Grounderwithfastball.Breaking_point = int(4 * mag4)
			data.Grounderwithfastball.Mental = int(0 * mag5)
		case "Popflywithfastball":
			data.Popflywithfastball.Power = int(11 * mag1)
			data.Popflywithfastball.Techni = int(7 * mag2)
			data.Popflywithfastball.Agile = int(7 * mag3)
			data.Popflywithfastball.Breaking_point = int(2 * mag4)
			data.Popflywithfastball.Mental = int(0 * mag5)
		case "Outwithbreakingball":
			data.Outwithbreakingball.Power = int(5 * mag1)
			data.Outwithbreakingball.Techni = int(12 * mag2)
			data.Outwithbreakingball.Agile = int(2 * mag3)
			data.Outwithbreakingball.Breaking_point = int(21 * mag4)
			data.Outwithbreakingball.Mental = int(0 * mag5)
		case "Grounderwithbreakingball":
			data.Grounderwithbreakingball.Power = int(2 * mag1)
			data.Grounderwithbreakingball.Techni = int(13 * mag2)
			data.Grounderwithbreakingball.Agile = int(3 * mag3)
			data.Grounderwithbreakingball.Breaking_point = int(20 * mag4)
			data.Grounderwithbreakingball.Mental = int(0 * mag5)
		case "Popflywithbreaking":
			data.Popflywithbreaking.Power = int(4 * mag1)
			data.Popflywithbreaking.Techni = int(11 * mag2)
			data.Popflywithbreaking.Agile = int(6 * mag3)
			data.Popflywithbreaking.Breaking_point = int(18 * mag4)
			data.Popflywithbreaking.Mental = int(0 * mag5)
		case "Doubleplay":
			data.Doubleplay.Power = int(3 * mag1)
			data.Doubleplay.Techni = int(10 * mag2)
			data.Doubleplay.Agile = int(10 * mag3)
			data.Doubleplay.Breaking_point = int(5 * mag4)
			data.Doubleplay.Mental = int(0 * mag5)
		default:
			continue
		}
	}
}

func Tohash(originalstr string) string {
	salt := os.Getenv("SALT")
	hashstr := []byte(originalstr + salt)

	hash_sha256 := sha256.Sum256(hashstr)

	return hex.EncodeToString(hash_sha256[:])
}

func isMatch(hash, userid string) bool {
	salt := os.Getenv("SALT")
	hashstr := []byte(userid + salt)
	hash_sha256 := sha256.Sum256(hashstr)

	if hash == hex.EncodeToString(hash_sha256[:]) {
		return true
	}
	return false
}
