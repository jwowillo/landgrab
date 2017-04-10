import 'package:angular2/core.dart';
import 'package:angular2/router.dart';

import 'package:landgrab/component/player_choice_form/component.dart';
import 'package:landgrab/model/player.dart';

/// GameFormComponent allows options to be selected before routing to the
/// GameComponent.
///
/// These options include what the Player 1 and Player 2 are and if the
/// GameComponent should prompt the user before continuing to the next turn.
@Component(
  selector: 'game-form',
  templateUrl: 'template.html',
  directives: const [PlayerChoiceFormComponent, ROUTER_DIRECTIVES],
)
class GameFormComponent {
  /// player1's PlayerID.
  final PlayerID player1 = PlayerID.player1;
  /// player2's PlayerID.
  final PlayerID player2 = PlayerID.player2;
}
