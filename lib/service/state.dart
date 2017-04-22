import 'dart:async';

import 'package:angular2/core.dart';

import 'package:landgrab/service/api.dart';
import 'package:landgrab/service/convert.dart';
import 'package:landgrab/model/player.dart';
import 'package:landgrab/model/state.dart';

/// StateService is used to get initial States for games with provided Players
/// and get the next States for given States.
@Injectable()
class StateService {
  final APIService _service;

  StateService(this._service);

  /// initial State for a game with the given Players.
  Future<State> initial(Player p1, Player p2) async {
    String s1 = encode(playerToMap(p1));
    String s2 = encode(playerToMap(p2));
    Map<String, dynamic> json =
        await _service.request('/new', query: {'player1': s1, 'player2': s2});
    return mapToState(json['data']);
  }

  /// next State for the given State.
  Future<State> next(State s) async {
    String serialized = encode(stateToMap(s));
    Map<String, dynamic> json =
        await _service.request('/next', query: {'state': serialized});
    return mapToState(json['data']);
  }
}
