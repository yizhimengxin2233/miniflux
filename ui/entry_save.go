// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package ui // import "miniflux.app/v2/ui"

import (
	"net/http"

	"miniflux.app/v2/http/request"
	"miniflux.app/v2/http/response/json"
	"miniflux.app/v2/integration"
	"miniflux.app/v2/model"
)

func (h *handler) saveEntry(w http.ResponseWriter, r *http.Request) {
	entryID := request.RouteInt64Param(r, "entryID")
	builder := h.store.NewEntryQueryBuilder(request.UserID(r))
	builder.WithEntryID(entryID)
	builder.WithoutStatus(model.EntryStatusRemoved)

	entry, err := builder.GetEntry()
	if err != nil {
		json.ServerError(w, r, err)
		return
	}

	if entry == nil {
		json.NotFound(w, r)
		return
	}

	settings, err := h.store.Integration(request.UserID(r))
	if err != nil {
		json.ServerError(w, r, err)
		return
	}

	go func() {
		integration.SendEntry(entry, settings)
	}()

	json.Created(w, r, map[string]string{"message": "saved"})
}
