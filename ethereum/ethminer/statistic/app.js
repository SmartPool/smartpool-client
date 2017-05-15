(function() {
    'use strict';

    angular
        .module('app', ['ngRoute', 'ui.bootstrap'])
        .constant('appConstants', {
            CONST_MAX_FILE_SIZE: 2000000,
            CONST_MAX_PROOF_SIZE: 2000000
        })
        .config(config)
        //.run(run)

    config.$inject = ['$routeProvider', '$locationProvider'];

    function config($routeProvider, $locationProvider) {
        $routeProvider
            .when('/', {
                controller: 'DashboardController',
                templateUrl: 'dashboard/dashboard.view.html',
                controllerAs: 'vm'
            })
            .when('/rig/:rigId', {
                controller: 'RigController',
                templateUrl: 'rig/rig.view.html',
                controllerAs: 'vm'
            })
        // .when('/submitMessage', {
        //     controller: 'SubmitMessageController',
        //     templateUrl: 'submitMessage/submitMessage.view.html',
        //     controllerAs: 'vm'
        // })

        // .when('/checkMessage', {
        //     controller: 'CheckMessageController',
        //     templateUrl: 'checkMessage/checkMessage.view.html',
        //     controllerAs: 'vm'
        // })

        .otherwise({ redirectTo: '/' });
    }


    //run.$inject = ['$rootScope', '$location', '$http'];

    // function run($rootScope, $location, $http) {
    //     if (window.WebSocket === undefined) {
    //         console.log("windows is not support websocket");
    //     } else {
    //         var socket = new WebSocket("ws://" + $location.$$host + ":" + $location.$$port + "/ws");

    //         socket.onopen = function() {
    //             console.log("Socket is open");
    //         };
    //         socket.onmessage = function(message) {
    //             var reponse = JSON.parse(message.data);
    //             $rootScope.$apply(function() {
    //                 $rootScope.currency = reponse.currency;
    //                 if( $rootScope.auth ){
    //                     $rootScope.globals.currentUser.balance = reponse.balance;
    //                     $rootScope.globals.currentUser.pending_tx = reponse.pending_tx;
    //                 }
    //             })

    //         }
    //         socket.onclose = function() {
    //             console.log("Socket is close");
    //         }
    //     }
    // }

})();
