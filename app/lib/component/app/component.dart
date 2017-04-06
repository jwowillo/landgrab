import 'dart:async';

import 'package:angular2/core.dart';

import 'package:landgrab/model/player.dart';
import 'package:landgrab/service/players_service.dart';

@Component(
  selector: 'app',
  templateUrl: 'template.html',
  providers: const [PlayersService],
)
class AppComponent implements OnInit {
  final PlayersService _service;

  AppComponent(this._service);

  @override
  Future ngOnInit() async {
    try {
      await _service.load();
    } catch (error) {
      print(error);
    }
  }

  List<Player> players() {
    return _service.players;
  }
}
