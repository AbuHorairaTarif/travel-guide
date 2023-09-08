{{template "partials/header.tpl" .}}
{{template "partials/nav.tpl" .}}

<div class="container">
    <h2 class="text-center">Hotel Search Results</h2>

    <div class="row">
        {{range .Results}}
        <div class="col-md-4 mb-4">
            <div class="card">
                <img src="{{.Image.URLTemplate | printf "https://dynamic-media-cdn.tripadvisor.com/media/photo-o/26/d1/94/e6/exterior.jpg?w=%d&h=%d&s=1" 300 200}}" class="card-img-top" alt="{{.Title}}">
                <div class="card-body">
                    <h5 class="card-title">{{.Title}}</h5>
                    <p class="card-text">{{.SecondaryText}}</p>
                </div>
            </div>
        </div>
        {{end}}
    </div>
</div>

{{template "partials/footer-lib.tpl" .}}
