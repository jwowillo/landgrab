import 'package:angular2/core.dart';

import 'package:landgrab/landgrab.dart';

/// mapToResource
Resource mapToResource(Map<String, dynamic> map) {
  return new Resource(map['name'], map['description'], map['structure']);
}

/// mapToAction
Action mapToAction(Map<String, dynamic> map) {
  return new Action(
      map['path'], mapToControllerDescription(map['description']));
}

/// mapToControllerDescription
ControllerDescription mapToControllerDescription(Map<String, dynamic> map) {
  MethodDescription getMethod;
  if (map.containsKey('get')) {
    getMethod = mapToMethodDescription('get', map['get']);
  }
  MethodDescription postMethod;
  if (map.containsKey('post')) {
    postMethod = mapToMethodDescription('post', map['post']);
  }
  MethodDescription putMethod;
  if (map.containsKey('put')) {
    putMethod = mapToMethodDescription('put', map['put']);
  }
  MethodDescription deleteMethod;
  if (map.containsKey('delete')) {
    deleteMethod = mapToMethodDescription('delete', map['delete']);
  }
  MethodDescription optionsMethod;
  if (map.containsKey('options')) {
    optionsMethod = mapToMethodDescription('options', map['options']);
  }
  MethodDescription headMethod;
  if (map.containsKey('head')) {
    headMethod = mapToMethodDescription('head', map['head']);
  }
  MethodDescription connectMethod;
  if (map.containsKey('connect')) {
    connectMethod = mapToMethodDescription('connect', map['connect']);
  }
  return new ControllerDescription(
      getMethod: getMethod,
      postMethod: postMethod,
      putMethod: putMethod,
      deleteMethod: deleteMethod,
      optionsMethod: optionsMethod,
      headMethod: headMethod,
      connectMethod: connectMethod);
}

/// mapToMethodDescription
MethodDescription mapToMethodDescription(
    String method, Map<String, dynamic> map) {
  Map<String, String> urlArguments;
  if (map.containsKey('urlArguments')) {
    urlArguments = map['urlArguments'];
  }
  Map<String, String> formArguments;
  if (map.containsKey('formArguments')) {
    formArguments = map['formArguments'];
  }
  String response;
  if (map.containsKey('response')) {
    response = map['response'];
  }
  String authentication;
  if (map.containsKey('authentication')) {
    authentication = map['authentication'];
  }
  String limiting;
  if (map.containsKey('limiting')) {
    limiting = map['limiting'];
  }
  return new MethodDescription(method,
      urlArguments: urlArguments,
      formArguments: formArguments,
      response: response,
      authentication: authentication,
      limiting: limiting);
}

/// mapToState converts a map of the form the API server returns for States
/// to a State.
State mapToState(Map<String, dynamic> map) {
  PlayerID current = stringToPlayerID(map['currentPlayer']);
  Player p1 = mapToPlayer(map['player1']);
  Player p2 = mapToPlayer(map['player2']);
  Set<Piece> player1Pieces = new Set();
  Set<Piece> player2Pieces = new Set();
  Map<Cell, Piece> cells = new Map();
  for (Map<String, dynamic> piece in map['pieces']) {
    Cell cell = new Cell(piece['cell'][0], piece['cell'][1]);
    Piece built = new Piece(piece['id'], piece['life'], piece['damage']);
    if (stringToPlayerID(piece['player']) == PlayerID.player1) {
      player1Pieces.add(built);
    }
    if (stringToPlayerID(piece['player']) == PlayerID.player2) {
      player2Pieces.add(built);
    }
    cells[cell] = built;
  }
  Rules rules = mapToRules(map['rules']);
  PlayerID winner = PlayerID.noPlayer;
  if (map.containsKey('winner')) {
    winner = stringToPlayerID(map['winner']);
  }
  return new State(rules, current, p1, p2, player1Pieces, player2Pieces, cells,
      winner: winner);
}

/// stateToMap converts a State to a Map of the form the API server expects.
Map<String, dynamic> stateToMap(State s) {
  Map<String, dynamic> map = {};
  map['currentPlayer'] = playerIDToString(s.currentPlayer);
  if (s.winner != PlayerID.noPlayer) {
    map['winner'] = playerIDToString(s.winner);
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
    piece['player'] = playerIDToString(s.playerForPiece(p));
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

List<Move> mapToMoves(Map<String, dynamic> m) {
  List<Move> moves = [];
  for (Map<String, dynamic> move in m['moves']) {
    Piece piece = new Piece(
        move['piece']['id'], move['piece']['life'], move['piece']['damage']);
    moves.add(new Move(stringToDirection(move['direction']), piece));
  }
  return moves;
}
