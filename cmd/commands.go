package cmd

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func runServer(_ *cobra.Command, _ []string) error {
	delay := 5 * time.Second
	logrus.Infof("introducing artificial delay for %s", delay)
	<-time.After(delay)

	port := viper.GetInt64("server.port")
	logrus.Infof("running server on port %d", port)

	router := mux.NewRouter()
	router.HandleFunc("/api/echo", func(w http.ResponseWriter, r *http.Request) {
		bodyData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logrus.WithError(err).Error("failed to read request body")
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		logrus.Infof("server received request with body %s, now echo back", string(bodyData))
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(bodyData)
	})

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler: router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return srv.ListenAndServe()
}

func runClient(_ *cobra.Command, _ []string) error {
	port := viper.GetInt64("server.port")
	logrus.Infof("running client to reach server on port %d", port)

	for range time.NewTicker(5 * time.Second).C {
		u := url.URL{
			Scheme:     "http",
			Host:       fmt.Sprintf("localhost:%d", port),
			Path:       "/api/echo",
		}

		logrus.Infof("dialing %s", u.String())
		reader := bytes.NewReader([]byte(`{ "todo": "supervisor" }`))
		resp, err := http.Post(u.String(), "application/json", reader)
		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			logrus.Warnf("server returned non OK code %d", resp.StatusCode)
			continue
		}

		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		logrus.Infof("received back echo message %s", string(respBody))
	}

	return nil
}
