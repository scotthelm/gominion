angular
  .module('app')
    .controller('aboutCtrl', ['$scope', ($scope) ->
      $scope.title = "About"
      $scope.items = ['some', 'other', 'stuff']
    ])

