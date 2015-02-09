angular
  .module('app.controllers', [])
    .controller('campaignsListCtrl', [
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
