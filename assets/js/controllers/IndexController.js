angular.module("supertests").controller("IndexController", [
    "$scope",
    "$http",
    function ($scope, $http) {
        var limit = 30;
        $scope.testes = [];
        $scope.loading = true;
        $scope.init = function () {
            $http.get("/l").then(function (json) {
                $scope.loading = false;
                $scope.testes = json.data;
            });
        };
    },
]);
