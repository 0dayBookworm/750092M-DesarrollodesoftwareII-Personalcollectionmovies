
$('#LoginForm')
    .bootstrapValidator({
        message: 'Este valor no es valido',
        feedbackIcons: {
            valid: 'glyphicon glyphicon-ok',
            invalid: 'glyphicon glyphicon-remove',
            validating: 'glyphicon glyphicon-refresh'
        },
        fields: {
            Username: {
                validators: {
                    notEmpty: {
                        message: 'El nombre de usuario es requerido'
                    }
                }
            },
            Password: {
                validators: {
                    notEmpty: {
                        message: 'La contraseña es requerida'
                    }
                }
            }
        }
    });
$('#LoginButton').click(function() {
    event.preventDefault();
    login();
});

function login() {
    // Obtenemos los datos del usuario.
    var LoginRequest = $('#LoginForm').serialize();
    var Title = "INICIO DE SESIÓN"
    $.ajax({
        type: 'POST',
        url: 'http://personalcollectionmovies-alobaton.c9users.io/login',
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
            // alert("Iniciaste sesión correctamente.");
            // Redireccionamos a la pagina en la que se encontraba. Exccepto si se encontraba en registro.
            if(window.location.href.indexOf("/register") > -1) {
                location.href = 'http://personalcollectionmovies-alobaton.c9users.io/';
            } else {
                location.href = window.location; 
            }
        }
    });
}

function checkEnter(e) {
    e = e || event;
    var txtArea = /textarea/i.test((e.target || e.srcElement).tagName);
    var isEnter = txtArea || (e.keyCode || e.which || e.charCode || 0) == 13;
    if (isEnter) { 
        login();
    }
}