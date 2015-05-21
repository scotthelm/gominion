app = angular.module('app.controllers')

app.controller 'campaignsListCtrl', app.listController('campaigns')
app.controller 'campaignDetailCtrl', app.detailController('campaigns')

app.controller 'spellTypesListCtrl', app.listController('spell_types')
app.controller 'spellTypesDetailCtrl', app.detailController('spell_types')

app.controller 'racesListCtrl', app.listController('races')
app.controller 'racesDetailCtrl', app.detailController('races')

app.controller 'proficiencyTypesListCtrl', app.listController('proficiency_types')
app.controller 'proficiencyTypesDetailCtrl', app.detailController('proficiency_types')

app.controller 'proficienciesListCtrl', app.listController('proficiencies')
app.controller 'proficienciesDetailCtrl', app.detailController('proficiencies', ['proficiency_types'])

app.controller 'skillsListCtrl', app.listController('skills')
app.controller 'skillsDetailCtrl', app.detailController('skills')

app.controller 'classesListCtrl', app.listController('classes')
app.controller 'classesDetailCtrl',
  app.detailController('classes',
  ['proficiencies', 'skills'],
  (scope, stateParams, apifactory, location, state) ->
    scope.addProficiency = (res, prof) ->
      proficiency = (prof) ->
        for index, p of scope.proficiencies.result
          if prof == p.id
            return p
      res.proficiencies.push proficiency prof
      console.log(res)

    scope.removeProficiency = (res, prof) ->
      res.proficiencies.splice(res.proficiencies.indexOf(prof), 1)
)
