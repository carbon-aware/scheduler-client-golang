// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package carbonaware

import (
	"context"
	"net/http"
	"time"

	"github.com/carbon-aware/scheduler-client-golang/internal/apijson"
	"github.com/carbon-aware/scheduler-client-golang/internal/requestconfig"
	"github.com/carbon-aware/scheduler-client-golang/option"
	"github.com/carbon-aware/scheduler-client-golang/packages/param"
	"github.com/carbon-aware/scheduler-client-golang/packages/respjson"
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
	// Any of "aws", "gcp", "azure", "ovh".
	Provider CloudZoneProvider `json:"provider,required"`
	// Any of "af-south-1", "ap-east-1", "ap-northeast-1", "ap-northeast-2",
	// "ap-northeast-3", "ap-south-1", "ap-south-2", "ap-southeast-2",
	// "ap-southeast-3", "ap-southeast-4", "ap-southeast-5", "ap-southeast-7",
	// "ca-central-1", "ca-west-1", "eu-central-1", "eu-central-2", "eu-north-1",
	// "eu-south-1", "eu-south-2", "eu-west-1", "eu-west-2", "eu-west-3",
	// "il-central-1", "me-central-1", "mx-central-1", "sa-east-1", "us-east-1",
	// "us-east-2", "us-west-1", "us-west-2", "asia-east1", "asia-northeast1",
	// "asia-southeast1", "europe-north1", "europe-west1", "europe-west3",
	// "europe-west4", "northamerica-northeast2", "southamerica-west1", "us-central1",
	// "us-east1", "us-east4", "us-west1", "us-west2", "australiacentral",
	// "australiaeast", "australiasoutheast", "brazilsouth", "canadacentral",
	// "canadaeast", "centralindia", "centralus", "chinaeast2", "chinanorth",
	// "chinanorth2", "chinanorth3", "eastasia", "eastus", "eastus2", "francecentral",
	// "germanywestcentral", "indonesiacentral", "israelcentral", "italynorth",
	// "japaneast", "japanwest", "koreacentral", "mexicocentral", "newzealandnorth",
	// "northcentralus", "northeurope", "norwayeast", "polandcentral",
	// "southafricanorth", "southcentralus", "southindia", "spaincentral",
	// "swedencentral", "switzerlandnorth", "uaenorth", "uksouth", "ukwest",
	// "westcentralus", "westeurope", "westus", "westus2", "westus3", "fr-strasbourg",
	// "pl-warsaw".
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
	CloudZoneProviderOvh   CloudZoneProvider = "ovh"
)

type CloudZoneRegion string

