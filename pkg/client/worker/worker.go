package worker

import (
	"github.com/9to6/aws-autoscaling-tester/pkg/client/config"
	"github.com/9to6/aws-autoscaling-tester/pkg/log"
	llog "github.com/sirupsen/logrus"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Worker struct {
	ticker *time.Ticker
	stop   chan bool
	conf   config.Config
	logger *llog.Entry
}

func NewWorker(conf config.Config, log *llog.Entry) *Worker {
	return &Worker{
		ticker: time.NewTicker(time.Duration(conf.Period) * time.Second),
		stop:   make(chan bool, 1),
		conf:   conf,
		logger: log.WithFields(llog.Fields{"worker": conf.Url}),
	}
}

func (w *Worker) StartWork() {
	go func() {
		for {
			select {
			case t := <-w.ticker.C:
				log.Logger.Info("Ticker executed: ", t)
				if err := w.request(); err != nil {
					log.Logger.Error(err)
				}
			case <-w.stop:
				break
			}
		}
	}()

}

func (w *Worker) setMode() {
	if w.conf.IncMode {
		if w.conf.ConnectionValue < int32(math.Exp2(31))-1 {
			w.conf.ConnectionValue += w.conf.IncValue
		} else {
			// change mode
			w.conf.IncMode = false
			w.conf.DecMode = true
			w.conf.DecValue = w.conf.IncValue
		}
	} else if w.conf.DecMode {
		if w.conf.ConnectionValue < 1 {
			// change mode
			w.conf.IncMode = true
			w.conf.DecMode = false
			w.conf.IncValue = w.conf.DecValue
		} else {
			w.conf.ConnectionValue -= w.conf.DecValue
		}
	}
}

func (w *Worker) request() error {
	w.setMode()
	w.logger.Info("start request")
	connectionValue := strconv.Itoa(int(w.conf.ConnectionValue))
	resp, err := http.PostForm(w.conf.Url, url.Values{"conn": []string{connectionValue}})
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
	w.stop <- true
	w.ticker.Stop()
}
