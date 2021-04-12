angular.module("supertests").controller("TesteController", [
    "$scope",
    "$uibModal",
    "$http",
    "$window",
    function ($scope, $uibModal, $http, $window) {
        var scope = $scope;

        /**
         * Status de carregamento do teste
         * @type {boolean}
         */
        $scope.loading = false;

        /**
         * Status de login do usuário no facebook
         * @type {boolean}
         */
        $scope.facebookReady = false;
        $scope.$watch(
            function () {
                return Facebook.isReady();
            },
            function (newVal) {
                $scope.facebookReady = true;
            }
        );

        /**
         * Chamado ao usuário clicar no botão INICIAR TESTE
         */
        $scope.start = function (guid) {
            if ($scope.facebookReady) {
                $scope.loading = true;
                $scope.goToQuiz(guid);
            } else {
                $scope.openModal();
            }
        };

        /**
         * Inicia a realização do teste
         * este método só é chamado se o usuário
         * estiver logado no facebook
         */
        $scope.goToQuiz = function (guid) {
            $window.location.href = "/t/" + guid + "/l";
        };

        /**
         * Realiza o login no site
         * @param guid
         */
        $scope.loginSite = function (guid) {
            $scope.loading = true;
            /* var fields = [
                'id', 'name', 'age_range', 'birthday', 'cover', 'email', 'favorite_teams',
                'favorite_athletes', 'gender', 'context'
            ];*/

            var fields = [
                "id",
                "name",
                "age_range",
                "cover",
                "email",
                "gender",
                "context",
            ];
            FB.api("/me?fields=" + fields.join(","), function (response) {
                var sucessoLogin = function (result) {
                    if (result.status) {
                        $scope.goToQuiz(guid);
                    } else {
                        $scope.loading = false;
                        alert(result.result);
                    }
                };

                var errorLogin = function () {
                    $scope.loading = false;
                    $scope.openModal();
                    alert("Ocorreu um erro ao fazer login, tente novamente.");
                };

                $http.post("/login", response).then(sucessoLogin, errorLogin);
            });
        };

        /**
         * Realiza o login no facebook via JS SDK
         */
        $scope.loginFacebook = function (guid) {
            FB.login(
                function (response) {
                    if (response.authResponse) {
                        $scope.closeModal();
                        $scope.loginSite(guid);
                    }
                },
                { scope: "email" }
            );
        };

        /**
         * Abre a modal com o botão de conectar ao facebook
         */
        $scope.openModal = function () {
            $uibModal.open({
                animation: false,
                templateUrl: "conectar_facebook.html",
                size: "sm",
                controller: function ($scope, $uibModalInstance) {
                    $scope.closeModal = function () {
                        $uibModalInstance.close();
                    };

                    $scope.loginFacebook = scope.loginFacebook;
                    scope.closeModal = $scope.closeModal;
                },
            });
        };

        /**
         * Carrega mais testes abaixo do teste atual
         */
        $scope.loadingtestes = true;
        $scope.tests = [];
        $scope.loadTests = function (active) {
            $http.get("/l").then(function (json) {
                $scope.loadingtestes = false;
                for (var index in json.data) {
                    if (json.data[index].id != active) {
                        $scope.tests.push(json.data[index]);
                    }
                }
            });
        };
    },
]);
