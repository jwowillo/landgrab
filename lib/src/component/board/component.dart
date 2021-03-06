import 'dart:collection';

import 'package:angular2/core.dart';

import 'package:landgrab/landgrab.dart';

/// BoardComponent contains a landgrab board with Players already chosen.
///
/// Another parameter, named wait, can be passed. If true, the board will prompt /// the user before proceeding to the next turn.
///
/// At the end of the board, a link will be provided to go back to the
/// GameFormComponent.
@Component(
  selector: 'board',
  templateUrl: 'template.html',
  styleUrls: const ['styles.css'],
  directives: const [MoveChoiceFormComponent],
  pipes: const [PlayerIDToStringPipe, NoSpacePipe],
)
class BoardComponent {
  @Input()
  Rules rules;

  /// state of the board.
  @Input()
  State state;

  @Input()
  List<Move> moves = [];

  @Output()
  final EventEmitter<Map<String, dynamic>> changed = new EventEmitter();

  List<List<Piece>> get grid {
    List<List<Piece>> g = [];
    if (rules == null) return g;
    for (int i = 0; i < rules.boardSize; i++) {
      List<Piece> row = [];
      for (int j = 0; j < rules.boardSize; j++) {
        if (state == null) {
          row.add(NO_PIECE);
        } else {
          row.add(state.pieceForCell(new Cell(i, j)));
        }
      }
      g.add(row);
    }
    return g;
  }

  Map<int, Set<Move>> get bucketedMoves {
    sortedSet() => new SplayTreeSet(
        (Move a, Move b) => a.direction.index.compareTo(b.direction.index));
    Map<int, Set<Move>> buckets = new SplayTreeMap();
    for (Move move in moves) {
      buckets.putIfAbsent(move.piece.id, sortedSet).add(move);
    }
    return buckets;
  }

  playerName(PlayerID id) {
    if (state == null) {
      return '';
    }
    if (id == PlayerID.player1) {
      return state.player1.name;
    }
    if (id == PlayerID.player2) {
      return state.player2.name;
    }
    return '';
  }

  /// isPiece returns true iff the Piece isn't a NO_PIECE.
  isPiece(Piece piece) => piece != NO_PIECE;

  /// isPlayer1 returns true iff the Piece is owned by PlayerID.player1.
  isPlayer1(Piece p) => state.playerForPiece(p) == PlayerID.player1;

  /// isPlayer1 returns true iff the Piece is owned by PlayerID.player2.
  isPlayer2(Piece p) => state.playerForPiece(p) == PlayerID.player2;

  /// isWinner returns true iff the PlayerID isn't PlayerID.noPlayer.
  isWinner(PlayerID id) => id != PlayerID.noPlayer;

  double get timeRemaining {
    return state.timeRemaining.inMilliseconds /
        Duration.MILLISECONDS_PER_SECOND.toDouble();
  }

  State _decidingState;

  Map<int, dynamic> _chosenMoves = {};

  emit(Move chosenMove) {
    if (_decidingState == null || _decidingState != state) {
      _decidingState = state;
      _chosenMoves = {};
    }
    for (Move move in moves) {
      if (chosenMove.piece.id == move.piece.id &&
          chosenMove.direction == move.direction) {
        Cell c = state.cellForPiece(move.piece);
        _chosenMoves[move.piece.id] = {
          'direction': directionToString(move.direction),
          'piece': {
            'id': move.piece.id,
            'player': playerIDToString(state.playerForPiece(move.piece)),
            'life': move.piece.life,
            'damage': move.piece.damage,
            'cell': [c.row, c.column]
          }
        };
      }
    }
    changed.emit({'moves': new List.from(_chosenMoves.values)});
  }
}
