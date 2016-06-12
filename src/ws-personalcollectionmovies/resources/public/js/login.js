$(document).ready(function() {
    $('#LoginButton').click(function(){
        // Cancels the form submission
        event.preventDefault();
        // Obtenemos los datos del usuario.
        var LoginRequest = $('#LoginForm').serialize();
        var Title ="Inicio de sesión"
        $.ajax({
            type: 'post',
            url: 'https://personalcollectionmovies-alobaton.c9users.io/login',
            data: LoginRequest,
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
                alert("Iniciaste sesión correctamente.");
                // Redireccionamos a la pagina de inicio.
                location.href = 'https://personalcollectionmovies-alobaton.c9users.io/';
            }
        });
    });
});