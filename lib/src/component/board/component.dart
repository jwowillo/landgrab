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

  // Need something that returns a list of lists of moves by player id.
  List<Map<String, dynamic>> get movesByPiece {
    Map<int, List<String>> byPiece = {};
    for (Move move in moves) {
      if (!byPiece.containsKey(move.piece.id)) {
        byPiece[move.piece.id] = [];
      }
      String str = '';
      switch (move.direction) {
        case Direction.north:
          str = 'north';
          break;
        case Direction.northEast:
          str = 'north-east';
          break;
        case Direction.east:
          str = 'east';
          break;
        case Direction.southEast:
          str = 'south-east';
          break;
        case Direction.south:
          str = 'south';
          break;
        case Direction.southWest:
          str = 'south-west';
          break;
        case Direction.west:
          str = 'west';
          break;
        case Direction.northWest:
          str = 'north-west';
          break;
      }
      byPiece[move.piece.id].add(str);
    }
    for (List<String> list in byPiece.values) {
      list.sort((String a, String b) => a.compareTo(b));
    }
    List<Map<String, dynamic>> out = [];
    for (int id in byPiece.keys) {
      out.add({'id': id, 'directions': byPiece[id]});
    }
    out.sort((Map<String, dynamic> a, Map<String, dynamic> b) =>
        a['id'].compareTo(b['id']));
    return out;
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

  Map<int, dynamic> _moves = {};

  emit(int id, String rawD) {
    if (_decidingState == null || _decidingState != state) {
      _decidingState = state;
      _moves = {};
    }
    Direction d;
    int dInt;
    switch (rawD) {
      case 'north':
        dInt = 0;
        d = Direction.north;
        break;
      case 'north-east':
        dInt = 1;
        d = Direction.northEast;
        break;
      case 'east':
        dInt = 2;
        d = Direction.east;
        break;
      case 'south-east':
        dInt = 3;
        d = Direction.southEast;
        break;
      case 'south':
        dInt = 4;
        d = Direction.south;
        break;
      case 'south-west':
        dInt = 5;
        d = Direction.southWest;
        break;
      case 'west':
        dInt = 6;
        d = Direction.west;
        break;
      case 'north-west':
        dInt = 7;
        d = Direction.northWest;
        break;
    }
    for (Move move in moves) {
      if (id == move.piece.id && d == move.direction) {
        int player = 0;
        if (state.playerForPiece(move.piece) == PlayerID.player1) {
          player = 0;
        }
        if (state.playerForPiece(move.piece) == PlayerID.player2) {
          player = 1;
        }
        Cell c = state.cellForPiece(move.piece);
        _moves[id] = {
          'direction': dInt,
          'piece': {
            'id': move.piece.id,
            'player': player,
            'life': move.piece.life,
            'damage': move.piece.damage,
            'cell': [c.row, c.column]
          }
        };
      }
    }
    changed.emit({'moves': new List.from(_moves.values)});
  }

  bool isChecked(int id, String direction) {
    if (_decidingState == null || _decidingState != state) {
      _decidingState = state;
      _moves = {};
    }
    int dInt;
    switch (direction) {
      case 'north':
        dInt = 0;
        break;
      case 'north-east':
        dInt = 1;
        break;
      case 'east':
        dInt = 2;
        break;
      case 'south-east':
        dInt = 3;
        break;
      case 'south':
        dInt = 4;
        break;
      case 'south-west':
        dInt = 5;
        break;
      case 'west':
        dInt = 6;
        break;
      case 'north-west':
        dInt = 7;
        break;
    }
    if (!_moves.containsKey(id)) {
      return false;
    }
    return _moves[id]['direction'] == dInt;
  }
}
