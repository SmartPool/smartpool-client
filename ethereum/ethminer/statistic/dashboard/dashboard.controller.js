(function() {
    'use strict';

    angular
        .module('app')
        .controller('DashboardController', DashboardController);

    DashboardController.$inject = ['$location', '$rootScope'];

    function DashboardController($location, $rootScope) {
        //farm info is update after a time
        var vm = this;
        vm.farm = {
            "reported_hashrate": 76678,
            "effective_hashrate": 138153,
            "mined_share": 47,
            "valid_share": 47,
            "rejected_share": 0,
            "active_workers": 168,
            "hash_rate": [
                ['x', 30, 50, 100, 230, 300, 310],
                ['Reported Hashrate', 30, 200, 100, 400, 150, 250],
                ['Effective Hashrate', 50, 20, 10, 40, 15, 25]
            ],
            "shares": [
                ['x', 30, 50, 100, 230, 300, 310],
                ['Mined shares', 30, 200, 100, 400, 150, 250],
                ['Valid shares', 50, 20, 10, 40, 15, 25],
                ['Reject shares', 50, 20, 10, 40, 15, 25]
            ],
            "activeWorkers": [
                ['x', 30, 50, 100, 230, 300, 310],
                ['Active workers', 30, 200, 100, 400, 150, 250],
            ],
            // "worker_list": [{
            //     "name": "cuong",
            //     "reported_hashrate": 123,
            //     "current_hashrate": 21,
            //     "avarage_hashrate": 234,
            //     "mined_shares": 43,
            //     "valid_shares": 23,
            //     "rejected_shares": 34
            // }, {
            //     "name": "vu",
            //     "reported_hashrate": 124,
            //     "current_hashrate": 21,
            //     "avarage_hashrate": 234,
            //     "mined_shares": 43,
            //     "valid_shares": 23,
            //     "rejected_shares": 34
            // }]
            "worker_list": [
                ["cuong", 123, 21, 234, 43, 23, 34],
                ["vu", 123, 21, 234, 43, 23, 34],
                ["vu", 123, 21, 234, 43, 23, 34],
                ["vu", 123, 21, 234, 43, 23, 34],
                ["vu", 123, 21, 234, 43, 23, 34],
                ["vu", 123, 21, 234, 43, 23, 34],
                ["vu", 123, 21, 234, 43, 23, 34],
                ["vu", 123, 21, 234, 43, 23, 34],
                ["vu", 123, 21, 234, 43, 23, 34]
            ]
        };
        vm.chart = null;
        vm.tableWorker = null;
        vm.searchWorker = searchWorker;
        vm.workerRate = [
            ['x', 30, 50, 100, 230, 300, 310],
            ['Reported Hashrate', 30, 200, 100, 400, 150, 250],
            ['Effective Hashrate', 50, 20, 10, 40, 15, 25]
        ];
        (function initController() {
            vm.hashrateChart = c3.generate({
                bindto: '#hashChart',
                data: {
                    x: 'x',
                    columns: vm.farm.hash_rate
                },
                axis: {
                    y: {
                        label: {
                            text: 'Hashrate [MH/s]',
                            position: 'outer-middle'
                        }
                    }
                },
                grid: {
                    y: {
                        show: true
                    }
                }
            });

            setTimeout(function() {
                vm.farm.hash_rate = [
                    ['x', 20, 60, 70, 100, 200, 300],
                    ['Reported Hashrate', 30, 200, 100, 400, 150, 250],
                    ['Effective Hashrate', 50, 20, 10, 40, 15, 25]
                ];
                vm.hashrateChart.load({
                    columns: vm.farm.hash_rate
                })
            }, 2000);

            vm.sharesChart = c3.generate({
                bindto: '#sharesChart',
                data: {
                    x: 'x',
                    columns: vm.farm.shares
                },
                axis: {
                    y: {
                        label: {
                            text: 'Shares',
                            position: 'outer-middle'
                        }
                    }
                },
                grid: {
                    y: {
                        show: true
                    }
                }
            });
            vm.activeWorkers = c3.generate({
                bindto: '#activeWorkers',
                data: {
                    x: 'x',
                    columns: vm.farm.activeWorkers
                },
                axis: {
                    y: {
                        label: {
                            text: 'Active workers',
                            position: 'outer-middle'
                        }
                    }
                },
                grid: {
                    y: {
                        show: true
                    }
                }
            });
            //data table 
            vm.tableWorker = $("#worker_table").DataTable({
                paging: false,
                info: false,
                stateSave: true,
                data: vm.farm.worker_list,
                columns: [
                    { title: "Worker" },
                    { title: "Reported Hashrate" },
                    { title: "Current Hashrate" },
                    { title: "Average Hashrate" },
                    { title: "Mined Shares (1h)" },
                    { title: "Valid Shares (1h)" },
                    { title: "Rejected Shares (1h)" }
                ],
                columnDefs: [{
                        // The `data` parameter refers to the data for the cell (defined by the
                        // `data` option, which defaults to the column being worked with, in
                        // this case `data: 0`.
                        "render": function(data, type, row) {
                            return '<button rel="workerChart" class="btn btn-default btn-xs">' + data + '</button>';
                        },
                        "targets": 0
                    },
                    // { "visible": false,  "targets": [ 3 ] }
                ]
            });

            $("[rel=workerChart]").popover({
                    trigger: 'focus',
                    placement: 'right',
                    html: 'true',
                    content: '<div id="workerChart" style="height: 100px;width: 800px;">Loading...</div>',
                    template: '<div class="popover"><div class="arrow"></div>' +
                        '<h3 class="popover-title"></h3><div class="popover-content">' +
                        '</div><div class="popover-footer"></div></div>'
                })
                .on('shown.bs.popover', function() {
                    //hide any visible comment-popover
                    $("[rel=workerChart]").not(this).popover('hide');
                    var $this = $(this);

                    var workerChart  = c3.generate({
                        bindto: '#workerChart',
                        data: {
                            x: 'x',
                            columns: [
                                ['x', 30, 50, 100, 230, 300, 310],
                                ['Reported Hashrate', 30, 200, 100, 400, 150, 250],
                                ['Effective Hashrate', 50, 20, 10, 40, 15, 25]
                            ]
                        },
                        axis: {
                                x: {
                                    show: false
                                },
                                y: {
                                    show: false
                                }
                            },
                    });

                    // $.getJSON(miner + "/" + $this.text() + "/hashrate", function(data) {
                    //     var workerChart = c3.generate({
                    //         bindto: '#workerChart',
                    //         data: {
                    //             json: data,
                    //             keys: {
                    //                 x: 'time',
                    //                 value: ['HR_W_SHORT', 'HR_W_LONG', 'RHR_W']
                    //             },
                    //             xFormat: '%Y-%m-%d %H:%M',
                    //             names: {
                    //                 HR_W_SHORT: 'Current Effective Hashrate',
                    //                 HR_W_LONG: 'Average Effective Hashrate',
                    //                 RHR_W: 'Reported Hashrate'
                    //             }
                    //         },
                    //         axis: {
                    //             x: {
                    //                 type: 'timeseries',
                    //                 tick: {
                    //                     format: '%Y-%m-%d %H:%M'
                    //                 },
                    //                 show: false
                    //             },
                    //             y: {
                    //                 show: false
                    //             }
                    //         },
                    //         grid: {
                    //             y: {
                    //                 show: false
                    //             }
                    //         },
                    //         point: {
                    //             show: false
                    //         }
                    //     });
                    // });
                });



            // $('#worker_table tr td:first-child button').on('click', function() {
            //     var tr = $(this).closest("tr");
            //     var data = vm.tableWorker.row(tr[0]).data();
            //     vm.workerRateChart = c3.generate({
            //         bindto: '#temp_workerrate',
            //         data: {
            //             x: 'x',
            //             columns: vm.workerRate
            //         },
            //         axis: {
            //             y: {
            //                 show: false
            //             },
            //             x: {
            //                 show: false
            //             }
            //         }
            //     });
            // });
        })();

        function searchWorker(name) {
            console.log(name);
        }
    }

})();
