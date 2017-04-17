import 'package:angular2/core.dart';

import 'package:landgrab/model/board.dart';
import 'package:landgrab/model/player.dart';
import 'package:landgrab/model/rules.dart';
import 'package:landgrab/model/state.dart';

/// mapToState converts a map of the form the API server returns for States
/// to a State.
State mapToState(Map<String, dynamic> map) {
  PlayerID current = PlayerID.noPlayer;
  if (map['currentPlayer'] == 1) {
    current = PlayerID.player1;
  }
  if (map['currentPlayer'] == 2) {
    current = PlayerID.player2;
  }
  Player p1 = mapToPlayer(map['player1']);
  Player p2 = mapToPlayer(map['player2']);
  Set<Piece> player1Pieces = new Set();
  Set<Piece> player2Pieces = new Set();
  Map<Cell, Piece> cells = new Map();
  for (Map<String, dynamic> piece in map['pieces']) {
    Cell cell = new Cell(piece['cell'][0], piece['cell'][1]);
    Piece built = new Piece(piece['id'], piece['life'], piece['damage']);
    if (piece['player'] == 1) {
      player1Pieces.add(built);
    }
    if (piece['player'] == 2) {
      player2Pieces.add(built);
    }
    cells[cell] = built;
  }
  Rules rules = mapToRules(map['rules']);
  PlayerID winner = PlayerID.noPlayer;
  if (map.containsKey('winner')) {
    if (map['winner'] == 1) {
      winner = PlayerID.player1;
    }
    if (map['winner'] == 2) {
      winner = PlayerID.player2;
    }
  }
  return new State(rules, current, p1, p2, player1Pieces, player2Pieces, cells,
      winner: winner);
}

/// stateToMap converts a State to a Map of the form the API server expects.
Map<String, dynamic> stateToMap(State s) {
  Map<String, dynamic> map = {};
  if (s.currentPlayer == PlayerID.player1) {
    map['currentPlayer'] = 1;
  }
  if (s.currentPlayer == PlayerID.player2) {
    map['currentPlayer'] = 2;
  }
  if (s.winner != PlayerID.noPlayer) {
    if (s.winner == PlayerID.player1) {
      map['winner'] = 1;
    }
    if (s.winner == PlayerID.player2) {
      map['winner'] = 2;
    }
  }
  map['rules'] = rulesToMap(s.rules);
  List<Map<String, dynamic>> pieces = [];
  Set<Piece> allPieces = new Set();
  allPieces.addAll(s.player1Pieces);
  allPieces.addAll(s.player2Pieces);
  for (Piece p in allPieces) {
    if (p.id == -1) {
      continue;
    }
    Map<String, dynamic> piece = {};
    piece['id'] = p.id;
    piece['damage'] = p.damage;
    piece['life'] = p.life;
    if (s.playerForPiece(p) == PlayerID.player1) {
      piece['player'] = 1;
    }
    if (s.playerForPiece(p) == PlayerID.player2) {
      piece['player'] = 2;
    }
    Cell c = s.cellForPiece(p);
    piece['cell'] = [c.row, c.column];
    pieces.add(piece);
  }
  map['pieces'] = pieces;
  map['player1'] = playerToMap(s.player1);
  map['player2'] = playerToMap(s.player2);
  return map;
}

/// mapToRules converts the Map to Rules.
Rules mapToRules(Map<String, dynamic> m) {
  return new Rules(
    new Duration(seconds: m['timerDuration']),
    m['pieceCount'],
    m['boardSize'],
    m['life'],
    m['damage'],
    m['lifeIncrease'],
    m['damageIncrease'],
  );
}

/// rulesToMap converts the Rules to a Map.
Map<String, dynamic> rulesToMap(Rules r) {
  return {
    'timerDuration': r.timerDuration.inSeconds,
    'pieceCount': r.pieceCount,
    'boardSize': r.boardSize,
    'life': r.life,
    'damage': r.damage,
    'lifeIncrease': r.lifeIncrease,
    'damageIncrease': r.damageIncrease,
  };
}

Map<String, dynamic> playerToMap(Player p) {
  Map<String, dynamic> map = {'name': p.name, 'description': p.description};
  if (p.arguments != null) {
    map['arguments'] = p.arguments;
  }
  return map;
}

Player mapToPlayer(Map<String, dynamic> m) {
  String description = '';
  if (m.containsKey('description')) {
    description = m['description'];
  }
  Map<String, dynamic> args = {};
  if (m.containsKey('arguments')) {
    args = m['arguments'];
  }
  return new Player(m['name'], description: description, arguments: args);
}
