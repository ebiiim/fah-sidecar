package main

import (
	"bytes"
	"encoding/json"
	"log"
)

func injectValue(j string, key string, val string) (string, error) {
	l := []map[string]interface{}{}
	if err := json.NewDecoder(bytes.NewBufferString(j)).Decode(&l); err != nil {
		log.Printf("injectValue: err=%v s=%s", err, j)
		return "", err
	}
	for _, v := range l {
		v[key] = val
	}
	p, err := json.Marshal(l)
	if err != nil {
		log.Printf("injectValue: err=%v", err)
		return "", err
	}
	return string(p), err
}

func allRunning(s string) bool {
	l := []map[string]interface{}{}
	if err := json.NewDecoder(bytes.NewBufferString(s)).Decode(&l); err != nil {
		log.Printf("allRunning: err=%v s=%s", err, s)
		return false
	}
	for idx, value := range l {
		state, ok := value["state"].(string)
		if !ok {
			log.Printf("allRunning: value[\"state\"] not ok (%v)", value)
			return false
		}
		if state != "RUNNING" {
			log.Printf("allRunning: not RUNNING but %v at %d", state, idx)
			return false
		}
	}
	return true
}
