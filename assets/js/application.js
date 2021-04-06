require("angular/angular.min.js");

angular.module('supertests', []);

const requireAll = (r) => { 
  r.keys().forEach(r); 
}

requireAll(require.context('./controllers/', true, /\.js$/));