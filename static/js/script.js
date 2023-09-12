// // Add an event listener to the category buttons
// document.querySelectorAll('.category-button').forEach(function(button) {
//     button.addEventListener('click', function() {
//         // Remove active class from all buttons
//         document.querySelectorAll('.category-button').forEach(function(btn) {
//             btn.classList.remove('active');
//         });

//         // Add active class to the clicked button
//         button.classList.add('active');

//         const category = button.getAttribute('data-category');
//         toggleInputFields(category);
//     });
// });

// // Function to toggle input fields based on category
// function toggleInputFields(category) {
//     // Get all input fields
//     const inputFields = document.querySelectorAll('.form-control');

//     // Hide all input fields
//     inputFields.forEach(function(input) {
//         input.style.display = 'none';
//     });

//     // Show input fields based on category
//     if (category === 'hotels') {
//         document.getElementById('location').style.display = 'block';
//         document.getElementById('check-in').style.display = 'block';
//         document.getElementById('check-out').style.display = 'block';
//         document.getElementById('guests').style.display = 'block';
//         document.getElementById('head-text').innerHTML = 'Are you looking for hotels ?';
//     } else if (category === 'restaurants') {
//         document.getElementById('location').style.display = 'block';
//         document.getElementById('head-text').innerHTML = 'Are you looking for Restaurants ? Hurry Up!!!';
//     } else if (category === 'flights') {
//         document.getElementById('location-from').style.display = 'block';
//         document.getElementById('location-to').style.display = 'block';
//         document.getElementById('departure-date').style.display = 'block';
//         document.getElementById('return-date').style.display = 'block';
//         document.getElementById('head-text').innerHTML = 'Are you looking for Air Trip ?';
//     }
// }

// // Initialize input fields based on the initial category (Hotels)
// toggleInputFields('hotels');

// // Add the 'active' class to the initial active button (Hotels)
// document.querySelector('.category-button[data-category="hotels"]').classList.add('active');


// Add an event listener to the category buttons
document.querySelectorAll('.category-button').forEach(function (button) {
    button.addEventListener('click', function () {
        // Remove active class from all buttons
        document.querySelectorAll('.category-button').forEach(function (btn) {
            btn.classList.remove('active');
        });

        // Add active class to the clicked button
        button.classList.add('active');

        const category = button.getAttribute('data-category');
        toggleInputFields(category);
    });
});

// Function to toggle input fields based on category
function toggleInputFields(category) {
    // Get all input fields
    const inputFields = document.querySelectorAll('.form-control');

    // Hide all input fields
    inputFields.forEach(function (input) {
        input.style.display = 'none';
    });

    // Show input fields based on category
    if (category === 'hotels') {
        document.getElementById('location').style.display = 'block';
        document.getElementById('check-in').style.display = 'block';
        document.getElementById('check-out').style.display = 'block';
        document.getElementById('guests').style.display = 'block';
        document.getElementById('head-text').innerHTML = 'Are you looking for hotels ?';
    } else if (category === 'restaurants') {
        document.getElementById('location').style.display = 'block';
        document.getElementById('head-text').innerHTML = 'Are you looking for Restaurants ? Hurry Up!!!';
    } else if (category === 'flights') {
        document.getElementById('location-from').style.display = 'block';
        document.getElementById('location-to').style.display = 'block';
        document.getElementById('departure-date').style.display = 'block';
        document.getElementById('return-date').style.display = 'block';
        document.getElementById('head-text').innerHTML = 'Are you looking for Air Trip ?';
    }
}

// Initialize input fields based on the initial category (Hotels)
toggleInputFields('hotels');

// Add the 'active' class to the initial active button (Hotels)
document.querySelector('.category-button[data-category="hotels"]').classList.add('active');

// Add an event listener to the submit button
document.getElementById('search-button').addEventListener('click', function () {
    // Check the active category
    const activeCategory = document.querySelector('.category-button.active').getAttribute('data-category');

    if (activeCategory === 'flights') {
        // Perform the flight search and update the card-items with flight data
        fetchFlightData().then(function (flightData) {
            updateCardItems(flightData);
        });
        // window.location.href = '/flights/return';
    }

    else if (activeCategory === 'restaurants') {
        // Perform the flight search and update the card-items with flight data
        
    }
    else if (activeCategory === 'hotels') {
        // Perform the flight search and update the card-items with flight data
        
    }
});

// Function to fetch flight data from the server
function fetchFlightData() {
    // Get the form values
    const locationFrom = document.getElementById("location-from").value;
    const locationTo = document.getElementById("location-to").value;
    const departureDate = document.getElementById("departure-date").value;
    const returnDate = document.getElementById("return-date").value;

    // Construct the URL
    const apiUrl = `/flights/return?location_from=${locationFrom}&location_to=${locationTo}&departure_date=${departureDate}&return_date=${returnDate}&page=1&country_flag=us`;

    // Fetch flight data from the server
    return fetch(apiUrl)
        .then((response) => {
            if (!response.ok) {
                throw new Error("Network response was not ok");
            }

            const contentType = response.headers.get("content-type");
            if (contentType && contentType.includes("application/json")) {
                return response.json(); // Parse the response as JSON
            } else {
                console.error("Received non-JSON response");
                return []; // Return an empty array or handle this case as needed
            }
        })
        .catch((error) => {
            console.error("There was a problem with the fetch operation:", error);
        });
}

  
  // Function to update flight search results in the DOM
function updateCardItems(flights) {
    const flightList = document.getElementById("flightResults");

    // Clear the existing flight list
    while (flightList.firstChild) {
        flightList.removeChild(flightList.firstChild);
    }

    // Loop through the flight data and create card items
    flights.forEach((flight, index) => {
        const cardItem = document.createElement("div");
        cardItem.classList.add("card");

        const cardHeader = document.createElement("div");
        cardHeader.classList.add("card-header");
        cardHeader.textContent = `Flight ${index + 1}`;

        const cardBody = document.createElement("div");
        cardBody.classList.add("card-body");

        const cardTitle = document.createElement("h5");
        cardTitle.classList.add("card-title");
        cardTitle.textContent = "Flight Details";

        const flightDetails = [
            `Flight Number: ${flight.flightNumber}`,
            `Departure Time: ${flight.departuredAt}`,
            `Arrival Time: ${flight.arrivedAt}`,
            `Duration: ${flight.duration} seconds`,
            `Origin: ${flight.origin.name}`,
            `Destination: ${flight.destination.name}`,
            `Operating Carrier: ${flight.operatingCarrier.name}`,
            `Marketing Carrier: ${flight.marketingCarrier.name}`,
            `Cabin Class: ${flight.cabinClassName}`,
            `Number of Technical Stops: ${flight.numberOfTechnicalStops}`
        ];

        flightDetails.forEach(detail => {
            const paragraph = document.createElement("p");
            paragraph.classList.add("card-text");
            paragraph.textContent = detail;
            cardBody.appendChild(paragraph);
        });

        cardItem.appendChild(cardHeader);
        cardItem.appendChild(cardBody);

        flightList.appendChild(cardItem);
    });
}

  
 
  
