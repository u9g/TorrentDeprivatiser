package main

import (
	"errors"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/zeebo/bencode"
)

// File available as part of the torrent
type File struct {
	Length int      `bencode:"length"`
	Md5sum string   `bencode:"md5sum"`
	Path   []string `bencode:"path"`
}

// MetaInfoData hold data about the download itself
type MetaInfoData struct {
	Length      int    `bencode:"length"`
	Md5sum      string `bencode:"md5sum"`
	Name        string `bencode:"name"`
	PieceLength int    `bencode:"piece length"`
	Pieces      string `bencode:"pieces"`
	Private     int    `bencode:"private"`
	Files       []File `bencode:"files"`
}

// MetaInfo contain .torrent file description.
type MetaInfo struct {
	Announce     string       `bencode:"announce"`
	AnnounceList [][]string   `bencode:"announce-list"`
	Comment      string       `bencode:"comment"`
	CreatedBy    string       `bencode:"created by"`
	CreationDate int          `bencode:"creation date"`
	Info         MetaInfoData `bencode:"info"`
	Encoding     string       `bencode:"encoding"`
}

// Torrent struct
type Torrent struct {
	Path string
	Data MetaInfo
	Hash []byte
}

var trackerList [][]string

func readTrackerList() error {
	bodyContent, err := ioutil.ReadFile(arguments.TrackersFile)
	if err != nil {
		return err
	}

	trackers := strings.Split(string(bodyContent), "\n")

	trackerList = make([][]string, len(trackers))
	for k, v := range trackers {
		trackerList[k] = []string{v}
	}

	return nil
}

func newTorrentFromFile(path string) (Torrent, error) {
	torrentContent, err := ioutil.ReadFile(path)
	if err != nil {
		return Torrent{}, errors.New("Failed to open torrent file: " + err.Error())
	}

	info := MetaInfo{}
	err = bencode.DecodeBytes(torrentContent, &info)
	if err != nil {
		return Torrent{}, errors.New("Failed to decode torrent file: " + err.Error())
	}

	return Torrent{Path: path, Data: info}, nil
}

func newFileFromTorrent(torrent *Torrent) error {
	data, err := bencode.EncodeBytes(torrent.Data)
	if err != nil {
		return errors.New("Failed to encode torrent file: " + err.Error())
	}

	err = ioutil.WriteFile(torrent.Path, data, 0644)
	if err != nil {
		return errors.New("Failed to write torrent file: " + err.Error())
	}

	return nil
}

func work(fileName string, worker *sync.WaitGroup) error {
	defer worker.Done()

	path := arguments.Input + "/" + fileName

	torrent, err := newTorrentFromFile(path)
	if err != nil {
		return err
	}

	torrent.Data.Info.Private = 0
	torrent.Data.Announce = trackerList[0][0]
	torrent.Data.AnnounceList = trackerList

	err = newFileFromTorrent(&torrent)
	if err != nil {
		return err
	}

	return nil
}
