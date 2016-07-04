function checkCookies(){
    // Verificamos que las cookies esten habilitadas.
    if (navigator.cookieEnabled != true) {
        alert("Debes habilitar las cookies.")
        // Se debe redireccionar a una pagina de error 404...
        
    }
}

// $(window).load(function() {
// 	'use strict';
// 	$('.loading-icon').delay(100).fadeOut();
// 	$('#Preloader').delay(200).fadeOut('slow');
// });