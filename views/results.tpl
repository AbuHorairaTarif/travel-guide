{{template "partials/header.tpl" .}}
{{template "partials/nav.tpl" .}}
<h1>Search Results</h1>

    <!-- Display the search parameters -->
    <p>Category: <span id="category"></span></p>
    <p>Location: <span id="location"></span></p>
    <p>Check-In Date: <span id="checkInDate"></span></p>
    <p>Check-Out Date: <span id="checkOutDate"></span></p>
    <p>Guests: <span id="guests"></span></p>



{{template "partials/footer-lib.tpl" .}}