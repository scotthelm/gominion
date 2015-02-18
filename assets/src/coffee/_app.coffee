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
    templateUrl: 'templates/campaigns_list.html?order_by&page&per_page&direction'
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
  ).state(
    'spell_types',
    url: '/spell_types'
    templateUrl: 'templates/spell_types_list.html'
    controller: 'spellTypesListCtrl'
  ).state(
    'spell_types_edit',
    url: '/spell_types/{id:int}'
    templateUrl: 'templates/spell_types_update.html'
    controller: 'spellTypesDetailCtrl'
  ).state(
    'spell_types_new',
    url: '/spell_types/new'
    templateUrl: 'templates/spell_types_update.html'
    controller: 'spellTypesCreationCtrl'
  )
]

app.directive 'pagination', () ->
  restrict: 'EA'
  scope:
    resources: '='
    resources_type: '='
    show_pages: '='
  controller: ($scope) ->
    this.current_page = resources.page
    this.previous_page = resources.page - 1
    this.pages = []
    this.next_page = resources.page + 1
    # noop
  templateUrl: 'templates/pagination.html'

