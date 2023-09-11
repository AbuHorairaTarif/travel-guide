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
    const apiUrl = `/flights/return?locationFrom=${locationFrom}&locationTo=${locationTo}&departureDate=${departureDate}&returnDate=${returnDate}`;
  
    // Fetch flight data from the server
    return fetch(apiUrl)
        .then((response) => {
            if (!response.ok) {
                throw new Error("Network response was not ok");
            }
            return response.text(); // Read response as text
        })
        .then((data) => {
            if (!data) {
                console.error("Received empty JSON response");
                return []; // Return an empty array or handle this case as needed
            }
            return JSON.parse(data); // Parse the JSON data
        })
        .catch((error) => {
            console.error("There was a problem with the fetch operation:", error);
        });
}
  
  // Function to update flight search results in the DOM
  function updateCardItems(flights) {
    const flightList = document.getElementById("flightResults");

  
    // Clear the existing flight list
    flightList.innerHTML = "";
  
    // Loop through the flight data and create card items
    flights.forEach((flight, index) => {
      const cardItem = document.createElement("div");
      cardItem.classList.add("card");
      cardItem.innerHTML = `
        <div class="card-header">
          Flight ${index + 1}
        </div>
        <div class="card-body">
          <h5 class="card-title">Flight Details</h5>
          <p class="card-text">Flight Number: ${flight.flightNumber}</p>
          <p class="card-text">Departure Time: ${flight.departuredAt}</p>
          <p class="card-text">Arrival Time: ${flight.arrivedAt}</p>
          <p class="card-text">Duration: ${flight.duration} seconds</p>
          <p class="card-text">Origin: ${flight.origin.name}</p>
          <p class="card-text">Destination: ${flight.destination.name}</p>
          <p class="card-text">Operating Carrier: ${flight.operatingCarrier.name}</p>
          <p class="card-text">Marketing Carrier: ${flight.marketingCarrier.name}</p>
          <p class="card-text">Cabin Class: ${flight.cabinClassName}</p>
          <p class="card-text">Number of Technical Stops: ${flight.numberOfTechnicalStops}</p>
        </div>
      `;
  
      flightList.appendChild(cardItem);
    });
  }
  
 
  