const (
	CloudZoneRegionAfSouth1               CloudZoneRegion = "af-south-1"
	CloudZoneRegionApEast1                CloudZoneRegion = "ap-east-1"
	CloudZoneRegionApNortheast1           CloudZoneRegion = "ap-northeast-1"
	CloudZoneRegionApNortheast2           CloudZoneRegion = "ap-northeast-2"
	CloudZoneRegionApNortheast3           CloudZoneRegion = "ap-northeast-3"
	CloudZoneRegionApSouth1               CloudZoneRegion = "ap-south-1"
	CloudZoneRegionApSouth2               CloudZoneRegion = "ap-south-2"
	CloudZoneRegionApSoutheast2           CloudZoneRegion = "ap-southeast-2"
	CloudZoneRegionApSoutheast3           CloudZoneRegion = "ap-southeast-3"
	CloudZoneRegionApSoutheast4           CloudZoneRegion = "ap-southeast-4"
	CloudZoneRegionApSoutheast5           CloudZoneRegion = "ap-southeast-5"
	CloudZoneRegionApSoutheast7           CloudZoneRegion = "ap-southeast-7"
	CloudZoneRegionCaCentral1             CloudZoneRegion = "ca-central-1"
	CloudZoneRegionCaWest1                CloudZoneRegion = "ca-west-1"
	CloudZoneRegionEuCentral1             CloudZoneRegion = "eu-central-1"
	CloudZoneRegionEuCentral2             CloudZoneRegion = "eu-central-2"
	CloudZoneRegionEuNorth1               CloudZoneRegion = "eu-north-1"
	CloudZoneRegionEuSouth1               CloudZoneRegion = "eu-south-1"
	CloudZoneRegionEuSouth2               CloudZoneRegion = "eu-south-2"
	CloudZoneRegionEuWest1                CloudZoneRegion = "eu-west-1"
	CloudZoneRegionEuWest2                CloudZoneRegion = "eu-west-2"
	CloudZoneRegionEuWest3                CloudZoneRegion = "eu-west-3"
	CloudZoneRegionIlCentral1             CloudZoneRegion = "il-central-1"
	CloudZoneRegionMeCentral1             CloudZoneRegion = "me-central-1"
	CloudZoneRegionMxCentral1             CloudZoneRegion = "mx-central-1"
	CloudZoneRegionSaEast1                CloudZoneRegion = "sa-east-1"
	CloudZoneRegionUsEast1KebabCase       CloudZoneRegion = "us-east-1"
	CloudZoneRegionUsEast2                CloudZoneRegion = "us-east-2"
	CloudZoneRegionUsWest1KebabCase       CloudZoneRegion = "us-west-1"
	CloudZoneRegionUsWest2KebabCase       CloudZoneRegion = "us-west-2"
	CloudZoneRegionAsiaEast1              CloudZoneRegion = "asia-east1"
	CloudZoneRegionAsiaNortheast1         CloudZoneRegion = "asia-northeast1"
	CloudZoneRegionAsiaSoutheast1         CloudZoneRegion = "asia-southeast1"
	CloudZoneRegionEuropeNorth1           CloudZoneRegion = "europe-north1"
	CloudZoneRegionEuropeWest1            CloudZoneRegion = "europe-west1"
	CloudZoneRegionEuropeWest3            CloudZoneRegion = "europe-west3"
	CloudZoneRegionEuropeWest4            CloudZoneRegion = "europe-west4"
	CloudZoneRegionNorthamericaNortheast2 CloudZoneRegion = "northamerica-northeast2"
	CloudZoneRegionSouthamericaWest1      CloudZoneRegion = "southamerica-west1"
	CloudZoneRegionUsCentral1             CloudZoneRegion = "us-central1"
	CloudZoneRegionUsEast1                CloudZoneRegion = "us-east1"
	CloudZoneRegionUsEast4                CloudZoneRegion = "us-east4"
	CloudZoneRegionUsWest1                CloudZoneRegion = "us-west1"
	CloudZoneRegionUsWest2                CloudZoneRegion = "us-west2"
	CloudZoneRegionAustraliacentral       CloudZoneRegion = "australiacentral"
	CloudZoneRegionAustraliaeast          CloudZoneRegion = "australiaeast"
	CloudZoneRegionAustraliasoutheast     CloudZoneRegion = "australiasoutheast"
	CloudZoneRegionBrazilsouth            CloudZoneRegion = "brazilsouth"
	CloudZoneRegionCanadacentral          CloudZoneRegion = "canadacentral"
	CloudZoneRegionCanadaeast             CloudZoneRegion = "canadaeast"
	CloudZoneRegionCentralindia           CloudZoneRegion = "centralindia"
	CloudZoneRegionCentralus              CloudZoneRegion = "centralus"
	CloudZoneRegionChinaeast2             CloudZoneRegion = "chinaeast2"
	CloudZoneRegionChinanorth             CloudZoneRegion = "chinanorth"
	CloudZoneRegionChinanorth2            CloudZoneRegion = "chinanorth2"
	CloudZoneRegionChinanorth3            CloudZoneRegion = "chinanorth3"
	CloudZoneRegionEastasia               CloudZoneRegion = "eastasia"
	CloudZoneRegionEastus                 CloudZoneRegion = "eastus"
	CloudZoneRegionEastus2                CloudZoneRegion = "eastus2"
	CloudZoneRegionFrancecentral          CloudZoneRegion = "francecentral"
	CloudZoneRegionGermanywestcentral     CloudZoneRegion = "germanywestcentral"
	CloudZoneRegionIndonesiacentral       CloudZoneRegion = "indonesiacentral"
	CloudZoneRegionIsraelcentral          CloudZoneRegion = "israelcentral"
	CloudZoneRegionItalynorth             CloudZoneRegion = "italynorth"
	CloudZoneRegionJapaneast              CloudZoneRegion = "japaneast"
	CloudZoneRegionJapanwest              CloudZoneRegion = "japanwest"
	CloudZoneRegionKoreacentral           CloudZoneRegion = "koreacentral"
	CloudZoneRegionMexicocentral          CloudZoneRegion = "mexicocentral"
	CloudZoneRegionNewzealandnorth        CloudZoneRegion = "newzealandnorth"
	CloudZoneRegionNorthcentralus         CloudZoneRegion = "northcentralus"
	CloudZoneRegionNortheurope            CloudZoneRegion = "northeurope"
	CloudZoneRegionNorwayeast             CloudZoneRegion = "norwayeast"
	CloudZoneRegionPolandcentral          CloudZoneRegion = "polandcentral"
	CloudZoneRegionSouthafricanorth       CloudZoneRegion = "southafricanorth"
	CloudZoneRegionSouthcentralus         CloudZoneRegion = "southcentralus"
	CloudZoneRegionSouthindia             CloudZoneRegion = "southindia"
	CloudZoneRegionSpaincentral           CloudZoneRegion = "spaincentral"
	CloudZoneRegionSwedencentral          CloudZoneRegion = "swedencentral"
	CloudZoneRegionSwitzerlandnorth       CloudZoneRegion = "switzerlandnorth"
	CloudZoneRegionUaenorth               CloudZoneRegion = "uaenorth"
	CloudZoneRegionUksouth                CloudZoneRegion = "uksouth"
	CloudZoneRegionUkwest                 CloudZoneRegion = "ukwest"
	CloudZoneRegionWestcentralus          CloudZoneRegion = "westcentralus"
	CloudZoneRegionWesteurope             CloudZoneRegion = "westeurope"
	CloudZoneRegionWestus                 CloudZoneRegion = "westus"
	CloudZoneRegionWestus2                CloudZoneRegion = "westus2"
	CloudZoneRegionWestus3                CloudZoneRegion = "westus3"
	CloudZoneRegionFrStrasbourg           CloudZoneRegion = "fr-strasbourg"
	CloudZoneRegionPlWarsaw               CloudZoneRegion = "pl-warsaw"
)

