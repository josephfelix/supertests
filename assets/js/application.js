require("angular/angular.min.js");

const requireAll = (r) => { 
  r.keys().forEach(r); 
}

requireAll(require.context('./controllers/', true, /\.js$/));