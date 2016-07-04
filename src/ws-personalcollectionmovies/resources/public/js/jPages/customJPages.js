$('.pagination').jPages({
    containerID: 'Movies',
    first: false,
    previous: 'Anterior',
    next: 'Siguiente',
    last: false,
    perPage: 10,
    startRange: 1,
    midRange: 5,
    endRange: 1,
    delay: 0,
    minHeight: false,
    callback: function(pages, items) {
        bootstrapPagination($('.pagination'));
    }
});

$('.pagination').jPages({
    containerID: 'Collections',
    first: false,
    previous: 'Anterior',
    next: 'Siguiente',
    last: false,
    perPage: 10,
    startRange: 1,
    midRange: 5,
    endRange: 1,
    delay: 0,
    minHeight: false,
    callback: function(pages, items) {
        bootstrapPagination($('.pagination'));
    }
});

function bootstrapPagination(element) {
    element.find('a,span').each(function() {
        if ($(this).parent('li').length) {
            $(this).parent('li').removeAttr('class');
        }
        else {
            $(this).wrap('<li></li>');
        }

        if ($(this).is('a')) {
            $(this).attr('href', '#');
        }
        if ($(this).is('span')) {
            $(this).parent('li').addClass('spacer');
        }


        if ($(this).hasClass('jp-current')) {
            $(this).parent('li').addClass('active');
        }
        if ($(this).hasClass('jp-disabled')) {
            $(this).parent('li').addClass('disabled');
        }
    });
}