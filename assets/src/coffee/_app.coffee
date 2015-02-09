app = angular.module 'app', [ 'ui.router', 'app.controllers', 'app.services' ]
app.config ['$urlRouterProvider', '$stateProvider', ($urlRouterProvider, $stateProvider) ->
  $urlRouterProvider.otherwise('/')
  $stateProvider.state(
    'home',
    url: '/'
    templateUrl: 'templates/home.html'
    controller: 'homeCtrl'
  ).state(
    'about',
    url: '/about'
    templateUrl: 'templates/about.html'
    controller: 'aboutCtrl'
  ).state(
    'campaigns',
    url: '/campaigns'
    templateUrl: 'templates/campaigns_new.html'
    controller: 'campaignsListCtrl'
  )
]
