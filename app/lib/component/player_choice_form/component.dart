import 'dart:async';

import 'package:angular2/core.dart';

import 'package:landgrab/model/player.dart';
import 'package:landgrab/pipe/player_id_to_string.dart';
import 'package:landgrab/pipe/to_lower_no_space.dart';
import 'package:landgrab/service/players_service.dart';

// TODO: Use @Output to set output values.

@Component(
  selector: 'player-choice-form',
  templateUrl: 'template.html',
  providers: const [PlayersService],
  pipes: const [PlayerIDToStringPipe, ToLowerNoSpace],
)
class PlayerChoiceFormComponent implements OnInit {
  @Input()
  PlayerID id;
  final PlayersService _service;

  PlayerChoiceFormComponent(this._service);

  @override
  Future ngOnInit() async {
    try {
      await _service.load();
    } catch (error) {
      print(error);
    }
  }

  Set<Player> players() => _service.players;
}
