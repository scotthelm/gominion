angular
  .module('app')
    .controller('aboutCtrl', ['$scope', ($scope) ->
      $scope.title = "About"
      $scope.items = ['some', 'other', 'stuff']
      $scope.matrix_values = [0,1,2,3,4,5,6,7]
    ])

