import 'dart:async';

import 'package:landgrab/landgrab.dart';

/// State encapsulates all the information necessary to represent a state of the
/// landgrab game.
class State {
  Timer _timer;

  /// rules of the game.
  final Rules rules;

  Duration _timeRemaining;

  /// currentPlayer in the game.
  final PlayerID currentPlayer;

  /// winner of the game.
  final PlayerID winner;

  /// player1 of the game.
  final Player player1;

  /// player2 of the game.
  final Player player2;

  /// player1Pieces is the set of all Pieces that belong to Player 1.
  final Set<Piece> player1Pieces;

  /// player2Pieces is the set of all Pieces that belong to Player 2.
  final Set<Piece> player2Pieces;

  /// _cells is a mapping of all Cells on the board to the Piece they contain.
  final Map<Cell, Piece> _cells;

  final Map<Piece, Cell> _pieces = {};

  State(this.rules, this.currentPlayer, this.player1, this.player2,
      this.player1Pieces, this.player2Pieces, this._cells,
      {this.winner: PlayerID.noPlayer}) {
    for (Cell c in _cells.keys) {
      _pieces[_cells[c]] = c;
    }
    _timeRemaining = rules.timerDuration;
  }

  /// pieceForCell returns the Piece a Cell contains within the State.
  ///
  /// If the Cell doesn't contain a Piece, NO_PIECE is returned.
  Piece pieceForCell(Cell c) {
    if (!_cells.containsKey(c)) {
      return NO_PIECE;
    }
    return _cells[c];
  }

  Cell cellForPiece(Piece p) {
    if (!_pieces.containsKey(p)) {
      return NO_CELL;
    }
    return _pieces[p];
  }

  /// playerForPiece returns the Player a Piece belongs to.
  ///
  /// Returns PlayerID.noPlayer if no Player within the State owns the Piece.
  PlayerID playerForPiece(Piece p) {
    if (player1Pieces.contains(p)) {
      return PlayerID.player1;
    }
    if (player2Pieces.contains(p)) {
      return PlayerID.player2;
    }
    return PlayerID.noPlayer;
  }

  startTimer() {
    if (_timer != null) return;
    _timer = new Timer.periodic(new Duration(milliseconds: 100), _decreaseTime);
  }

  stopTimer() {
    if (_timer == null) return;
    _timer.cancel();
  }

  Duration get timeRemaining {
    if (_timeRemaining.isNegative) {
      return new Duration();
    }
    return _timeRemaining;
  }

  _decreaseTime(Timer t) {
    if (_timeRemaining.isNegative) {
      stopTimer();
    }
    _timeRemaining = _timeRemaining - new Duration(milliseconds: 100);
  }
}
