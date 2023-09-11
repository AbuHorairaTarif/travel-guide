<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Flight Search Results</title>
</head>
<body>
    <h1>Flight Search Results</h1>

    {{if .FlightResults}}
        <ul>
            {{range .FlightResults}}
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

    
</body>
</html>
