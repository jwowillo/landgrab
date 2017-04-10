import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:angular2/core.dart';

import 'package:landgrab/service/config.dart';
import 'package:landgrab/model/player.dart';

/// PlayersService fetches all available Players from the API server.
@Injectable()
class PlayersService {
  /// players available to be chosen.
  ///
  /// The Set is empty until load is called to initialize the class.
  final players = new Set<Player>();

  /// load the players from the API server.
  ///
  /// This can be called multiple times but the originally loaded Players will
  /// remain.
  Future load() async {
    if (players.isNotEmpty) return;
    var raw = await HttpRequest.getString(Config.API_URL + '/players');
    var json = JSON.decode(raw);
    for (Map<String, String> player in json['data']['players']) {
      players
          .add(new Player(player['name'], description: player['description']));
    }
  }
}
