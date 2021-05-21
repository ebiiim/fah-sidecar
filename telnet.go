package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
)

func StartTelnetInfoCh(host, port string, interval time.Duration) <-chan string {
	ch := make(chan string)

	caller := &BufCaller{}
	caller.In = make(chan string)

	go func() {
		err := telnet.DialToAndCall(fmt.Sprintf("%s:%s", host, port), caller)
		log.Printf("GetInfoByTelnet: %v", err)
	}()

	go func() {
		// wait welcome message "Welcome to the FAHClient command server."
		for {
			<-time.After(100 * time.Millisecond)
			if strings.Contains(caller.Out.String(), "FAHClient") {
				break
			}
		}
		caller.Out.Reset()

		// queue-info loop
		for {
			caller.In <- "queue-info"
			s := caller.Out.String()
			// expects:
			//   PyON 1 units
			//   [ { ... } ]
			//   ---
			if !strings.Contains(s, "PyON") || !strings.Contains(s, "---") {
				log.Println("INFO: No data; continue")
				<-time.After(100 * time.Millisecond)
				continue
			}
			caller.Out.Reset()

			// send objects array
			begin := strings.Index(s, "[")
			end := strings.Index(s, "]")
			if begin == -1 || end == -1 || begin > end {
				log.Printf("ERROR: no valid JSON in result (%s)\n", s)
				caller.Out.Reset()
				continue
			}

			ch <- s[begin : end+1]
			<-time.After(interval)
		}
	}()

	return ch
}

type BufCaller struct {
	In  chan string
	Out bytes.Buffer
	Err bytes.Buffer
}

func (c *BufCaller) CallTELNET(_ telnet.Context, w telnet.Writer, r telnet.Reader) {
	go func(writer io.Writer, reader io.Reader) {
		var buffer [1]byte
		p := buffer[:]
		for {
			n, err := reader.Read(p)
			if n <= 0 && nil == err {
				continue
			} else if n <= 0 && nil != err {
				break
			}
			oi.LongWrite(writer, p)
		}
	}(&c.Out, r)

	for {
		cmd := <-c.In + "\r\n"
		buf := bytes.NewBufferString(cmd)
		p := buf.Bytes()
		n, err := oi.LongWrite(w, p)
		if nil != err {
			break
		}
		if expected, actual := int64(len(p)), n; expected != actual {
			err := fmt.Errorf("tried sending %d bytes, but actually only sent %d bytes", expected, actual)
			fmt.Fprint(&c.Err, err.Error())
			return
		}
	}
	time.Sleep(3 * time.Millisecond)
}
