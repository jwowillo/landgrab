/// Piece in a landgrab game.
class Piece {
  /// id of the Piece.
  final int id;

  /// life the Piece has left.
  final int life;

  /// damage the Piece does.
  final int damage;

  /// Piece constructor sets the id, life, and damage of the Piece.
  const Piece(this.id, this.life, this.damage);
}

/// Cell which may contain a Piece.
class Cell {
  /// row of the Cell.
  final int row;

  /// column of the Cell.
  final int column;

  /// Cell constructor sets the row and column of the Cell.
  const Cell(this.row, this.column);
}

/// NO_PIECE_ID is used when a int representing a Piece's ID is expected but the
/// Piece doesn't exist.
final int NO_PIECE_ID = -1;

/// NO_PIECE is used when a Piece is expected but doesn't exist.
final Piece NO_PIECE = const Piece(NO_PIECE_ID, -1, -1);

/// NO_CELL is used when a Cell is expected but doesn't exist.
final Cell NO_CELL = const Cell(-1, -1);
