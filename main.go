// Copyright (c) 2021 Red Hat, Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func main() {
	start()
}

func OkHandler(w http.ResponseWriter, r *http.Request) {
	repo := r.FormValue("repo")
	path := r.FormValue("path")
	contents, err := getFileContents(context.TODO(), repo, path, "", func(url string) {
		http.Redirect(w, r, url, http.StatusFound)
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
	} else {
		_, _ = io.Copy(w, contents)
	}
}

func start() {
	router := mux.NewRouter()

	router.HandleFunc("/", OkHandler).Methods("GET")

	err := http.ListenAndServe(fmt.Sprintf(":%d", 8000), router)
	if err != nil {
		zap.L().Error("failed to start the HTTP server", zap.Error(err))
	}
}
