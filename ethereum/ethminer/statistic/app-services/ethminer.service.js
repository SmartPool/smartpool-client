(function () {
    'use strict';

    angular
        .module('app')
        .factory('EthminerService', EthminerService);

    EthminerService.$inject = ['$http'];
    function EthminerService($http) {
        var service = {};

        service.GetAdvanceInfo = GetAdvanceInfo;
        service.GetConfigInfo = GetConfigInfo;

        return service;

        function GetAdvanceInfo() {
            return $http.post('/api/ethminer/advanceinfo').then(handleSuccess, handleError('Error get information'));
        }
        function GetConfigInfo() {
            return $http.get('/status').then(handleSuccess, handleError('Error get information'));
        }


        // private functions

        function handleSuccess(res) {
            return res.data;
        }

        function handleError(error) {
            return function () {
                return { success: false, message: error };
            };
        }
    }
})();
