import 'package:landgrab/landgrab.dart';

enum Direction {
  north,
  northEast,
  east,
  southEast,
  south,
  southWest,
  west,
  northWest
}

String directionToString(Direction d) {
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
  }
}

Direction stringToDirection(String x) {
  switch (x) {
    case 'north':
      return Direction.north;
    case 'north-east':
      return Direction.northEast;
    case 'east':
      return Direction.east;
    case 'south-east':
      return Direction.southEast;
    case 'south':
      return Direction.south;
    case 'south-west':
      return Direction.southWest;
    case 'west':
      return Direction.west;
    case 'north-west':
      return Direction.northWest;
  }
}

class Move {
  final Direction direction;

  final Piece piece;

  const Move(this.direction, this.piece);
}
