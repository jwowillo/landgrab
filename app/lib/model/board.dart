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

  /// hashCode is the Piece's ID.
  int get hashCode => id;

  /// == returns trie iff the other Piece has the same field values.
  bool operator ==(Piece other) =>
      other.id == id && other.life == life && other.damage == damage;
}

/// Cell which may contain a Piece.
class Cell {
  /// row of the Cell.
  final int row;

  /// column of the Cell.
  final int column;

  /// Cell constructor sets the row and column of the Cell.
  const Cell(this.row, this.column);

  /// hashCode is a combination of the Cell's row and column.
  int get hashCode => (row << 16) ^ column;

  /// == returns true iff the other Cell  has the same field values.
  bool operator ==(Cell other) => other.row == row && other.column == column;
}

/// NO_PIECE is used when a Piece is expected but doesn't exist.
final Piece NO_PIECE = const Piece(-1, -1, -1);

/// NO_CELL is used when a Cell is expected but doesn't exist.
final Cell NO_CELL = const Cell(-1, -1);
