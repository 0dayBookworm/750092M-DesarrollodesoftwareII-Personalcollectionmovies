$(document).ready(function() {
    $('#LogoutButton').click(function(){
        // Cancels the form submission
        event.preventDefault();
        // Obtenemos los datos del usuario.
        var LogoutRequest = "";
        var Title ="CIERRE DE SESIÓN"
        $.ajax({
            type: 'post',
            url: 'http://personalcollectionmovies-alobaton.c9users.io/logout',
            data: LogoutRequest,
            dataType: 'json'
        }).success(function(response) {
            if (response.Status === '999') {
                // Mostramos el o los mensajes de error.
                bootbox.dialog({
                    message: response.Message,
                    title: Title,
                    buttons: {
                        close: {
                            label: "Cerrar",
                            className: "btn-warning",
                            callback: function() {}
                        }
                    }
                });

            }
            else {
                // Mensaje de alerta de pruebas, debera ser removido antes de exponer el demo.
                alert("Se ha cerrado la sessión.");
                // Redireccionamos a la pagina de inicio.
                location.href = 'http://personalcollectionmovies-alobaton.c9users.io/';
            }
        });
    });
});