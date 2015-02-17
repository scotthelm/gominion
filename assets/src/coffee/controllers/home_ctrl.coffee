angular
  .module('app')
    .controller('homeCtrl', ['$scope', ($scope) ->
      $scope.title = "Home"
      $scope.items = ['this', 'that', 'the other thing']
    ])

