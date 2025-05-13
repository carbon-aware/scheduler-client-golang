# carbonaware

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go">carbonaware</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go#GetResponse">GetResponse</a>

Methods:

- <code title="get /">client.<a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go#CarbonawareService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go">carbonaware</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go#GetResponse">GetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Schedule

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go">carbonaware</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go#CloudZoneParam">CloudZoneParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go">carbonaware</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go#CloudZone">CloudZone</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go">carbonaware</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go#ScheduleOption">ScheduleOption</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go">carbonaware</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go#ScheduleNewResponse">ScheduleNewResponse</a>

Methods:

- <code title="post /v0/schedule/">client.Schedule.<a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go#ScheduleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go">carbonaware</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go#ScheduleNewParams">ScheduleNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go">carbonaware</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go#ScheduleNewResponse">ScheduleNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Regions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go">carbonaware</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go#RegionListResponse">RegionListResponse</a>

Methods:

- <code title="get /v0/regions/">client.Regions.<a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go#RegionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go">carbonaware</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go#RegionListResponse">RegionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go">carbonaware</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go#HealthCheckResponse">HealthCheckResponse</a>

Methods:

- <code title="get /health">client.Health.<a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go#HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go">carbonaware</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/carbonaware-go#HealthCheckResponse">HealthCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
