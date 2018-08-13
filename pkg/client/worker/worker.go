package worker

import (
	"io/ioutil"
	"github.com/9to6/aws-autoscaling-tester/pkg/log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Worker struct {
	ticker          *time.Ticker
	connectionValue int32
	url             string
}

func NewWorker(durationSec, connectionValue int32, url string) *Worker {
	return &Worker{
		ticker:          time.NewTicker(time.Duration(durationSec) * time.Second),
		connectionValue: connectionValue,
		url:             url,
	}
}

func (w *Worker) StartWork() {
	// ctx, cancelFunc := context.WithCancel(context.Background())
	// shutdown(cancelFunc)
	go func() {
		for t := range w.ticker.C {
			log.Logger.Info("Ticker executed: ", t)
			if err := w.request(); err != nil {
				log.Logger.Error(err)
			}
		}
	}()

}

func (w *Worker) request() error {
	connectionValue := strconv.Itoa(int(w.connectionValue))
	resp, err := http.PostForm(w.url, url.Values{"conn": []string{connectionValue}})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Response 체크.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	str := string(respBody)
	log.Logger.Info("resp", str)
	return nil
}

func (w *Worker) StopWork() {
	w.ticker.Stop()
}
