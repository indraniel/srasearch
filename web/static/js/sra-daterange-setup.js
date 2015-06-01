/**
* @version: 0.0.1
* @author: Indraniel Das https://github.com/indraniel
* @copyright: Copyright (c) 2015 Indraniel Das
* @license: Licensed under the BSD license. See http://choosealicense.com/licenses/bsd-2-clause/
*/

$(function() {
 
    var rangeStart = $('#range-start-time').attr("datetime");
    var rangeEnd   = $('#range-end-time').attr("datetime");

    if (rangeStart && rangeEnd) {
        $('#reportrange span').html(moment(rangeStart, "YYYY-MM-DD").format('MMMM D, YYYY') + ' - ' + moment(rangeEnd, "YYYY-MM-DD").format('MMMM D, YYYY'));
        $('#searchStart').val(moment(rangeStart, "YYYY-MM-DD").format("YYYY-MM-DD"));
        $('#searchEnd').val(moment(rangeEnd, "YYYY-MM-DD").format("YYYY-MM-DD"));
    }
    else {
        $('#reportrange span').html(moment().subtract(7, 'days').format('MMMM D, YYYY') + ' - ' + moment().format('MMMM D, YYYY'));
        $('#searchStart').val(moment().subtract(7, 'days').format("YYYY-MM-DD"));
        $('#searchEnd').val(moment().format("YYYY-MM-DD"));
    }

    $('#reportrange').daterangepicker({
        format: 'YYYY-MM-DD',
        startDate: moment().subtract(29, 'days'),
        endDate: moment(),
        minDate: '2007-01-01',
        maxDate: '2020-12-31',
//        dateLimit: { days: 60 },
        showDropdowns: true,
        showWeekNumbers: true,
        timePicker: false,
        timePickerIncrement: 1,
        timePicker12Hour: true,
        ranges: {
           'Today': [moment(), moment()],
           'Yesterday': [moment().subtract(1, 'days'), moment().subtract(1, 'days')],
           'Last 7 Days': [moment().subtract(6, 'days'), moment()],
           'Last 30 Days': [moment().subtract(29, 'days'), moment()],
           'This Month': [moment().startOf('month'), moment().endOf('month')],
           'Last Month': [moment().subtract(1, 'month').startOf('month'), moment().subtract(1, 'month').endOf('month')],
           'Last 3 Months': [moment().subtract(3, 'month').startOf('month'), moment().subtract(1, 'month').endOf('month')],
           'Last 6 Months': [moment().subtract(6, 'month').startOf('month'), moment().subtract(1, 'month').endOf('month')],
           'Last Year': [moment().subtract(12, 'month').startOf('month'), moment().subtract(1, 'month').endOf('month')],
        },
        opens: 'left',
        drops: 'down',
        buttonClasses: ['btn', 'btn-sm'],
        applyClass: 'btn-primary',
        cancelClass: 'btn-default',
        separator: ' to ',
        locale: {
            applyLabel: 'Confirm',
            cancelLabel: 'Cancel',
            fromLabel: 'From',
            toLabel: 'To',
            customRangeLabel: 'Custom',
            daysOfWeek: ['Su', 'Mo', 'Tu', 'We', 'Th', 'Fr','Sa'],
            monthNames: ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'],
            firstDay: 1
        }
    }, function(start, end, label) {
        console.log(start.toISOString(), end.toISOString(), label);
        $('#reportrange span').html(start.format('MMMM D, YYYY') + ' - ' + end.format('MMMM D, YYYY'));
        $('#searchStart').val(start.format('YYYY-MM-DD'));
        $('#searchEnd').val(end.format('YYYY-MM-DD'));
    });
 
});
