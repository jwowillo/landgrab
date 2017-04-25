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

class Move {
  final Direction direction;

  final Piece piece;

  const Move(this.direction, this.piece);
}
