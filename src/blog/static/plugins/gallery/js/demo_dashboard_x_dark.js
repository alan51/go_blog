$(function(){        
    /* reportrange 
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
     end reportrange */   
    
    /* Line dashboard chart */
    Morris.Line({
      element: 'x-dashboard-line-dark',
      data: [
        { y: '2014-10-10', a: 2},
        { y: '2014-10-11', a: 4},
        { y: '2014-10-12', a: 7},
        { y: '2014-10-13', a: 5},
        { y: '2014-10-14', a: 6},
        { y: '2014-10-15', a: 9},
        { y: '2014-10-16', a: 18},
        { y: '2014-10-17', a: 16},
        { y: '2014-10-18', a: 15},
        { y: '2014-10-19', a: 13},
        { y: '2014-10-20', a: 18},
        { y: '2014-10-21', a: 14},
      ],
      xkey: 'y',
      ykeys: ['a'],
      labels: ['Views'],
      resize: true,
      pointSize: 2,
      lineWidth: '5px',
      hideHover: 'auto',
      pointStrokeColors: ['#95B75D'],
      xLabels: 'day',
      gridTextSize: '11px',      
      lineColors: ['#95B75D'],
      gridLineColor: '#445664',
      gridTextColor: '#94ABBA',
      gridTextWeight: 'bold'
    });   
    /* EMD Line dashboard chart */

});

