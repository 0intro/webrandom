package main

import (
	"fmt"
	"net/http"
	"time"
)

type LogRecord struct {
	http.ResponseWriter

	addr                  string
	time                  time.Time
	method, uri, protocol string
	status                int
	size                  int64
	duration              time.Duration
}

func (r *LogRecord) String() string {
	time := r.time.Format("2006-01-02 03:04:05")
	duration := r.duration.Nanoseconds() / 1000000
	return fmt.Sprintf("%s %s %s %s %s %d %d %d", time, r.addr, r.method, r.uri, r.protocol, r.status, r.size, duration)
}

func (r *LogRecord) Write(b []byte) (int, error) {
	n, err := r.ResponseWriter.Write(b)
	r.size += int64(n)
	return n, err
}

func (r *LogRecord) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		record := &LogRecord{
			ResponseWriter: w,
			addr:           r.RemoteAddr,
			method:         r.Method,
			uri:            r.RequestURI,
			protocol:       r.Proto,
			status:         http.StatusOK,
		}

		startTime := time.Now()
		handler.ServeHTTP(record, r)
		stopTime := time.Now()

		record.time = stopTime
		record.duration = stopTime.Sub(startTime)
		fmt.Println(record)
	})
}
