require("angular/angular.min.js");
require("angularjs-ui-bootstrap/ui-bootstrap.min.js");
require("angularjs-facebook/lib/angular-facebook.js");

angular.module('supertests', ['ui.bootstrap', 'facebook'])
  .config(function(FacebookProvider) {
     FacebookProvider.init('258828887510894');
  });

const requireAll = (r) => { 
  r.keys().forEach(r); 
}

requireAll(require.context('./controllers/', true, /\.js$/));