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
    const location_from = document.getElementById("location-from").value;
    const location_to = document.getElementById("location-to").value;
    const departure_date = document.getElementById("departure-date").value;
    const return_date = document.getElementById("return-date").value;
    const data = null;

const xhr = new XMLHttpRequest();
// xhr.withCredentials = true;

// xhr.addEventListener('readystatechange', function () {
// 	if (this.readyState === this.DONE) {
// 		console.log(this.responseText);
// 	}
// });

// xhr.open('GET', 'https://localhost:8080/flights/return?location_from=Luanda%2C%20Angola&location_to=Paris%2C%20France&departure_date=2023-10-01&return_date=2023-10-20&page=1&country_flag=us');
// xhr.setRequestHeader('X-RapidAPI-Key', '16e4aeb170mshafe993eedee5cdep154eb3jsn8d818903b1df');
// xhr.setRequestHeader('X-RapidAPI-Host', 'booking-com13.p.rapidapi.com');

xhr.open('GET', '/flights/return?location_from=' + encodeURIComponent(location_from) +
        '&location_to=' + encodeURIComponent(location_to) +
        '&departure_date=' + encodeURIComponent(departure_date) +
        '&return_date=' + encodeURIComponent(return_date), true)+"&page=1&country_flag=us";

xhr.setRequestHeader('X-RapidAPI-Key', '16e4aeb170mshafe993eedee5cdep154eb3jsn8d818903b1df');
xhr.setRequestHeader('X-RapidAPI-Host', 'booking-com13.p.rapidapi.com');


        xhr.onreadystatechange = function () {
            if (xhr.readyState === 4 && xhr.status === 200) {
                console.log(this.responseText);
            }
        };
    
        xhr.send(data);
        // Perform the flight search and update the card-items with flight data
        // fetchFlightData().then(function (flightData) {
        //     updateCardItems(flightData);
        // });
        // window.location.href = '/flights/return';

    // const locationFrom = document.getElementById("location-from").value;
    // const locationTo = document.getElementById("location-to").value;
    // const departureDate = document.getElementById("departure-date").value;
    // const returnDate = document.getElementById("return-date").value;

    // // Construct the URL
    // const apiUrl = `/flights/return?location_from=${locationFrom}&location_to=${locationTo}&departure_date=${departureDate}&return_date=${returnDate}&page=1&country_flag=us`;

    // // Fetch flight data from the server
    // fetch(apiUrl)
    //     .then((response) => {
    //         if (!response.ok) {
    //             throw new Error("Network response was not ok");
    //         }

    //         const contentType = response.headers.get("content-type");
    //         if (contentType && contentType.includes("application/json")) {
    //             return response.json(); // Parse the response as JSON
    //         } else {
    //             console.error("Received non-JSON response");
    //             return []; // Return an empty array or handle this case as needed
    //         }
    //     })
    //     .catch((error) => {
    //         console.error("There was a problem with the fetch operation:", error);
    //     });
    }
});

// Function to fetch flight data from the server
function fetchFlightData() {
    // Get the form values
//     const locationFrom = document.getElementById("location-from").value;
//     const locationTo = document.getElementById("location-to").value;
//     const departureDate = document.getElementById("departure-date").value;
//     const returnDate = document.getElementById("return-date").value;

//     const data = null;

// const xhr = new XMLHttpRequest();
// xhr.withCredentials = true;

// xhr.addEventListener('readystatechange', function () {
// 	if (this.readyState === this.DONE) {
// 		console.log(this.responseText);
// 	}
// });

// xhr.open('GET', 'https://booking-com13.p.rapidapi.com/flights/return?location_from=Luanda%2C%20Angola&location_to=Paris%2C%20France&departure_date=2023-10-01&return_date=2023-10-20&page=1&country_flag=us');
// xhr.setRequestHeader('X-RapidAPI-Key', '16e4aeb170mshafe993eedee5cdep154eb3jsn8d818903b1df');
// xhr.setRequestHeader('X-RapidAPI-Host', 'booking-com13.p.rapidapi.com');

// xhr.send(data);


    
    // xhr.open('GET', '/flights/return?location_from=' + encodeURIComponent(locationFrom) +
    //     '&location_to=' + encodeURIComponent(locationTo) +
    //     '&departure_date=' + encodeURIComponent(departureDate) +
    //     '&return_date=' + encodeURIComponent(returnDate), true);

    //     xhr.onreadystatechange = function () {
    //         if (xhr.readyState === 4 && xhr.status === 200) {
    //             document.getElementById('search-results').innerHTML = xhr.responseText;
    //         }
    //     };
    
    //     xhr.send(data);
    // Construct the URL
    const apiUrl = `/flights/return?location_from=${location_from}&location_to=${location_to}&departure_date=${departure_date}&return_date=${return_date}&page=1&country_flag=us`;

    // Fetch flight data from the server
    // fetch(apiUrl)
    //     .then((response) => {
    //         if (!response.ok) {
    //             throw new Error("Network response was not ok");
    //         }

    //         const contentType = response.headers.get("content-type");
    //         if (contentType && contentType.includes("application/json")) {
    //             return response.json(); // Parse the response as JSON
    //         } else {
    //             console.error("Received non-JSON response");
    //             return []; // Return an empty array or handle this case as needed
    //         }
    //     })
    //     .catch((error) => {
    //         console.error("There was a problem with the fetch operation:", error);
    //     });
}

// // Function to update flight search results in the DOM
// function updateCardItems(flights) {
//     const flightList = document.getElementById("flightResults");

//     // Clear the existing flight list
//     while (flightList.firstChild) {
//         flightList.removeChild(flightList.firstChild);
//     }

//     // Loop through the flight data and create card items
//     flights.forEach((flight, index) => {
//         const cardItem = document.createElement("div");
//         cardItem.classList.add("card");

//         const cardHeader = document.createElement("div");
//         cardHeader.classList.add("card-header");
//         cardHeader.textContent = `Flight ${index + 1}`;

//         const cardBody = document.createElement("div");
//         cardBody.classList.add("card-body");

//         const cardTitle = document.createElement("h5");
//         cardTitle.classList.add("card-title");
//         cardTitle.textContent = "Flight Details";

//         const flightDetails = [
//             `Flight Number: ${flight.origin.countryName}`,
//             `Departure Time: ${flight.departuredAt}`,
//             `Arrival Time: ${flight.departureDate}`,
//             `Duration: ${flight.flightDetails} seconds`,
//             `Origin: ${flight.origin.name}`,
//             `Destination: ${flight.destination.name}`,
//             `Cabin Class: ${flight.cabinClassName}`
//         ];

//         flightDetails.forEach(detail => {
//             const paragraph = document.createElement("p");
//             paragraph.classList.add("card-text");
//             paragraph.textContent = detail;
//             cardBody.appendChild(paragraph);
//         });

//         cardItem.appendChild(cardHeader);
//         cardItem.appendChild(cardBody);

//         flightList.appendChild(cardItem);
//     });
// }
