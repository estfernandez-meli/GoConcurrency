package main

import "sync"

type Mail struct {
	Domain      string
	Host        string
	Port        int
	Username    string
	Password    string
	Enryption   string
	FromAddress string
	FromName    string
	Wait        *sync.WaitGroup
	MailerChan  chan Message
	ErrorChan   chan error
	DoneChan    chan bool
}

type Message struct {
	From         string
	FromName     string
	To           string
	Subject      string
	Attachhments []string
	Data         interface{}
	DataMap      map[string]interface{}
	Template     string
}
