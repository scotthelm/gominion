app = angular.module( 'app.controllers', [])

app.listController = (resource_type) ->
  [
    '$scope',
    'ApiFactory',
    '$location',
    '$state',
    ($scope, ApiFactory, $location, $state) ->
      self = this
      self.resourceType = resource_type

      $scope.editResource = (resourceId) ->
        $location.path "/#{self.resourceType}/#{resourceId}"

      $scope.deleteResource = (resourceId) ->
        ApiFactory.provider("/api/#{self.resourceType}/:id").delete(id: resourceId).$promise.then () ->
          $state.transitionTo($state.current, {}, {reload: true});

      $scope.showNewResource = () ->
        $location.path "/#{self.resourceType}/new"

      $scope.resources = []
      ApiFactory.provider("/api/#{self.resourceType}").query().$promise.then (data) ->
        if data != null
          $scope.resources = data
        else
          []
  ]

app.creationController = (resource_type) ->
  [
    '$scope',
    'ApiFactory',
    '$location',
    ($scope, ApiFactory, $location) ->
      self = this
      self.resourceType = resource_type

      $scope.createNewResource = () ->
        ApiFactory
          .provider("/api/#{self.resourceType}")
          .create($scope.resource).$promise.then () ->
            $location.path "/#{self.resourceType}"

      $scope.resource =  {}

      $scope.cancel = () ->
        $location.path "/#{self.resourceType}"
  ]

 app.detailController = (resource_type) ->
  [
    '$scope',
    '$stateParams',
    'ApiFactory',
    '$location',
    ($scope, $stateParams, ApiFactory, $location) ->
      self = this
      self.resourceType = resource_type

      $scope.updateResource = ()->
        ApiFactory
          .provider("/api/#{self.resourceType}/#{$scope.resource.id}")
          .update($scope.resource)
          .$promise
          .then () ->
            $location.path "/#{self.resourceType}"

      $scope.cancel = ()->
        $location.path "/#{self.resourceType}"

      $scope.resource = ApiFactory.provider("/api/#{self.resourceType}/:id").show(id: $stateParams.id)
  ]
