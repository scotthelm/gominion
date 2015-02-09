app = angular.module 'app', [ 'ui.router' ]
app.config ['$urlRouterProvider', '$stateProvider', ($urlRouterProvider, $stateProvider) ->
  $urlRouterProvider.otherwise('/')
  $stateProvider.state(
    'home',
    url: '/'
    templateUrl: 'templates/home.html'
  ).state(
    'about',
    url: '/about'
    templateUrl: 'templates/about.html'
  )
]
