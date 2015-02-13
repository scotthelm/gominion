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
    templateUrl: 'templates/campaigns_list.html'
    controller: 'campaignsListCtrl'
  ).state(
    'campaigns_edit',
    url: '/campaigns/{id:int}'
    templateUrl: 'templates/campaign_update.html'
    controller: 'campaignDetailCtrl'
  ).state(
    'campaigns_new',
    url: '/campaigns/new'
    templateUrl: 'templates/campaign_update.html'
    controller: 'campaignCreationCtrl'
  )
]
