$('#ViewListAdd').click(function() {
   $('#PlacesModal').modal('show'); 
});
// Remove
$('#ViewListRemove').click(function() {
    alert("Non implemented yet");
});
$('#Places').click(function() {
    // Deshabilitamos el boton.
    document.getElementById("Places").disabled = true;
    // Obtenemos la lista en la cual pondremos las peliculas de la busqueda.
    var ul = document.getElementById("PlacesResult");
    // Removemos los elementos anteriores que la lista tenga.
    while (ul.firstChild) {
        ul.removeChild(ul.firstChild);
    }
    // Obtenemos los datos que pondremos. ($body)
    var title = "SELECCONA UNA UBICACIÓN";
    var PlacesFormData = $('#PlacesForm').serialize();
    $.ajax({
        type: 'POST',
        url: 'http://personalcollectionmovies-alobaton.c9users.io/places',
        data: PlacesFormData,
        dataType: 'json'
    }).success(function(response) {
        document.getElementById("Places").disabled = false;
         if (response.status === "OK") {
             
            var searchResult = response.predictions;
            $.each(searchResult, function(idA, item) {
                var place = item;
                
                // Creamos el item de lista.
                var li = document.createElement("li");
                li.setAttribute("class", "list-group-item");
                
                var a = document.createElement("a");
                // Seteamos el enlace de busqueda.
                var busqueda = window.location.search;
                a.id = idA;
                a.setAttribute("name", idA)
                a.setAttribute("onclick", 'addViewListPost("'+busqueda+'", "'+place.description+ '","'+idA+'")');
                // a.setAttribute("data-toggle", "popover");
                // a.setAttribute("title", "PELÍCULAS VISTAS");
                // a.setAttribute("data-content", "");
                 // Creamos el icono.
                var i = document.createElement("i");
                i.setAttribute("class", "col-sm-1 glyphicon glyphicon-map-marker");

                a.appendChild(i)
                
                // Creamos el identificador de la pelicula.
                var divPlaceInfo = document.createElement("small");

                var description = document.createElement("p");
                description.setAttribute("class", "row");
                description.appendChild(document.createTextNode(place.description));
                divPlaceInfo.appendChild(description);

                a.appendChild(divPlaceInfo);

                //a.appendChild(divRow);
                // Adjuntamos el encale a el item de la lista.
                li.appendChild(a);
                // Adjuntamos el elemento a la lista.
                ul.appendChild(li);
            });
         } else{
             bootbox.dialog({
                message: "No se obtuvieron resultados para la busqueda.",
                title: title,
                buttons: {
                    close: {
                        label: "Cerrar",
                        className: "btn-warning",
                        callback: function() {}
                    }
                }
            });
         }
    });
});


function addViewListPost(pBusqueda, pPlace, pID){
    var queryDict = {}
    pBusqueda.substr(1).split("&").forEach(function(item) {queryDict[item.split("=")[0]] = item.split("=")[1]})
    var json = '{"ID": "'+queryDict["ID"]+'", "Place":"'+pPlace+'"}'; 
    var request =JSON.parse(json);
    var title = "PELÍCULAS VISTAS"
    $.ajax({
        type: 'post',
        url: 'http://personalcollectionmovies-alobaton.c9users.io/account/viewlist/add',
        data: request,
        dataType: 'json'
    }).success(function(response) {
        if (response.Status === '999') {
            
            bootbox.dialog({
                message: response.Message,
                title: title,
                buttons: {
                    close: {
                    label: "Cerrar",
                    className: "btn-default",
                        back: function() {
                        }
                    }
                }
            });
        }
        else {
            bootbox.dialog({
                message: response.Message,
                title: title,
                buttons: {
                    close: {
                        label: "Cerrar",
                        className: "btn-success",
                        callback: function() {
                            window.location.replace(window.location);
                        }
                    }
                }
            });
        }
    });
}