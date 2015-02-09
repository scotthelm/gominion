(function() {
  var app;

  app = angular.module('app', ['ui.router', 'app.controllers', 'app.services']);

  app.config([
    '$urlRouterProvider', '$stateProvider', function($urlRouterProvider, $stateProvider) {
      $urlRouterProvider.otherwise('/');
      return $stateProvider.state('home', {
        url: '/',
        templateUrl: 'templates/home.html',
        controller: 'homeCtrl'
      }).state('about', {
        url: '/about',
        templateUrl: 'templates/about.html',
        controller: 'aboutCtrl'
      }).state('campaigns', {
        url: '/campaigns',
        templateUrl: 'templates/campaigns_new.html',
        controller: 'campaignsListCtrl'
      });
    }
  ]);

}).call(this);

(function() {
  angular.module('app').controller('aboutCtrl', [
    '$scope', function($scope) {
      $scope.title = "About";
      return $scope.items = ['some', 'other', 'stuff'];
    }
  ]);

}).call(this);

(function() {
  var services;

  services = angular.module('app.services', ['ngResource']);

  services.factory('CampaignsFactory', [
    '$resource', function($resource) {
      return $resource('/api/campaigns', {}, {
        query: {
          method: 'GET',
          isArray: true
        },
        create: {
          method: 'POST'
        }
      });
    }
  ]);

  services.factory('CampaignFactory', [
    '$resource', function($resource) {
      return $resource('/api/campaign/:id', {}, {
        show: {
          method: 'GET'
        },
        update: {
          method: 'PUT',
          params: {
            id: '@id'
          }
        },
        "delete": {
          method: 'DELETE',
          params: {
            id: '@id'
          }
        }
      });
    }
  ]);

}).call(this);

(function() {
  angular.module('app.controllers', []).controller('campaignsListCtrl', [
    '$scope', 'CampaignsFactory', 'CampaignFactory', '$location', function($scope, CampaignsFactory, CampaignFactory, $location) {
      $scope.editCampaign = function(campaignId) {
        return $location.path("/campaigns/" + campaignId);
      };
      $scope.deleteCampaign = function(campaignId) {
        return CampaignFactory["delete"]({
          id: campaignId
        });
      };
      $scope.newCampaign = function() {
        return $location.path("/campaigns/new");
      };
      return $scope.campaigns = CampaignsFactory.query();
    }
  ]);

}).call(this);

(function() {
  angular.module('app').controller('homeCtrl', [
    '$scope', function($scope) {
      $scope.title = "Home";
      return $scope.items = ['this', 'that', 'the other thing'];
    }
  ]);

}).call(this);
