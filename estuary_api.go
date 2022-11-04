package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const BASE_URL = "https://api.estuary.tech"

func GetAllMiners() (*Miners, error) {

	resp, err := http.Get(BASE_URL + "/public/miners")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := UnmarshalMiners(body)
	if err != nil {
		return nil, err
	}

	return &result, err
}

func GetMinerStats(addr string) (*MinerStats, error) {
	resp, err := http.Get(BASE_URL + "/public/miners/stats/" + addr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := UnmarshalMinerStats(body)
	if err != nil {
		return nil, err
	}

	return &result, err
}

type Miners []Miner

func UnmarshalMiners(data []byte) (Miners, error) {
	var r Miners
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Miners) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Miner struct {
	Addr            string  `json:"addr"`
	Name            string  `json:"name"`
	Suspended       bool    `json:"suspended"`
	Version         string  `json:"version"`
	SuspendedReason *string `json:"suspendedReason,omitempty"`
}

func UnmarshalMinerStats(data []byte) (MinerStats, error) {
	var r MinerStats
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MinerStats) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MinerStats struct {
	Miner           string    `json:"miner"`
	Name            string    `json:"name"`
	Version         string    `json:"version"`
	UsedByEstuary   bool      `json:"usedByEstuary"`
	DealCount       int64     `json:"dealCount"`
	ErrorCount      int64     `json:"errorCount"`
	Suspended       bool      `json:"suspended"`
	SuspendedReason string    `json:"suspendedReason"`
	ChainInfo       ChainInfo `json:"chainInfo"`
}

type ChainInfo struct {
	PeerID    string   `json:"peerId"`
	Addresses []string `json:"addresses"`
	Owner     string   `json:"owner"`
	Worker    string   `json:"worker"`
}
