{{template "partials/header.tpl" .}}
{{template "partials/nav.tpl" .}}

    <h1 class="text-center">Flight Search Results</h1>

    {{ if .Flights }}
    <div class="mt-3">
        <h2 class ="text-center text-info">
            Showing {{ .Length }} flights
        </h2>
        {{ $limit := 1 }}
        {{ range $index, $item := .Flights.Routes }}
            {{ if lt $index $limit }}
            <h2 class="text-center text-warning">from {{ $item.Origin.CityCode }} to {{ $item.Destination.CityCode }} and vice versa</h2>
            {{ end }}
        {{ end }}
    </div>
    <p class="text-center">Price per passenger includes tax and fees</p>

    {{ $Departure_date := .departure_date }}
    {{ $Return_date := .return_date }}
        
    
    {{ range .Flights.Flights }}
     <div class="card">
    <div class="row">
              <div class="col col-12">
               
          <div class="card-body">
          
                    <div class="my-5"><p>ID: {{ .ID }}</p>
                        <p>Departure Date: {{ $Departure_date }}</p>
                        <p>Return Date: {{ $Return_date }}</p>
                    </div>
                        
                    <p class="card-text">
                      
                      {{ $limit := 1 }}
                        {{ range $index, $item := .TravelerPrices }}
                            {{ if lt $index $limit }}
                            <div class="text-end text-danger" style="font-size:24px;">{{ $item.Price.Price.Currency.Code }} {{ $item.Price.Price.Value }}</div>
                            <div class="text-end">Average Price Per Passenger</div>
                            {{ end }}
                        {{ end }}
                     
                      <div class="mt-2 text-end">
                        <a class=" btn btn-warning" href="{{ .ShareableURL }}">Flight Details</a>
                      </div>
                      <div class="mt-3">
                        <div class="row">
                          
                          <div class="col text-right text-end">
                            <small class="BC">Free cancellation within 24 hours of booking completion</small>
                          </div>
                      </div>

                      </div>
                    </p>
                  </div>
                </div>
              </div>
        {{end}}
{{ else }}
    <h3 class="text-center text-danger">No flights available</h3>
{{ end }}
{{template "partials/footer-lib.tpl" .}}