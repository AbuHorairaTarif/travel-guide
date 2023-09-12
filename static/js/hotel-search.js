
$(document).ready(function () {
    $("#search-button").click(function () {
        var destination = $("#location").val();
        var checkin = $("#check-in").val();
        var checkout = $("#check-out").val();

        $.ajax({
            type: "GET",
            url: "/search/hotellist", // URL of your Beego controller
            data: {
                destination: destination,
                "t-start": checkin,
                "t-end": checkout
            },
            success: function (data) {
                // Handle the data received from the controller
                // 'data' contains the image URLs in JSON format
                if (data.length > 0) {
                    // Clear any existing images in the container
                    $("#imageContainer").empty();
                    
                    // Iterate through the image URLs
                    for (var i = 0; i < data.length; i++) {
                        var imageUrl = data[i];
                        console.log("Image URL " + (i + 1) + ": " + imageUrl);

                        // Create an <img> element and set its 'src' attribute
                        var imgElement = $("<img>");
                        imgElement.attr("src", imageUrl);
                        imgElement.attr("alt", "Hotel Image");

                        // Append the <img> element to the 'imageContainer'
                        $("#imageContainer").append(imgElement);
                    }
                } else {
                    console.log("No image URLs found.");
                    // Optionally, you can display a message indicating no images were found.
                }
            },
            error: function () {
                alert("Error while fetching image URLs.");
            },
        });
    });
});
