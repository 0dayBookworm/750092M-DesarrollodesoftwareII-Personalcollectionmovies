$(document).ready(function() {
     // Validaciones del formulario.
    $('#ProfileForm').bootstrapValidator({
            message: 'Este valor no es valido',
            feedbackIcons: {
                valid: 'glyphicon glyphicon-ok',
                invalid: 'glyphicon glyphicon-remove',
                validating: 'glyphicon glyphicon-refresh'
            },
            fields: {
                Email: {
                    validators: {
                        notEmpty: {
                            message: 'El Email es requerido'
                        },
                        emailAddress: {
                            message: 'El correo electronico no es valido'
                        }
                    }
                },
                FirstName: {
                    validators: {
                        notEmpty: {
                            message: 'El primer nombre es requerido'
                        }
                    }
                },
                LastName: {
                    validators: {
                        notEmpty: {
                            message: 'Los apellidos son requeridos'
                        }
                    }
                },
                BirthDate: {
                    validators: {
                        notEmpty: {
                            message: 'La fecha de nacimiento es requerida'
                        },
                        date: {
                            format: 'YYYY-MM-DD',
                            message: 'La fecha de nacimiento no es valida'
                        }
                    }
                },
                Gender: {
                    validators: {
                        notEmpty: {
                            message: 'El genero es requerido'
                        }
                    }
                }
            }
    });
    // UpdateProfile
    $('#UpdateProfile').click(function(){
        // Cancels the form submission
        event.preventDefault();
        // Estructura de la peticion.
        var UpdateProfileRequest = $('#ProfileForm').serialize();
        var title ="Perfil";
        // Send data to back-end
        $.ajax({
            type: 'post',
            url: 'https://personalcollectionmovies-alobaton.c9users.io/account/profile/update',
            data: UpdateProfileRequest,
            dataType: 'json'
        }).success(function(response) {
            if (response.Status === '999') {
                // Mostramos el o los mensajes de error.
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
            else {
                bootbox.dialog({
                    message: response.Message,
                    title: title,
                    buttons: {
                        close: {
                            label: "Cerrar",
                            className: "btn-success",
                            callback: function() {
                                location.href = 'https://personalcollectionmovies-alobaton.c9users.io/account/profile';
                            }
                        }
                    }
                });
            }
        });
    });
    
    // ChangePassword
    $('#ChangePassword').click(function(){
        // Cancels the form submission
        event.preventDefault();
        // Estructura de la peticion.
        var ChangePasswordRequest = $('#ChangePasswordForm').serialize();
        var Title ="Cambiar Contraseña";
        // Send data to back-end
        $.ajax({
            type: 'post',
            url: 'https://personalcollectionmovies-alobaton.c9users.io/account/profile/password',
            data: ChangePasswordRequest,
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
                bootbox.dialog({
                    message: response.Message,
                    title: Title,
                    buttons: {
                        close: {
                            label: "Cerrar",
                            className: "btn-success",
                            callback: function() {
                                // Si el proceso de registro se llevo a cabo correctamente redireccionamos a /.
                                location.href = 'https://personalcollectionmovies-alobaton.c9users.io/account/profile';
                            }
                        }
                    }
                });

            }
        });
    });
    // Validaciones contraseña.
    $('#ChangePasswordForm').bootstrapValidator({
            message: 'Este valor no es valido',
            feedbackIcons: {
                valid: 'glyphicon glyphicon-ok',
                invalid: 'glyphicon glyphicon-remove',
                validating: 'glyphicon glyphicon-refresh'
            },
            fields: {
                Password:{
                    validators: {
                        notEmpty: {
                            message: 'Contraseña actual requerida'
                        },
                        stringLength: {
                            min: 8,
                            message: 'El password debe contener al menos 8 caracteres'
                        }
                    }
                },
                NewPassword:{
                    validators: {
                        notEmpty: {
                                message: 'Nueva contraseña requerida'
                            },
                        stringLength: {
                            min: 8,
                            message: 'El password debe contener al menos 8 caracteres'
                        }
                    }
                },
                ConfirmPassword:{
                    validators: {
                        notEmpty: {
                                message: 'Confirmacion de contraseña requerida'
                            },
                        identical: {
                            field: 'NewPassword',
                            message: 'La contraseña y su confirmación no son iguales'
                        }
                    }
                }
            }
        });
    
});
