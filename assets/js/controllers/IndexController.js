angular.module("supertests").controller("IndexController", [
    "$scope",
    "$http",
    "$window",
    "Facebook",
    function ($scope, $http, $window, Facebook) {
        var limit = 30;

        $scope.tests = [];

        $scope.loading = true;

        $scope.init = function () {
            $http.get("/l").then(function (json) {
                $scope.loading = false;
                $scope.tests = json.data;
            });
        };

        /**
         * Realiza o login no facebook via JS SDK
         */
        $scope.loginFacebook = function () {
            Facebook.login(
                function (response) {
                    if (response.authResponse) {
                        $scope.loginSite();
                    }
                },
                { scope: "email" }
            );
        };

        /**
         * Realiza o login no site
         * @param guid
         */
        $scope.loginSite = function (guid) {
            var fields = ["id", "name", "email", "picture"];

            Facebook.api("/me?fields=" + fields.join(","), function (response) {
                var loginSuccess = function (result) {
                    if (!result.status) {
                        alert(result.result);
                    } else {
                        $window.location.reload();
                    }
                };

                var loginError = function () {
                    alert("Ocorreu um erro ao fazer login, tente novamente.");
                };

                var userData = {
                    id: response.id,
                    name: response.name,
                    email: response.email,
                    photo: response.picture.data.url,
                };

                $http.post("/login", userData).then(loginSuccess, loginError);
            });
        };
    },
]);
