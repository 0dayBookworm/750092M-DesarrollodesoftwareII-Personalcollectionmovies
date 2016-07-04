$(document).ready(function() {
    // Add
    $('#WatchListAdd').click(function() {
        // Cancels the form submission
        event.preventDefault();
        // var URLactual = window.location;
        // alert(URLactual);
        // Estructura de la peticion: Contiene el ID de la pelicula.
        var busqueda = window.location.search;
        var title ="Peliculas por ver"
        // Send data to back-end
        $.ajax({
            url: 'https://personalcollectionmovies-alobaton.c9users.io/account/watchlist/add'+busqueda
        }).success(function(response) {
            if (response.Status === '999') {
                bootbox.dialog({
                    message: response.Message,
                    title: title,
                    buttons: {
                        close: {
                            label: "Cerrar",
                            className: "btn-default",
                            callback: function() {
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
    });
    // Remove
    $('#WatchListRemove').click(function() {
        alert("Non implemented yet");
    });
});


// function removeChildren(item) {
//     if ( item.hasChildNodes() )
//     {
//         while ( item.childNodes.length > 0 )
//         {
//             item.removeChild( item.firstChild );
//         }
//     }
// }
