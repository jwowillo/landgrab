import 'package:angular2/core.dart';

import 'package:landgrab/model/player.dart';
import 'package:landgrab/pipe/player_id_to_string.dart';
import 'package:landgrab/pipe/no_space.dart';

/// PlayerChoiceFormComponent allows a Player to be chosen from a set of
/// possible Players.
///
/// The Players are retrieved from the API server and displayed with radio
/// buttons showing their names and descriptions. A PlayerID must be passed so
/// the component knows which user the Player is representing.
@Component(
  selector: 'player-choice-form',
  templateUrl: 'template.html',
  pipes: const [PlayerIDToStringPipe, NoSpacePipe],
)
class PlayerChoiceFormComponent {
  /// id the chosen Player will have.
  @Input()
  PlayerID id;

  @Input()
  Set<Player> players = new Set();

  /// chosen emits Players whenever a new Player is chosen.
  @Output()
  final EventEmitter<Player> chosen = new EventEmitter();

  String url = '';

  emit(Player p) {
    if (p.name == 'api') {
      p.arguments['url'] = url;
    }
    chosen.emit(p);
  }
}
