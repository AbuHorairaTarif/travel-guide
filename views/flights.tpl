{{template "partials/header.tpl" .}}
{{template "partials/nav.tpl" .}}

    <h1 class="text-center">Flight Search Results</h1>

    {{if .flightResults.Flights}}
        <ul>
            {{range .flightResults.Flights}}
                <li>
                    <strong>Flight ID:</strong> {{.id}}<br>
                    <strong>Departure Date:</strong> {{.data.departureDate}}<br>
                    <strong>Return Date:</strong> {{.data.returnDate}}<br>
                    <strong>Location From:</strong> {{.data.locationFrom}}<br>
                    <strong>Location To:</strong> {{.data.locationTo}}<br>
                    <!-- Add more flight details as needed -->
                </li>
            {{end}}
        </ul>
    {{else}}
        <p>No flight results found.</p>
    {{end}}
{{template "partials/footer-lib.tpl" .}}