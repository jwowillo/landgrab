import 'dart:async';

import 'package:angular2/core.dart';

import 'package:landgrab/service/api.dart';
import 'package:landgrab/model/player.dart';

/// PlayersService fetches all available Players from the API server.
@Injectable()
class PlayersService {
  /// players from the API server.
  ///
  /// This can be called multiple times but the originally loaded Players will
  /// remain.
  Future<Set<Player>> players() async {
    Set<Player> players = new Set();
    Map<String, dynamic> json = await api('/players');
    for (Map<String, String> player in json['data']['players']) {
      players.add(new Player(
        player['name'],
        description: player['description'],
      ));
    }
    return players;
  }
}
