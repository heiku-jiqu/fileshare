package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"runtime/metrics"
	"time"

	t "github.com/heiku-jiqu/fileshare/web/template"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServerFS(t.Website))
	mux.HandleFunc("/login", unimplemented)
	mux.HandleFunc("/healthcheck", Healthcheck)
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", NewFilesRouter()))

	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	slog.Info("Listening and serving...")
	log.Fatal(s.ListenAndServe())
}

func NewFilesRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{user}/files/", unimplemented)
	mux.HandleFunc("POST /{user}/file/", unimplemented)    // initiate new upload
	mux.HandleFunc("PUT /{user}/file/{id}", unimplemented) // complete upload?
	return logger(mux)
}
func Healthcheck(w http.ResponseWriter, r *http.Request) {
	descs := metrics.All()
	// Create a sample for each metric.
	samples := make([]metrics.Sample, len(descs))
	for i := range samples {
		samples[i].Name = descs[i].Name
	}

	// Sample the metrics. Re-use the samples slice if you can!
	metrics.Read(samples)

	// Iterate over all results.
	for _, sample := range samples {
		// Pull out the name and value.
		name, value := sample.Name, sample.Value

		// Handle each sample.
		switch value.Kind() {
		case metrics.KindUint64:
			fmt.Fprintf(w, "%s: %d\n", name, value.Uint64())
		case metrics.KindFloat64:
			fmt.Fprintf(w, "%s: %f\n", name, value.Float64())
		case metrics.KindFloat64Histogram:
			// The histogram may be quite large, so let's just pull out
			// a crude estimate for the median for the sake of this example.
			fmt.Fprintf(w, "%s: %f\n", name, medianBucket(value.Float64Histogram()))
		case metrics.KindBad:
			// This should never happen because all metrics are supported
			// by construction.
			panic("bug in runtime/metrics package!")
		default:
			// This may happen as new metrics get added.
			//
			// The safest thing to do here is to simply log it somewhere
			// as something to look into, but ignore it for now.
			// In the worst case, you might temporarily miss out on a new metric.
			fmt.Fprintf(w, "%s: unexpected metric Kind: %v\n", name, value.Kind())
		}
	}
}

func medianBucket(h *metrics.Float64Histogram) float64 {
	total := uint64(0)
	for _, count := range h.Counts {
		total += count
	}
	thresh := total / 2
	total = 0
	for i, count := range h.Counts {
		total += count
		if total >= thresh {
			return h.Buckets[i]
		}
	}
	panic("should not happen")
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ip     = r.RemoteAddr
			method = r.Method
			url    = r.URL.String()
			proto  = r.Proto
		)
		next.ServeHTTP(w, r)
		slog.Info("Request served.", "protocol", proto, "method", method, "url", url, "ip", ip, "matched", r.Pattern)
	})

}

func unimplemented(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "unimplemented")
}
