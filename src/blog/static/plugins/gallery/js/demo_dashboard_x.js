$(function(){        
    /* reportrange */
    if($("#reportrange").length > 0){   
        $("#reportrange").daterangepicker({                    
            ranges: {
               'Today': [moment(), moment()],
               'Yesterday': [moment().subtract(1, 'days'), moment().subtract(1, 'days')],
               'Last 7 Days': [moment().subtract(6, 'days'), moment()],
               'Last 30 Days': [moment().subtract(29, 'days'), moment()],
               'This Month': [moment().startOf('month'), moment().endOf('month')],
               'Last Month': [moment().subtract(1, 'month').startOf('month'), moment().subtract(1, 'month').endOf('month')]
            },
            opens: 'left',
            buttonClasses: ['btn btn-default'],
            applyClass: 'btn-small btn-primary',
            cancelClass: 'btn-small',
            format: 'MM.DD.YYYY',
            separator: ' to ',
            startDate: moment().subtract('days', 29),
            endDate: moment()            
          },function(start, end) {
              $('#reportrange span').html(start.format('MMMM D, YYYY') + ' - ' + end.format('MMMM D, YYYY'));
        });
        
        $("#reportrange span").html(moment().subtract('days', 29).format('MMMM D, YYYY') + ' - ' + moment().format('MMMM D, YYYY'));
    }
    /* end reportrange */   
    
    /* Line dashboard chart */
    Morris.Line({
      element: 'x-dashboard-line',
      data: [
        { y: '2014-10-10', a: 2,b: 4},
        { y: '2014-10-11', a: 4,b: 6},
        { y: '2014-10-12', a: 7,b: 10},
        { y: '2014-10-13', a: 5,b: 7},
        { y: '2014-10-14', a: 6,b: 9},
        { y: '2014-10-15', a: 9,b: 12},
        { y: '2014-10-16', a: 18,b: 20},
        { y: '2014-10-17', a: 16,b: 18},
        { y: '2014-10-18', a: 15,b: 17},
        { y: '2014-10-19', a: 13,b: 15},
        { y: '2014-10-20', a: 18,b: 20},
        { y: '2014-10-21', a: 14,b: 18},
      ],
      xkey: 'y',
      ykeys: ['a','b'],
      labels: ['Sales','Event'],
      resize: true,      
      pointSize: '8px',
      //pointStrokeColors: '#DEE4EC',
      xLabels: 'day',
      gridTextSize: '11px',
      lineColors: ['#95B75D','#E34724'],
      gridLineColor: '#95ABBB',
      gridTextColor: '#95ABBB',
      gridTextWeight: 'bold'
    });   
    /* EMD Line dashboard chart */

});

