var angularApp = angular.module('angularApp',[]);

angularApp.controller('FirstNameListCtrl', function ($scope) {
	$scope.firstnames = [
    {'name': 'Dogely',
     'votes': '6'},
    {'name': 'Hello Doge',
     'votes': '3'},
    {'name': 'wow',
     'votes': '9'}
  ];
})



angularApp.controller('LastNameListCtrl', function ($scope) {
	$scope.firstnames = [
    {'name': 'Doge',
     'votes': '12'},
    {'name': 'Dogezilla™',
     'votes': '4'},
    {'name': 'Da-moon',
     'votes': '9'}
  ];
})