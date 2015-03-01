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
    url: '/campaigns?order_by&page&per_page&direction'
    templateUrl: 'templates/campaigns/list.html'
    controller: 'campaignsListCtrl'
  ).state(
    'campaigns_edit',
    url: '/campaigns/{id:int}'
    templateUrl: 'templates/campaigns/update.html'
    controller: 'campaignDetailCtrl'
  ).state(
    'campaigns_new',
    url: '/campaigns/new'
    templateUrl: 'templates/campaigns/update.html'
    controller: 'campaignDetailCtrl'
  ).state(
    'spell_types',
    url: '/spell_types?order_by&page&per_page&direction'
    templateUrl: 'templates/spell_types/list.html'
    controller: 'spellTypesListCtrl'
  ).state(
    'spell_types_edit',
    url: '/spell_types/{id:int}'
    templateUrl: 'templates/spell_types/update.html'
    controller: 'spellTypesDetailCtrl'
  ).state(
    'spell_types_new',
    url: '/spell_types/new'
    templateUrl: 'templates/spell_types/update.html'
    controller: 'spellTypesDetailCtrl'
  ).state(
    'races',
    url: '/races?order_by&page&per_page&direction'
    templateUrl: 'templates/races/list.html'
    controller: 'racesListCtrl'
  ).state(
    'races_edit',
    url: '/races/{id:int}'
    templateUrl: 'templates/races/update.html'
    controller: 'racesDetailCtrl'
  ).state(
    'races_new',
    url: '/races/new'
    templateUrl: 'templates/races/update.html'
    controller: 'racesDetailCtrl'
  ).state(
    'proficiency_types',
    url: '/proficiency_types?order_by&page&per_page&direction'
    templateUrl: 'templates/proficiency_types/list.html'
    controller: 'proficiencyTypesListCtrl'
  ).state(
    'proficiency_types_edit',
    url: '/proficiency_types/{id:int}'
    templateUrl: 'templates/proficiency_types/update.html'
    controller: 'proficiencyTypesDetailCtrl'
  ).state(
    'proficiency_types_new',
    url: '/proficiency_types/new'
    templateUrl: 'templates/proficiency_types/update.html'
    controller: 'proficiencyTypesDetailCtrl'
  ).state(
    'proficiencies',
    url: '/proficiencies?order_by&page&per_page&direction'
    templateUrl: 'templates/proficiencies/list.html'
    controller: 'proficienciesListCtrl'
  ).state(
    'proficiencies_edit',
    url: '/proficiencies/{id:int}'
    templateUrl: 'templates/proficiencies/update.html'
    controller: 'proficienciesDetailCtrl'
  ).state(
    'proficiencies_new',
    url: '/proficiencies/new'
    templateUrl: 'templates/proficiencies/update.html'
    controller: 'proficienciesDetailCtrl'
  ).state(
    'skills',
    url: '/skills?order_by&page&per_page&direction'
    templateUrl: 'templates/skills/list.html'
    controller: 'skillsListCtrl'
  ).state(
    'skills_edit',
    url: '/skills/{id:int}'
    templateUrl: 'templates/skills/update.html'
    controller: 'skillsDetailCtrl'
  ).state(
    'skills_new',
    url: '/skills/new'
    templateUrl: 'templates/skills/update.html'
    controller: 'skillsDetailCtrl'
  ).state(
    'classes',
    url: '/classes?order_by&page&per_page&direction'
    templateUrl: 'templates/classes/list.html'
    controller: 'classesListCtrl'
  ).state(
    'classes_edit',
    url: '/classes/{id:int}'
    templateUrl: 'templates/classes/update.html'
    controller: 'classesDetailCtrl'
  ).state(
    'classes_new',
    url: '/classes/new'
    templateUrl: 'templates/classes/update.html'
    controller: 'classesDetailCtrl'
  )
]

app.directive 'pagination', ['$state', ($state) ->
  restrict: 'EA'
  scope:
    apiService: '=apiService'
    resourceType: '=resourceType'
  controller: ($scope) ->
    $scope.paginator = 
      total_records: 0
      total_pages: 0
      current_page: 1
      current_chapter: 1
      per_page: 10
      pages_per_chapter: 5
      total_chapters: 1
      chapter: []
      resource_type: $scope.resourceType
      interval: {}
      order_by: "name"
      direction: "asc"
    $scope.apiService.$promise.then (data) ->
      total_records = data.total_records
      per_page = data.per_page
      $scope.paginator.total_pages = Math.ceil(total_records / per_page)
      $scope.paginator.current_page = parseInt(data.page)
      $scope.paginator.per_page = parseInt(data.per_page)
      $scope.paginator.total_records = parseInt(data.total_records)
      $scope.paginate($scope.paginator)
    $scope.paginate = (p) ->
      p.total_chapters = Math.ceil(p.total_pages / p.pages_per_chapter)
      p.current_chapter = Math.ceil(p.current_page / p.pages_per_chapter)
      p.interval =
        lbound:
          (p.current_chapter * p.pages_per_chapter) - p.pages_per_chapter + 1
        rbound:
          if p.total_pages >= p.current_chapter * p.pages_per_chapter
            p.current_chapter * p.pages_per_chapter + 1
          else
            p.total_pages + 1
      (p.chapter.push({page: i, current: i == p.current_page}) ) for i in [p.interval.lbound...p.interval.rbound]
          

    $scope.goToState = (page) ->
      $state.go($scope.paginator.resource_type,
        page: page,
        location: true
      )

  templateUrl: 'templates/pagination.html'
]

