(function() {
  var app;

  app = angular.module('app', ['ui.router']);

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
      });
    }
  ]);

}).call(this);

(function() {
  angular.module('app').controller('aboutCtrl', [
    '$scope', function($scope) {
      $scope.title = "About";
      return $scope.items = ['other', 'stuff', 'in here'];
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
