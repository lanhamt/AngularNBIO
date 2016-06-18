'use strict';

var ANBIOApp = angular.module('ANBIOApp', ['ngRoute', 'ngMaterial', 'ngResource'])

ANBIOApp.config(['$routeProvider',
    function ($routeProvider) {
        $routeProvider.
            when('/', {
                templateUrl: 'components/dataView/dataViewTemplate.html',
                controller: 'DataViewController'
            }).
            otherwise({
                redirectTo: '/'
            });
    }]);

ANBIOApp.controller('MainController', ['$scope', '$rootScope', '$location', '$resource', '$http',
    function ($scope, $rootScope, $location, $resource, $http) {
        
    }]);
