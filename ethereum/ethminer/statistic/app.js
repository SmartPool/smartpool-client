(function() {
    'use strict';

    angular
        .module('app', ['ngRoute', 'ui.bootstrap'])
        .constant('appConstants', {
            CONST_FRESH_FARM_DATA: 10000,
            CONST_FRESH_RIG_DATA: 10000
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
            .when('/rig/:rigId/:rigIp', {
                controller: 'RigController',
                templateUrl: 'rig/rig.view.html',
                controllerAs: 'vm'
            })
            .otherwise({ redirectTo: '/' });
    }
})();
