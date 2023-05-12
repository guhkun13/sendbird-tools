package service

import "os"

const EvmInfoWelcomeMessage = `Assalamu'alaikum Reseller,

Selamat datang di channel Evermos Info!

Kini, kamu makin mudah dapatkan info menarik terkait Evermos. Info yang akan didapatkan tentunya beragam, dan membantu perjuangan ikhtiarmu. Jadi, sering-sering ya intip info channel ini supaya tidak ketinggalan kabar terkini.

Salam Sungkem dari Kami,
Seluruh Tim Evermos`

type MigratedUserSendbird struct {
	UserID string
	// FullName string
}

type MigratedUserSendbirdList []MigratedUserSendbird

type SendLog struct {
	Index    int
	Request  interface{}
	Response interface{}
}

type BlastWelcomeMessageRequest struct {
	Users   MigratedUserSendbirdList
	LogFile *os.File
}