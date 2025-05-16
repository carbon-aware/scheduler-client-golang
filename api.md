# carbonaware

Response Types:

- <a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang">carbonaware</a>.<a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang#GetResponse">GetResponse</a>

Methods:

- <code title="get /">client.<a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang#CarbonawareService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang">carbonaware</a>.<a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang#GetResponse">GetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Schedule

Params Types:

- <a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang">carbonaware</a>.<a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang#CloudZoneParam">CloudZoneParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang">carbonaware</a>.<a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang#CloudZone">CloudZone</a>
- <a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang">carbonaware</a>.<a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang#ScheduleOption">ScheduleOption</a>
- <a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang">carbonaware</a>.<a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang#ScheduleNewResponse">ScheduleNewResponse</a>

Methods:

- <code title="post /v0/schedule/">client.Schedule.<a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang#ScheduleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang">carbonaware</a>.<a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang#ScheduleNewParams">ScheduleNewParams</a>) (<a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang">carbonaware</a>.<a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang#ScheduleNewResponse">ScheduleNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Regions

Response Types:

- <a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang">carbonaware</a>.<a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang#RegionListResponse">RegionListResponse</a>

Methods:

- <code title="get /v0/regions/">client.Regions.<a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang#RegionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang">carbonaware</a>.<a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang#RegionListResponse">RegionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang">carbonaware</a>.<a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang#HealthCheckResponse">HealthCheckResponse</a>

Methods:

- <code title="get /health">client.Health.<a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang#HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang">carbonaware</a>.<a href="https://pkg.go.dev/github.com/carbon-aware/scheduler-client-golang#HealthCheckResponse">HealthCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
