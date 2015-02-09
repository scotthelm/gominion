angular
  .module('app')
    .controller('aboutCtrl', ['$scope', ($scope) ->
      $scope.title = "About"
      $scope.items = ['other', 'stuff', 'in here']
    ])

