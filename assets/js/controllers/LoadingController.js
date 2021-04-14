angular.module("supertests").controller("LoadingController", [
    "$scope",
    "$http",
    "$window",
    "$timeout",
    function ($scope, $http, $window, $timeout) {
        $scope.message = "Por favor, aguarde...";

        $scope.makeTest = function (guid) {
            var success = function (result) {
                $window.location.href = "/t/" + guid + "/r/" + result.hash;
            };

            var failure = function () {};

            $timeout(function () {
                $scope.message = "Analisando seu perfil...";

                $timeout(function () {
                    $scope.message = "Gerando resultado...";

                    $http.post("/t/" + guid + "/m").then(success, failure);
                }, 3000);
            }, 1000);
        };
    },
]);
