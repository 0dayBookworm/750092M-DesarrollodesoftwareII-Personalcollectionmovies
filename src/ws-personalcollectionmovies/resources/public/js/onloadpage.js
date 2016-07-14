function checkCookies(){
    // Verificamos que las cookies esten habilitadas.
    if (navigator.cookieEnabled != true) {
        alert("Debes habilitar las cookies.")
        // Se debe redireccionar a una pagina de error 404...
        
    }
}
$(window).load(function() {
	'use strict';
	$('.loading-icon').delay(200).fadeOut();
	$('#MainPreloader').delay(400).fadeOut('slow');
});

$(window).load(function() {
	'use strict';
	$('.loading-icon').delay(200).fadeOut();
	$('#MapPreloader').delay(200).fadeOut('slow');
	$('#MapContainer').delay(800).fadeIn('slow');
});

