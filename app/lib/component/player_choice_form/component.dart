import 'dart:async';

import 'package:angular2/core.dart';

import 'package:landgrab/model/player.dart';
import 'package:landgrab/pipe/player_id_to_string.dart';
import 'package:landgrab/pipe/to_lower_no_space.dart';
import 'package:landgrab/service/players_service.dart';

// TODO: Use @Output to set output values.

/// PlayerChoiceFormComponent allows a Player to be chosen from a set of
/// possible Players.
///
/// The Players are retrieved from the API server and displayed with radio
/// buttons showing their names and descriptions. A PlayerID must be passed so
/// the component knows which user the Player is representing.
@Component(
  selector: 'player-choice-form',
  templateUrl: 'template.html',
  providers: const [PlayersService],
  pipes: const [PlayerIDToStringPipe, ToLowerNoSpace],
)
class PlayerChoiceFormComponent implements OnInit {
  /// id the chosen Player will have.
  @Input()
  PlayerID id;

  /// _service to retrieve Players from.
  final PlayersService _service;

  /// PlayerChoiceFormComponent initializes the PlayersService.
  PlayerChoiceFormComponent(this._service);

  /// ngOnInit loads the PlayersService so the Players are available.
  ///
  /// An error is logged to the console if this fais.
  @override
  Future ngOnInit() async {
    try {
      await _service.load();
    } catch (error) {
      print(error);
    }
  }

  /// players is a getter for the PlayersService's Players for use from the
  /// template.
  Set<Player> players() => _service.players;
}
