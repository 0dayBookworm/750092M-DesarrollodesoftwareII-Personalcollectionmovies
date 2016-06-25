$(document).ready(function() {
    $('#ViewList').click(function(event){event.preventDefault();$('#Movies .item').removeClass('grid-group-item');$('#Movies .item').addClass('list-group-item');});
    $('#ViewGrid').click(function(event){event.preventDefault();$('#Movies .item').removeClass('list-group-item');$('#Movies .item').addClass('grid-group-item');});
});