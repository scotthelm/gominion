app = angular.module('app.controllers', [])

app.controller('campaignsListCtrl', [
  '$scope',
  'CampaignsFactory',
  'CampaignFactory',
  '$location',
  ($scope, CampaignsFactory, CampaignFactory, $location) ->

    $scope.editCampaign = (campaignId) ->
      $location.path "/campaigns/#{campaignId}"

    $scope.deleteCampaign = (campaignId) ->
      CampaignFactory.delete id: campaignId

    $scope.newCampaign = () ->
      $location.path "/campaigns/new"

    $scope.campaigns = CampaignsFactory.query()
  ])

app.controller('campaignCreationCtrl', [
  '$scope',
  'CampaignsFactory',
  '$location',
  ($scope, CampaignsFactory, $location) ->
    $scope.createNewCampaign = () ->
      CampaignsFactory.newCampaign($scope.user)
      $location.path "/campaigns"
  ])

app.controller('campaignDetailCtrl', [
  '$scope',
  '$stateParams',
  'CampaignFactory',
  '$location',
  ($scope, $stateParams, CampaignFactory, $location) ->
    $scope.updateCampaign = ()->

      debugger
      CampaignFactory.update($scope.campaign)
      $location.path "/campaigns"

    $scope.cancel = ()->
      $location.path "/campaigns"

    $scope.campaign = CampaignFactory.show(id: $stateParams.id)
  ])
