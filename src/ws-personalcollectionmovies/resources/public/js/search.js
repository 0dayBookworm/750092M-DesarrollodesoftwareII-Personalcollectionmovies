$('#Search').click(function() {
    // Deshabilitamos el boton.
    document.getElementById("Search").disabled = true;
    // Obtenemos la lista en la cual pondremos las peliculas de la busqueda.
    var ul = document.getElementById("SearchResult");
    // Removemos los elementos anteriores que la lista tenga.
    while (ul.firstChild) {
        ul.removeChild(ul.firstChild);
    }
    // Obtenemos los datos que pondremos. ($body)
    var title = "BÚSQUEDA";
    var searchFormData = $('#SearchForm').serialize();
    $.ajax({
        type: 'get',
        url: 'https://personalcollectionmovies-alobaton.c9users.io/movie/search',
        data: searchFormData,
        dataType: 'json'
    }).success(function(response) {
        document.getElementById("Search").disabled = false;
        // Validamos que se obtuvo respuesta.
        // Así se accede a los campos del JSon: response.Response
        if (response.total_results > 0) {
            // Obtenemos el objeto del resultado.
            var searchResult = response.Results;
            $.each(searchResult, function(i, item) {
                // Damos formato y acción respectiva a los elementos de la busqueda.
                var movie = item;
                // alert(movie.Title);
                // Creamos el item de lista.
                var li = document.createElement("li");
                li.setAttribute("class", "list-group-item");
                // Creamos el enlace.
                var a = document.createElement("a");
                // Seteamos el enlace de busqueda.
                a.setAttribute("href", "https://personalcollectionmovies-alobaton.c9users.io/movie?ID=" + movie.ID);
                // Damos formato al enlace.

                // Contenedor principal.
                // var divRow = document.createElement("div");
                // divRow.setAttribute("class", "row");

                // Creamos el icono.
                var i = document.createElement("i");
                i.setAttribute("class", "col-sm-1 glyphicon glyphicon-film");

                a.appendChild(i)

                // Creamos el identificador de la pelicula.
                var divMovieInfo = document.createElement("small");

                var title = document.createElement("p");
                title.setAttribute("class", "row");
                title.appendChild(document.createTextNode(movie.Title));
                divMovieInfo.appendChild(title);

                a.appendChild(divMovieInfo);

                //a.appendChild(divRow);
                // Adjuntamos el encale a el item de la lista.
                li.appendChild(a);
                // Adjuntamos el elemento a la lista.
                ul.appendChild(li);
            });
            //alert("Llega aquí");
            $('#SearchModal').modal('show');
        }
        else {
            bootbox.dialog({
                message: response.Message,
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
