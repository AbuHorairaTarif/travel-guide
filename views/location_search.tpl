{{template "partials/header.tpl" .}}

{{template "partials/nav.tpl" .}}

<h1 class="text-center">Location Search Results</h1>

{{if .Results}}
<div class="container">
    <div class="row">
        {{range .Results}}
        <div class="col-md-4 mb-4">
            <div class="card">
                <div class="card-body">
                    <p class="card-title">{{.Title}}</p>
                    <p class="card-text">Geo ID: {{.GeoID}}</p>
                    <p class="card-text">Document ID: {{.DocumentID}}</p>
                    <p class="card-text">Tracking Items: {{.TrackingItems}}</p>
                    <p class="card-text">Secondary Text: {{.SecondaryText}}</p>
                </div>
            </div>
        </div>
        {{end}}
    </div>
</div>
{{else}}
<div class="container">
    <p>No results found.</p>
</div>
{{end}}

{{template "partials/footer-lib.tpl" .}}
