app = angular.module('app.controllers', [])

app.controller 'campaignsListCtrl', [
  '$scope',
  'ApiFactory',
  '$location',
  '$state',
  ($scope, ApiFactory, $location, $state) ->

    $scope.editCampaign = (campaignId) ->
      $location.path "/campaigns/#{campaignId}"

    $scope.deleteCampaign = (campaignId) ->
      ApiFactory.provider('/api/campaigns/:id').delete(id: campaignId).$promise.then () ->
        $state.transitionTo($state.current, {}, {reload: true});

    $scope.showNewCampaign = () ->
      $location.path "/campaigns/new"

    $scope.campaigns = []
    ApiFactory.provider('/api/campaigns').query().$promise.then (data) ->
      if data != null
        $scope.campaigns = data
      else
        []
  ] 

app.controller('campaignCreationCtrl', [
  '$scope',
  'ApiFactory',
  '$location',
  ($scope, ApiFactory, $location) ->

    $scope.createNewCampaign = () ->
      ApiFactory
        .provider('/api/campaigns/:id')
        .create($scope.campaign).$promise.then () ->
          $location.path "/campaigns"

    $scope.campaign =
      id: ""
      name: ""
      description: ""

    $scope.cancel = () ->
      $location.path "/campaigns"
  ])

app.controller('campaignDetailCtrl', [
  '$scope',
  '$stateParams',
  'ApiFactory',
  '$location',
  ($scope, $stateParams, ApiFactory, $location) ->

    $scope.updateCampaign = ()->
      ApiFactory.provider('/api/campaigns/:id').update($scope.campaign).$promise.then () ->
        $location.path "/campaigns"

    $scope.cancel = ()->
      $location.path "/campaigns"

    $scope.campaign = ApiFactory.provider('/api/campaigns/:id').show(id: $stateParams.id)
  ])
