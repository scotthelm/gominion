app = angular.module('app.controllers')

app.controller 'campaignsListCtrl', app.listController('campaigns')
app.controller 'campaignCreationCtrl', app.creationController('campaigns')
app.controller 'campaignDetailCtrl', app.detailController('campaigns')

app.controller 'spellTypesListCtrl', app.listController('spell_types')
app.controller 'spellTypesCreationCtrl', app.creationController('spell_types')
app.controller 'spellTypesDetailCtrl', app.detailController('spell_types')

app.controller 'racesListCtrl', app.listController('races')
app.controller 'racesCreationCtrl', app.creationController('races')
app.controller 'racesDetailCtrl', app.detailController('races')

app.controller 'proficiencyTypesListCtrl', app.listController('proficiency_types')
app.controller 'proficiencyTypesCreationCtrl', app.creationController('proficiency_types')
app.controller 'proficiencyTypesDetailCtrl', app.detailController('proficiency_types')
