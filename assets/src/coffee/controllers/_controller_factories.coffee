app = angular.module( 'app.controllers', [])

app.location_defaults = 
  order_by: 'name'
  direction: 'asc'
  page: 1
  per_page: 8

app.listController = (resource_type) ->
  [
    '$scope',
    'ApiFactory',
    '$location',
    '$state',
    ($scope, ApiFactory, $location, $state) ->
      self = this
      self.resourceType = resource_type

      $scope.resourceType = self.resourceType

      $scope.editResource = (resourceId) ->
        $location.path "/#{self.resourceType}/#{resourceId}"

      $scope.deleteResource = (resourceId) ->
        ApiFactory
          .provider("/api/#{self.resourceType}/:id")
          .delete(id: resourceId)
          .$promise.then () ->
            $state.transitionTo($state.current, app.location_defaults, {reload: true});

      $scope.showNewResource = () ->
        $location.path "/#{self.resourceType}/new"

      $scope.resources = {}
      $scope.apiService = ApiFactory
        .provider("/api/#{self.resourceType}")
        .query(
          order_by: $location.search().order_by
          direction: $location.search().direction
          page: $location.search().page
          per_page: $location.search().per_page
        )
      $scope.apiService.$promise.then (data) ->
        if data != null
          $scope.resources = data
        else
          $scope.resources = {}
  ]

 app.detailController = (resource_type, preloads, extension) ->
  [
    '$scope',
    '$stateParams',
    'ApiFactory',
    '$location',
    '$state',
    ($scope, $stateParams, ApiFactory, $location, $state) ->
      self = this
      self.resourceType = resource_type
      preloads = [] unless preloads

      $scope.createNewResource = () ->
        ApiFactory
          .provider("/api/#{self.resourceType}")
          .create($scope.resource).$promise.then () ->
            $state.transitionTo(self.resourceType, app.location_defaults)

      $scope.updateResource = ()->
        ApiFactory
          .provider("/api/#{self.resourceType}/#{$scope.resource.id}")
          .update($scope.resource)
          .$promise
          .then () ->
            $state.transitionTo(self.resourceType, app.location_defaults)
      (
        $scope[pre] =
        ApiFactory
          .provider("/api/#{pre}")
          .query(
            order_by: 'name'
            page: 1
            per_page: 200
          )
        $scope[pre].$promise.then (data) ->
          $scope["#{pre}_data"] = data

      ) for pre in preloads


      $scope.cancel = ()->
        $state.transitionTo(self.resourceType, app.location_defaults)

      $scope.resource = ApiFactory.provider("/api/#{self.resourceType}/:id").show(id: $stateParams.id)
      if extension
        extension($scope, $stateParams, ApiFactory, $location, $state)
  ]
