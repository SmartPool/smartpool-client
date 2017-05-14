(function() {
    'use strict';

    angular
        .module('app')
        .controller('RigController', RigController);

    RigController.$inject = ['$location', '$rootScope', '$http', '$scope', 'EthminerService','$routeParams'];

    function RigController($location, $rootScope, $http, $scope, EthminerService,$routeParams) {
        var vm = this;
        vm.rigId = $routeParams.rigId;
        vm.roundHashRate = roundHashRate;
        vm.roundShares = roundShares;
        vm.applyShortPeriod = applyShortPeriod;
        vm.applyLongPeriod = applyLongPeriod;
        vm.applyWorker = applyWorker;
        vm.showAdvanceInfo = showAdvanceInfo;
        vm.applyOverall = applyOverall;
        vm.advance = {
            "load": false,
            "flag": false,
        }
        vm.config = {};

        vm.farm = {
            "short_duration": {
                "duration_in_hour": 1,
                "hash_rate": {
                    "effective_hashrate_avarage": 0,
                    "reported_hashrate_avarage": 0,
                    "chart": [

                    ],
                },
                "shares": {
                    "mined_share_avarage": 0,
                    "valid_share_avarage": 0,
                    "rejected_share_avarage": 0,
                    "chart": [

                    ],
                }
            },
            "long_duration": {
                "duration_in_hour": 1,
                "hash_rate": {
                    "effective_hashrate_avarage": 0,
                    "reported_hashrate_avarage": 0,
                    "chart": [

                    ],
                },
                "shares": {
                    "mined_share_avarage": 0,
                    "valid_share_avarage": 0,
                    "rejected_share_avarage": 0,
                    "chart": [

                    ],
                }
            },
            "overall": {

            }
        };
        vm.shortHashrateChart = c3.generate({
            bindto: '#shortHashChart',
            data: {
                x: 'x',
                columns: vm.farm.short_duration.hash_rate.chart
            },
            axis: {
                x: {
                    type: 'timeseries',
                    tick: {
                        format: '%Y-%m-%d %H:%M'
                    },
                    show: false
                },
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
        vm.longHashrateChart = c3.generate({
            bindto: '#longHashChart',
            data: {
                x: 'x',
                columns: vm.farm.long_duration.hash_rate.chart
            },
            axis: {
                x: {
                    type: 'timeseries',
                    tick: {
                        format: '%Y-%m-%d %H:%M'
                    },
                    show: false
                },
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
        vm.shortSharesChart = c3.generate({
            bindto: '#shortSharesChart',
            data: {
                x: 'x',
                columns: vm.farm.short_duration.shares.chart
            },
            axis: {
                x: {
                    type: 'timeseries',
                    tick: {
                        format: '%Y-%m-%d %H:%M'
                    },
                    show: false
                },
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
        vm.longSharesChart = c3.generate({
            bindto: '#longSharesChart',
            data: {
                x: 'x',
                columns: vm.farm.long_duration.shares.chart
            },
            axis: {
                x: {
                    type: 'timeseries',
                    tick: {
                        format: '%Y-%m-%d %H:%M'
                    },
                    show: false
                },
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

        if (window.WebSocket === undefined) {
            console.log("windows is not support websocket");
        } else {
            var socket = new WebSocket("ws://" + $location.$$host + ":" + $location.$$port + "/ws/rig");

            socket.onopen = function() {
                console.log("Socket is open");
            };
            socket.onmessage = function(message) {
                var response = JSON.parse(JSON.parse(message.data));
                //reperate data
                //vm.$apply(function() {
                $scope.$apply(function() {
                    vm.applyShortPeriod(response);
                    vm.applyLongPeriod(response);
                    vm.applyOverall(response);
                    vm.applyWorker(response);
                })

                //})
                //console.log(response);
            }
            socket.onclose = function() {
                console.log("Socket is close");
            }
        }

        function applyShortPeriod(response) {
            vm.farm.short_duration.duration_in_hour = response.short_window_duration / 3600;
            var pointTotal = response.short_window_duration / response.period_duration;
            var totalEffectiveHashRate = 0;
            var totalReportedHashRate = 0;
            var reportedChart = ['Reported Hashrate'];
            var effectiveChart = ['Effective Hashrate'];

            var totalMinedShare = 0;
            var totalValidShare = 0;
            var totalRejectedShare = 0;
            var minedChart = ['Mined Shares'];
            var validChart = ['Valid Shares'];
            var rejectedChart = ['Rejected Shares'];

            var xChart = ['x'];
            $.each(response.short_window_sample, function(key, val) {
                xChart.push(key * response.period_duration * 1000);

                reportedChart.push(vm.roundHashRate(val.reported_hashrate / 1000000));
                effectiveChart.push(vm.roundHashRate(val.effective_hashrate / 1000000));
                totalEffectiveHashRate += val.effective_hashrate;
                totalReportedHashRate += val.reported_hashrate;

                minedChart.push(val.mined_share);
                validChart.push(val.valid_share);
                rejectedChart.push(val.rejected_share);
                totalMinedShare += val.mined_share;
                totalValidShare += val.valid_share;
                totalRejectedShare += val.rejected_share;
            })
            vm.farm.short_duration.hash_rate.chart = [xChart, reportedChart, effectiveChart];
            vm.farm.short_duration.hash_rate.effective_hashrate_avarage = vm.roundHashRate(totalEffectiveHashRate / pointTotal / 1000000);
            vm.farm.short_duration.hash_rate.reported_hashrate_avarage = vm.roundHashRate(totalReportedHashRate / pointTotal / 1000000);

            vm.farm.short_duration.shares.chart = [xChart, minedChart, validChart, rejectedChart];
            vm.farm.short_duration.shares.mined_share_avarage = vm.roundShares(totalMinedShare / pointTotal);
            vm.farm.short_duration.shares.valid_share_avarage = vm.roundShares(totalValidShare / pointTotal);
            vm.farm.short_duration.shares.rejected_share_avarage = vm.roundShares(totalRejectedShare / pointTotal);

            //calculate share percent
            vm.farm.short_duration.shares.valid_share_percent = vm.roundShares(vm.farm.short_duration.shares.valid_share_avarage / vm.farm.short_duration.shares.mined_share_avarage * 100);
            vm.farm.short_duration.shares.rejected_share_percent = vm.roundShares(vm.farm.short_duration.shares.rejected_share_avarage / vm.farm.short_duration.shares.mined_share_avarage * 100);

            //load chart
            vm.shortHashrateChart.load({
                columns: vm.farm.short_duration.hash_rate.chart
            })
            vm.shortSharesChart.load({
                columns: vm.farm.short_duration.shares.chart
            })
        }

        function applyLongPeriod(response) {
            vm.farm.long_duration.duration_in_hour = response.long_window_duration / 3600;
            var pointTotal = response.long_window_duration / response.period_duration;
            var totalEffectiveHashRate = 0;
            var totalReportedHashRate = 0;
            var reportedChart = ['Reported Hashrate'];
            var effectiveChart = ['Effective Hashrate'];

            var totalMinedShare = 0;
            var totalValidShare = 0;
            var totalRejectedShare = 0;
            var minedChart = ['Mined Shares'];
            var validChart = ['Valid Shares'];
            var rejectedChart = ['Rejected Shares'];

            var xChart = ['x'];
            $.each(response.long_window_sample, function(key, val) {
                xChart.push(key * response.period_duration * 1000);
                reportedChart.push(vm.roundHashRate(val.reported_hashrate / 1000000));
                effectiveChart.push(vm.roundHashRate(val.effective_hashrate / 1000000));
                totalEffectiveHashRate += val.effective_hashrate;
                totalReportedHashRate += val.reported_hashrate;

                minedChart.push(val.mined_share);
                validChart.push(val.valid_share);
                rejectedChart.push(val.rejected_share);
                totalMinedShare += val.mined_share;
                totalValidShare += val.valid_share;
                totalRejectedShare += val.rejected_share;
            })
            vm.farm.long_duration.hash_rate.chart = [xChart, reportedChart, effectiveChart];
            vm.farm.long_duration.hash_rate.effective_hashrate_avarage = vm.roundHashRate(totalEffectiveHashRate / pointTotal / 1000000);
            vm.farm.long_duration.hash_rate.reported_hashrate_avarage = vm.roundHashRate(totalReportedHashRate / pointTotal / 1000000);

            vm.farm.long_duration.shares.chart = [xChart, minedChart, validChart, rejectedChart];
            vm.farm.long_duration.shares.mined_share_avarage = vm.roundShares(totalMinedShare / pointTotal);
            vm.farm.long_duration.shares.valid_share_avarage = vm.roundShares(totalValidShare / pointTotal);
            vm.farm.long_duration.shares.rejected_share_avarage = vm.roundShares(totalRejectedShare / pointTotal);

            //calculate percent
            vm.farm.long_duration.shares.valid_share_percent = vm.roundShares(vm.farm.long_duration.shares.valid_share_avarage / vm.farm.long_duration.shares.mined_share_avarage * 100);
            vm.farm.long_duration.shares.rejected_share_percent = vm.roundShares(vm.farm.long_duration.shares.rejected_share_avarage / vm.farm.long_duration.shares.mined_share_avarage * 100);

            //load chartl
            vm.longHashrateChart.load({
                columns: vm.farm.long_duration.hash_rate.chart
            })
            vm.longSharesChart.load({
                columns: vm.farm.long_duration.shares.chart
            })
        }
        function applyOverall(response){
            vm.farm.overall.effective_hashrate = vm.roundHashRate(response.overall.effective_hashrate/1000000);
            vm.farm.overall.reported_hashrate = vm.roundHashRate(response.overall.reported_hashrate/1000000);
            vm.farm.overall.mined_share = response.overall.mined_share;
            vm.farm.overall.valid_share = response.overall.valid_share;
            vm.farm.overall.rejected_share = response.overall.rejected_share;
            vm.farm.overall.verified_share = response.overall.verified_share;
            vm.farm.overall.pending_share = response.overall.pending_share;
            if (vm.farm.overall.mined_share > 0 ){
                vm.farm.overall.valid_share_percent = vm.roundShares(vm.farm.overall.valid_share / vm.farm.overall.mined_share * 100);
                vm.farm.overall.rejected_share_percent = vm.roundShares(vm.farm.overall.rejected_share / vm.farm.overall.mined_share * 100);
            }
        }
        function applyWorker(response) {
            $.each(response.short_window_sample, function(key, val) {

            })
        }

        function showAdvanceInfo() {
            if (vm.advance.load) {
                vm.advance.flag = !vm.advance.flag;
                return;
            }
            EthminerService.GetAdvanceInfo()
                .then(function(response) {
                    vm.advance.x = response.x;
                    vm.advance.y = response.y;
                    vm.advance.load = true;
                    vm.advance.flag = true;
                });
        }
        (function initController() {
            EthminerService.GetConfigInfo()
                .then(function(response) {
                    vm.config = response;
                });
        })();

        function roundHashRate(rate) {
            return Math.round(rate * 100) / 100;
        }

        function roundShares(shares) {
            return Math.round(shares * 100) / 100;
        }
        //simulate like websocker
        // (function initController() {

        // })();



        //farm info is update after a time



        // vm.farm = {
        //     "reported_hashrate": 76678,
        //     "effective_hashrate": 138153,
        //     "mined_share": 47,
        //     "valid_share": 47,
        //     "rejected_share": 0,
        //     "active_workers": 168,
        //     "hash_rate": [
        //         ['x', 30, 50, 100, 230, 300, 310],
        //         ['Reported Hashrate', 30, 200, 100, 400, 150, 250],
        //         ['Effective Hashrate', 50, 20, 10, 40, 15, 25]
        //     ],
        //     "shares": [
        //         ['x', 30, 50, 100, 230, 300, 310],
        //         ['Mined shares', 30, 200, 100, 400, 150, 250],
        //         ['Valid shares', 50, 20, 10, 40, 15, 25],
        //         ['Reject shares', 50, 20, 10, 40, 15, 25]
        //     ],
        //     "activeWorkers": [
        //         ['x', 30, 50, 100, 230, 300, 310],
        //         ['Active workers', 30, 200, 100, 400, 150, 250],
        //     ],
        //     // "worker_list": [{
        //     //     "name": "cuong",
        //     //     "reported_hashrate": 123,
        //     //     "current_hashrate": 21,
        //     //     "avarage_hashrate": 234,
        //     //     "mined_shares": 43,
        //     //     "valid_shares": 23,
        //     //     "rejected_shares": 34
        //     // }, {
        //     //     "name": "vu",
        //     //     "reported_hashrate": 124,
        //     //     "current_hashrate": 21,
        //     //     "avarage_hashrate": 234,
        //     //     "mined_shares": 43,
        //     //     "valid_shares": 23,
        //     //     "rejected_shares": 34
        //     // }]
        //     "worker_list": [
        //         ["cuong", 123, 21, 234, 43, 23, 34],
        //         ["vu", 123, 21, 234, 43, 23, 34],
        //         ["vu", 123, 21, 234, 43, 23, 34],
        //         ["vu", 123, 21, 234, 43, 23, 34],
        //         ["vu", 123, 21, 234, 43, 23, 34],
        //         ["vu", 123, 21, 234, 43, 23, 34],
        //         ["vu", 123, 21, 234, 43, 23, 34],
        //         ["vu", 123, 21, 234, 43, 23, 34],
        //         ["vu", 123, 21, 234, 43, 23, 34]
        //     ]
        // };
        // vm.chart = null;
        // vm.tableWorker = null;
        // vm.searchWorker = searchWorker;
        // vm.workerRate = [
        //     ['x', 30, 50, 100, 230, 300, 310],
        //     ['Reported Hashrate', 30, 200, 100, 400, 150, 250],
        //     ['Effective Hashrate', 50, 20, 10, 40, 15, 25]
        // ];
        // (function initController() {
        //     vm.hashrateChart = c3.generate({
        //         bindto: '#hashChart',
        //         data: {
        //             x: 'x',
        //             columns: vm.farm.hash_rate
        //         },
        //         axis: {
        //             y: {
        //                 label: {
        //                     text: 'Hashrate [MH/s]',
        //                     position: 'outer-middle'
        //                 }
        //             }
        //         },
        //         grid: {
        //             y: {
        //                 show: true
        //             }
        //         }
        //     });

        //     setTimeout(function() {
        //         vm.farm.hash_rate = [
        //             ['x', 20, 60, 70, 100, 200, 300],
        //             ['Reported Hashrate', 30, 200, 100, 400, 150, 250],
        //             ['Effective Hashrate', 50, 20, 10, 40, 15, 25]
        //         ];
        //         vm.hashrateChart.load({
        //             columns: vm.farm.hash_rate
        //         })
        //     }, 2000);

        //     vm.sharesChart = c3.generate({
        //         bindto: '#sharesChart',
        //         data: {
        //             x: 'x',
        //             columns: vm.farm.shares
        //         },
        //         axis: {
        //             y: {
        //                 label: {
        //                     text: 'Shares',
        //                     position: 'outer-middle'
        //                 }
        //             }
        //         },
        //         grid: {
        //             y: {
        //                 show: true
        //             }
        //         }
        //     });
        //     vm.activeWorkers = c3.generate({
        //         bindto: '#activeWorkers',
        //         data: {
        //             x: 'x',
        //             columns: vm.farm.activeWorkers
        //         },
        //         axis: {
        //             y: {
        //                 label: {
        //                     text: 'Active workers',
        //                     position: 'outer-middle'
        //                 }
        //             }
        //         },
        //         grid: {
        //             y: {
        //                 show: true
        //             }
        //         }
        //     });
        //     //data table 
        //     vm.tableWorker = $("#worker_table").DataTable({
        //         paging: false,
        //         info: false,
        //         stateSave: true,
        //         data: vm.farm.worker_list,
        //         columns: [
        //             { title: "Worker" },
        //             { title: "Reported Hashrate" },
        //             { title: "Current Hashrate" },
        //             { title: "Average Hashrate" },
        //             { title: "Mined Shares (1h)" },
        //             { title: "Valid Shares (1h)" },
        //             { title: "Rejected Shares (1h)" }
        //         ],
        //         columnDefs: [{
        //                 // The `data` parameter refers to the data for the cell (defined by the
        //                 // `data` option, which defaults to the column being worked with, in
        //                 // this case `data: 0`.
        //                 "render": function(data, type, row) {
        //                     return '<button rel="workerChart" class="btn btn-default btn-xs">' + data + '</button>';
        //                 },
        //                 "targets": 0
        //             },
        //             // { "visible": false,  "targets": [ 3 ] }
        //         ]
        //     });

        //     $("[rel=workerChart]").popover({
        //             trigger: 'focus',
        //             placement: 'right',
        //             html: 'true',
        //             content: '<div id="workerChart" style="height: 100px;width: 800px;">Loading...</div>',
        //             template: '<div class="popover"><div class="arrow"></div>' +
        //                 '<h3 class="popover-title"></h3><div class="popover-content">' +
        //                 '</div><div class="popover-footer"></div></div>'
        //         })
        //         .on('shown.bs.popover', function() {
        //             //hide any visible comment-popover
        //             $("[rel=workerChart]").not(this).popover('hide');
        //             var $this = $(this);

        //             var workerChart = c3.generate({
        //                 bindto: '#workerChart',
        //                 data: {
        //                     x: 'x',
        //                     columns: [
        //                         ['x', 30, 50, 100, 230, 300, 310],
        //                         ['Reported Hashrate', 30, 200, 100, 400, 150, 250],
        //                         ['Effective Hashrate', 50, 20, 10, 40, 15, 25]
        //                     ]
        //                 },
        //                 axis: {
        //                     x: {
        //                         show: false
        //                     },
        //                     y: {
        //                         show: false
        //                     }
        //                 },
        //             });

        //             // $.getJSON(miner + "/" + $this.text() + "/hashrate", function(data) {
        //             //     var workerChart = c3.generate({
        //             //         bindto: '#workerChart',
        //             //         data: {
        //             //             json: data,
        //             //             keys: {
        //             //                 x: 'time',
        //             //                 value: ['HR_W_SHORT', 'HR_W_LONG', 'RHR_W']
        //             //             },
        //             //             xFormat: '%Y-%m-%d %H:%M',
        //             //             names: {
        //             //                 HR_W_SHORT: 'Current Effective Hashrate',
        //             //                 HR_W_LONG: 'Average Effective Hashrate',
        //             //                 RHR_W: 'Reported Hashrate'
        //             //             }
        //             //         },
        //             //         axis: {
        //             //             x: {
        //             //                 type: 'timeseries',
        //             //                 tick: {
        //             //                     format: '%Y-%m-%d %H:%M'
        //             //                 },
        //             //                 show: false
        //             //             },
        //             //             y: {
        //             //                 show: false
        //             //             }
        //             //         },
        //             //         grid: {
        //             //             y: {
        //             //                 show: false
        //             //             }
        //             //         },
        //             //         point: {
        //             //             show: false
        //             //         }
        //             //     });
        //             // });
        //         });



        //     // $('#worker_table tr td:first-child button').on('click', function() {
        //     //     var tr = $(this).closest("tr");
        //     //     var data = vm.tableWorker.row(tr[0]).data();
        //     //     vm.workerRateChart = c3.generate({
        //     //         bindto: '#temp_workerrate',
        //     //         data: {
        //     //             x: 'x',
        //     //             columns: vm.workerRate
        //     //         },
        //     //         axis: {
        //     //             y: {
        //     //                 show: false
        //     //             },
        //     //             x: {
        //     //                 show: false
        //     //             }
        //     //         }
        //     //     });
        //     // });
        // })();

        // function searchWorker(name) {
        //     console.log(name);
        // }
    }

})();
