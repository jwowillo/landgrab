import 'package:landgrab/model/board.dart';
import 'package:landgrab/model/player.dart';

/// State encapsulates all the information necessary to represent a state of the
/// landgrab game.
class State {
  /// currentPlayer in the game.
  final PlayerID currentPlayer;

  /// winner of the game.
  final PlayerID winner;

  /// player1 of the game.
  final Player player1;

  /// player2 of the game.
  final Player player2;

  /// boardSize of the game is the n value of the nxn game board.
  final int boardSize;

  /// _player1Pieces is the set of all Pieces that belong to Player 1.
  final Set<Piece> _player1Pieces;

  /// _player2Pieces is the set of all Pieces that belong to Player 2.
  final Set<Piece> _player2Pieces;

  /// _cells is a mapping of all Cells on the board to the Piece they contain.
  final Map<Cell, Piece> _cells;

  const State(this.currentPlayer, this.player1, this.player2,
      this._player1Pieces, this._player2Pieces, this._cells, this.boardSize,
      {this.winner: PlayerID.noPlayer});

  /// pieceForCell returns the Piece a Cell contains within the State.
  ///
  /// If the Cell doesn't contain a Piece, NO_PIECE is returned.
  Piece pieceForCell(Cell c) {
    if (!_cells.containsKey(c)) {
      return NO_PIECE;
    }
    return _cells[c];
  }

  /// playerForPiece returns the Player a Piece belongs to.
  ///
  /// Returns PlayerID.noPlayer if no Player within the State owns the Piece.
  PlayerID playerForPiece(Piece p) {
    if (_player1Pieces.contains(p)) {
      return PlayerID.player1;
    }
    if (_player2Pieces.contains(p)) {
      return PlayerID.player2;
    }
    return PlayerID.noPlayer;
  }
}
