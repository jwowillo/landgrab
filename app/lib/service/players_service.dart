import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:angular2/core.dart';

import 'package:landgrab/config/api.dart';
import 'package:landgrab/model/player.dart';

@Injectable()
class PlayersService {
  final players = new Set<Player>();

  Future load() async {
    if (players.isNotEmpty) return;
    var raw = await HttpRequest.getString(APIConfig.URL + '/players');
    var json = JSON.decode(raw);
    for (Map<String, String> player in json['data']['players']) {
      players.add(new Player(player['name'], player['description']));
    }
  }
}
