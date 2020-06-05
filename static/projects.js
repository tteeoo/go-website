var projs = new Vue({
    el: "#projs",
    data: {
	projects: []
    }
})

var xhr = new XMLHttpRequest();
xhr.open("GET", "/api/projects", true)
xhr.onload = function (e) {
    if (xhr.readyState === 4) {
	if (xhr.status === 200) {
	    projs.projects = JSON.parse(xhr.responseText)
	}
    }
}

xhr.send(null);
