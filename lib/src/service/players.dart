import 'dart:async';

import 'package:angular2/core.dart';

import 'package:landgrab/landgrab.dart';

/// PlayersService fetches all available Players from the API server.
@Injectable()
class PlayersService {

  final APIService _service;

  static Set<Player> _players = new Set();

  PlayersService(this._service);

  /// players from the API server.
  ///
  /// This can be called multiple times but the originally loaded Players will
  /// remain.
  Future<Set<Player>> players() async {
    if (!_players.isEmpty) {
      return _players;
    }
    Map<String, dynamic> json = await _service.request('/players');
    for (Map<String, String> player in json['data']['players']) {
      _players.add(mapToPlayer(player));
    }
    return _players;
  }
}
