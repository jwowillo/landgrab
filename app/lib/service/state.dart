import 'dart:async';
import 'dart:convert';

import 'package:angular2/core.dart';

import 'package:landgrab/service/api.dart';
import 'package:landgrab/service/convert.dart';
import 'package:landgrab/model/player.dart';
import 'package:landgrab/model/state.dart';

/// StateService is used to get initial States for games with provided Players
/// and get the next States for given States.
@Injectable()
class StateService {
  /// initial State for a game with the given Players.
  Future<State> initial(Player p1, Player p2) async {
    Map<String, dynamic> json =
        await api('/new', query: {'player1': p1.name, 'player2': p2.name});
    return mapToState(json['data']);
  }

  /// Next State for the given State.
  Future<State> next(State s) async {
    String serialized =
        Uri.encodeQueryComponent(new JsonEncoder().convert(stateToMap(s)));
    Map<String, dynamic> json =
        await api('/next', query: {'state': serialized});
    return mapToState(json['data']);
  }
}
