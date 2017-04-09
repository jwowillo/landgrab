import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:angular2/core.dart';

import 'package:landgrab/service/config.dart';
import 'package:landgrab/model/board.dart';
import 'package:landgrab/model/player.dart';
import 'package:landgrab/model/state.dart';

/// StateService is used to get initial States for games with provided Players
/// and get the next States for given States.
@Injectable
class StateService {
  /// initial State for a game with the given Players.
  Future<State> initial(Player p1, Player p2) async {
    Map<String, dynamic> json =
        await _api('/new', {'player1': p1.name, 'player2': p2.name});
    return _mapToState(json['data']);
  }

  /// Next State for the given State.
  Future<State> next(State s) async {
    String serialized =
        Uri.encodeQueryComponent(JSONEncoder.convert(_stateToMap(s)));
    String serialized = '';
    Map<String, dynamic> json = await _api('/next', {'state': serialized});
    return _mapToState(json['data']);
  }

  /// _mapToState converts a map of the form the API server returns for States
  /// to a State.
  State _mapToState(Map<String, dynamic> map) {
    PlayerID current = PlayerID.noPlayer;
    if (map['currentPlayer'] == 1) {
      current = PlayerID.player1;
    }
    if (map['currentPlayer'] == 2) {
      current = PlayerID.player2;
    }
    Player p1 =
        new Player(map['player1']['name'], map['player1']['description']);
    Player p2 =
        new Player(map['player2']['name'], map['player2']['description']);
    Set<Piece> player1Pieces = new Set();
    Set<Piece> player2Pieces = new Set();
    Map<Cell, Piece> cells = new Map();
    for (String id in map['pieces'].keys) {
      Map<String, dynamic> piece = map['pieces'][id];
      Cell cell = new Cell(piece['cell'][0], piece['cell'][1]);
      Piece built = new Piece(int.parse(id), piece['life'], piece['damage']);
      if (piece['player'] == 1) {
        player1Pieces.add(built);
      }
      if (piece['player'] == 2) {
        player2Pieces.add(built);
      }
      cells[cell] = built;
    }
    int size = map['boardSize'];
    PlayerID winner = PlayerID.noPlayer;
    if (map.containsKey('winner')) {
      if (map['winner'] == 1) {
        winner = PlayerID.player1;
      }
      if (map['winner'] == 2) {
        winner = PlayerID.player2;
      }
    }
    return new State(current, p1, p2, player1Pieces, player2Pieces, cells, size,
        winner: winner);
  }

  /// _stateToMap converts a State to a Map of the form the API server expects.
  Map<String, dynamic> _stateToMap(State s) {
    // TODO: Implement.
    return {};
  }

  /// _api makes a request to the path at the API server with the given query
  /// string.
  Future<Map<String, dynamic>> _api(
      String path, Map<String, String> query) async {
    Uri uri = new Uri(path: Config.API_URL + path, queryParameteters: query);
    String raw = await HttpRequest.getString(uri.toString());
    return JSON.decode(raw);
  }
}
