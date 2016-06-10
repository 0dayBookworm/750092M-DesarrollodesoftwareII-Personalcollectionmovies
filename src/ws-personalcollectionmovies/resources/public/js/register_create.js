$(document).ready(function() {
    $('#Useraccount')
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
                    },
                    regexp: {
                        regexp: /^[A-Za-z][A-Za-z0-9]*$/,
                        message: 'El nombre de usuario sólo puede contener caracteres alfabéticos, numericos y puntos'
                    }
                },
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
                Password: {
                    validators: {
                        notEmpty: {
                            message: 'La contraseña es requerida'
                        },
                        stringLength: {
                            min: 1,
                            message: 'El password debe contener al menos 8 caracteres'
                        }
                    }
                },
                ConfirmPassword: {
                    validators: {
                        notEmpty: {
                            message: 'La confirmacion de contraseña es requerida'
                        },
                        identical: {
                            field: 'Password',
                            message: 'La contraseña y su confirmación no son iguales'
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
                },
                TermsAndConditions: {
                    validators: {
                        notEmpty: {
                            message: 'Debes aceptar los terminos y condiciones'
                        }
                    }
                }
            }
        })
    $("#CreateUseraccount").click(function() {
        // Cancels the form submission
        event.preventDefault();
        // Estructura de la peticion.
        var RegistrationRequest = $('#Useraccount').serialize();
        // Send data to back-end
        $.ajax({
            type: 'post',
            url: 'https://personalcollectionmovies-alobaton.c9users.io/register/create',
            data: RegistrationRequest,
            dataType: 'json'
        }).success(function(response) {
            if (response.Status === '999') {
                // Mostramos el o los mensajes de error.
                bootbox.dialog({
                    message: response.Message,
                    title: "Registro",
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
                    title: "Registro",
                    buttons: {
                        close: {
                            label: "Cerrar",
                            className: "btn-success",
                            callback: function() {
                                // Si el proceso de registro se llevo a cabo correctamente redireccionamos a /.
                                location.href = 'https://personalcollectionmovies-alobaton.c9users.io/';
                            }
                        }
                    }
                });

            }
        });
    });
});
