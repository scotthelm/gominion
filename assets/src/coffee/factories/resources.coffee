services = angular.module('app.services', ['ngResource'])

services.factory 'ApiFactory', ['$resource', ($resource) ->
  provider: (url, name) ->
    $resource(
      url,
      {}
      query:
        method: 'GET'
        params:
          order_by: '@order_by'
          direction: '@direction'
          page: '@page'
          per_page: '@per_page'
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

