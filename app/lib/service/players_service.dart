import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:angular2/core.dart';

import 'package:landgrab/model/player.dart';

@Injectable()
class PlayersService {
  final players = <Player>[];

  Future load() async {
    if (players.isNotEmpty) return;
    var rawData = await HttpRequest.getString('{{ api }}/players');
    var data = JSON.decode(rawData);
    for (var player in data['data']['players']) {
      players.add(new Player(player['name'], player['description']));
    }
  }
}
