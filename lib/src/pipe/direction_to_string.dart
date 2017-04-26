import 'package:angular2/angular2.dart';

import 'package:landgrab/landgrab.dart';

@Pipe(name: 'directionToString')
class DirectionToStringPipe extends PipeTransform {
  String transform(Direction d) {
    switch (d) {
      case Direction.north:
        return 'north';
      case Direction.northEast:
        return 'north-east';
      case Direction.east:
        return 'east';
      case Direction.southEast:
        return 'south-east';
      case Direction.south:
        return 'south';
      case Direction.southWest:
        return 'south-west';
      case Direction.west:
        return 'west';
      case Direction.northWest:
        return 'north-west';
      default:
        return '';
    }
  }
}
