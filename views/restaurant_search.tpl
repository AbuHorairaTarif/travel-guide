{{template "partials/header.tpl" .}}
{{template "partials/nav.tpl" .}}

<!-- Restaurant Search Results -->
<section class="search-results">
    <div class="container">
        <h1 class="text-center">Restaurant Search Results</h1>
        <div class="row">
            {{range .Results}}
            <div class="col-md-4 mb-4">
                <div class="card">
                    <img src="{{.Image.URLTemplate | printf .Image.URLTemplate .Image.maxWidth .Image.maxHeight}}" class="card-img-top" alt="Restaurant Image">
                    <div class="card-body">
                        <h5 class="card-title">{{.Title}}</h5>
                        <p class="card-text">{{.SecondaryText}}</p>
                    </div>
                </div>
            </div>
            {{end}}
        </div>
    </div>
</section>

{{template "partials/footer-lib.tpl" .}}
