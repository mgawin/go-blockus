app.factory('backendService', function($http) {
  var backendService = {
    getGame: function(gid) {
      var promise = $http.get('https://golang-mgawin.c9.io/status?id='+gid).then(function (response) {
        return response.data;
      });
      return promise;
    },
    init: function() {
     
      var promise = $http.get('https://golang-mgawin.c9.io/new').then(function (response) {
        return response.data;
      });
      return promise;
    },
    getMoves: function(gid,pid,bid,rotates){
      var promise = $http.get('https://golang-mgawin.c9.io/moves?gid='+gid+'&pid='+pid+'&bid='+bid+'&rotates='+rotates).then(function (response) {
        return response.data;
      
      })
    
      return promise;
    },
    doMove: function(gid,pid,bid,rotates,x,y){
      var promise = $http.get('https://golang-mgawin.c9.io/move?gid='+gid+'&pid='+pid+'&bid='+bid+'&rotates='+rotates+'&x='+x+'&y='+y).then(function (response) {
        return response.data;
      
      })
    
      return promise;
    }
    
      
    
    
    
  };
  return backendService;
});
