document.addEventListener("DOMContentLoaded", function () {
    // Get the category buttons and search inputs
    const categoryButtons = document.querySelectorAll(".category-button");
    const searchButton = document.getElementById("search-button");

    // Add click event listeners to category buttons
    categoryButtons.forEach((button) => {
        button.addEventListener("click", () => {
            const selectedCategory = button.getAttribute("data-category");
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

        // Make an AJAX request to the Beego controller
        const xhr = new XMLHttpRequest();
        xhr.open("GET", `/searchLocation?query=${location}`, true);
        xhr.onreadystatechange = function () {
            if (xhr.readyState === XMLHttpRequest.DONE) {
                if (xhr.status === 200) {
                    // Load the location_search.tpl template with the search results
                    document.body.innerHTML = xhr.responseText;
                } else {
                    // Handle the error
                    alert("Error: Unable to retrieve search results.");
                }
            }
        };
        xhr.send();
    });
});
