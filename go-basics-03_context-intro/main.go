package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	r := Search(ctx, "golang")

	if r.Err != nil {
		log.Println(r.Err)
	} else {
		defer r.Body.Close()
		fmt.Fprintf(os.Stdout, "%v\n%v\n", r.Source, r.Status)
		io.Copy(os.Stdout, r.Body)
	}
}

func Search(ctx context.Context, keyword string) *Result {
	bingURL := bingSearchURL(keyword)
	sogouURL := sogouSearchURL(keyword)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	bingRequest, err := http.NewRequestWithContext(
		ctx, http.MethodGet, bingURL.String(), nil)
	if err != nil {
		log.Panic(err)
	}

	sogouRequest, err := http.NewRequestWithContext(
		ctx, http.MethodGet, sogouURL.String(), nil)
	if err != nil {
		log.Panic(err)
	}

	ch := make(chan *Result)

	go func() {
		ch <- NewResult(http.DefaultClient.Do(bingRequest))
	}()

	go func() {
		ch <- NewResult(http.DefaultClient.Do(sogouRequest))

	}()

	// Wait for the concurrent search result.
	// Earlier one wins.
	r := <-ch
	cancel()
	// Cleanup the failure.
	failure := <-ch
	if failure.Body != nil {
		io.Copy(io.Discard, failure.Body)
		failure.Body.Close()
	}

	return r
}

// bingSearchURL returns a cn.bing.com search URL.
func bingSearchURL(keyword string) *url.URL {
	u, err := url.Parse("https://cn.bing.com/search")
	if err != nil {
		log.Panic(err)
	}
	q := make(url.Values)
	q.Add("q", keyword)
	u.RawQuery = q.Encode()
	return u
}

// sogouSearchURL returns a sogou.com search URL.
func sogouSearchURL(keyword string) *url.URL {
	u, err := url.Parse("https://sogou.com/web")
	if err != nil {
		log.Panic(err)
	}
	q := make(url.Values)
	q.Add("query", keyword)
	u.RawQuery = q.Encode()
	return u
}

// The result of a search request.
type Result struct {
	Source string
	Status string
	Body   io.ReadCloser
	Err    error
}

func NewResult(r *http.Response, err error) *Result {
	if err != nil {
		return &Result{Err: err}
	}
	return &Result{Source: r.Request.URL.Host,
		Status: r.Status,
		Body:   r.Body}
}
