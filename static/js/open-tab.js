function openTab(tabName) {
    var i;
    var tabContent = document.getElementsByClassName("tab-content");
    var tabItems = document.getElementsByClassName("tab-item");
    
    for (i = 0; i < tabContent.length; i++) {
        tabContent[i].style.display = "none";
        tabItems[i].classList.remove("active");
    }
    
    document.getElementById(tabName).style.display = "block";
    document.querySelector('[onclick="openTab(\'\')"]').classList.add("active");
}