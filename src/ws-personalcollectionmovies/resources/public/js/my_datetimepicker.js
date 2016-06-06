$( document ).ready(function() {
    $('#Date').datepicker({
                    format: "dd/mm/yyyy",
                    defaultViewDate: { year: 2000, month: 01, day: 01 },
                    language: "es",
                    orientation: "bottom left",
                    clearBtn: true,
                    multidate: false,
                    autoclose: true
                });
});