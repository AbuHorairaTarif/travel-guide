{{template "partials/header.tpl" .}}

{{template "partials/nav.tpl" .}}



<header>
  <div class="centered-div-container">
  
        <div class="centered-div">
            <div class="container-fluid">
            <form  action="/flights/return" method="get">
            <h2 class= "text-center text-warning" id="head-text">Are Looking For Hotel ?</h2>
            <div class="row">
            
            <div class="col col-md-12">
                
            <!-- Multipurpose search box with Bootstrap 5 classes -->
            <div class="input-group">
            <div class="btn-group" role="group">
       
            <span class="bg-warning px-6 py-2 rounded-pill">
            <button type="button" class="btn btn-outline-secondary category-button" data-category="hotels"><i class="fa fa-home" aria-hidden="true"></i>Hotels & Homes</button>
            <button type="button" class="btn btn-outline-secondary category-button" data-category="restaurants"><i class="fa fa-hotel" aria-hidden="true"></i>Restaurants</button>
            <button type="button" class="btn btn-outline-secondary category-button" data-category="flights"><i class="fa fa-fighter-jet" aria-hidden="true"></i> Flights</button>
            </span>
            </div>
            <div class="input-group input-group-t">
    
        <input type="text" class="form-control" id="location" placeholder="Enter a location">
        <input type="date" class="form-control" id="check-in" placeholder="Check-In Date">
        <input type="date" class="form-control" id="check-out" placeholder="Check-Out Date">
        <input type="number" class="form-control" id="guests" placeholder="Guests">

<!-- Additional input fields for flights -->
        <input type="text" class="form-control" id="location-from" placeholder="Source Location" style="display: none;" required>
        <input type="text" class="form-control" id="location-to" placeholder="Destination Location" style="display: none;" required>
        <input type="date" class="form-control" id="departure-date" placeholder="Departure Date" style="display: none;" required>
        <input type="date" class="form-control" id="return-date" placeholder="Return Date" style="display: none;" required>

        <button class="btn btn-primary rounded-pill mx-1" type="submit" id="search-button"><i class="fa fa-search" aria-hidden="true"></i></button>

    </div>
</div>
                </div>
            </div>
        </div>
        </form>
        </div>
    </div>  
    <div id="search-results">
    </div>
    </header>
  


{{template "partials/card-item.tpl" .}}
{{template "partials/pricing-section.tpl" .}}
{{template "partials/footer-section.tpl" .}}

{{template "partials/footer-lib.tpl" .}}