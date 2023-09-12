{{template "partials/header.tpl" .}}


{{template "partials/nav.tpl" .}}

<section class="container mt-3" id="imageContainer">
    <div class="row row-cols-1 row-cols-md-2 row-cols-lg-3 g-4 bedroomVillacard">
        {{ range.Hotel.Data }}
        <div class="col">
          <div class="card h-100">
            <div class="img-card">
                <img src="https://cf.bstatic.com{{ .BasicPropertyData.Photos.Main.Urls.ImgURL}}" class="card-img-top"
                alt="Hollywood Sign on The Hill Lazy-loaded"  />

                <span class="btn btn-light btn-money btn-price">From <i class="fa-solid fa-bangladeshi-taka-sign"></i> {{.PriceDisplayInfoIrene.DisplayPrice.AmountPerStay.AmountRounded}}</i></span>
                <div class="btn-mores">
                    <span class="btn btn-light opacity-75 btn-love btn-love"><i class="fa-regular fa-heart"></i></span>
                    <span class="btn btn-light opacity-75 btn-map btn-map"><i class="fa-solid fa-location-dot"></i></span>
                </div>
            
            </div>
            
            <div class="card-body">
              <h6 class="card-title">{{.DisplayName.Title}}</h6>
              <p class="card-text">
                <i class="fa-solid fa-star"></i>
                <i class="fa-solid fa-star"></i>
                <i class="fa-solid fa-star"></i>
                <i class="fa-solid fa-star"></i>
                <i class="fa-regular fa-star"></i>
                683 Reviews
                <h6>Sleeps 85</h6>
                <a href="#" class="textstyle">Arizona > Lake Havasu City</a>
                <div class="text-end">
                  <a href="/search/hotel/details/?id_detail={{.IDDetail}}" class=" btn btn-warning">View Details</a>
                </div>
                <div class="mt-3">
                  <div class="row">
                    <div class="col-auto">
                      <small class="mt-2"><i class="fa-solid fa-check"></i> View  <i class="fa-solid fa-check"></i> Child Friendly</small>
                    </div>
                    <div class="col text-right text-end">
                      <small class="BC">BC - 367639</small>
                    </div>
                </div>
    
                </div>
              </p>
            </div>
          </div>
        </div>
            {{ end }}
        </div>
</section>

{{template "partials/footer-lib.tpl" .}}