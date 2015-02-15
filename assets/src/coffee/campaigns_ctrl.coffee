app = angular.module('app.controllers', [])

app.controller 'campaignsListCtrl', [
  '$scope',
  'CampaignsFactory',
  'CampaignFactory',
  '$location',
  '$state',
  ($scope, CampaignsFactory, CampaignFactory, $location, $state) ->

    $scope.editCampaign = (campaignId) ->
      $location.path "/campaigns/#{campaignId}"

    $scope.deleteCampaign = (campaignId) ->
      CampaignFactory.delete(id: campaignId).$promise.then () ->
        $state.transitionTo($state.current, {}, {reload: true});

    $scope.showNewCampaign = () ->
      $location.path "/campaigns/new"

    $scope.campaigns = []
    CampaignsFactory.query().$promise.then (data) ->
      if data != null
        $scope.campaigns = data
      else
        []
  ] 

app.controller('campaignCreationCtrl', [
  '$scope',
  'CampaignsFactory',
  '$location',
  ($scope, CampaignsFactory, $location) ->

    $scope.createNewCampaign = () ->
      CampaignsFactory.create($scope.campaign).$promise.then () ->
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
  'CampaignFactory',
  '$location',
  ($scope, $stateParams, CampaignFactory, $location) ->

    $scope.updateCampaign = ()->
      CampaignFactory.update($scope.campaign).$promise.then () ->
        $location.path "/campaigns"

    $scope.cancel = ()->
      $location.path "/campaigns"

    $scope.campaign = CampaignFactory.show(id: $stateParams.id)
  ])
