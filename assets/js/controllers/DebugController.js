angular.module('supertests')
    .controller('DebugController', function ($scope, $http) {
        $scope.message = '';
        $scope.coordinates = {
            top: 0,
            left: 0
        };
        var imageCoordinates = {left: 0, top: 0, width: 0, height: 0, naturalWidth: 0, naturalHeight: 0};
        $scope.init = function () {
            var image = document.querySelector('#imagem');
            var rect = image.getBoundingClientRect();
            imageCoordinates.left = rect.left;
            imageCoordinates.top = rect.top;
            imageCoordinates.width = image.width;
            imageCoordinates.height = image.height;
            image.addEventListener('load', function () {
                imageCoordinates.naturalWidth = image.naturalWidth;
                imageCoordinates.naturalHeight = image.naturalHeight;
            });
        };
        $scope.countPixels = function ($event) {
            var naturalClickPosX = $event.pageX - imageCoordinates.left;
            var naturalClickPosY = $event.pageY - imageCoordinates.top;
            var scaleX = imageCoordinates.width / imageCoordinates.naturalWidth;
            var scaleY = imageCoordinates.height / imageCoordinates.naturalHeight;

            var calcLeft = Math.ceil(scaleX * naturalClickPosX);
            var calcTop = Math.ceil(scaleY * naturalClickPosY);
            //var naturalClickPosX = Math.ceil((imageCoordinates.width / imageCoordinates.naturalWidth) * $event.x);
            //var naturalClickPosY = Math.ceil((imageCoordinates.height / imageCoordinates.naturalHeight) * $event.y);


            $scope.coordinates = {
                left: naturalClickPosX,
                top: naturalClickPosY
            };
        };

        $scope.enviarTeste = function(guid) {
            if (confirm('Tem certeza que deseja fazer isso?\nATENÇÃO: Ao enviar, este teste ficará visível para todos os usuários.')) {
                $scope.message = 'Enviando...';

                var success = function() {
                    $scope.message = 'Enviado com sucesso!';
                };

                var failure = function() {
                    $scope.message = 'Erro ao enviar teste.';
                };

                $http.post('/debug/' + guid + '/enviar')
                    .then(success, failure);
            }
        };
    });