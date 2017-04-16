angular.module('HelloWorldApp', ['ngWebSocket'])
   .controller('HelloWorldController', function($scope, $websocket) {

       window.dataStream = $websocket('ws://127.0.0.1:8080/ws');

        var collection = [];


        $scope.nodes = [
        {name : "Node1", role: "cassandra", hdds: [{name : "xvdf", used: 50}, {name : "xvgf", used: 70}]},
        {name : "Node2", role: "hector", hdds: [{name : "xvgf", used: 70}]},
        {name : "Node3", hdds: [{name: "xkvf", used: 10}]}
        ]


       $scope.greeting = [];


        dataStream.onMessage(function(message) {

//            myMessage = JSON.parse(message.data)
           //  $scope.greeting = message.data // myMessage.hello

            $scope.greeting.push(message.data)
            console.log(message.data)


        });


        setInterval(function(){

             dataStream.send("heyyyyyyyyy")
        }, 15000)


        dataStream.send("heyyyyyyyyy")

});