// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package carbonaware

import (
	"context"
	"net/http"
	"time"

	"github.com/stainless-sdks/carbonaware-go/internal/apijson"
	"github.com/stainless-sdks/carbonaware-go/internal/requestconfig"
	"github.com/stainless-sdks/carbonaware-go/option"
	"github.com/stainless-sdks/carbonaware-go/packages/param"
	"github.com/stainless-sdks/carbonaware-go/packages/respjson"
)

// ScheduleService contains methods and other services that help with interacting
// with the carbonaware-scheduler API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScheduleService] method instead.
type ScheduleService struct {
	Options []option.RequestOption
}

// NewScheduleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScheduleService(opts ...option.RequestOption) (r ScheduleService) {
	r = ScheduleService{}
	r.Options = opts
	return
}

// Schedule
func (r *ScheduleService) New(ctx context.Context, body ScheduleNewParams, opts ...option.RequestOption) (res *ScheduleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/schedule/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CloudZone struct {
	// Any of "aws", "gcp", "azure".
	Provider CloudZoneProvider `json:"provider,required"`
	// Any of "us-east-1", "us-west-1", "eu-central-1", "ap-southeast-2",
	// "us-central1", "eastus", "eastus2", "southcentralus", "westus2", "westus3",
	// "northeurope", "swedencentral", "uksouth", "westeurope", "centralus",
	// "francecentral", "germanywestcentral", "italynorth", "norwayeast",
	// "polandcentral", "eastus2euap", "eastusstg", "northcentralus", "westus",
	// "centraluseuap", "westcentralus", "francesouth", "germanynorth", "norwaywest",
	// "ukwest".
	Region CloudZoneRegion `json:"region,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Provider    respjson.Field
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloudZone) RawJSON() string { return r.JSON.raw }
func (r *CloudZone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CloudZone to a CloudZoneParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CloudZoneParam.Overrides()
func (r CloudZone) ToParam() CloudZoneParam {
	return param.Override[CloudZoneParam](r.RawJSON())
}

type CloudZoneProvider string

const (
	CloudZoneProviderAws   CloudZoneProvider = "aws"
	CloudZoneProviderGcp   CloudZoneProvider = "gcp"
	CloudZoneProviderAzure CloudZoneProvider = "azure"
)

type CloudZoneRegion string

const (
	CloudZoneRegionUsEast1            CloudZoneRegion = "us-east-1"
	CloudZoneRegionUsWest1            CloudZoneRegion = "us-west-1"
	CloudZoneRegionEuCentral1         CloudZoneRegion = "eu-central-1"
	CloudZoneRegionApSoutheast2       CloudZoneRegion = "ap-southeast-2"
	CloudZoneRegionUsCentral1         CloudZoneRegion = "us-central1"
	CloudZoneRegionEastus             CloudZoneRegion = "eastus"
	CloudZoneRegionEastus2            CloudZoneRegion = "eastus2"
	CloudZoneRegionSouthcentralus     CloudZoneRegion = "southcentralus"
	CloudZoneRegionWestus2            CloudZoneRegion = "westus2"
	CloudZoneRegionWestus3            CloudZoneRegion = "westus3"
	CloudZoneRegionNortheurope        CloudZoneRegion = "northeurope"
	CloudZoneRegionSwedencentral      CloudZoneRegion = "swedencentral"
	CloudZoneRegionUksouth            CloudZoneRegion = "uksouth"
	CloudZoneRegionWesteurope         CloudZoneRegion = "westeurope"
	CloudZoneRegionCentralus          CloudZoneRegion = "centralus"
	CloudZoneRegionFrancecentral      CloudZoneRegion = "francecentral"
	CloudZoneRegionGermanywestcentral CloudZoneRegion = "germanywestcentral"
	CloudZoneRegionItalynorth         CloudZoneRegion = "italynorth"
	CloudZoneRegionNorwayeast         CloudZoneRegion = "norwayeast"
	CloudZoneRegionPolandcentral      CloudZoneRegion = "polandcentral"
	CloudZoneRegionEastus2euap        CloudZoneRegion = "eastus2euap"
	CloudZoneRegionEastusstg          CloudZoneRegion = "eastusstg"
	CloudZoneRegionNorthcentralus     CloudZoneRegion = "northcentralus"
	CloudZoneRegionWestus             CloudZoneRegion = "westus"
	CloudZoneRegionCentraluseuap      CloudZoneRegion = "centraluseuap"
	CloudZoneRegionWestcentralus      CloudZoneRegion = "westcentralus"
	CloudZoneRegionFrancesouth        CloudZoneRegion = "francesouth"
	CloudZoneRegionGermanynorth       CloudZoneRegion = "germanynorth"
	CloudZoneRegionNorwaywest         CloudZoneRegion = "norwaywest"
	CloudZoneRegionUkwest             CloudZoneRegion = "ukwest"
)

// The properties Provider, Region are required.
type CloudZoneParam struct {
	// Any of "aws", "gcp", "azure".
	Provider CloudZoneProvider `json:"provider,omitzero,required"`
	// Any of "us-east-1", "us-west-1", "eu-central-1", "ap-southeast-2",
	// "us-central1", "eastus", "eastus2", "southcentralus", "westus2", "westus3",
	// "northeurope", "swedencentral", "uksouth", "westeurope", "centralus",
	// "francecentral", "germanywestcentral", "italynorth", "norwayeast",
	// "polandcentral", "eastus2euap", "eastusstg", "northcentralus", "westus",
	// "centraluseuap", "westcentralus", "francesouth", "germanynorth", "norwaywest",
	// "ukwest".
	Region CloudZoneRegion `json:"region,omitzero,required"`
	paramObj
}

func (r CloudZoneParam) MarshalJSON() (data []byte, err error) {
	type shadow CloudZoneParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CloudZoneParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScheduleOption struct {
	Co2Intensity float64   `json:"co2_intensity,required"`
	Time         time.Time `json:"time,required" format:"date-time"`
	Zone         CloudZone `json:"zone,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Co2Intensity respjson.Field
		Time         respjson.Field
		Zone         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduleOption) RawJSON() string { return r.JSON.raw }
func (r *ScheduleOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScheduleNewResponse struct {
	CarbonSavings ScheduleNewResponseCarbonSavings `json:"carbon_savings,required"`
	Ideal         ScheduleOption                   `json:"ideal,required"`
	MedianCase    ScheduleOption                   `json:"median_case,required"`
	NaiveCase     ScheduleOption                   `json:"naive_case,required"`
	Options       []ScheduleOption                 `json:"options,required"`
	WorstCase     ScheduleOption                   `json:"worst_case,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CarbonSavings respjson.Field
		Ideal         respjson.Field
		MedianCase    respjson.Field
		NaiveCase     respjson.Field
		Options       respjson.Field
		WorstCase     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduleNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ScheduleNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScheduleNewResponseCarbonSavings struct {
	VsMedianCase float64 `json:"vs_median_case,required"`
	VsNaiveCase  float64 `json:"vs_naive_case,required"`
	VsWorstCase  float64 `json:"vs_worst_case,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		VsMedianCase respjson.Field
		VsNaiveCase  respjson.Field
		VsWorstCase  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduleNewResponseCarbonSavings) RawJSON() string { return r.JSON.raw }
func (r *ScheduleNewResponseCarbonSavings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScheduleNewParams struct {
	Duration string `json:"duration,required" format:"duration"`
	// List of time windows to schedule (start and end must be in the future)
	Windows    []ScheduleNewParamsWindow `json:"windows,omitzero,required"`
	Zones      []CloudZoneParam          `json:"zones,omitzero,required"`
	NumOptions param.Opt[int64]          `json:"num_options,omitzero"`
	paramObj
}

func (r ScheduleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ScheduleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScheduleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties End, Start are required.
type ScheduleNewParamsWindow struct {
	End   time.Time `json:"end,required" format:"date-time"`
	Start time.Time `json:"start,required" format:"date-time"`
	paramObj
}

func (r ScheduleNewParamsWindow) MarshalJSON() (data []byte, err error) {
	type shadow ScheduleNewParamsWindow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScheduleNewParamsWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
