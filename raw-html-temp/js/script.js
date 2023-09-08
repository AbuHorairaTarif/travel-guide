document.addEventListener("DOMContentLoaded", function () {
    // Get the category buttons and search inputs
    const categoryButtons = document.querySelectorAll(".category-button");
    const searchInputs = document.querySelectorAll(".search-inputs input");
    const searchButton = document.getElementById("search-button");

    // Initialize with the "Flights" category
    let selectedCategory = "flights";
    toggleCategory(selectedCategory);

    // Add click event listeners to category buttons
    categoryButtons.forEach((button) => {
        button.addEventListener("click", () => {
            selectedCategory = button.getAttribute("data-category");
            toggleCategory(selectedCategory);
        });
    });

    // Function to toggle search inputs based on the selected category
    function toggleCategory(category) {
        // Remove the "active" class from all category buttons
        categoryButtons.forEach((button) => {
            button.classList.remove("active");
        });
        
        // Add the "active" class to the selected category button
        document.querySelector(`[data-category="${category}"]`).classList.add("active");
    }

    // Add click event listener to the "Search" button
    searchButton.addEventListener("click", () => {
        const location = document.getElementById("location").value;
        const checkInDate = document.getElementById("check-in").value;
        const checkOutDate = document.getElementById("check-out").value;
        const guests = document.getElementById("guests").value;

        // Simulate the search functionality (replace with actual API calls)
        let message = `Searching for ${selectedCategory}:\n`;
        message += `Location: ${location}\n`;
        message += `Check-In Date: ${checkInDate}\n`;
        message += `Check-Out Date: ${checkOutDate}\n`;
        message += `Guests: ${guests}`;

        alert(message);
    });
});
