document.getElementById("loading").style.display = "block";

var projs = new Vue({
    el: "#projs",
    data: {
	projects: []
    }
});

var xhr = new XMLHttpRequest();
xhr.open("GET", "/api/projects", true);
xhr.onload = function (e) {
    if (xhr.readyState === 4) {
	if (xhr.status === 200) {
	    document.getElementById("loading").style.display = "none";
	    document.getElementById("projs").style.display = "flex";
	    projs.projects = JSON.parse(xhr.responseText);
	}
    }
}

xhr.send(null);
