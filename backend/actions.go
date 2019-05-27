package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func getDatabase() (*gorm.DB, error) {
	const addr = "postgresql://maxroach@localhost:26257/test?sslmode=disable"
	db, err := gorm.Open("postgres", addr)
	return db, err
}

/*Funcion para crear la llave RSA*/
func createRSAKey(w http.ResponseWriter, r *http.Request) {
	var post Llave
	json.NewDecoder(r.Body).Decode(&post)
	if post.NombreLlave == "" {
		log.Fatal("err")
	}
	// Connect to the "bank" database as the "maxroach" user.
	db, err := getDatabase()
	checkError(err)
	defer db.Close()

	bitSize := 2048
	//Genero la llaveRSA
	llavePrivada := generateRSAKey(bitSize)
	//Guardo en bytes la llave privada y la llave publica
	llavePrivadaBytes := PrivateKeyToBytes(llavePrivada)
	llavePublicaBytes := PublicKeyToBytes(&llavePrivada.PublicKey)
	//Encrypto la llave privada
	llavePrivadaEncryptedGCM := GCMEncrypt(llavePrivadaBytes)
	//Creo el registro en la base de datos
	if err := db.Create(&Llave{NombreLlave: post.NombreLlave, LlavePublica: llavePublicaBytes, LlavePrivada: llavePrivadaEncryptedGCM}).Error; err != nil {
		checkErrorResponse(err, w)
	}
	var llaves []Llave
	db.Find(&llaves)
	responseKeys(w, 200, llaves)
}

func getRSAKeys(w http.ResponseWriter, r *http.Request) {
	db, err := getDatabase()
	checkError(err)
	defer db.Close()
	var llaves []Llave
	db.AutoMigrate(&Llave{})
	if err := db.Find(&llaves).Error; err != nil {
		checkErrorResponse(err, w)
	}
	checkErrorResponse(err, w)
	responseKeys(w, 200, llaves)
}

func getRSAKey(w http.ResponseWriter, r *http.Request) {
	NombreLlave := chi.URLParam(r, "NombreLlave")
	db, err := getDatabase()
	checkErrorResponse(err, w)
	defer db.Close()
	var llave []Llave
	if err := db.Find(&llave, Llave{NombreLlave: NombreLlave}).Error; err != nil {
		checkErrorResponse(err, w)
	}
	checkErrorResponse(err, w)
	responseKeys(w, 200, llave)
}

func encryptPlainText(w http.ResponseWriter, r *http.Request) {
	var entrada Entrada
	json.NewDecoder(r.Body).Decode(&entrada)
	const addr = "postgresql://maxroach@localhost:26257/bank?sslmode=disable"
	db, err := getDatabase()
	checkErrorResponse(err, w)
	defer db.Close()
	IdLlaveInt, err := strconv.ParseInt(entrada.IdLlave, 10, 64)

	var llave Llave
	//Se realiza la consulta a la base de datos
	if err := db.First(&llave, Llave{IdLlave: IdLlaveInt}).Error; err != nil {
		checkErrorResponse(err, w)
	}
	//Se convierte los bytes a la llave privada
	publicKeyBytes := BytesToPublicKey(llave.LlavePublica)
	TextoPlanoBytes := []byte(entrada.TextoPlano)
	resultadoEncryptado, err := EncryptWithPublicKey(TextoPlanoBytes, publicKeyBytes)
	checkErrorResponse(err, w)
	sEnc := base64.StdEncoding.EncodeToString(resultadoEncryptado)
	response(w, 200, string(sEnc))
}

func decryptPlainText(w http.ResponseWriter, r *http.Request) {
	var entrada Entrada
	json.NewDecoder(r.Body).Decode(&entrada)
	db, err := getDatabase()
	checkErrorResponse(err, w)
	defer db.Close()
	IdLlaveInt, err := strconv.ParseInt(entrada.IdLlave, 10, 64)
	var llave Llave
	//Se realiza la consulta a la base de datos
	if err := db.First(&llave, Llave{IdLlave: IdLlaveInt}).Error; err != nil {
		checkErrorResponse(err, w)
	}
	sDec, _ := base64.StdEncoding.DecodeString(entrada.TextoPlano)
	LlavePrivada := GCMDecrypt(llave.LlavePrivada)
	privateKey := BytesToPrivateKey(LlavePrivada)
	resultadoDecryptado, err := DecryptWithPrivateKey(sDec, privateKey)
	checkErrorResponse(err, w)
	response(w, 200, string(resultadoDecryptado))
}
