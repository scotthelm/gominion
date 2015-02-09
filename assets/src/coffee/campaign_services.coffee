services = angular.module('app.services', ['ngResource'])

services.factory 'CampaignsFactory', ['$resource', ($resource) ->
  $resource(
    '/api/campaigns',
    {},
    query:
      method: 'GET'
      isArray: true
    create:
      method: 'POST'
  )
]

services.factory 'CampaignFactory', ['$resource', ($resource) ->
  $resource(
    '/api/campaign/:id',
    {},
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
