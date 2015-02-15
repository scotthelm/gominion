services = angular.module('app.services', ['ngResource'])

services.factory 'ApiFactory', ['$resource', ($resource) ->
  provider: (url, name) ->
    $resource(
      url,
      {}
      query:
        method: 'GET'
        isArray: true
      create:
        method: 'POST'
      show:
        method: 'GET'
      update:
        method: 'PUT'
        params:
          id: '@id'
      delete:
        method: 'DELETE'
        params:
          id: '@id'
    )
]

