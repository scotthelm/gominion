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
app.controller 'classesDetailCtrl', app.detailController('classes', ['proficiencies', 'skills'])
