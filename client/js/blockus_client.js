angular.module('blockusApp', [])
  .controller('GameController', ['$scope', function($scope) {
    $scope.blocks = [{
        "val": 1,
        "shape": [[2]]
      }, {
        "val": 2,
        "shape": [[2, 2]]
      }, {
        "val": 3,
        "shape": [[2, 2], [0, 2]]
      },
      {
        "val": 3,
        "shape": [[2, 2, 2]]
      }, {
        "val": 4,
        "shape": [[2, 2], [2, 2]]
      }, {
        "val": 4,
        "shape": [[0, 2, 0], [2, 2, 2]]
      },
      {
        "val": 4,
        "shape": [[2, 2, 2, 2]]
      }, {
        "val": 4,
        "shape": [[0, 0, 2], [2, 2, 2]]
      }, {
        "val": 4,
        "shape": [[0, 2, 2], [2, 2, 0]]
      },
      {
        "val": 5,
        "shape": [[2, 0, 0, 0], [2, 2, 2, 2]]
      }, {
        "val": 5,
        "shape": [[0, 2, 0], [0, 2, 0], [2, 2, 2]]
      }, {
        "val": 5,
        "shape": [[2, 0, 0],
    [2, 0, 0], [2, 2, 2]]
      }, {
        "val": 5,
        "shape": [[0, 2, 2, 2], [2, 2, 0, 0]]
      }, {
        "val": 5,
        "shape": [[0, 0, 2], [2, 2, 2], [2, 0, 0]]
      },
      {
        "val": 5,
        "shape": [[2, 2, 2, 2, 2]]
      }, {
        "val": 5,
        "shape": [[2, 0], [2, 2], [2, 2]]
      }, {
        "val": 5,
        "shape": [[0, 2, 2],
    [2, 2, 0], [2, 0, 0]]
      }, {
        "val": 5,
        "shape": [[2, 2], [2, 0], [2, 2]]
      }, {
        "val": 5,
        "shape": [[0, 2, 2], [2, 2, 0], [0, 2, 0]]
      },
      {
        "val": 5,
        "shape": [[0, 2, 0], [2, 2, 2], [0, 2, 0]]
      }, {
        "val": 5,
        "shape": [[0, 2, 0, 0], [2, 2, 2, 2]]
      }]
//Draw board
    var canvas = new fabric.Canvas('c1', {
      selection: false
    });

    var grid = 24;
    var shiftx = 96;
    var shifty = 48;

    for (var i = 0; i < 15; i++) {
      canvas.add(new fabric.Line([shiftx + i * grid, shifty, shiftx + i * grid, shifty + 14 * grid], {
        stroke: '#ccc',
        selectable: false
      }));
      canvas.add(new fabric.Line([shiftx, shifty + i * grid, shiftx + 14 * grid, shifty + i * grid], {
        stroke: '#ccc',
        selectable: false
      }))
    }

    canvas.add(new fabric.Rect({
      left: shiftx + 4 * grid,
      top: shifty + 9 * grid,
      width: grid,
      height: grid,
      fill: '#ccc',
      opacity: 0.5,
      selectable: false
    }));

    canvas.add(new fabric.Rect({
      left: shiftx + 9 * grid,
      top: shifty + 4 * grid,
      width: grid,
      height: grid,
      fill: '#ccc',
      opacity: 0.5,
      selectable: false
    }));
//Draw elements
    var y = grid;
    var x = 450;

    $scope.blocks.forEach(function(element, index) {

      x = x + 120;
      if (x > 450 + 600) {
        x = 450 + 120
        y = y + 80;
      }

      var group = new fabric.Group([], {
        hasControls: false,
        hasBorders: false,
      });



      for (var j = 0; j < element.shape.length; j++) {

        for (var i = 0; i < element.shape[0].length; i++) {

          if (element.shape[j][i] > 0) {

            var shadow = {
              color: 'rgba(0,0,0,0.4)',
              blur: 20,
              offsetX: 10,
              offsetY: 10,
              opacity: 0.2,
              fillShadow: true,
              strokeShadow: true
            }

            var rect = new fabric.Rect({
              originX: 'left',
              originY: 'top',
              left: grid * i,
              top: grid * j,
              fill: "orange",
              stroke: "#000",
              width: grid,
              height: grid,
              strokeWidth: 2,
              opacity: .8
            });


            rect.setShadow(shadow);
            group.addWithUpdate(rect);

          }
        }

      }
      group.left = x;
      group.top = y;
      canvas.add(group);



    });



    canvas.on('object:moving', function(options) {

      if ((options.target.left <= 430) && (options.target.top <= 430 + shifty) && (options.target.left > shiftx - 5 * grid)) {
        options.target.set({
          left: Math.round(options.target.left / grid) * grid,
          top: Math.round(options.target.top / grid) * grid
        });
        //remove shadow
        options.target.forEachObject(function(e) {
          e.setShadow(null);

        });

      }
    });

    canvas.on('mouse:down', function(options) {

      options.target.angle+=90;

    });





  }]);