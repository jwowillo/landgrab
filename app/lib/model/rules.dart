/// Rules a landgrab game is played by.
class Rules {
  /// timerDuration each Player can take while making a move.
  final Duration timerDuration;

  /// pieceCount for each Player.
  final int pieceCount;

  /// boardSize is the size of the board.
  final int boardSize;

  /// life of each Piece initially.
  final int life;

  /// damage of each Piece initially.
  final int damage;

  /// lifeIncrease each time a Piece levels up.
  final int lifeIncrease;

  /// damageIncrease each time a Piece levels up.
  final int damageIncrease;

  /// Rules constructor initializes all fields.
  const Rules(this.timerDuration, this.pieceCount, this.boardSize, this.life,
      this.damage, this.lifeIncrease, this.damageIncrease);
}
