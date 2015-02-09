(function() {
  var app;

  app = angular.module('app', ['ui.router']);

  app.config([
    '$urlRouterProvider', '$stateProvider', function($urlRouterProvider, $stateProvider) {
      $urlRouterProvider.otherwise('/');
      return $stateProvider.state('home', {
        url: '/',
        templateUrl: 'templates/home.html'
      }).state('about', {
        url: '/about',
        templateUrl: 'templates/about.html'
      });
    }
  ]);

}).call(this);
