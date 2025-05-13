// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package carbonaware

import (
	"context"
	"net/http"

	"github.com/carbon-aware/scheduler-client-golang/internal/apijson"
	"github.com/carbon-aware/scheduler-client-golang/internal/requestconfig"
	"github.com/carbon-aware/scheduler-client-golang/option"
	"github.com/carbon-aware/scheduler-client-golang/packages/respjson"
)

// RegionService contains methods and other services that help with interacting
// with the carbonaware-scheduler API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegionService] method instead.
type RegionService struct {
	Options []option.RequestOption
}

// NewRegionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRegionService(opts ...option.RequestOption) (r RegionService) {
	r = RegionService{}
	r.Options = opts
	return
}

// Returns list of available regions.
func (r *RegionService) List(ctx context.Context, opts ...option.RequestOption) (res *RegionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/regions/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RegionListResponse struct {
	Regions []CloudZone `json:"regions,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Regions     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegionListResponse) RawJSON() string { return r.JSON.raw }
func (r *RegionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
