<div class="row">
    <div class="col-md-10 offset-md-1">
        <div class="card shadow">
            <div class="card-body"> 
                <form class="container mb-5">
                    <div class="input-group py-3">
                        <div class="input-group-append bg-white">
                            <button class="btn py-3 rounded-0 rounded-start border-secondary border-end-0 search-border" type="button" id="search-button"><i class="fas fa-search"></i></button>
                        </div>
                        <input type="text" class="form-control py-3 border border-secondary border-start-0 search-border" placeholder="Enter a Destination or Property" aria-label="Search" aria-describedby="search-button">
                    </div>
                    <input type="hidden" id="date-range-slider">
                   <div class="row">
                    <div class="col-7">
                        <div class="t-datepicker">
                            <div class="t-check-in bg-white py-2"></div>
                            <div class="t-check-out bg-white py-2"></div>
                          </div>        
                       </div>
                       <div class="col">
                        <div class="input-group">
                            <!-- Input field for number of guests -->
                            <div class="input-group-append bg-white">
                                <button class="btn py-3 rounded-0 rounded-start border-secondary border-end-0 search-border" type="button" id="search-button"><i class="fa-solid fa-users"></i></button>
                            </div>
                            <input class="form-control py-3 border-secondary search-border border-start-0"  placeholder="Guests" aria-label="Guests" id="guests-input">
                        </div>
                       </div>
                   </div>
                   <div class="bundle">
                    <span class="btn btn-danger text-light rounded-0 p-0 px-1">Bundle & Save</span>
                    <span class="ms-5 text-primary fw-bold">+ Add a Hotel</span>
                   </div>
                       
                    <button class="btn btn-primary search-btn-travel shadow">SEARCH</button>
                </form>
              
            </div>
        </div>
    </div>
</div>

