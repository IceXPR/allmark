// Copyright 2014 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package viewmodel

import (
	"sort"
)

type Base struct {
	RepositoryName        string `json:"repositoryName"`
	RepositoryDescription string `json:"repositoryDescription"`

	Type  string `json:"type"`
	Level int    `json:"level"`
	Route string `json:"route"`
	Alias string `json:"alias"`

	ParentRoute string `json:"parentRoute"`

	BaseUrl  string `json:"baseUrl"`
	PrintUrl string `json:"printUrl"`
	JsonUrl  string `json:"jsonUrl"`
	RtfUrl   string `json:"rtfUrl"`

	PageTitle   string `json:"pageTitle"`
	Title       string `json:"title"`
	Description string `json:"description"`

	LanguageTag      string `json:"languageTag"`
	CreationDate     string `json:"creationdate"`
	LastModifiedDate string `json:"lastmodifieddate"`

	LiveReloadEnabled bool
}

type SortBaseModelBy func(model1, model2 *Base) bool

func (by SortBaseModelBy) Sort(models []*Base) {
	sorter := &baseModelSorter{
		models: models,
		by:     by,
	}

	sort.Sort(sorter)
}

type baseModelSorter struct {
	models []*Base
	by     SortBaseModelBy
}

func (sorter *baseModelSorter) Len() int {
	return len(sorter.models)
}

func (sorter *baseModelSorter) Swap(i, j int) {
	sorter.models[i], sorter.models[j] = sorter.models[j], sorter.models[i]
}

func (sorter *baseModelSorter) Less(i, j int) bool {
	return sorter.by(sorter.models[i], sorter.models[j])
}