// The properties Provider, Region are required.
type CloudZoneParam struct {
	// Any of "aws", "gcp", "azure", "ovh".
	Provider CloudZoneProvider `json:"provider,omitzero,required"`
	// Any of "af-south-1", "ap-east-1", "ap-northeast-1", "ap-northeast-2",
	// "ap-northeast-3", "ap-south-1", "ap-south-2", "ap-southeast-2",
	// "ap-southeast-3", "ap-southeast-4", "ap-southeast-5", "ap-southeast-7",
	// "ca-central-1", "ca-west-1", "eu-central-1", "eu-central-2", "eu-north-1",
	// "eu-south-1", "eu-south-2", "eu-west-1", "eu-west-2", "eu-west-3",
	// "il-central-1", "me-central-1", "mx-central-1", "sa-east-1", "us-east-1",
	// "us-east-2", "us-west-1", "us-west-2", "asia-east1", "asia-northeast1",
	// "asia-southeast1", "europe-north1", "europe-west1", "europe-west3",
	// "europe-west4", "northamerica-northeast2", "southamerica-west1", "us-central1",
	// "us-east1", "us-east4", "us-west1", "us-west2", "australiacentral",
	// "australiaeast", "australiasoutheast", "brazilsouth", "canadacentral",
	// "canadaeast", "centralindia", "centralus", "chinaeast2", "chinanorth",
	// "chinanorth2", "chinanorth3", "eastasia", "eastus", "eastus2", "francecentral",
	// "germanywestcentral", "indonesiacentral", "israelcentral", "italynorth",
	// "japaneast", "japanwest", "koreacentral", "mexicocentral", "newzealandnorth",
	// "northcentralus", "northeurope", "norwayeast", "polandcentral",
	// "southafricanorth", "southcentralus", "southindia", "spaincentral",
	// "swedencentral", "switzerlandnorth", "uaenorth", "uksouth", "ukwest",
	// "westcentralus", "westeurope", "westus", "westus2", "westus3", "fr-strasbourg",
	// "pl-warsaw".
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
