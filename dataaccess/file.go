// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dataaccess

import (
	"fmt"
	"github.com/andreaskoch/allmark2/common"
	"github.com/andreaskoch/allmark2/common/route"
)

// A File represents a file ressource that is associated with an Item.
type File struct {
	route *route.Route

	hashProvider    common.HashProviderFunc
	contentProvider common.ContentProviderFunc
}

func NewFile(route *route.Route, hashProvider common.HashProviderFunc, contentProvider common.ContentProviderFunc) (*File, error) {
	return &File{
		route: route,

		hashProvider:    hashProvider,
		contentProvider: contentProvider,
	}, nil
}

func (file *File) String() string {
	return fmt.Sprintf("%s", file.route)
}

func (file *File) Route() *route.Route {
	return file.route
}

func (file *File) HashProvider() common.HashProviderFunc {
	return file.hashProvider
}

func (file *File) ContentProvider() common.ContentProviderFunc {
	return file.contentProvider
}
