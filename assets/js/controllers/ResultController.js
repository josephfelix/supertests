angular.module("supertests").controller("ResultController", [
    "$scope",
    "$http",
    function ($scope, $http) {
        /**
         * Compartilha o resultado no facebook
         */
        $scope.shareFacebook = function () {
            var url = window.location.href;

            if (url.indexOf("?") != -1) {
                url = url + "&fb=1";
            } else {
                url = url + "?fb=1";
            }

            FB.ui(
                {
                    method: "share",
                    display: "popup",
                    href: url,
                    hashtag: "#TestesWeb",
                },
                function (response) {}
            );
        };

        /**
         * Carrega mais testes abaixo do teste atual
         */
        $scope.loading = true;
        $scope.tests = [];
        $scope.loadTests = function (active) {
            $http.get("/l").then(function (json) {
                $scope.loading = false;
                for (var index in json.data) {
                    if (json.data[index].id != active) {
                        $scope.tests.push(json.data[index]);
                    }
                }
            });
        };
    },
]);
