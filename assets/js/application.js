require("angular/angular.min.js");
require("angularjs-ui-bootstrap/ui-bootstrap.min.js");

angular.module('supertests', ['ui.bootstrap']);

const requireAll = (r) => { 
  r.keys().forEach(r); 
}

requireAll(require.context('./controllers/', true, /\.js$/));