var app = angular.module("app", ['ngRoute', 'ngResource'], function($routeProvider){
	$routeProvider.when("/", {
		templateUrl: "/tasks.html",
		controller: "TasksCtl"
	});
});

app.controller("MainCtl", function($scope){
	
})

app.controller("TasksCtl", function($scope, $resource){
	var Task = $resource("/tasks/:id", {id: '@id'}, {});

	$scope.tasks = Task.query();
})